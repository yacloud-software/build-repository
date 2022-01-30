package diskscanner

import (
	"flag"
	"fmt"
	"golang.conradwood.net/go-easyops/linux"
	"golang.conradwood.net/go-easyops/utils"
	"io/ioutil"
	"os"
	"strconv"
	"time"
)

var (
	debug       = flag.Bool("diskscanner_debug", false, "diskscanner debug mode")
	backup      = flag.Bool("diskscanner_backup", true, "run backups of everything regularly and prior to archiving")
	sleep       = flag.Int("diskscanner_sleep", 60, "amount of `seconds` between checks of diskspace")
	max_runtime = flag.Int("diskscanner_max_runtime", 600, "amount of `seconds` before rsync is forcibly killed")
	unclean     = true
)

type DiskScanner struct {
	lastRun time.Time
	running bool
	Dir     string
	MaxSize int64 // MegaBytes
	ch      chan int
	Builds  *BuildDir
}

func NewDiskScanner() *DiskScanner {
	res := &DiskScanner{
		ch:      make(chan int, 100),
		MaxSize: 1024 * 100, // max 100G
	}
	return res
}
func (d *DiskScanner) Start() {
	d.Trigger()
	go d.loop()
}
func (d *DiskScanner) loop() {
	go d.find()
	for {
		time.Sleep(time.Duration(*sleep) * time.Second)
		if len(d.ch) > 0 && d.running {
			continue
		}
		if len(d.ch) > 1 { // one queued already? do not queue another
			continue
		}
		if time.Since(d.lastRun) < (time.Duration(*sleep) * time.Second) {
			// has it ran meanwhile? if so, don't trigger
			continue
		}
		d.ch <- 1
	}
}
func (d *DiskScanner) Unclean() {
	unclean = true
}
func (d *DiskScanner) find() {
	var err error
	for {
		d.running = false
		<-d.ch
		d.running = true
		if d.Dir == "" {
			fmt.Printf("[diskscanner] No dir set!\n")
			continue
		}
		if *backup && unclean {
			err = d.rsync()
			if err != nil {
				fmt.Printf("[diskscanner] Failed to archive (rsync): %s\n", err)
				continue
			}
			unclean = false
		}
		if *debug {
			fmt.Printf("[diskscanner] calculating size...\n")
		}

		d.Builds, err = d.calc()
		if err != nil {
			fmt.Printf("[diskscanner] Failed: %s\n", err)
			continue
		}
		if *debug {
			for _, b := range d.Builds.repos {
				fmt.Printf("[diskscanner] Repo: %s (%d branches, %d versions, %16d bytes)\n", b.name, len(b.branches), len(b.Versions()), b.Size())
			}
		}
		maxBytes := uint64(d.MaxSize * 1024 * 1024)
		if d.Builds.Size() < maxBytes {
			continue
		}
		fmt.Printf("[diskscanner] Too big (%d Gb)\n", d.MaxSize/1024)
		versions := d.Builds.Archivable()
		fmt.Printf("[diskscanner] %d versions to archive\n", len(versions))
		for i, v := range versions {
			if d.Builds.Size() < maxBytes {
				break
			}
			err = sync_to_archive(v)
			if err != nil {
				fmt.Printf("Error syncing: %s\n", utils.ErrorString(err))
				break
			}
			fmt.Printf("[diskscanner] %3d. Version %d in %s (%v) (size=%dGb)\n", i, v.version, v.Path(), v.Created(), d.Builds.Size()/1024/1024/1024)
			err = os.RemoveAll(v.Path())
			if err != nil {
				fmt.Printf("[diskscanner] Failed to remove version (%s): %s", v.Path(), err)
				continue
			}
			v.deleted = true
		}
	}
}

func (d *DiskScanner) rsync() error {
	if *debug {
		fmt.Printf("[diskscanner] Running backup of %s...\n", d.Dir)
	}
	l := linux.New()
	l.SetRuntime(*max_runtime)
	foo, err := l.SafelyExecute([]string{"rsync", "-pra", d.Dir, "rsync://johnsmith/buildrepo/"}, nil)
	if err != nil {
		fmt.Println(foo)
		fmt.Printf("[diskscanner] Failed: %s\n", err)
		return err
	}
	return nil
}
func (d *DiskScanner) Trigger() {
	d.ch <- 1
}
func (d *DiskScanner) calc() (*BuildDir, error) {
	res := &BuildDir{root: d.Dir}

	//repos
	files, err := ioutil.ReadDir(res.root)
	if err != nil {
		return nil, err
	}
	for _, f := range files {
		n := f.Name()
		r := &Repo{name: n, builddir: res}
		res.repos = append(res.repos, r)
	}

	// branches
	for _, r := range res.repos {
		files, err := ioutil.ReadDir(fmt.Sprintf("%s/%s", res.root, r.name))
		if err != nil {
			return nil, err
		}
		for _, f := range files {
			b := &Branch{name: f.Name(), repo: r}
			r.branches = append(r.branches, b)
		}
	}

	// versions
	for _, r := range res.repos {
		for _, b := range r.branches {
			files, err := ioutil.ReadDir(fmt.Sprintf("%s/%s/%s", res.root, r.name, b.name))
			if err != nil {
				return nil, err
			}
			for _, f := range files {
				n := f.Name()
				if n == "latest" {
					continue
				}
				u, err := strconv.Atoi(n)
				if err != nil {
					fmt.Printf("[diskscanner] Not a version: %s (filename: %s/%s/%s/%s)\n", err, res.root, r.name, b.name, n)
					continue
				}
				v := &Version{version: u, branch: b}
				b.versions = append(b.versions, v)
			}
		}
	}

	return res, nil
}

func sync_to_archive(v *Version) error {
	fmt.Printf("[diskscanner] syncing %s, %s\n", v.branch.repo.builddir, v.branch.name)
	return nil
}

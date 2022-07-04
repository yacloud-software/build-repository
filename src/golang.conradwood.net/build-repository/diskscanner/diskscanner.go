package diskscanner

import (
	"flag"
	"fmt"
	"golang.conradwood.net/go-easyops/authremote"
	"golang.conradwood.net/go-easyops/linux"
	"golang.conradwood.net/go-easyops/prometheus"
	"golang.conradwood.net/go-easyops/utils"
	ba "golang.yacloud.eu/apis/buildrepoarchive"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"time"
)

var (
	domain_id      = flag.String("domain_id", "default_buildrepo_service", "domain id for buildrepo")
	do_remove      = flag.Bool("do_remove", true, "if true actually delete archived stuff")
	debug          = flag.Bool("diskscanner_debug", false, "diskscanner debug mode")
	backup         = flag.Bool("diskscanner_backup", true, "run backups of everything regularly and prior to archiving")
	sleep          = flag.Int("diskscanner_sleep", 60, "amount of `seconds` between checks of diskspace")
	sleep_fail     = flag.Duration("diskscanner_sleep_fail", time.Duration(60)*time.Minute, "sleep  between checks of diskspace, in fail mode")
	max_runtime    = flag.Int("diskscanner_max_runtime", 600, "amount of `seconds` before rsync is forcibly killed")
	do_enable      = flag.Bool("diskscanner_enable", true, "if false, do not run diskscanner")
	unclean        = true
	sl             = utils.NewSlidingAverage()
	fail_mode      = false // if true, sleep a long time and be unhappy
	prom_fail_mode = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "buildrepo_diskscanner_fail_mode",
			Help: "V=1 UNIT=none DESC=gauge indicating if in failmode",
		},
	)
	prom_disk_size = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "buildrepo_diskscanner_size",
			Help: "V=1 UNIT=decbytes DESC=gauge indicating size of archive",
		},
	)
	prom_syncs = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "buildrepo_diskscanner_total_syncs",
			Help: "V=1 UNIT=none DESC=total sync attempts",
		},
	)
	prom_sync_fails = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "buildrepo_diskscanner_failed_syncs",
			Help: "V=1 UNIT=none DESC=total failed attempts",
		},
	)
)

func init() {
	prometheus.MustRegister(prom_fail_mode, prom_disk_size, prom_syncs, prom_sync_fails)
}

type DiskScanner struct {
	lastRun time.Time
	running bool
	Dir     string
	MaxSize int64 // MegaBytes
	ch      chan int
	Builds  *BuildDir
	sl      *utils.SlidingAverage
}

func NewDiskScanner() *DiskScanner {
	res := &DiskScanner{
		ch:      make(chan int, 100),
		MaxSize: 1024 * 100, // max 100G
	}
	return res
}
func printRate() {
	r := float64(sl.GetCounter(1)) / float64(sl.GetCounter(0)) * 100
	fmt.Printf("Total syncs: %d, Failed syncs: %d, Percent:%0.2f\n", sl.GetCounter(0), sl.GetCounter(1), r)
}
func (d *DiskScanner) Start() {
	d.Trigger()
	go d.loop()
}
func (d *DiskScanner) Unfail() {
	if fail_mode {
		fmt.Printf("[diskscanner] unfailed by user\n")
	} else {
		fmt.Printf("[diskscanner] user request to unfail, but not in fail_mode\n")
	}
	fail_mode = false
}
func sleep_while_fail() {
	if !fail_mode {
		return
	}
	started := time.Now()
	for {
		fmt.Printf("[diskscanner] diskscanner failed. fail_sleep_mode\n")
		time.Sleep(time.Duration(10) * time.Second)
		dur := time.Since(started)
		if dur > *sleep_fail {
			break
		}
		if !fail_mode {
			break
		}
	}
}

func (d *DiskScanner) loop() {
	go d.find()
	for {
		time.Sleep(time.Duration(*sleep) * time.Second)
		sleep_while_fail()
		if sl.GetCounter(1) == sl.GetCounter(0) && sl.GetCounter(0) > 0 {
			fail_mode = true
			prom_fail_mode.Set(1)
		} else {
			prom_fail_mode.Set(0)
			fail_mode = false
		}

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
		if !*do_enable {
			continue
		}
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
		prom_disk_size.Set(float64(d.Builds.Size()))
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
			prom_syncs.Inc()
			err = sync_to_archive(v)
			sl.Add(0, 1)
			if err != nil {
				prom_sync_fails.Inc()
				sl.Add(1, 1)
				fmt.Printf("Error syncing: %s\n", utils.ErrorString(err))
				break
			}
			printRate()
			fmt.Printf("[diskscanner] %3d. Version %d in %s (%v) (size=%dGb)\n", i, v.version, v.Path(), v.Created(), d.Builds.Size()/1024/1024/1024)
			if *do_remove {
				err = os.RemoveAll(v.Path())
				if err != nil {
					fmt.Printf("[diskscanner] Failed to remove version (%s): %s", v.Path(), err)
					continue
				}
				v.deleted = true
			}
		}
		printRate()

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
	bc := ba.GetBuildRepoArchiveClient()
	key := fmt.Sprintf("%s/%s/%d", v.branch.repo.name, v.branch.name, v.version)
	ctx := authremote.ContextWithTimeout(time.Duration(60) * time.Second)
	srv, err := bc.Upload(ctx)
	if err != nil {
		return err
	}
	err = srv.Send(&ba.UploadRequest{DomainID: *domain_id, Key: key})
	if err != nil {
		return err
	}

	dirs := []string{
		fmt.Sprintf("artefacts/%s/%s/%d", v.branch.repo.name, v.branch.name, v.version),
		fmt.Sprintf("metadata/%s/%s/%d", v.branch.repo.name, v.branch.name, v.version),
	}
	for _, local_dir := range dirs {
		// fmt.Printf("[diskscanner] Uploading %s...\n", local_dir)

		err = upload_dir(srv, local_dir)
		if err != nil {
			fmt.Printf("[diskscanner] Upload dir %s failed: %s\n", local_dir, utils.ErrorString(err))
			return err
		}
	}
	_, err = srv.CloseAndRecv()
	if err != io.EOF {
		fmt.Printf("[diskscanner] closeandrecv() failed: %s\n", utils.ErrorString(err))
		return err
	}

	fmt.Printf("[diskscanner] syncing %s, %s, %s, version=%d\n", v.branch.repo.builddir.root, v.branch.repo.name, v.branch.name, v.version)
	return nil
}
func upload_dir(srv ba.BuildRepoArchive_UploadClient, dir string) error {
	fdir := "/srv/build-repository/" + dir
	sym, err := isSymLink(fdir)
	if err != nil || sym {
		return err
	}
	files, err := ioutil.ReadDir(fdir)
	if err != nil {
		fmt.Printf("[diskscanner] readdir \"%s\" failed: %s\n", dir, err)
		return err
	}
	for _, f := range files {
		if f.IsDir() {
			// files first
			continue
		}
		ffname := dir + "/" + f.Name()
		//fmt.Printf("Uploading file %s...\n", ffname)
		err := upload_file(srv, ffname)
		if err != nil {
			fmt.Printf("[diskscanner] upload of file \"%s\" failed: %s\n", ffname, err)
			return err
		}
	}
	for _, f := range files {
		if !f.IsDir() {
			// dirs only
			continue
		}
		ffname := dir + "/" + f.Name()

		err := upload_dir(srv, ffname)
		if err != nil {
			fmt.Printf("[diskscanner] upload of dir \"%s\" failed: %s\n", ffname, err)
			return err
		}
	}
	return nil
}
func isSymLink(filename string) (bool, error) {
	// if file is a symlink, don't sync it
	fileInfo, err := os.Lstat(filename)
	if err != nil {
		fmt.Printf("[diskscanner] stat() of file \"%s\" failed: %s\n", filename, err)
		return false, err
	}

	if fileInfo.Mode()&os.ModeSymlink == os.ModeSymlink {
		fmt.Printf("File %s is a symbolic link, skipped.\n", filename)
		return true, nil
	}
	return false, nil
}
func upload_file(srv ba.BuildRepoArchive_UploadClient, filename string) error {
	fullfile := "/srv/build-repository/" + filename
	sym, err := isSymLink(fullfile)
	if err != nil || sym {
		return err
	}
	f, err := os.Open(fullfile)
	if err != nil {
		fmt.Printf("[diskscanner] open() of file \"%s\" failed: %s\n", fullfile, err)
		return err
	}
	defer f.Close()
	buf := make([]byte, 1024)
	for {
		n, err := f.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Printf("[diskscanner] read() of file \"%s\" failed: %s\n", filename, err)
			return err
		}
		upr := &ba.UploadRequest{
			Filename: filename,
			Data:     buf[:n],
		}
		err = srv.Send(upr)
		if err != nil {
			fmt.Printf("[diskscanner] send() of file \"%s\" failed: %s\n", filename, err)
			return err
		}
	}
	return nil
}

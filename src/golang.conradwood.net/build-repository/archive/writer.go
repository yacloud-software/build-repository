package archive

import (
	"fmt"
	"golang.conradwood.net/go-easyops/utils"
	"os"
	"path/filepath"
)

type writer struct {
	cur_filename string
	targetdir    string
	cur_file     *os.File
	bytes        uint64
}

func (w *writer) NewFile(filename string) error {
	if w.cur_filename == filename {
		return nil
	}
	err := w.Close()
	if err != nil {
		return err
	}
	w.bytes = 0
	w.cur_filename = filename
	fullfile := w.targetdir + filename
	p := filepath.Dir(fullfile)
	if !utils.FileExists(p) {
		os.MkdirAll(p, 0777)
	}
	f, err := utils.OpenWriteFile(fullfile)
	if err != nil {
		return err
	}
	w.cur_file = f
	fmt.Printf("Opened %s\n", fullfile)
	return nil
}
func (w *writer) Write(buf []byte) error {
	if w.cur_file == nil {
		return fmt.Errorf("data received with no filename")
	}
	n, err := w.cur_file.Write(buf)
	if err != nil {
		return err
	}
	if n != len(buf) {
		return fmt.Errorf("write error: %d != %d", len(buf), n)
	}
	w.bytes = w.bytes + uint64(len(buf))
	return nil
}
func (w *writer) Close() error {
	if w.cur_file == nil {
		return nil
	}
	fmt.Printf("Closed \"%s\" (%d bytes)\n", w.cur_filename, w.bytes)
	err := w.cur_file.Close()
	if err != nil {
		return err
	}
	w.cur_file = nil
	return nil
}

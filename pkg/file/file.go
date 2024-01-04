package file

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type File struct {
	path       string
	Href, Name string
	IsDir      bool
	ModTime    time.Time
	Size       int64
}

const (
	_ = 1 << (10 * iota)
	KB
	MB
	GB
)

func FormatSize(size int64) string {
	switch {
	case size < KB:
		return fmt.Sprintf("%d B", size)
	case size < MB:
		return fmt.Sprintf("%.2f KB", float64(size)/float64(KB))
	case size < GB:
		return fmt.Sprintf("%.2f MB", float64(size)/float64(MB))
	default:
		return fmt.Sprintf("%.2f GB", float64(size)/float64(GB))
	}
}

func DirContent(path string) ([]File, error) {
	var c []File
	pathes, err := filepath.Glob(path + "/*")
	if err != nil {
		return nil, err
	}
	for _, p := range pathes {
		file, err := New(p)
		if err != nil {
			return nil, err
		}
		c = append(c, *file)
	}
	return c, nil
}

func (f *File) stat() error {
	stat, err := os.Stat(f.path)
	if err != nil {
		return fmt.Errorf("get file info: %w for file %s", err, f.path)
	}
	f.IsDir = stat.IsDir()
	f.ModTime = stat.ModTime()
	f.Size = stat.Size()
	return nil
}

func (f *File) href() error {
	base, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("get current folder: %w", err)
	}
	f.Href = strings.ReplaceAll(f.path, base, "")
	return nil
}

func New(_path string) (*File, error) {
	f := File{
		path: _path,
		Name: filepath.Base(_path),
	}
	if err := f.stat(); err != nil {
		return nil, err
	}
	if err := f.href(); err != nil {
		return nil, err
	}
	return &f, nil
}

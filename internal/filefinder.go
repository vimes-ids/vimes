package internal

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"os"
	"path/filepath"
	"time"
)

type File struct {
	name     string
	size     int
	md5hash  string
	modified time.Time
}

type FileFinder struct {
}

func NewFileFinder() *FileFinder {
	return &FileFinder{}
}

func (list *FileFinder) FindFiles() []File {
	found := []File{}

	// @TODO: Make sure we look at the configuration to see what paths to scan.
	// Regexes probably?
	err := filepath.Walk("/etc", func(path string, info os.FileInfo, err error) error {
		if err == nil {
			if !info.IsDir() {
				md5hash, err := list.md5sum(path)
				if err != nil {
					panic(err)
				}

				found = append(found, File{
					name:     path,
					size:     int(info.Size()),
					md5hash:  md5hash,
					modified: info.ModTime(),
				})

			}
		}

		return nil
	})

	if err != nil {
		panic(err)
	}

	for _, file := range found {
		println(file.name, " -> ", file.md5hash)
	}

	return found
}

func (list *FileFinder) md5sum(filePath string) (string, error) {
	file, err := os.Open(filePath)

	if err != nil {
		return "", err
	}
	defer file.Close()

	hash := md5.New()

	if _, err := io.Copy(hash, file); err != nil {
		return "", err
	}

	return hex.EncodeToString(hash.Sum(nil)), nil
}

package files

import (
	"context"
	"io/fs"
	"os"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
)

const DEFAULT_DIR_PERMISSION fs.FileMode = 0755
const DEFAULT_FILE_PERMISSION fs.FileMode = 0644

type (
	DirectoryWriter interface {
		MkdirAll(string, fs.FileMode) error // *PathError
	}

	DirectoryRemover interface {
		RemoveAll(string) error // *PathError
	}

	FileWriter interface {
		WriteFile(string, []byte, fs.FileMode) error
	}

	FileRemover interface {
		Remove(string) error // *PathError
	}

	FileReader interface {
		ReadDir(string) ([]fs.DirEntry, error)
		ReadFile(string) ([]byte, error)
	}

	Writer interface {
		DirectoryWriter
		FileWriter
	}

	Remover interface {
		FileRemover
		DirectoryRemover
	}

	Generator interface {
		Writer
		FileReader
		Remover
	}

	File struct {
		// Name of the file. i.e. main.go, utils
		Name string `json:"name" yaml:"name"`

		// Content of the file (used if the file is not a directory)
		Content string `json:"content,omitempty" yaml:"content,omitempty"`

		// Defines whether the file is a directory type
		IsDirectory bool `json:"is_directory,omitempty" yaml:"is_directory,omitempty"`

		// If the file is a directory type, this is the set of files contained within it.
		// If the file is not a directory type, setting this field has not effect.
		Files []File `json:"files,omitempty" yaml:"files,omitempty"`
	}

	FileGenerator struct{}
)

// ......................................
// ......................................
// ......................................

func (F *File) Formatted() string { return F.Name }

func NewDirectory(name string, files ...File) File {
	return File{
		IsDirectory: true,
		Name:        name,
		Files:       files,
	}
}

func (f *File) Create(writer Writer, path ...string) (res []string) {
	name := strings.Join(append(path, f.Name), "/")

	if !f.IsDirectory {
		writer.WriteFile(name, []byte(f.Content), DEFAULT_FILE_PERMISSION)
		res = append(res, f.Name)
	} else {
		_, cancel := context.WithTimeout(context.TODO(), 30*time.Second)
		defer cancel()

		writer.MkdirAll(name, DEFAULT_DIR_PERMISSION)
		res = append(res, name)

		for _, file := range f.Files {
			res = append(res, file.Create(writer, append(path, f.Name)...)...)
		}
	}

	return res
}

func (f *File) Remove(remover Remover) error {
	if f.IsDirectory {
		return remover.RemoveAll(f.Name)
	}
	return remover.Remove(f.Name)
}

func (f *File) AddFiles(files ...File) { f.Files = append(f.Files, files...) }

func (f *File) Count() (count int) {
	count += len(f.Files)
	for _, dir := range f.Files {
		count += dir.Count()
	}
	return count
}

// ......................................
// ......................................
// ......................................

func (fg FileGenerator) WriteFile(name string, data []byte, perm fs.FileMode) error {
	log.Debug("Creating file", name)
	return os.WriteFile(name, data, perm)
}

func (fg FileGenerator) MkdirAll(path string, perm fs.FileMode) error {
	log.Debug("Creating directory", path)
	return os.MkdirAll(path, perm)
}

func (fg FileGenerator) Remove(name string) error {
	log.Debug("Removing", name)
	return os.Remove(name)
}

func (fg FileGenerator) RemoveAll(pathname string) error {
	log.Debug("Removing", pathname)
	return os.RemoveAll(pathname)
}

func (FileGenerator) ReadFile(name string) ([]byte, error) { return os.ReadFile(name) }

func (FileGenerator) ReadDir(name string) ([]fs.DirEntry, error) { return os.ReadDir(name) }

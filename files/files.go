package files

import (
	"context"
	"fmt"
	"io"
	"io/fs"
	"os"
	"strings"
	"time"

	"github.com/oleoneto/go-toolkit/logger"
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
		logger.LogWriter
	}

	Remover interface {
		FileRemover
		DirectoryRemover
		logger.LogWriter
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

	NullFileGenerator struct{}
)

// ......................................
// ......................................
// ......................................

func (F *File) Formatted() string {
	return F.Name
}

func NewDirectory(name string, files ...File) File {
	return File{
		IsDirectory: true,
		Name:        name,
		Files:       files,
	}
}

func (F *File) Create(writer Writer, path ...string) (res []string) {
	name := strings.Join(append(path, F.Name), "/")

	if !F.IsDirectory {
		writer.WriteFile(name, []byte(F.Content), DEFAULT_FILE_PERMISSION)
		res = append(res, F.Name)
	} else {
		_, cancel := context.WithTimeout(context.TODO(), 30*time.Second)
		defer cancel()

		writer.MkdirAll(name, DEFAULT_DIR_PERMISSION)
		res = append(res, name)

		for _, file := range F.Files {
			res = append(res, file.Create(writer, append(path, F.Name)...)...)
		}
	}

	return res
}

func (F *File) Remove(remover Remover) error {
	if F.IsDirectory {
		return remover.RemoveAll(F.Name)
	}

	return remover.Remove(F.Name)
}

func (F *File) AddFiles(files ...File) {
	F.Files = append(F.Files, files...)
}

func (F *File) Count() (count int) {
	count += len(F.Files)

	for _, dir := range F.Files {
		count += dir.Count()
	}

	return count
}

// ......................................
// ......................................
// ......................................

func (FileGenerator) Log(content any, w io.Writer, template string) {
	logger.NewDefaultLogger().Log(content, w, template)
}

func (G FileGenerator) WriteFile(name string, data []byte, perm fs.FileMode) error {
	G.Log(fmt.Sprintf(`Creating file %v`, name), os.Stdout, "")
	return os.WriteFile(name, data, perm)
}

func (G FileGenerator) MkdirAll(path string, perm fs.FileMode) error {
	G.Log(fmt.Sprintf(`Creating directory %v`, path), os.Stdout, "")
	return os.MkdirAll(path, perm)
}

func (G FileGenerator) Remove(name string) error {
	G.Log(fmt.Sprintf(`Removing %v`, name), os.Stdout, "")
	return os.Remove(name)
}

func (G FileGenerator) RemoveAll(pathname string) error {
	G.Log(fmt.Sprintf(`Removing %v`, pathname), os.Stdout, "")
	return os.RemoveAll(pathname)
}

func (FileGenerator) ReadFile(name string) ([]byte, error) {
	return os.ReadFile(name)
}

func (FileGenerator) ReadDir(name string) ([]fs.DirEntry, error) {
	return os.ReadDir(name)
}

// ......................................
// ......................................
// ......................................

func (NullFileGenerator) Log(content any, w io.Writer, template string) {
	logger.NewDefaultLogger().Log(content, w, template)
}

func (G NullFileGenerator) WriteFile(name string, data []byte, perm fs.FileMode) error {
	G.Log(fmt.Sprintf(`Creating file %v`, name), os.Stdout, "")
	return nil
}

func (G NullFileGenerator) MkdirAll(path string, perm fs.FileMode) error {
	G.Log(fmt.Sprintf(`Creating directory %v`, path), os.Stdout, "")
	return nil
}

func (G NullFileGenerator) Remove(name string) error {
	G.Log(fmt.Sprintf(`Removing %v`, name), os.Stdout, "")
	return nil
}

func (G NullFileGenerator) RemoveAll(pathname string) error {
	G.Log(fmt.Sprintf(`Removing %v`, pathname), os.Stdout, "")
	return nil
}

func (NullFileGenerator) ReadFile(name string) ([]byte, error) {
	return []byte{}, nil
}

func (NullFileGenerator) ReadDir(name string) ([]fs.DirEntry, error) {
	return []fs.DirEntry{}, nil
}

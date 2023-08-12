package entity

import "strings"

type File struct {
	Name    string // File name. If file name is "*", it indicates that the object is a directory, not a file.
	DirPath string // Directory that file contains
}

func (f *File) FullPath() string {
	return strings.Join([]string{f.DirPath, f.Name}, "/")
}

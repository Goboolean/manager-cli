package entity

type File struct {
	Name string // File name. If file name is "*", it indicates that the object is a directory, not a file.
	Path string // Directory that file contains
}

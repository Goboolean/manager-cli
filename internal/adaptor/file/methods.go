package file

import (
	"errors"

	"github.com/Goboolean/manager-cli/internal/domain/entity"
)

// This method get file list from a directory
func (a *FileAdaptor) GetFileList(dir string) ([]entity.File, error) {
	names, err := a.file.GetFileNameList(dir)
	if err != nil {
		return []entity.File{}, err
	}

	res := make([]entity.File, len(names))

	for _, name := range names {
		res = append(res, entity.File{
			Name:    name,
			DirPath: dir,
		})
	}

	return res, nil
}

// This method remove file from local storage
func (a *FileAdaptor) RemoveFile(target entity.File) error {
	if target.Name == "*" {
		return a.file.RemoveFileOrDir(target.DirPath)
	} else {
		return a.file.RemoveFileOrDir(target.FullPath())
	}
}

// This method calculate hash of file
// YOU MUST USE THIS METHOD WHEN YOU WANT TO CALCULATE HASH OF A FILE
func (a *FileAdaptor) CalculateFileHash(target entity.File) (string, error) {

	if target.Name == "*" {
		return "", errors.New("file: can't calculate hash for directory")
	}

	return a.file.CalculateXxhashChecksum(target.FullPath())
}

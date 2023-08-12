package out

import "github.com/Goboolean/manager-cli/internal/domain/entity"

type FileOperatorPort interface {
	//This method get file list from a directory
	GetFileList(dir string) ([]entity.File, error)
	//This method remove file from local storage
	RemoveFile(target entity.File) error
	//This method calculate hash of file
	CalculateFileHash(target entity.File) (string, error)
}

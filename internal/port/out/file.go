package out

import (
	"context"

	"github.com/Goboolean/manager-cli/internal/domain/entity"
)

type FileOperatorPort interface {
	//This method check directory is exist or not
	IsDirExist(ctx context.Context, dir string) (bool, error)
	//This method get file list from a directory
	GetFileList(ctx context.Context, dir string) ([]entity.File, error)
	//This method create directory
	CreateDirectory(ctx context.Context, dir string) error
	//This method remove file from local storage
	RemoveFile(ctx context.Context, target entity.File) error
	//This method calculate hash of file
	CalculateFileHash(ctx context.Context, target entity.File) (string, error)
}

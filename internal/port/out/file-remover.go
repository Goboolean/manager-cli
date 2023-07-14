package out

import "github.com/Goboolean/manager-cli/internal/domain/entity"

type FileRemover interface {
	RemoveFile(target entity.FileManager) error
}

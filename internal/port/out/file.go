package out

import "github.com/Goboolean/manager-cli/internal/domain/entity"

type FilePort interface {
	RemoveFile(target entity.FileManager) error
}

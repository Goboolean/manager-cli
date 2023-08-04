package out

import "github.com/Goboolean/manager-cli/internal/domain/entity"

// Defines Port for file operation ex: create, remove, copy, move
type FilePort interface {
	RemoveFile(target entity.FileManager) error
}

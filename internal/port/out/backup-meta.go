package out

import "github.com/Goboolean/manager-cli/internal/domain/entity"

type BackupMetaPort interface {
	//This method stores backup metadata to metadata repository which can be mysql, radius so on...
	StoreBackupMeta(meta entity.BackupMeta, target entity.File) error
	//This method gets backup metadata from metadata repository which can be mysql, radius so on...
	ImportBackupMetaFromFile(target entity.File) (entity.BackupMeta, error)
}

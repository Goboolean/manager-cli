package out

import "github.com/Goboolean/manager-cli/internal/domain/entity"

type BackupMetaPort interface {
	//This method stores backup metadata to metadata repository which can be mysql, radius so on...
	StoreBackupMeta(meta entity.BackupMeta, outDir string) error
	//This method gets backup metadata from metadata repository which can be mysql, radius so on...
	ImportBackupMetaFromFile(file entity.File) (entity.BackupMeta, error)
}

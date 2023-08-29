package backupMeta

import (
	"context"
	"encoding/json"

	"github.com/Goboolean/manager-cli/internal/domain/entity"
)

// This method stores backup metadata to metadata repository which can be mysql, radius so on...
func (a *BackupMetaAdaptor) StoreBackupMeta(ctx context.Context, meta entity.BackupMeta, target entity.File) error {

	metaJson, err := json.Marshal(meta)
	if err != nil {
		return err
	}

	return a.file.StoreJsonToFile(metaJson, target.FullPath())
}

// This method gets backup metadata from metadata repository which can be mysql, radius so on...
func (a *BackupMetaAdaptor) ImportBackupMetaFromFile(ctx context.Context, target entity.File) (entity.BackupMeta, error) {
	var res entity.BackupMeta

	err := a.file.DecodeJsonFromFile(target.FullPath(), &res)
	return res, err
}

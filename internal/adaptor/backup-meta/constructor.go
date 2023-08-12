package backupMeta

import fileInf "github.com/Goboolean/manager-cli/infrastructure/file"

type BackupMetaAdaptor struct {
	file *fileInf.FileInfra
}

func New(fileOperator *fileInf.FileInfra) *BackupMetaAdaptor {
	return &BackupMetaAdaptor{
		file: fileOperator,
	}
}

package file

import fileInf "github.com/Goboolean/manager-cli/internal/infrastructure/file"

type FileAdaptor struct {
	file        *fileInf.FileInfra
	hashVersion string
}

// DON'T CHANGE VERSION EXCEPT FOR CHANGING HASHING LOGIC
func New(fileOperator *fileInf.FileInfra) *FileAdaptor {
	return &FileAdaptor{
		file:        fileOperator,
		hashVersion: "v1",
	}
}

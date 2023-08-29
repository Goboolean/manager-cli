package fileInf

type anyPointer interface{}

type FileInfra struct {
}

func New() *FileInfra {
	return &FileInfra{}
}

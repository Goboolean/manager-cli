package transmissionMock

import "github.com/Goboolean/manager-cli/internal/domain/entity"

type TransmissionAdaptorMock struct {
	dirToFilesVirtual map[string][]entity.File
}

func New() *TransmissionAdaptorMock {
	return &TransmissionAdaptorMock{
		dirToFilesVirtual: make(map[string][]entity.File),
	}
}

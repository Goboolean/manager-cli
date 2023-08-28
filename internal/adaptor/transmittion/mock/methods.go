package transmissionMock

import "github.com/Goboolean/manager-cli/internal/domain/entity"

func (a *TransmissionAdaptorMock) CreateRemoteDir(dir string) error {
	a.dirToFilesVirtual[dir] = make([]entity.File, 0)
	return nil
}

func (a *TransmissionAdaptorMock) TransmitDataToRemote(localFile entity.File, remoteDir string) error {
	a.dirToFilesVirtual[remoteDir] = append(a.dirToFilesVirtual[remoteDir], localFile)
	return nil
}

package backup

import "github.com/Goboolean/manager-cli/internal/port/out"

type BackupService struct {
	tradeRepo   out.TradeRepositoryPort
	transmit    out.DataTransmitterPort
	fileRemover out.FilePort
}

// TODO: Find good name for field and parm
func New(tradeRepoPort out.TradeRepositoryPort, transmitter out.DataTransmitterPort, fileRemover out.FilePort) *BackupService {
	return &BackupService{
		tradeRepo:   tradeRepoPort,
		transmit:    transmitter,
		fileRemover: fileRemover,
	}
}

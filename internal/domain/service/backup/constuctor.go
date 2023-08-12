package backup

import "github.com/Goboolean/manager-cli/internal/port/out"

// format date to yyyy-mm-dd_hh:mm:ss
const toolTimeFormatString = "2006-01-02_15:04:05"

const backupTypeFull = "f"
const backupTypeDifferential = "d"
const backupTypeSpecific = "s"

type BackupService struct {
	txCreator    out.TransactionCreator
	tradeDumper  out.TradeDumperPort
	metadataRepo out.MetadataRepositoryPort
	transmitter  out.DataTransmitterPort
	fileRemover  out.FileOperatorPort
	backUpDir    string
}

// TODO: Find good name for field and parm
func New(
	transactionCreator out.TransactionCreator,
	tradeRepoPort out.TradeDumperPort,
	metadataRepoPort out.MetadataRepositoryPort,
	transmitter out.DataTransmitterPort,
	fileRemover out.FileOperatorPort,
	outDir string) *BackupService {

	return &BackupService{
		txCreator:    transactionCreator,
		tradeDumper:  tradeRepoPort,
		metadataRepo: metadataRepoPort,
		transmitter:  transmitter,
		fileRemover:  fileRemover,
		backUpDir:    outDir,
	}
}

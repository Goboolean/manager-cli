package backup

import "github.com/Goboolean/manager-cli/internal/port/out"

const toolTimeFormatString = "2006-01-02_15h04m05s"

const backupTypeFull = "f"
const backupTypeDifferential = "d"
const backupTypeSpecific = "s"

const hashVer = "v1"

type BackupService struct {
	txCreator      out.TransactionCreator
	tradeDumper    out.TradeDumperPort
	metadataRepo   out.MetadataRepositoryPort
	transmitter    out.DataTransmitterPort
	fileOperator   out.FileOperatorPort
	backupMetaPort out.BackupMetaPort
	backUpDir      string
}

// TODO: Find good name for field and parm
func New(
	transactionCreator out.TransactionCreator,
	tradeRepoPort out.TradeDumperPort,
	metadataRepoPort out.MetadataRepositoryPort,
	transmitter out.DataTransmitterPort,
	fileRemover out.FileOperatorPort,
	backupMeta out.BackupMetaPort,
	outDir string) *BackupService {

	return &BackupService{
		txCreator:      transactionCreator,
		tradeDumper:    tradeRepoPort,
		metadataRepo:   metadataRepoPort,
		transmitter:    transmitter,
		fileOperator:   fileRemover,
		backupMetaPort: backupMeta,
		backUpDir:      outDir,
	}
}

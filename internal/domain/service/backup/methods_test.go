package backup_test

import (
	"os"
	"testing"

	fileInf "github.com/Goboolean/manager-cli/infrastructure/file"
	backupMeta "github.com/Goboolean/manager-cli/internal/adaptor/backup-meta"
	"github.com/Goboolean/manager-cli/internal/adaptor/file"
	productMetaRepoMock "github.com/Goboolean/manager-cli/internal/adaptor/product-meta-repo/mock"
	tradeRepoDumpMock "github.com/Goboolean/manager-cli/internal/adaptor/trade-repo-dump/mock"
	transactionCreatorMock "github.com/Goboolean/manager-cli/internal/adaptor/transaction-manager/transaction-creator/mock"
	transmissionMock "github.com/Goboolean/manager-cli/internal/adaptor/transmittion/mock"
	"github.com/Goboolean/manager-cli/internal/domain/service/backup"
	"github.com/stretchr/testify/assert"
)

var instance backup.BackupService
var fobj *file.FileAdaptor
var outDir string

func setUp() {

	fileinf := fileInf.New()

	fobj = file.New(fileinf)
	outDir = `/home/lsjtop10/backup`

	instance = *backup.New(
		transactionCreatorMock.New(),
		tradeRepoDumpMock.New(),
		productMetaRepoMock.New(),
		transmissionMock.New(),
		fobj,
		backupMeta.New(fileinf),
		outDir,
	)
}

func TestMain(m *testing.M) {
	setUp()
	m.Run()
	tearDown()
}

func TestBackupToLocal(t *testing.T) {

	//arrange

	os.RemoveAll(outDir)
	os.MkdirAll(outDir, os.ModePerm)

	var err error
	//act
	err = instance.BackupTradeFull()

	//assert
	assert.NoError(t, err)
	assert.Condition(t, func() (success bool) {
		res, err := fobj.GetFileList(outDir)
		if len(res) > 0 && err == nil {
			return true
		}
		return false
	})
}

func tearDown() {

}

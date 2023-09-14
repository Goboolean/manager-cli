package backup_test

import (
	"context"
	"os"
	"testing"

	backupMeta "github.com/Goboolean/manager-cli/internal/adaptor/backup-meta"
	"github.com/Goboolean/manager-cli/internal/adaptor/file"
	productMetaRepoMock "github.com/Goboolean/manager-cli/internal/adaptor/product-meta-repo/mock"
	tradeRepoDumpMock "github.com/Goboolean/manager-cli/internal/adaptor/trade-repo-dump/mock"
	transactionCreatorMock "github.com/Goboolean/manager-cli/internal/adaptor/transaction-manager/transaction-creator/mock"
	transmissionMock "github.com/Goboolean/manager-cli/internal/adaptor/transmission/mock"
	"github.com/Goboolean/manager-cli/internal/domain/service/backup"
	fileInf "github.com/Goboolean/manager-cli/internal/infrastructure/file"
	"github.com/stretchr/testify/assert"
)

var instance backup.BackupService
var fobj *file.FileAdaptor
var outDir string
var ctx context.Context

func setUp() {

	fileinf := fileInf.New()

	fobj = file.New(fileinf)
	outDir = `/home/lsjtop10/backup`

	ctx = context.Background()
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

	t.Run("BackupTradeFull", func(t *testing.T) {
		//arrange
		os.RemoveAll(outDir)
		os.MkdirAll(outDir, os.ModePerm)

		var err error
		//act
		err = instance.BackupTradeFull(ctx)

		//assert
		assert.NoError(t, err)
		assert.Condition(t, func() (success bool) {
			res, err := fobj.GetFileList(ctx, outDir)
			if len(res) > 0 && err == nil {
				return true
			}
			return false
		})
	})

	t.Run("BackupTradeFull without out dir", func(t *testing.T) {
		//arrange
		os.RemoveAll(outDir)

		var err error
		//act
		err = instance.BackupTradeFull(ctx)

		//assert
		assert.NoError(t, err)
		assert.Condition(t, func() (success bool) {
			res, err := fobj.GetFileList(ctx, outDir)
			if len(res) > 0 && err == nil {
				return true
			}
			return false
		})
	})

}

func tearDown() {

}

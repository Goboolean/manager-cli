package inject_test

import (
	"context"
	"os"
	"strings"
	"testing"

	"github.com/Goboolean/manager-cli/inject"
	"github.com/Goboolean/manager-cli/util/env"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	env.Init()
	m.Run()
}

func TestBackupToLocal(t *testing.T) {

	t.Run("Backup trade full", func(t *testing.T) {
		//arrange
		expectedBaseOutDir := "/home/lsjtop10/backup"
		os.RemoveAll(expectedBaseOutDir)
		os.MkdirAll(expectedBaseOutDir, os.ModePerm)

		ctx := context.TODO()
		instance, err := inject.InitCommandAdaptor()
		if err != nil {
			panic(err)
		}

		//act
		err = instance.BackupTrade(ctx, "full", false)

		//assert
		assert.NoError(t, err)
		assert.DirExists(t, expectedBaseOutDir)

		dirInfo, err := os.ReadDir(expectedBaseOutDir)
		assert.NoError(t, err)
		assert.True(t, len(dirInfo) == 1)
		assert.FileExists(t, strings.Join([]string{expectedBaseOutDir, dirInfo[0].Name(), "metadata.json"}, "/"))
		os.Chdir(strings.Join([]string{expectedBaseOutDir, dirInfo[0].Name(), "goboolean-stock"}, "/"))
	})

	t.Run("Backup trade full", func(t *testing.T) {
		//arrange
		expectedBaseOutDir := "/home/lsjtop10/backup"
		os.RemoveAll(expectedBaseOutDir)
		os.MkdirAll(expectedBaseOutDir, os.ModePerm)

		ctx := context.TODO()
		instance, err := inject.InitCommandAdaptor()
		if err != nil {
			panic(err)
		}

		//act
		err = instance.BackupProduct(ctx, "stock.apple.usa", "full", false)

		//assert
		assert.NoError(t, err)
		assert.DirExists(t, expectedBaseOutDir)

		dirInfo, err := os.ReadDir(expectedBaseOutDir)
		assert.NoError(t, err)
		assert.True(t, len(dirInfo) == 1)
		assert.FileExists(t, strings.Join([]string{expectedBaseOutDir, dirInfo[0].Name(), "metadata.json"}, "/"))
		os.Chdir(strings.Join([]string{expectedBaseOutDir, dirInfo[0].Name(), "goboolean-stock"}, "/"))
		assert.Condition(t, func() bool {
			fInfo, _ := os.ReadDir(strings.Join([]string{expectedBaseOutDir, dirInfo[0].Name()}, "/"))
			return len(fInfo) == 2
		})
	})

}

func TestStatus(t *testing.T) {
	testProductId := "stock.apple.usa"

	t.Run("Update status", func(t *testing.T) {
		ctx := context.TODO()
		// arrange
		instance, err := inject.InitCommandAdaptor()
		if err != nil {
			panic(err)
		}

		/*err = instance.UpdateStatus(ctx, "stock.apple.usa", "-rst")
		if err != nil {
			panic(err)
		}*/

		// act
		err = instance.UpdateStatus(ctx, testProductId, "=rst")
		// assert
		assert.NoError(t, err)
		res, _ := instance.GetStatus(ctx, testProductId)
		assert.Equal(t, res, "rst")
	})

	t.Run("Update status with invalid", func(t *testing.T) {
		ctx := context.TODO()
		// arrange
		instance, err := inject.InitCommandAdaptor()
		if err != nil {
			panic(err)
		}

		/*err = instance.UpdateStatus(ctx, "stock.apple.usa", "-rst")
		if err != nil {
			panic(err)
		}*/

		// act
		err = instance.UpdateStatus(ctx, "=st", testProductId)
		// assert
		assert.Error(t, err)
	})

	t.Run("Remove status", func(t *testing.T) {
		ctx := context.TODO()
		// arrange
		instance, err := inject.InitCommandAdaptor()
		if err != nil {
			panic(err)
		}

		instance.UpdateStatus(ctx, testProductId, "=rst")
		/*err = instance.UpdateStatus(ctx, "stock.apple.usa", "-rst")
		if err != nil {
			panic(err)
		}*/

		// act
		err = instance.UpdateStatus(ctx, testProductId, "-st")
		// assert
		assert.NoError(t, err)
		res, _ := instance.GetStatus(ctx, testProductId)
		assert.Equal(t, "r--", res)
	})

	t.Run("Remove status with invaild", func(t *testing.T) {
		ctx := context.TODO()
		// arrange
		instance, err := inject.InitCommandAdaptor()
		if err != nil {
			panic(err)
		}

		instance.UpdateStatus(ctx, testProductId, "=rst")
		/*err = instance.UpdateStatus(ctx, "stock.apple.usa", "-rst")
		if err != nil {
			panic(err)
		}*/

		// act
		err = instance.UpdateStatus(ctx, testProductId, "-r")
		// assert
		assert.Error(t, err)
	})

	t.Run("Update status with num: 0", func(t *testing.T) {
		ctx := context.TODO()
		// arrange
		instance, err := inject.InitCommandAdaptor()
		if err != nil {
			panic(err)
		}

		/*err = instance.UpdateStatus(ctx, "stock.apple.usa", "-rst")
		if err != nil {
			panic(err)
		}*/

		// act
		err = instance.UpdateStatus(ctx, testProductId, "0")
		// assert
		assert.NoError(t, err)
		res, _ := instance.GetStatus(ctx, testProductId)
		assert.Equal(t, res, "")
	})

	t.Run("Update status with num: 5", func(t *testing.T) {
		ctx := context.TODO()
		// arrange
		instance, err := inject.InitCommandAdaptor()
		if err != nil {
			panic(err)
		}

		/*err = instance.UpdateStatus(ctx, "stock.apple.usa", "-rst")
		if err != nil {
			panic(err)
		}*/

		// act
		err = instance.UpdateStatus(ctx, testProductId, "5")
		// assert
		assert.NoError(t, err)
		res, _ := instance.GetStatus(ctx, testProductId)
		assert.Equal(t, res, "r-t")
	})

	t.Run("Update status with num expected error", func(t *testing.T) {
		ctx := context.TODO()
		// arrange
		instance, err := inject.InitCommandAdaptor()
		if err != nil {
			panic(err)
		}

		/*err = instance.UpdateStatus(ctx, "stock.apple.usa", "-rst")
		if err != nil {
			panic(err)
		}*/

		// act
		err = instance.UpdateStatus(ctx, testProductId, "3")
		// assert
		assert.Error(t, err)
	})

}

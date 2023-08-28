package tradeRepoDumpMock

import (
	"os"
	"path"
	"strings"
	"time"

	"github.com/Goboolean/manager-cli/internal/domain/entity"
)

func createDummyFile(fList []entity.File) error {
	for i := range fList {
		f, err := os.Create(fList[i].FullPath())
		defer f.Close()

		if err != nil {
			return err
		}

		f.WriteString("hello world")
	}
	return nil
}

// This method dumps trade data of specific product created before time
func (a *TradeDumpAdaptorMock) DumpProductBefore(id string, outDir string, date time.Time) ([]entity.File, error) {

	var fmgr []entity.File
	if err := os.MkdirAll(strings.Join([]string{outDir, a.database}, "/"), 0777); err != nil {
		return nil, err
	}

	fmgr = append(fmgr, entity.File{
		Name:    strings.Join([]string{id, ".bson.gz"}, ""),
		DirPath: strings.Join([]string{outDir, a.database}, "/"),
	})

	fmgr = append(fmgr, entity.File{
		Name:    strings.Join([]string{id, ".metadata.json"}, ""),
		DirPath: strings.Join([]string{outDir, a.database}, "/"),
	})

	if err := createDummyFile(fmgr); err != nil {
		return nil, err
	}

	return fmgr, nil
}

// This method dumps trade data of specific product created between time\
func (a *TradeDumpAdaptorMock) DumpProductBetween(id string, outDir string, from, to time.Time) ([]entity.File, error) {
	var fmgr []entity.File
	os.MkdirAll(path.Base(outDir), 0777)

	fmgr = append(fmgr, entity.File{
		Name:    strings.Join([]string{id, ".bson.gz"}, ""),
		DirPath: strings.Join([]string{outDir, a.database}, "/"),
	})

	fmgr = append(fmgr, entity.File{
		Name:    strings.Join([]string{id, ".metadata.json"}, ""),
		DirPath: strings.Join([]string{outDir, a.database}, "/"),
	})

	if err := createDummyFile(fmgr); err != nil {
		return nil, err
	}

	return fmgr, nil
}

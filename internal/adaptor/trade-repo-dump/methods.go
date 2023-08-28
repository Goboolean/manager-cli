package tradeRepoDump

import (
	"fmt"
	"strings"
	"time"

	"github.com/Goboolean/manager-cli/internal/domain/entity"
	mongoInf "github.com/Goboolean/manager-cli/internal/infrastructure/mongo"
)

// This method dumps trade data of specific product created before time
func (a *TradeDumpAdaptor) DumpProductBefore(id string, outDir string, date time.Time) ([]entity.File, error) {

	q := fmt.Sprintf(`'{startTime:{"$lte":%d}}'`, date.Unix())

	err := mongoInf.ExecMongodump(
		[]string{
			"--gzip",
			strings.Join([]string{"--host", a.Host}, "="),
			strings.Join([]string{"--port", a.Port}, "="),
			strings.Join([]string{"--username", a.User}, "="),
			strings.Join([]string{"--password", a.PassWord}, "="),
			strings.Join([]string{"--authenticationDatabase", a.Database}, "="),
			strings.Join([]string{"--db", a.Database}, "="),
			strings.Join([]string{"--out", outDir}, "="),
			strings.Join([]string{"--collection", id}, "="),
			strings.Join([]string{"--query=", q}, "="),
		})

	if err != nil {
		return nil, err
	}

	fmgr := make([]entity.File, 2)
	fmgr = append(fmgr, entity.File{
		Name:    strings.Join([]string{id, ".bson.gz"}, ""),
		DirPath: strings.Join([]string{outDir, a.Database}, "/"),
	})
	fmgr = append(fmgr, entity.File{
		Name:    strings.Join([]string{id, ".metadata.json"}, ""),
		DirPath: strings.Join([]string{outDir, a.Database}, "/"),
	})

	return fmgr, nil
}

// This method dumps trade data of specific product created between time\
func (a *TradeDumpAdaptor) DumpProductBetween(id string, outDir string, from, to time.Time) ([]entity.File, error) {

	q := fmt.Sprintf(`'{"startTime":{"$gt":%d,"lte":%d}}'`, from.Unix(), to.Unix())

	err := mongoInf.ExecMongodump(
		[]string{
			"--gzip",
			strings.Join([]string{"--host", a.Host}, "="),
			strings.Join([]string{"--port", a.Port}, "="),
			strings.Join([]string{"--username", a.User}, "="),
			strings.Join([]string{"--password", a.PassWord}, "="),
			strings.Join([]string{"--authenticationDatabase", a.Database}, "="),
			strings.Join([]string{"--db", a.Database}, "="),
			strings.Join([]string{"--out", outDir}, "="),
			strings.Join([]string{"--collection", id}, "="),
			strings.Join([]string{"--query=", q}, "="),
		})

	if err != nil {
		return nil, err
	}

	fmgr := make([]entity.File, 2)
	fmgr = append(fmgr, entity.File{
		Name:    strings.Join([]string{id, ".bson.gz"}, ""),
		DirPath: strings.Join([]string{outDir, a.Database}, "/"),
	})

	fmgr = append(fmgr, entity.File{
		Name:    strings.Join([]string{id, ".metadata.json"}, ""),
		DirPath: strings.Join([]string{outDir, a.Database}, "/"),
	})

	return fmgr, nil
}

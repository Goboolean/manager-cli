package tradeRepoDump

import (
	"fmt"
	"strings"
	"time"

	mongoInf "github.com/Goboolean/manager-cli/infrastructure/mongo"
	"github.com/Goboolean/manager-cli/internal/domain/entity"
)

// This method dumps trade data of specific product
func (a *TradeDumpAdaptor) DumpProduct(id string) (entity.File, error) {

	targetDir := strings.Join([]string{a.baseOutDir, time.Now().Format("2006/01/02")}, "/")

	err := mongoInf.ExecMongodump(
		[]string{
			strings.Join([]string{"--host", a.Host}, "="),
			strings.Join([]string{"--port", a.Port}, "="),
			strings.Join([]string{"--username", a.User}, "="),
			strings.Join([]string{"--password", a.PassWord}, "="),
			strings.Join([]string{"--authenticationDatabase", a.Database}, "="),
			strings.Join([]string{"--db", a.Database}, "="),
			strings.Join([]string{"--out", targetDir}, "="),
			strings.Join([]string{"--collection", id}, "="),
		})

	if err != nil {
		return entity.File{}, err
	}

	return entity.File{Name: "*", Path: strings.Join([]string{targetDir, a.Database}, "/")}, nil
}

// This method dumps trade data of specific product created before time
func (a *TradeDumpAdaptor) DumpProductBefore(id string, date time.Time) (entity.File, error) {
	q := fmt.Sprintf(`'{startTime:{"$lt":%d}}'`, date.Unix())

	targetDir := strings.Join([]string{a.baseOutDir, time.Now().Format("2006/01/02")}, "/")

	err := mongoInf.ExecMongodump(
		[]string{
			strings.Join([]string{"--host", a.Host}, "="),
			strings.Join([]string{"--port", a.Port}, "="),
			strings.Join([]string{"--username", a.User}, "="),
			strings.Join([]string{"--password", a.PassWord}, "="),
			strings.Join([]string{"--authenticationDatabase", a.Database}, "="),
			strings.Join([]string{"--db", a.Database}, "="),
			strings.Join([]string{"--out", targetDir}, "="),
			strings.Join([]string{"--collection", id}, "="),
			strings.Join([]string{"--query=", q}, "="),
		})

	if err != nil {
		return entity.File{}, err
	}

	return entity.File{Name: "*", Path: strings.Join([]string{targetDir, a.Database}, "/")}, nil
}

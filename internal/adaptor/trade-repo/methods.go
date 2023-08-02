package tradeRepo

import (
	"fmt"
	"strings"
	"time"

	mongoInf "github.com/Goboolean/manager-cli/infrastructure/mongo"
	"github.com/Goboolean/manager-cli/internal/domain/entity"
)

// This method dumps trade data of specific product
func (a *TradeRepoAdaptor) DumpProduct(id string) (entity.File, error) {

	targetDir := strings.Join([]string{a.baseOutDir, time.Now().Format("2006/01/02")}, "/")

	err := mongoInf.ExecMongodump(
		[]string{
			strings.Join([]string{"--host", a.config.Host}, "="),
			strings.Join([]string{"--port", a.config.Port}, "="),
			strings.Join([]string{"--username", a.config.User}, "="),
			strings.Join([]string{"--password", a.config.PassWord}, "="),
			strings.Join([]string{"--authenticationDatabase", a.config.Database}, "="),
			strings.Join([]string{"--db", a.config.Database}, "="),
			strings.Join([]string{"--out", targetDir}, "="),
			strings.Join([]string{"--collection", id}, "="),
		})

	if err != nil {
		return entity.File{}, err
	}

	return entity.File{Name: "*", Path: strings.Join([]string{targetDir, a.config.Database}, "/")}, nil
}

// This method dumps trade data of specific product created before time
func (a *TradeRepoAdaptor) DumpProductBefore(id string, date time.Time) (entity.File, error) {
	q := fmt.Sprintf(`'{startTime:{"$lt":%d}}'`, date.Unix())

	targetDir := strings.Join([]string{a.baseOutDir, time.Now().Format("2006/01/02")}, "/")

	err := mongoInf.ExecMongodump(
		[]string{
			strings.Join([]string{"--host", a.config.Host}, "="),
			strings.Join([]string{"--port", a.config.Port}, "="),
			strings.Join([]string{"--username", a.config.User}, "="),
			strings.Join([]string{"--password", a.config.PassWord}, "="),
			strings.Join([]string{"--authenticationDatabase", a.config.Database}, "="),
			strings.Join([]string{"--db", a.config.Database}, "="),
			strings.Join([]string{"--out", targetDir}, "="),
			strings.Join([]string{"--collection", id}, "="),
			strings.Join([]string{"--query=", q}, "="),
		})

	if err != nil {
		return entity.File{}, err
	}

	return entity.File{Name: "*", Path: strings.Join([]string{targetDir, a.config.Database}, "/")}, nil
}

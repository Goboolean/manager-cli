package tradeRepoDump

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/Goboolean/manager-cli/internal/domain/entity"
	mongoInf "github.com/Goboolean/manager-cli/internal/infrastructure/mongo"
)

func (a *TradeDumpAdaptor) dumpWithQuery(ctx context.Context, id string, outDir string, query string) ([]entity.File, error) {

	err := mongoInf.ExecMongodump(
		[]string{
			strings.Join([]string{"--uri", a.connUri}, "="),
			strings.Join([]string{"--db", a.database}, "="),
			//"--quiet",
			strings.Join([]string{"--collection", id}, "="),
			strings.Join([]string{"--out", outDir}, "="),
			strings.Join([]string{"--query", query}, "="),
			"--gzip",
		})

	if err != nil {
		return nil, err
	}

	fmgr := make([]entity.File, 2)
	fmgr[0] = entity.File{
		Name:    strings.Join([]string{id, ".bson.gz"}, ""),
		DirPath: strings.Join([]string{outDir, a.database}, "/"),
	}
	fmgr[1] = entity.File{
		Name:    strings.Join([]string{id, ".metadata.json.gz"}, ""),
		DirPath: strings.Join([]string{outDir, a.database}, "/"),
	}

	return fmgr, nil
}

// This method dumps trade data of specific product created before time
func (a *TradeDumpAdaptor) DumpProductBefore(ctx context.Context, id string, outDir string, date time.Time) ([]entity.File, error) {

	q := fmt.Sprintf(`{"startTime":{"$lte":%d}}`, date.Unix())

	return a.dumpWithQuery(ctx, id, outDir, q)
}

// This method dumps trade data of specific product created between time\
func (a *TradeDumpAdaptor) DumpProductBetween(ctx context.Context, id string, outDir string, from, to time.Time) ([]entity.File, error) {
	q := fmt.Sprintf(`{"startTime":{"$gt":%d,"lte":%d}}`, from.Unix(), to.Unix())

	return a.dumpWithQuery(ctx, id, outDir, q)
}

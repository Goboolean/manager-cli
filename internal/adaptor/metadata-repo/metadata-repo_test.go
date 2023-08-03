package metadataRepo_test

import (
	"context"
	"log"
	"os"
	"testing"

	metadataRepo "github.com/Goboolean/manager-cli/internal/adaptor/metadata-repo"
	"github.com/Goboolean/manager-cli/internal/domain/entity"
	"github.com/Goboolean/shared/pkg/rdbms"
	"github.com/Goboolean/shared/pkg/resolver"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var metaRepoAdaptor *metadataRepo.MetadataRepositoryAdaptor

func TestMain(m *testing.M) {

	os.Chdir("/home/lsjtop10/projects/goboolean/manager-cli")
	godotenv.Load()

	metaRepoAdaptor = metadataRepo.New(rdbms.NewDB(
		&resolver.ConfigMap{
			"USER":     os.Getenv("PSQL_USER"),
			"PASSWORD": os.Getenv("PSQL_PASS"),
			"HOST":     os.Getenv("PSQL_HOST"),
			"PORT":     os.Getenv("PSQL_PORT"),
			"DATABASE": os.Getenv("PSQL_DATABASE"),
		}))

	code := m.Run()
	//metaRepoAdaptor.Close()

	os.Exit(code)
}

// DO NOT EXECUTE THIS TEST ON THE PRODUCTION DB BECAUSE IT WILL CAUSE DATA POLLUTION
func TestInsertMetaRollback(t *testing.T) {

	var err error

	session, err := metaRepoAdaptor.CreateTxSession(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	err = metaRepoAdaptor.StoreProductMeta(
		session,
		entity.ProductMeta{
			Id:          "stock.apple.usa",
			Name:        "apple",
			Code:        "AAPL",
			Exchange:    "nasdaq",
			Description: "",
			Type:        "stock",
			Location:    "usa",
		})

	if err != nil {
		log.Fatal(err)
	}

	metaRepoAdaptor.Rollback(session)

	session, err = metaRepoAdaptor.CreateTxSession(context.Background())

	if err != nil {
		panic(err)
	}

	if str, _ := metaRepoAdaptor.GetProductId(session, "AAPL"); str == "stock.apple.usa" {
		t.Error("Fail!!!")
	}

	metaRepoAdaptor.Commit(session)
}

// WARN: DO NOT EXECUTE THIS TEST ON THE PRODUCTION DB BECAUSE IT WILL CAUSE DATA POLLUTION
// This test inserts mock data and checks if the data is inserted correctly to search Id by symbol
func TestInsertMetaCommitted(t *testing.T) {

	var err error

	session, err := metaRepoAdaptor.CreateTxSession(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	err = metaRepoAdaptor.StoreProductMeta(
		session,
		entity.ProductMeta{
			Id:          "stock.apple.usa",
			Name:        "apple",
			Code:        "AAPL",
			Exchange:    "nasdaq",
			Description: "",
			Type:        "stock",
			Location:    "usa",
		})

	if err != nil {
		t.Error(err)
	}

	metaRepoAdaptor.Commit(session)

	session, err = metaRepoAdaptor.CreateTxSession(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	var str string
	str, err = metaRepoAdaptor.GetProductId(session, "AAPL")
	if err != nil {
		t.Error(err)
	}

	if str != "stock.apple.usa" {
		t.Error("Fail to insert")
	}
	metaRepoAdaptor.Commit(session)
}

func TestGetProductMeta(t *testing.T) {

	session, err := metaRepoAdaptor.CreateTxSession(context.Background())
	defer metaRepoAdaptor.Commit(session)

	meta, err := metaRepoAdaptor.GetProductMeta(session, "stock.apple.usa")

	if err != nil {
		t.Error(err)
	}

	if meta.Id != "stock.apple.usa" ||
		meta.Code != "AAPL" ||
		meta.Name != "apple" ||
		meta.Exchange != "nasdaq" ||
		meta.Description != "" ||
		meta.Type != "stock" ||
		meta.Location != "usa" {
		t.Error("Fail to get product meta")
	}

}

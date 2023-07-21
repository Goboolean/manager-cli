package metadataRepo_test

import (
	"context"
	"log"
	"os"
	"testing"

	metadataRepo "github.com/Goboolean/manager-cli/internal/adaptor/metadata-repo"
	"github.com/Goboolean/manager-cli/internal/domain/entity"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var metaRepoAdaptor *metadataRepo.MetadataRepositoryAdaptor

func TestMain(m *testing.M) {

	os.Chdir("/home/lsjtop10/projects/goboolean/manager-cli")
	godotenv.Load()

	metaRepoAdaptor = metadataRepo.New()
	code := m.Run()
	//metaRepoAdaptor.Close()

	os.Exit(code)
}

// DO NOT EXECUTE THIS TEST ON THE PRODUCTION DB BECAUSE IT WILL CAUSE DATA POLLUTION
func TestInsertMetaRollback(t *testing.T) {

	var err error

	err = metaRepoAdaptor.Begin(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	err = metaRepoAdaptor.StoreProductMeta(
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

	metaRepoAdaptor.Rollback()

	if str, _ := metaRepoAdaptor.GetProductId("AAPL"); str == "stock.apple.usa" {
		t.Error("Fail!!!")
	}
}

// WARN: DO NOT EXECUTE THIS TEST ON THE PRODUCTION DB BECAUSE IT WILL CAUSE DATA POLLUTION
// This test inserts mock data and checks if the data is inserted correctly to search Id by symbol
func TestInsertMetaCommitted(t *testing.T) {

	var err error

	err = metaRepoAdaptor.Begin(context.Background())
	if err != nil {
		t.Error(err)
	}

	err = metaRepoAdaptor.StoreProductMeta(
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

	metaRepoAdaptor.Commit()

	metaRepoAdaptor.Begin(context.Background())

	var str string
	str, err = metaRepoAdaptor.GetProductId("AAPL")
	if err != nil {
		t.Error(err)
	}

	if str != "stock.apple.usa" {
		t.Error("Fail to insert")
	}
	metaRepoAdaptor.Commit()
}

func TestGetProductMeta(t *testing.T) {

	metaRepoAdaptor.Begin(context.Background())
	defer metaRepoAdaptor.Commit()

	meta, err := metaRepoAdaptor.GetProductMeta("stock.apple.usa")

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

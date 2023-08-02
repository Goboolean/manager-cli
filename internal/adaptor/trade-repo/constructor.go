package tradeRepo

import (
	"os"

	"github.com/Goboolean/shared/pkg/mongo"
	"github.com/Goboolean/shared/pkg/resolver"
)

type TradeRepoAdaptor struct {
	db           *mongo.DB
	query        *mongo.Queries
	transactions map[int]resolver.Transactioner
	config       dbConfig
	baseOutDir   string
}

type dbConfig struct {
	Host     string
	Port     string
	User     string
	PassWord string
	Database string
}

func New() *TradeRepoAdaptor {
	dbcfg := dbConfig{
		Host:     os.Getenv("MONGO_HOST"),
		User:     os.Getenv("MONGO_USER"),
		Port:     os.Getenv("MONGO_PORT"),
		PassWord: os.Getenv("MONGO_PASS"),
		Database: os.Getenv("MONGO_DATABASE"),
	}

	dbInstance := mongo.NewDB(&resolver.ConfigMap{
		"USER":     dbcfg.User,
		"PASSWORD": dbcfg.PassWord,
		"HOST":     dbcfg.Host,
		"PORT":     dbcfg.Port,
		"DATABASE": dbcfg.Database,
	})

	return &TradeRepoAdaptor{
		db:           dbInstance,
		query:        mongo.New(dbInstance),
		config:       dbcfg,
		transactions: make(map[int]resolver.Transactioner),
		baseOutDir:   os.Getenv("MONGODUMP_OUTDIR"),
	}
}

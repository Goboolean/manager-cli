package tradeRepoDump

import "github.com/Goboolean/shared/pkg/resolver"

type TradeDumpAdaptor struct {
	User     string
	PassWord string
	Host     string
	Port     string
	Database string

	baseOutDir string
}

func New(c *resolver.ConfigMap, outDir string) *TradeDumpAdaptor {

	user, err := c.GetStringKey("USER")
	if err != nil {
		panic(err)
	}

	password, err := c.GetStringKey("PASSWORD")
	if err != nil {
		panic(err)
	}

	host, err := c.GetStringKey("HOST")
	if err != nil {
		panic(err)
	}

	port, err := c.GetStringKey("PORT")
	if err != nil {
		panic(err)
	}

	database, err := c.GetStringKey("DATABASE")
	if err != nil {
		panic(err)
	}

	return &TradeDumpAdaptor{
		User:     user,
		PassWord: password,
		Host:     host,
		Port:     port,
		Database: database,

		baseOutDir: outDir,
	}
}

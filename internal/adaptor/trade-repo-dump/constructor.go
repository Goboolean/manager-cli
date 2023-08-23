package tradeRepoDump

import "github.com/Goboolean/shared/pkg/resolver"

type TradeDumpAdaptor struct {
	User     string
	PassWord string
	Host     string
	Port     string
	Database string
}

func New(c *resolver.ConfigMap, outDir string) (*TradeDumpAdaptor, error) {

	user, err := c.GetStringKey("USER")
	if err != nil {
		return nil, err
	}

	password, err := c.GetStringKey("PASSWORD")
	if err != nil {
		return nil, err
	}

	host, err := c.GetStringKey("HOST")
	if err != nil {
		return nil, err
	}

	port, err := c.GetStringKey("PORT")
	if err != nil {
		return nil, err
	}

	database, err := c.GetStringKey("DATABASE")
	if err != nil {
		return nil, err
	}

	return &TradeDumpAdaptor{
		User:     user,
		PassWord: password,
		Host:     host,
		Port:     port,
		Database: database,
	}, nil
}

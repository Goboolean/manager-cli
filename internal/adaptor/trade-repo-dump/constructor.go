package tradeRepoDump

import (
	"fmt"

	"github.com/Goboolean/shared/pkg/resolver"
)

type TradeDumpAdaptor struct {
	connUri  string
	database string
}

func New(c *resolver.ConfigMap) (*TradeDumpAdaptor, error) {

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

	uri := fmt.Sprintf(
		"mongodb://%s:%s@%s:%s/?authSource=%s&directConnection=true",
		user, password, host, port, database)

	return &TradeDumpAdaptor{
		connUri:  uri,
		database: database,
	}, nil
}

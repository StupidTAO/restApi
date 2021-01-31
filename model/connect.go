package model

import (
	"crypto/tls"
	"github.com/arangodb/go-driver"
	dhttp "github.com/arangodb/go-driver/http"
	"github.com/test/restApi/conf"
)

func getDBClient() (driver.Client, error) {
	cfg, err := conf.LoadConfig()
	if err != nil {
		return nil, err
	}
	dbCfg, _ := conf.GetDBCfg(cfg)

	conn, err := dhttp.NewConnection(dhttp.ConnectionConfig{
		//Endpoints: []string{"http://localhost:8529"},
		Endpoints: []string{dbCfg.Url},
		TLSConfig: &tls.Config{},
	})
	if err != nil {
		return nil, err
	}

	c, err := driver.NewClient(driver.ClientConfig{
		Connection: conn,
		Authentication: driver.BasicAuthentication(dbCfg.User, dbCfg.Password),
	})
	if err != nil {
		return nil, err
	}
	return c, nil
}

package main

import (
	"github.com/bigstth/isekai-shop-api/config"
	"github.com/bigstth/isekai-shop-api/databases"
	"github.com/bigstth/isekai-shop-api/server"
)

func main() {
	conf := config.ConfigGetting()
	db := databases.NewPostgresDatabase(conf.Database)
	server := server.NewEchoServer(conf, db.Connect())

	server.Start()
}

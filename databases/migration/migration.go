package main

import (
	"fmt"

	"github.com/bigstth/isekai-shop-api/config"
	"github.com/bigstth/isekai-shop-api/databases"
)

func main() {
	conf := config.ConfigGetting()
	db := databases.NewPostgresDatabase(conf.Database)

	fmt.Println(db.ConnectionGetting())
}

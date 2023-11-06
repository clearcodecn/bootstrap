package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/xormplus/xorm"
	"tools/core/config"
)

var (
	db *xorm.Engine
)

func InitDatabase() {
	cfg := config.Get()
	engine, err := xorm.NewEngine(cfg.Driver, cfg.Dsn)
	if err != nil {
		panic(err)
	}
	if err := engine.Ping(); err != nil {
		panic(err)
	}
	if err := engine.Sync2(new(ExampleTable)); err != nil {
		panic(err)
	}
}

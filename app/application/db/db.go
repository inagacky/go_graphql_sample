package db

import (
	"github.com/inagacky/go_graphql_sample/app/util"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func Init() (*gorm.DB, error) {
	var err error

	// 環境変数から取得(お試しなのでdefaultvalue直書)
	dbUser := util.Getenv("GO_GRAPHQL_SAMPLE_DB_USER", "root")
	dbPass := util.Getenv("GO_GRAPHQL_SAMPLE_DB_PASS", "sample")
	dbName := util.Getenv("GO_GRAPHQL_SAMPLE_DB_NAME", "sample")
	dbHostName := util.Getenv("GO_GRAPHQL_SAMPLE_DB_HOSTNAME", "127.0.0.1")
	dbPort := util.Getenv("GO_GRAPHQL_SAMPLE_DB_PORT", "3306")
	protocol := "tcp(" + dbHostName + ":" + dbPort + ")"

	// データソースの定義
	dataSource := dbUser + ":" + dbPass + "@" + protocol + "/" + dbName + "?parseTime=true&charset=utf8"
	DB, err = gorm.Open("mysql", dataSource)
	if err != nil {
		return nil, err
	}

	return GetDB(), nil
}

// DBを返却
func GetDB() *gorm.DB {
	return DB
}

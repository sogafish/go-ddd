package models

import (
  "fmt"
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/mysql"
  "log"
  "go-go/config"
)

var Db *gorm.DB

func init() {
  var err error
  dbConnectInfo := fmt.Sprintf(
    `%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local`,
    config.Config.DbUserName,
    config.Config.DbUserPassword,
    config.Config.DbHost,
    config.Config.DbPort,
    config.Config.DbName,
  )

  // configから読み込んだ情報を元に、データベースに接続します
  Db, err = gorm.Open(config.Config.DbDriverName, dbConnectInfo)
  if err != nil {
    log.Fatalln(err)
  } else {
    fmt.Println("Successfully connect database..")
  }

  // 接続したデータベースにitemsテーブルを作成します
  Db.Set("gorm:table_options", "ENGINE = InnoDB").AutoMigrate(&Item{})
  if err != nil {
    fmt.Println(err.Error())
  } else {
    fmt.Println("Successfully created table..")
  }
}

func GetAllItems(items *[]Item) {
  Db.Find(&items)
}

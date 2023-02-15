package Databases

import (
	"dousheng-backend/Models"
	"encoding/json"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"io/ioutil"
	"log"
	"os"
)

var DB *gorm.DB

func InitDatabase() {
	//想到开闭原则，写在里面以免向外界暴露
	type DBconfig struct {
		Account  string `json:"account"`
		Password string `json:"password"`
		Host     string `json:"host"`
		Port     int    `json:"port"`
		Database string `json:"database"`
		Timeout  string `json:"timeout"`
	}

	conf, _ := os.Open("./confs/database.json")
	defer conf.Close() //执行完毕后关闭连接
	value, _ := ioutil.ReadAll(conf)
	var conn DBconfig
	json.Unmarshal([]byte(value), &conn)

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=true&loc=Local",
		conn.Account, conn.Password, conn.Host, conn.Port, conn.Database)
	//连接MYSQL, 获得DB类型实例，用于后面的数据库读写操作
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Errorf("failed creating database:%w", err))
	}
	log.Println("init database success")
	db.AutoMigrate(&Models.User{})
	DB = db
}

func DatabaseSession() *gorm.DB {
	return DB.Session(&gorm.Session{PrepareStmt: true})
}

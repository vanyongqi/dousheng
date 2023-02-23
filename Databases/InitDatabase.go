package Databases

import (
	"dousheng-backend/Models"
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"io/ioutil"
	"os"
)

var mysqldb *gorm.DB

type DBconfig struct {
	Account  string `json:"account"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Database string `json:"database"`
	Timeout  string `json:"timeout"`
}

func InitDatabase() {

	conf, _ := os.Open("Configs/mysql.json")
	defer conf.Close() //执行完毕后关闭连接
	value, _ := ioutil.ReadAll(conf)
	var conn DBconfig
	json.Unmarshal([]byte(value), &conn)
	//conn.Account = "root"
	//conn.Password = "123456"
	//conn.Host = "127.0.0.1"
	//conn.Port = 3306
	//conn.Database = "dousheng"
	//fmt.Println("用户:", conn.Account, "密码:", conn.Password, "主机地址 :", conn.Host, "端口:", conn.Port, "数据库名称", conn.Database)

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=true&loc=Local",
		conn.Account, conn.Password, conn.Host, conn.Port, conn.Database)
	//连接MYSQL, 获得DB类型实例，用于后面的数据库读写操作
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("failed creating database:%w", err)
		logrus.Error("failed creating database:%w", err)
		return
	}
	db.AutoMigrate(&Models.User{}, &Models.Video{}, &Models.Comment{})
	mysqldb = db
}

func DatabaseSession() *gorm.DB {
	return mysqldb.Session(&gorm.Session{PrepareStmt: true})
}

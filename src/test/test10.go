
package main

import (
	"github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/Sirupsen/logrus"
)

func check(db_type string, connect string, sql string) (error){
	if db_type == "postgres"
		db,err := gorm.Open("postgres", "host=47.110.154.127 port=30172 user=postgres dbname=hairtail sslmode=disable password=huansi@2017")
		if err != nil {
			logrus.Errorln(err)
			logrus.Fatalln("database ping attempts failed")
		}
		logrus.Infof("连接成功了")
		if err != nil{
			return err
		}

		_,err = db.Exec(sql)
		if err != nil{
			return err
		}
	} else if db_type == "mysql"{
		db,err := gorm.Open("mysql", "host=116.62.213.56 port=51603 user=root dbname=dougo sslmode=disable password=huansi@2017")
		if err != nil {
			logrus.Errorln(err)
			logrus.Fatalln("database ping attempts failed")
		}
		logrus.Infof("连接成功了")
		if err != nil{
			return err
		}

		_,err = db.Exec(sql)
		if err != nil{
			return err
		}
	} else if db_type == "sqlite"{
		db,err := gorm.Open("sqlite3", "test.db")
		if err != nil {
			logrus.Errorln(err)
			logrus.Fatalln("database ping attempts failed")
		}
		logrus.Infof("连接成功了")
		if err != nil{
			return err
		}

		_,err = db.Exec(sql)
		if err != nil{
			return err
		}
	}
}

func main() {
	check("postgres", "123", "select * from pipeline")

}

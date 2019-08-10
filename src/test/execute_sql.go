package main

import (
	"github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/sirupsen/logrus"
)

func check(connect string, sql string) (error){
	db,err := gorm.Open("postgres", connect)
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

func main() {

}

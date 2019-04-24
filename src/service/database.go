package service

import (
	"log"
	"fmt"
	"time"

        _ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/jinzhu/gorm"
	"github.com/Sirupsen/logrus"
)

func DB(db_type, host, port, user, pwd, name, table, column string, rows_count, timeout int) error {
	if timeout == 0 {
		timeout = 30
	}
	connec_str := ""
	if db_type == "postgres" {
		connec_str = "host="+host+" port="+port+" user="+user+" dbname="+name+" password="+pwd+" sslmode=disable"
	}
	if db_type == "mysql" {
		connec_str = user+":"+pwd+"@tcp("+host+":"+port+")/"+name
	}
	db,err := connec_db(db_type, connec_str)
	if err != nil {
		logrus.Errorln(err)
		logrus.Fatalln("database ping attempts failed")
		return nil
	}

	ticker:=time.NewTicker(time.Second * 5)

	go func() {
	    for _=range ticker.C {
		//rows,err := db.Query("select * from "+table+" LIMIT 10")
		rows,_ := db.Raw("select * from "+table+" LIMIT 10").Rows()
		//if err != nil{
		//	return err
		//}
		cols, _ := rows.Columns()
		defer rows.Close()
		for rows.Next() {
			//m := make(map[string]interface{})
			//if err := rows.Scan(&m); err != nil {
			//	log.Printf("err :", err)
			//	return
			//	// ERROR: sql: expected X destination arguments in Scan, not 1
			//	//return err
			//}
			//log.Printf("testm:", m)

			//log.Printf("Data :", string(data))

			columns := make([]interface{}, len(cols))
			columnPointers := make([]interface{}, len(cols))
			for i, _ := range columns {
			    columnPointers[i] = &columns[i]
			}

			// Scan the result into the column pointers...
			if err := rows.Scan(columnPointers...); err != nil {
				log.Printf("err :", err)
				return
			    //return err
			}

			// Create our map, and retrieve the value for each column from the pointers slice,
			// storing it in the map with the name of the column as the key.
			m := make(map[string]interface{})
			for i, colName := range cols {
			    val := columnPointers[i].(*interface{})
			    m[colName] = *val
			}

			// Outputs: map[columnName:value columnName2:value2 columnName3:value3 ...] 
			fmt.Print(m)
		}
	    }
	}()

	return nil
}


func connec_db(db_type string, connect string) (*gorm.DB, error){
	if db_type == "postgres"{
		db2,err := gorm.Open("postgres", connect)
		if err != nil {
			logrus.Infof("erororor")
			logrus.Errorln(err)
			logrus.Fatalln("database ping attempts failed")
		}
		logrus.Infof("连接成功了")
		logrus.Infof("99999999999999999999999")
		return db2, err

		//_,err = db.Exec(sql)
		//if err != nil{
		//	return err
		//}
	} else if db_type == "mysql"{
		//db,err := gorm.Open("mysql", "host=116.62.213.56 port=51603 user=root dbname=dougo password=huansi@2017")
		db2,err := gorm.Open("mysql", connect)

		if err != nil {
			logrus.Errorln(err)
			logrus.Fatalln("database ping attempts failed")
		}
		logrus.Infof("连接成功了")
		return db2, err

		//_,err = db.Exec(sql)
		//if err != nil{
		//	return err
		//}
	}
	//else if db_type == "sqlite"{
	//	db,err := gorm.Open("sqlite3", "test.db")
	//	if err != nil {
	//		logrus.Errorln(err)
	//		logrus.Fatalln("database ping attempts failed")
	//	}
	//	logrus.Infof("连接成功了")
	//	if err != nil{
	//		return err
	//	}

	//	_,err = db.Exec(sql)
	//	if err != nil{
	//		return err
	//	}
	//}
	return nil,nil
}

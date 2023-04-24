package mysql

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type ConnectionInfoMS struct{
	User string
	Pass string
	Host string
	Port string
	DbName string
}


type MySqlDb struct{
	db *sql.DB
}

func (d *MySqlDb) Connect(c *ConnectionInfoMS)(error){
	driver:="mysql"
	//user:= c.User //"root"
	//pass:= c.Pass //"123456789"
	//dbName:=c.DbName // "sistema"

	conn,err:= sql.Open(driver, c.User+":"+c.Pass+"@tcp("+ c.Host +")/"+c.DbName)

	if(err != nil){
		//panic(err.Error())
		return err
	}
	d.db  = conn
	return nil
}

func (d *MySqlDb)Query(query string)(*sql.Rows, error){
	return d.db.Query(query)
}

func (d *MySqlDb)Close(){
	d.db.Close()
}
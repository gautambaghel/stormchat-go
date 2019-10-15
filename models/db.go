package models

import (
	"github.com/astaxie/beego/orm"
	// Required for database connection
	_ "github.com/lib/pq"
)

var ormObject orm.Ormer

// ConnectToDb - Initializes the ORM and Connection to the postgres DB
func ConnectToDb() {
	orm.RegisterDriver("postgres", orm.DRPostgres)
	orm.RegisterDataBase("default", "postgres", "user=default password=1234 dbname=stormchat_dev  host=localhost sslmode=disable")
	orm.RegisterModel(new(Users))
	ormObject = orm.NewOrm()
}

// GetOrmObject - Getter function for the ORM object with which we can query the DB
func GetOrmObject() orm.Ormer {
	return ormObject
}

package databaseconn

import (
	"fmt"
	"log"
	"time"

	"github.com/ahmadirfaan/project-go/app"
	"github.com/ahmadirfaan/project-go/models/database"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitDb() *gorm.DB {
	app := app.Init()
	maxIdleConn := app.Config.DBMaxIdleConnections
	maxConn := app.Config.DBMaxConnections
	maxLifetimeConn := app.Config.DBMaxLifetimeConnections
	db_user := app.Config.DBUsername
	db_pass := app.Config.DBPassword
	db_host := app.Config.DBHost
	db_port := app.Config.DBPort
	db_database := app.Config.DBName
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", db_user, db_pass, db_host, db_port, db_database)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger:                 logger.Default,
		SkipDefaultTransaction: true,
	})
	if err != nil {
		panic(err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	sqlDB.SetMaxIdleConns(maxIdleConn)
	sqlDB.SetMaxOpenConns(maxConn)
	sqlDB.SetConnMaxLifetime(time.Duration(maxLifetimeConn))

	InitCreateTable(db)
	log.Println("database connect success")
	return db

}

func InitCreateTable(db *gorm.DB) {

	db.Debug().AutoMigrate(
		&database.Role{},
		&database.Agent{},
		&database.Customer{},
		&database.User{},
		&database.ServiceTypeTransaction{},
		&database.TransactionType{},
		&database.Transactions{},
	)

	/**
	  Untuk menjalankan pertama kali supaya tersimpan data role dan admin
	*/
	//roleRepo := repositories.NewRoleRepository(db)
	//roleAdmin := database.Role{
	//	Role: "Admin",
	//}
	//roleCustomer := database.Role{
	//	Role: "Customer",
	//}
	//roleRepo.Save(roleAdmin)
	//roleRepo.Save(roleCustomer)

}

package database

import(
	"os"
	"gorm.io/driver/postgres"
  	"gorm.io/gorm"
)

func InitDB() (*gorm.DB, error){
	var err error
	connectpostgres := os.Getenv("POSTGRES_CONNECT_STRING")
	DB, err := gorm.Open(postgres.Open(connectpostgres), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	if err != nil{
		panic(err)
	}

	return DB, nil
}
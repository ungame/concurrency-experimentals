package configs

import "fmt"

func GetMongoDbDsn() string {
	return fmt.Sprintf("mongodb://%s:%s@127.0.0.1:27017/%s", "root", "root", "mydb")
}

func GetPostgresDsn() string {
	return fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s", "root", "root", "mydb", "disable")
}

func GetMysqlDsn() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", "root", "root", "localhost", 3306, "mydb")
}

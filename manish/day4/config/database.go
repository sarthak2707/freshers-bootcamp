package config

import "github.com/jinzhu/gorm"

var DB *gorm.DB

type DBConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
}

func BuildDBConfig() *DBConfig {
	return &DBConfig{
		Host:     "localhost",
		Port:     "3306",
		Username: "root",
		Password: "Vijay@7392",
		DBName:   "ecommerce",
	}
}

func DbURL(dbConfig *DBConfig) string {
	return dbConfig.Username + ":" + dbConfig.Password + "@tcp(" + dbConfig.Host + ":" + dbConfig.Port + ")/" + dbConfig.DBName + "?charset=utf8&parseTime=True&loc=Local"
}

package config

// DBConfig represents db configuration
import (
	"fmt"
	"go-project/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

type DBConfig struct {
	Driver   string
	Host     string
	Port     int
	User     string
	DBName   string
	Password string
}

func BuildDBConfig() *DBConfig {
	dbConfig := DBConfig{
		Host:     "localhost",
		Port:     3308,
		User:     "root",
		Password: "12345678",
		DBName:   "test",
	}
	return &dbConfig
}

func DbURL(db *DBConfig) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		db.User,
		db.Password,
		db.Host,
		db.Port,
		db.DBName,
	)
}

func DbConnection() {
	DB, err = gorm.Open(mysql.Open(DbURL(BuildDBConfig())), &gorm.Config{})
	if err != nil {
		fmt.Println("Status:", err)
	}
	DB.AutoMigrate(&models.Role{}, &models.User{})
}

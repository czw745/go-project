package Config

// DBConfig represents db configuration
import (
	"fmt"

	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

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
		Driver:   "mysql",
		Host:     "localhost",
		Port:     3308,
		User:     "root",
		Password: "12345678",
		DBName:   "test",
	}
	return &dbConfig
}

func DbDriver(db *DBConfig) string {
	return fmt.Sprintf(db.Driver)
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

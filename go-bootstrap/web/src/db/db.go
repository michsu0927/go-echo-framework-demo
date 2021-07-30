package db

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB
var err error

//Init initial Database connection
func Init() {
	// Db is database Object
	//db connection
	//var err error https://gorm.io/zh_CN/docs/gorm_config.html
	//ms-sql dsn
	dsn := "sqlserver://username:pwd@hostorip?database=database+timeout=30"
	db, err = gorm.Open(sqlserver.Open(dsn), &gorm.Config{PrepareStmt: true, SkipDefaultTransaction: true})
	if err != nil {
		fmt.Println(db.Error)
		//panic("failed to connect database")
	}
	//db = Database

	// show  sql string on console 這一段是把 sql 丟到 console
	// newLogger := logger.New(
	// 	log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
	// 	logger.Config{
	// 		SlowThreshold: time.Second, // Slow SQL threshold
	// 		LogLevel:      logger.Info, // Log level
	// 		Colorful:      false,       // Disable color
	// 	},
	// )

	now := time.Now()
	formatted := fmt.Sprintf("%d-%02d-%02dT%02d_%02d_%02d", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
	f, err := os.OpenFile("log/"+formatted+"-log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()

	// show  sql string on console 這一段是把 sql 丟到 file
	newLogger := logger.New(
		log.New(f, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      logger.Info, // Log level
			Colorful:      false,       // Disable color
		},
	)

	db.Logger = newLogger
	// show sql string end
}

//Manager Return Database Struct
func Manager() *gorm.DB {
	return db
}

/*usage
Db := db.Manager()
rows, err := Db.Table("user").Select("player_id as playerID,password as Password,encrypt as Encrypt,status as Status,id as ID").Where("player_id = ?", un).Rows()
if err != nil {
	panic("failed to connect database")
}
defer rows.Close() //要記得 close
playerID := 0
password := ""
encrypt := ""
status := 0
ID := 0
for rows.Next() {
	err = rows.Scan(&playerID, &password, &encrypt, &status, &ID)
	if err != nil {
		fmt.Printf("Scan failed,err:%v\n", err)
		panic("Scan failed\n")
	}
}
*/

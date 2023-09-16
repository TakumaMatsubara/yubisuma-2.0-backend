package app

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Animal struct {
	ID            string
	Name          string
	NumFinger     int
	Skill         string
	SkillDesc     string
	HandUrl       string
	UpFingerUrl   string
	DownFingerUrl string
}


// GameState holds the state of the Yubisuma game
type GameState struct {
	numMyFinger 	int    
	numNpcFinger    int
	mySkill   		string
	mySkillDesc		string
	usedMySkill		bool
	npcSkill		string
	npcSkillDesc    string
	usedNpcSkill	bool
	activeInk		bool
	numMyMaxFinger  int
	numNpcMaxFinger int
}

// app/db.go (新しく作成するファイルもしくはapp内の適当なファイルに追加)

var db *gorm.DB

func InitDB() {
	if err := godotenv.Load(); err != nil {
		log.Println("Error loading .env")
	}

	user := os.Getenv("SQL_USER")
	password := os.Getenv("SQL_PASSWORD")
	host := os.Getenv("SQL_HOST")
	port := os.Getenv("SQL_PORT")
	database := os.Getenv("SQL_DATABASE")

	var err error

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, database)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("can't connect to mysql server. "+
			"SQL_USER=%s, "+
			"SQL_PASSWORD=%s, "+
			"SQL_HOST=%s, "+
			"SQL_PORT=%s, "+
			"SQL_DATABASE=%s, "+
			"error=%+v",
			user, password, host, port, database, err)
	}
}

func GetAnimals() ([]Animal, error) {
	var animals []Animal

	if err := db.Find(&animals).Error; err != nil {
		log.Printf("Error while fetching all animals: %v\n", err)
		return nil, err
	}

	return animals, nil
}

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
	NumMyFinger     int    `json:"numMyFinger"`
	NumNpcFinger    int    `json:"numNpcFinger"`
	MySkill         string `json:"mySkill"`
	UsedMySkill     bool   `json:"usedMySkill"`
	NpcSkill        string `json:"npcSkill"`
	NpcSkillDesc    string `json:"npcSkillDesc"`
	UsedNpcSkill    bool   `json:"usedNpcSkill"`
	ActiveInk       bool   `json:"activeInk"`
	NumMyMaxFinger  int    `json:"numMyMaxFinger"`
	NumNpcMaxFinger int    `json:"numNpcMaxFinger"`
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

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, database)
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Unable to connect to the database: %v", err)
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

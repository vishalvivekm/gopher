package main

import (
	"fmt"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Player struct {
	ID       uint `gorm:"primaryKey"`
	Name     string
	LastName string `gorm:"column:last_name"`
	Team     string
}

func main() {
	// connects to an in-memory SQLite DB and then populates 'players' with data!
	// SQLite in-memory databases are databases stored entirely in memory, not on disk.
	//Use the special data source filename :memory: to create an in-memory database.
	// When the connection is closed, the database is deleted.
	//When using :memory: , each connection creates its own database

	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	err = db.AutoMigrate(&Player{})
	if err != nil {
		log.Fatal(err)
	}
	populateDatabase(db)

	// Create the firstPlayer and lastPlayer variables of the Player type below:
	var firstPlayer, lastPlayer Player

	// Get the First and Last players from the database:
	db.First(&firstPlayer)
	db.Last(&lastPlayer)

	// Print the Name, LastName and Team of the firstPlayer and lastPlayer below:
	fmt.Println(firstPlayer.Name, firstPlayer.LastName, firstPlayer.Team)
	fmt.Println(lastPlayer.Name, lastPlayer.LastName, lastPlayer.Team)
}
func populateDatabase(db *gorm.DB) {
	players := []Player{
		{
			ID:       1,
			Name:     "Virat",
			LastName: "Kohli",
			Team:     "India",
		},
		{
			ID:       2,
			Name:     "Mitchel",
			LastName: "Stark",
			Team:     "Australia",
		},
	}
	for _, player := range players {
		db.Create(&player)
	}
}

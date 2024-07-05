package main

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

type Artist struct {
	ArtistId uint `gorm:"column:artistid"`
	Name     string
}

func main() {

	db, err := gorm.Open(sqlite.Open("chinook.db"), &gorm.Config{}) // gorm.Config: https://gorm.io/docs/gorm_config.html
	if err != nil {
		log.Fatal(err)
	}
	var firstArtist, takeArtist, lastArtist Artist
	db.First(&firstArtist) // SELECT * FROM artists ORDER BY artistid LIMIT 1;
	db.Take(&takeArtist)   // SELECT * FROM artists LIMIT 1;
	db.Last(&lastArtist)   // SELECT * FROM artists ORDER BY artistid DESC LIMIT 1;

	fmt.Println(firstArtist)
	fmt.Println(takeArtist)
	fmt.Println(lastArtist)

	var artists []Artist

	db.Find(&artists) // select * from artists;

	fmt.Println("\nartists: ")
	for _, artist := range artists {
		fmt.Println(artist)
	}
}

// resources:
/*
https://gorm.io/docs/

*/

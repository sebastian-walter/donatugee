package main

import (
	"time"

	"github.com/jinzhu/gorm"
	"os"
	"fmt"
)

type Application struct {
	ApplicationID int
	Created       time.Time
	Modified      time.Time
}

type Donator struct {
	DonatorID  int
	Challenges []Challenge
	Name       string
	Profile    string
	Image      string
	Created    time.Time
	Modified   time.Time

	ChallengeRefer uint
}

type Techfugee struct {
	TechfugeeID  int
	Applications []Application
	Name         string
	Email        string
	Created      time.Time
	Modified     time.Time

	ApplicationsRefer uint
}

type Challenge struct {
	ChallengeID  int
	Applications []Application
	Name         string
	Image        string
	Description  string
	Created      time.Time
	Modified     time.Time

	ApplicationRefer uint
}

type Donatugee struct {
	db *gorm.DB
}

func OpenDatabase() (db *DB, err error) {
	if os.Getenv("ENV") == "production" {
		db, err := gorm.Open("postgres",
			fmt.Sprintf("host=%s user=%s dbname=%s sslmode=enable password=%s",
				os.Getenv("P_HOST"),
				os.Getenv("P_USER"),
				os.Getenv("P_DB"),
				os.Getenv("P_PW")))
		return db, err
	} else {
		db, err := gorm.Open("sqlite3", "db.sqlite")
		return db, err
	}
}

func NewDonatugee() (*Donatugee, error) {
	db, err := OpenDatabase()
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(&Donatugee{})
	return &Donatugee{
		db: db,
	}, nil
}

func (d *Donatugee) IntializeDB() []error {
	errs := d.db.AutoMigrate(&Application{}).GetErrors()
	if len(errs) != 0 {
		return errs
	}
	errs = d.db.AutoMigrate(&Techfugee{}).GetErrors()
	if len(errs) != 0 {
		return errs
	}

	errs = d.db.AutoMigrate(&Challenge{}).GetErrors()
	if len(errs) != 0 {
		return errs
	}

	// errs = d.db.Model(&Donator{}).Related(&Item{}).GetErrors()
	// if len(errs) != 0 {
	// 	return errs
	// }

	return nil
}

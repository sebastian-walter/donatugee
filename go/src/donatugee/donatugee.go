package main

import (
	"time"

	"github.com/jinzhu/gorm"
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

func NewDonatugee() (*Donatugee, error) {
	db, err := gorm.Open("sqlite3", "db.sqlite")
	if err != nil {
		return nil, err
	}
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

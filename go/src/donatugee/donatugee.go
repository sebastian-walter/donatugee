package main

import (
	"time"

	"fmt"
	"os"

	"github.com/jinzhu/gorm"
)

type Application struct {
	gorm.Model
	ApplicationID uint
	TechfugeeID   uint `sql:"type: integer REFERENCES techfugees(id)"`
	ChallengeID   uint `sql:"type: integer REFERENCES challenges(id)"`

	Created  time.Time
	Modified time.Time
}

type Donator struct {
	gorm.Model
	Challenges []Challenge `gorm:"ForeignKey:ID"`

	Name     string
	Profile  string
	Image    string
	Created  time.Time
	Modified time.Time
}

type Techfugee struct {
	gorm.Model
	Applications  []Application
	Name          string
	Email         string
	Skills        string
	Authenticated string
	Created       time.Time
	Modified      time.Time
}

type Challenge struct {
	gorm.Model
	ChallengeID  uint
	DonatorID    uint `sql:"type: integer REFERENCES donators(id)"`
	Applications []Application
	Name         string
	Image        string
	Description  string
	Created      time.Time
	Modified     time.Time
}

type Donatugee struct {
	db *gorm.DB
}

func OpenDatabase(dbname string) (db *gorm.DB, err error) {
	if os.Getenv("ENV") == "production" {
		db, err := gorm.Open("postgres",
			fmt.Sprintf("host=%s user=%s dbname=%s sslmode=require password=%s",
				os.Getenv("P_HOST"),
				os.Getenv("P_USER"),
				os.Getenv("P_DB"),
				os.Getenv("P_PW")))
		return db, err
	} else {
		db, err := gorm.Open("sqlite3", dbname)
		db.Exec("PRAGMA foreign_keys = ON;")
		return db, err
	}
}

func NewDonatugee(dbname string) (*Donatugee, error) {
	db, err := OpenDatabase(dbname)
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(&Donatugee{})
	return &Donatugee{
		db: db,
	}, nil
}

func (d *Donatugee) Challenges() ([]Challenge, error) {
	return []Challenge{}, nil
}

func (d *Donatugee) InsertTechfugee(name, email, skills string) []error {
	techfugee := Techfugee{
		Name:     name,
		Email:    email,
		Skills:   skills,
		Created:  time.Now(),
		Modified: time.Now(),
	}

	return d.db.Create(&techfugee).GetErrors()
}

func (d *Donatugee) IntializeDB() []error {
	errs := d.db.CreateTable(&Techfugee{}, &Donator{}, &Challenge{}, &Application{}).GetErrors()
	if len(errs) != 0 {
		return errs
	}

	return nil
}

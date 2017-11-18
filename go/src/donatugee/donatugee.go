package main

import (
	"time"

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

func NewDonatugee(dbname string) (*Donatugee, error) {
	db, err := gorm.Open("sqlite3", dbname)
	db.Exec("PRAGMA foreign_keys = ON;")
	if err != nil {
		return nil, err
	}
	return &Donatugee{
		db: db,
	}, nil
}

func (d *Donatugee) GetChallenges() ([]Challenge, error) {
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

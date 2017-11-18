package main

import (
	"time"

	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Applicaton struct {
	ApplicationID int
	Created       time.Time
	Modified      time.Time
}

type Donator struct {
	DonatorID int
	Items     []Item `gorm:"ForeignKey:ItemRefer"`
	Name      string
	Profile   string
	Image     string
	Created   time.Time
	Modified  time.Time

	ItemRefer uint
}

type Techfugee struct {
	TechfugeeID  int
	Applications []Applicaton `gorm:"ForeignKey:ApplicationRefer"`
	Name         string
	Email        string
	Created      time.Time
	Modified     time.Time

	ApplicationsRefer uint
}

type Item struct {
	ID           int
	Applications []Applicaton `gorm:"ForeignKey:ApplicationRefer"`
	Name         string
	Image        string
	Description  string
	Created      time.Time
	Modified     time.Time

	ApplicatonRefer uint
}

type Donatugee struct {
	db *gorm.DB
}

type Server struct {
	donatugee *Donatugee
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
	errs := d.db.Model(&Applicaton{}).GetErrors()
	if len(errs) != 0 {
		return errs
	}
	errs = d.db.Model(&Techfugee{}).Related(&Applicaton{}).GetErrors()
	if len(errs) != 0 {
		return errs
	}

	// errs = d.db.Model(&Item{}).Related(&Applicaton{}).GetErrors()
	// if len(errs) != 0 {
	// 	return errs
	// }

	// errs = d.db.Model(&Donator{}).Related(&Item{}).GetErrors()
	// if len(errs) != 0 {
	// 	return errs
	// }

	return nil
}

func main() {
	d, err := NewDonatugee()
	if err != nil {
		panic(err)
	}

	errs := d.IntializeDB()
	if len(errs) != 0 {
		panic(fmt.Sprintf("%v", errs))
	}
}

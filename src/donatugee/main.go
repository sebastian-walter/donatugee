package main

import (
	"time"

	"os"

	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	_ "github.com/jinzhu/gorm/dialects/postgres"
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

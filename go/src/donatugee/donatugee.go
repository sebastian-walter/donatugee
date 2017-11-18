package main

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Application struct {
	gorm.Model
	ApplicationID uint
	DonatorID     uint
	ChallengeID   uint

	Created  time.Time
	Modified time.Time
}

type Donator struct {
	gorm.Model
	DonatorID  uint
	Challenges []Challenge `gorm:"ForeignKey:DonatorID"`

	Name     string
	Profile  string
	Image    string
	Created  time.Time
	Modified time.Time
}

type Techfugee struct {
	gorm.Model
	TechfugeeID  uint
	Applications []Application `gorm:"ForeignKey:TechfugeeID"`
	Name         string
	Email        string
	Created      time.Time
	Modified     time.Time
}

type Challenge struct {
	gorm.Model
	ChallengeID  uint
	DonatorID    uint
	Applications []Application `gorm:"ForeignKey:ChallengeID"`
	Name         string
	Image        string
	Description  string
	Created      time.Time
	Modified     time.Time
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

func (d *Donatugee) GetChallenges() ([]Challenge, error) {
	return []Challenge{}, nil
}

func (d *Donatugee) IntializeDB() []error {
	// .AddForeignKey("applications_refer", "application(application_id)", "RESTRICT", "RESTRICT")
	errs := d.db.AutoMigrate(&Techfugee{}).GetErrors()
	if len(errs) != 0 {
		return errs
	}

	errs = d.db.AutoMigrate(&Donator{}).GetErrors()
	if len(errs) != 0 {
		return errs
	}

	errs = d.db.AutoMigrate(&Challenge{}).GetErrors()
	if len(errs) != 0 {
		return errs
	}

	// errs = d.db.AutoMigrate(&Application{}).AddForeignKey("techfugee_id", "techfugees(techfugee_id)", "RESTRICT", "RESTRICT").GetErrors()
	// if len(errs) != 0 {
	// 	return errs
	// }

	// errs = d.db.AutoMigrate(&Challenge{}).GetErrors()
	// if len(errs) != 0 {
	// 	return errs
	// }

	// errs = d.db.Model(&Donator{}).Related(&Item{}).GetErrors()
	// if len(errs) != 0 {
	// 	return errs
	// }

	return nil
}

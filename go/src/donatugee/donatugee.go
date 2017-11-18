package main

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
)

type Application struct {
	gorm.Model
	ApplicationID uint
	TechfugeeID   uint `sql:"type: integer REFERENCES techfugees(id)"`
	ChallengeID   uint `sql:"type: integer REFERENCES challenges(id)"`
}

type Donator struct {
	gorm.Model
	Challenges []Challenge `gorm:"ForeignKey:ID"`

	Name    string
	Profile string
	Image   string
}

type Techfugee struct {
	gorm.Model
	Applications  []Application
	Name          string
	Email         string
	Skills        string
	Authenticated string
}

type Challenge struct {
	gorm.Model
	ChallengeID  uint
	DonatorID    uint `sql:"type: integer REFERENCES donators(id)"`
	Applications []Application
	Name         string
	Image        string
	Description  string
}

type Donatugee struct {
	db *gorm.DB
}

func OpenDatabase(dbname string) (db *gorm.DB, err error) {
	if os.Getenv("DB") == "postgres" {
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

	return &Donatugee{
		db: db,
	}, nil
}

func (d *Donatugee) Techfugees() ([]Techfugee, []error) {
	var techfugees []Techfugee
	errs := d.db.Debug().Find(&techfugees).GetErrors()
	return techfugees, errs

}

func (d *Donatugee) Challenges() ([]Challenge, error) {
	return []Challenge{}, nil
}

func (d *Donatugee) Techfugee(id string) (Techfugee, []error) {
	var techfugee Techfugee
	errs := d.db.First(&techfugee, "id = ?", id).GetErrors()
	return techfugee, errs
}

func (d *Donatugee) Challenge(id string) (Challenge, []error) {
	var challenge Challenge
	errs := d.db.First(&challenge, "id = ?", id).GetErrors()
	return challenge, errs
}

func (d *Donatugee) UpdateTechfugeeSkills(techfugee Techfugee, skills string) (Techfugee, []error) {
	techfugee.Skills = skills
	errs := d.db.Save(&techfugee).GetErrors()
	return techfugee, errs
}

func (d *Donatugee) InsertTechfugee(name, email, skills string) (Techfugee, []error) {
	techfugee := Techfugee{}
	errs := d.db.Where(&Techfugee{}, "email = ?", email).GetErrors()
	if len(errs) > 0 {
		return techfugee, errs
	}

	if techfugee.Email == email {
		return techfugee, nil
	}

	techfugee = Techfugee{
		Name:   name,
		Email:  email,
		Skills: skills,
	}

	return techfugee, d.db.Create(&techfugee).GetErrors()
}

func (d *Donatugee) IntializeDB() []error {
	errs := d.db.AutoMigrate(&Techfugee{}, &Donator{}, &Challenge{}, &Application{}).GetErrors()
	if len(errs) != 0 {
		return errs
	}

	return nil
}

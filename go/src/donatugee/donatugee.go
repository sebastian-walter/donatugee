package main

import (
	"fmt"
	"os"

	"strconv"

	"github.com/jinzhu/gorm"
)

type Application struct {
	gorm.Model
	ApplicationID uint
	TechfugeeID   uint `sql:"type: integer REFERENCES techfugees(id)"`
	ChallengeID   uint `sql:"type: integer REFERENCES challenges(id)"`
	Accepted      bool
}

type Donator struct {
	gorm.Model
	Challenges []Challenge `gorm:"ForeignKey:ID"`

	Name    string
	Email   string
	Profile string
	Image   string
}

type Techfugee struct {
	gorm.Model
	Applications  []Application
	Name          string
	Email         string
	Skills        string
	City          string
	Introduction  string
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
	errs := d.db.Preload("Applications").Find(&techfugees).GetErrors()
	return techfugees, errs

}

func (d *Donatugee) UpdateAuth(id string, passed string) (Techfugee, []error) {
	var techfugee Techfugee
	newID, err := strconv.Atoi(id)
	if err != nil {
		return techfugee, []error{err}
	}
	errs := d.db.First(&techfugee, "id = ?", newID).GetErrors()
	if len(errs) > 0 {
		return techfugee, errs
	}

	techfugee.Authenticated = passed
	return techfugee, d.db.Save(&techfugee).GetErrors()
}

func (d *Donatugee) ChallengesByTechfugee(idTechfugee string) ([]Challenge, []error) {
	var applications []Application
	errs := d.db.Find(&applications, "techfugee_id = ?", idTechfugee).GetErrors()
	if len(errs) != 0 {
		return nil, errs
	}

	ids := []uint{}
	for _, e := range applications {
		ids = append(ids, e.ChallengeID)
	}

	var challenges []Challenge
	errs = d.db.Find(&challenges, "id IN (?)", ids).GetErrors()
	return challenges, errs

}

func (d *Donatugee) Challenges() ([]Challenge, []error) {
	var challenges []Challenge
	errs := d.db.Preload("Applications").Find(&challenges).GetErrors()
	return challenges, errs
}

func (d *Donatugee) Techfugee(id string) (Techfugee, []error) {
	var techfugee Techfugee
	newID, _ := strconv.Atoi(id)
	errs := d.db.Preload("Applications").First(&techfugee, "id = ?", newID).GetErrors()
	return techfugee, errs
}

func (d *Donatugee) LoginTechfugee(email string) (Techfugee, []error) {
	var techfugee Techfugee
	errs := d.db.Preload("Applications").First(&techfugee, "email = ?", email).GetErrors()
	return techfugee, errs
}

func (d *Donatugee) Challenge(id string) (Challenge, []error) {
	var challenge Challenge
	newID, err := strconv.Atoi(id)
	if err != nil {
		return Challenge{}, []error{err}
	}

	errs := d.db.Preload("Applications").First(&challenge, "id = ?", newID).GetErrors()
	return challenge, errs
}

func (d *Donatugee) Donator(id string) (Donator, []error) {
	var donator Donator
	newID, err := strconv.Atoi(id)
	if err != nil {
		return donator, []error{err}
	}

	errs := d.db.First(&donator, "id = ?", newID).GetErrors()
	return donator, errs
}

func (d *Donatugee) UpdateTechfugeeSkills(techfugee Techfugee, skills string) (Techfugee, []error) {
	techfugee.Skills = skills
	errs := d.db.Save(&techfugee).GetErrors()
	return techfugee, errs
}

func (d *Donatugee) UpdateTechfugee(id, city, introduction string) (Techfugee, []error) {
	var techfugees []Techfugee

	errs := d.db.Find(&techfugees, "id = ?", id).GetErrors()
	if len(errs) > 0 {
		return Techfugee{}, errs
	}

	if len(techfugees) == 0 {
		return Techfugee{}, []error{fmt.Errorf("no such techfugee: %s", id)}
	}

	techfugee := techfugees[0]
	techfugee.City = city
	techfugee.Introduction = introduction
	return techfugee, d.db.Save(&techfugee).GetErrors()
}

func (d *Donatugee) InsertTechfugee(name, email, skills string) (Techfugee, []error) {
	var techfugees []Techfugee
	errs := d.db.Find(&techfugees, "email = ?", email).GetErrors()
	if len(errs) != 0 {
		return Techfugee{}, errs
	}

	if len(techfugees) > 0 {
		return techfugees[0], nil
	}

	techfugee := Techfugee{
		Name:   name,
		Email:  email,
		Skills: skills,
	}

	return techfugee, d.db.Create(&techfugee).GetErrors()
}

func (d *Donatugee) InsertApplication(techfugee, challenge string) (Application, []error) {
	newID1, _ := strconv.Atoi(techfugee)
	newID2, _ := strconv.Atoi(challenge)

	var applications []Application
	errs := d.db.Find(&applications, "techfugee_id = ? AND challenge_id = ?", newID1, newID2).GetErrors()
	if len(errs) > 0 {
		return Application{}, errs
	}

	if len(applications) > 0 {
		return applications[0], []error{fmt.Errorf("exists already")}
	}

	application := Application{
		TechfugeeID: uint(newID1),
		ChallengeID: uint(newID2),
	}

	return application, d.db.Create(&application).GetErrors()
}

func (d *Donatugee) InsertDonator(name, email, profile, image string) (Donator, []error) {
	donator := Donator{}
	errs := d.db.Where(&donator, "email = ?", email).GetErrors()
	if len(errs) > 0 {
		return donator, errs
	}

	if donator.Email == email {
		return donator, nil
	}

	donator = Donator{
		Name:    name,
		Email:   email,
		Profile: profile,
		Image:   image,
	}

	return donator, d.db.Create(&donator).GetErrors()
}

func (d *Donatugee) IntializeDB() []error {
	errs := d.db.AutoMigrate(&Techfugee{}, &Donator{}, &Challenge{}, &Application{}).GetErrors()
	if len(errs) != 0 {
		return errs
	}

	return nil
}

func (d *Donatugee) InsertChallenge(idDonator, name, description string) (Challenge, []error) {
	id, err := strconv.ParseUint(idDonator, 10, 64)
	if err != nil {
		return Challenge{}, []error{err}
	}

	challenge := Challenge{
		DonatorID:   uint(id),
		Name:        name,
		Description: description,
	}

	errs := d.db.Create(&challenge).GetErrors()
	return challenge, errs
}

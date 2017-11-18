package main

import (
	"fmt"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func main() {
	d, err := NewDonatugee("db.sqlite")
	if err != nil {
		panic(err)
	}

	errs := d.IntializeDB()
	if len(errs) != 0 {
		panic(fmt.Sprintf("%v", errs))
	}

	s := NewServer(d)
	err = s.start()
	if err != nil {
		panic(err)
	}
}

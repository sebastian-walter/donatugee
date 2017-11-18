package main

import _ "github.com/jinzhu/gorm/dialects/sqlite"
import _ "github.com/jinzhu/gorm/dialects/postgres"

func main() {
	d, err := NewDonatugee()
	if err != nil {
		panic(err)
	}

	// errs := d.IntializeDB()
	// if len(errs) != 0 {
	// 	panic(fmt.Sprintf("%v", errs))
	// }

	s := NewServer(d)
	err = s.start()
	if err != nil {
		panic(s)
	}
}

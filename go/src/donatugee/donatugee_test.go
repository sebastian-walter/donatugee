package main

import (
	"os"
	"testing"
)

const DBName = "test.sqlite"

func TestWriteAndReadTechfugee(t *testing.T) {
	d, err := NewDonatugee(DBName)
	if err != nil {
		t.Fatalf("new donatugee: %v", err)
	}
	defer os.Remove(DBName)

	errs := d.IntializeDB()
	if len(errs) != 0 {
		t.Fatalf("initialize: %v", errs)
	}

	errs = d.InsertTechfugee("foo", "bar", "foobar")
	if len(errs) != 0 {
		t.Fatalf("insert: %v", errs)
	}

	techfugee := Techfugee{}
	errs = d.db.First(&techfugee).GetErrors()
	if len(errs) != 0 {
		t.Fatalf("first: %v", errs)
	}

	if techfugee.Name != "foo" {
		t.Fatalf("%s != %s", "foo", techfugee.Name)
	}

	techfugees, errs := d.Techfugees()
	if len(errs) != 0 {
		t.Fatalf("query: %v", errs)
	}

	if len(techfugees) < 0 {
		t.Fatalf("%d != %d", 0, len(techfugees))
	}
}

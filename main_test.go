package main

import "testing"

func TestLoadStructFields(t *testing.T) {
	fields, err := loadStructFields("foo.go", "Foo")
	if err != nil {
		t.Fatal(err)
	}

	if len(fields) != 3 {
		t.Errorf("Expected 2 fields, got %d", len(fields))
	}

	if !fields[0].Equal(TableField{Name: "ID", DBName: "id"}) {
		t.Errorf("Unexpected fields[0], got '%s'", fields[0])
	}

	if !fields[1].Equal(TableField{Name: "Name", DBName: "name"}) {
		t.Errorf("Unexpected fields[1], got '%s'", fields[1])
	}

	if !fields[2].Equal(TableField{Name: "Mail", DBName: "email"}) {
		t.Errorf("Unexpected fields[2], got '%s'", fields[2])
	}
}

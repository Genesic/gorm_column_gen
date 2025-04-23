package main

type Foo struct {
	ID   string
	Name string
	Mail string `gorm:"uniqueIndex;column:email"`
}

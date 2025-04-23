# gorm_column_gen

use gorm default naming strategy(snake_case) to generate column names

## install

```bash
go install github.com/Genesic/gorm_column_gen@latest
```


## usage

```bash
schema_column_parser -source foo.go -destination gen_foo.go -struct Foo -package model
```


## example

destination file
```go
foo.go

package model

type Foo struct {
    ID string
    Name string
    Mail int `gorm:"column:email"`
}
```

generated file
```go
gen_foo.go

package model

const (
    Foo_ID = "id"
    Foo_Name = "name"
    Foo_Mail = "email"
)
```


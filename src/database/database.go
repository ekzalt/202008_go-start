package main

// here is the example of using gorm https://github.com/go-gorm/gorm
type User struct {
	// auto-populate columns: id, created_at, updated_at, deleted_at
	gorm.Model

	// set column manually
	Username string `sql:"type:VARCHAR(255);not null;unique"`

	// set default value
	Lastname string `sql:"DEFAULT:'Smith'"`

	// custom column name instead of default snake_case format
	FullName string `gorm:"column:full_name"`

	// do not ctreate this column and do not fill this field
	Role string `sql:"-"`

	Salary int64
}

// TableName - set custom table name
func (u *User) TableName() string {
	return "users"
}

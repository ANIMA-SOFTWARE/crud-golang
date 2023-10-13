package types

import (
	"database/sql"
	"fmt"
	"log"
	"reflect"
)

type Users struct {
	_Rows       map[int]TableRow
	_Fields     []string
	_NextPageNo int
}

func (u *Users) FirstRow() TableRow {
	return u._Rows[0]
}

func (u *Users) Rows() map[int]TableRow {
	return u._Rows
}

func (u *Users) Fields() []string {
	return u.Row().FieldNames()
}

func (u *Users) Field(index int) string {
	return u.Row().FieldNames()[index]
}

func (u *Users) TableName() string {
	return "users"
}

func (u *Users) Row() TableRow {
	return &User{}
}

func (u *Users) MapFromSqlRows(rows *sql.Rows) {

	UsersMap := make(map[int]TableRow)
	mapID := 0
	for rows.Next() {

		User := &User{}

		err := rows.Scan(&User.ID, &User.Name, &User.Email)
		if err != nil {
			log.Fatal(err)
		}

		UsersMap[mapID] = User
		mapID++

	}

	u._Rows = UsersMap

}

func (u *Users) MapFromTableRows(rows []TableRow) {

	UsersMap := make(map[int]TableRow)
	for i := 0; i < len(rows); i++ {
		UsersMap[i] = rows[i]
	}
	u._Rows = UsersMap
}

func (u *Users) NextPageNumber() int {
	return u._NextPageNo
}

func (u *Users) SetPageNumber(PageNumber int) {
	u._NextPageNo = PageNumber + 1
}

type User struct {
	ID    int    `json:"id" default:"0"`
	Name  string `json:"name" default:"test"`
	Email string `json:"email" default:"testemail"`
}

func (u *User) GetID() int {
	return u.ID
}

func (u *User) Table() Table {
	return &Users{}
}

func (u *User) Value(index int) string {
	values := reflect.ValueOf(*u)
	v := values.Field(index).Interface()
	return fmt.Sprintf("%v", v)
}

func (u *User) Field(index int) string {
	t := reflect.TypeOf(*u)
	return t.Field(index).Name
}

func (u User) FieldNames() []string {
	t := reflect.TypeOf(u)

	fieldNameSlice := make([]string, t.NumField())

	for i := range fieldNameSlice {
		fieldNameSlice[i] = t.Field(i).Name
	}

	return fieldNameSlice
}

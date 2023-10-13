package types

import "database/sql"

type TableTypes struct {
	users  Users
	orders Orders
}

type Table interface {
	FirstRow() TableRow
	Rows() map[int]TableRow
	Fields() []string
	Field(int) string
	TableName() string
	Row() TableRow
	MapFromSqlRows(rows *sql.Rows)
	MapFromTableRows(rows []TableRow)
	NextPageNumber() int
	SetPageNumber(int)
}

type TableRow interface {
	GetID() int
	Table() Table
	FieldNames() []string
	Value(int) string
	Field(int) string
}

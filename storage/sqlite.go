package storage

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"main/types"
	"main/util"
)

type SqliteStorage struct {
	db  *sql.DB
	ctx context.Context
}

func NewSQLiteStorage(db *sql.DB, ctx context.Context) *SqliteStorage {

	return &SqliteStorage{
		db:  db,
		ctx: ctx,
	}
}

func (r *SqliteStorage) Migrate() {
	query := `
    CREATE TABLE IF NOT EXISTS users(
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name TEXT NOT NULL,
        email TEXT NOT NULL
    );
    `
	_, err := r.db.Exec(query)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("users table created")

	trans, _ := r.db.Begin()

	query = `INSERT INTO users (ID,Name,Email) VALUES (0,"Scott` + util.IntToStr(0) + `", "sl193@pm.me");`

	_, err = trans.Exec(query)

	if err != nil {
		log.Fatal(err)
	}

	for i := 1; i < 10000; i++ {

		query = `INSERT INTO users (Name,Email) VALUES ("Scott` + util.IntToStr(i) + `", "sl193@pm.me");`

		_, err = trans.Exec(query)

		if err != nil {
			log.Fatal(err)
		}

	}

	trans.Commit()

	fmt.Println("users table records created")

}

func (s *SqliteStorage) GetAll(t types.Table) types.Table {

	//send query
	rows, err := s.db.Query(`SELECT * FROM $1`, t.TableName())

	//panic if error with query execution
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	//table := t.New()
	t.MapFromSqlRows(rows)

	return t

}

func (s *SqliteStorage) Get(id int, t types.Table) types.Table {

	//send query

	stmt, _ := s.db.Prepare(`SELECT * FROM ` + t.TableName() + ` WHERE ID = ?`)
	rows, err := stmt.Query(id)

	if err != nil {
		log.Fatal(err)
	}

	t.MapFromSqlRows(rows)

	return t
}

func (s *SqliteStorage) GetBySearch(SearchString string, t types.Table) types.Table {

	//send query

	fmt.Println(SearchString)
	var query string = `SELECT * FROM ` + t.TableName() + ` WHERE `

	fmt.Println(t.Fields())

	for i := 0; i < len(t.Fields()); i++ {
		query += t.Field(i) + `= '` + SearchString + `' OR `
	}

	query = query[0 : len(query)-3]

	fmt.Println(query)

	rows, err := s.db.Query(query)
	//s.db.Query(`SELECT * FROM user WHERE user MATCH ?;`, SearchString)

	if err != nil {
		log.Fatal(err)
	}

	t.MapFromSqlRows(rows)

	return t

}

func (s *SqliteStorage) GetPage(PageNumber int, t types.Table) types.Table {

	//send query

	tableName := t.TableName()
	offset := "0"
	if PageNumber != 0 {
		offset = util.IntToStr((PageNumber) * 50)
	}
	rows, err := s.db.Query("SELECT * FROM " + tableName + " LIMIT 50 OFFSET " + offset + ";")
	fmt.Println("SELECT * FROM " + tableName + " LIMIT 50 OFFSET " + offset + ";")
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	//table := t.New()
	fmt.Println(t)
	t.MapFromSqlRows(rows)
	t.SetPageNumber(PageNumber)

	return t

}

func (s *SqliteStorage) Delete(row types.TableRow, t types.Table) types.Table {

	_, err := s.db.Exec(`DELETE FROM $1 WHERE ID = $2`, t.TableName(), row.GetID())

	if err != nil {
		log.Fatal(err)
	}

	rows, _ := s.db.Query(`SELECT * FROM $1`, t.TableName())

	//table := t.New()
	t.MapFromSqlRows(rows)

	return t
}

func (s *SqliteStorage) Append(row types.TableRow, t types.Table) types.Table {

	//create append query
	var query string = `UPDATE ` + t.TableName() + ` SET `

	for i := 1; i < len(row.FieldNames()); i++ {
		query += row.Field(i) + ` = '` + row.Value(i) + `', `
	}

	query = query[0 : len(query)-2]

	query += ` WHERE ID = ` + row.Value(0)

	_, err := s.db.Exec(query)

	if err != nil {
		log.Fatal(err)
	}

	//table := t.New()
	rows := []types.TableRow{}
	rows = append(rows, row)
	t.MapFromTableRows(rows)

	return t
}

func (s *SqliteStorage) Create(t types.Table) types.Table {

	stmt, _ := s.db.Prepare("INSERT INTO " + t.TableName() + "(name,email) VALUES(?,?)")
	res, err := stmt.Exec("TestName", "TestEmail")

	if err != nil {
		log.Fatal(err)
	}

	nextID, err := res.LastInsertId()

	if err != nil {
		log.Fatal(err)
	}

	stmt, _ = s.db.Prepare("SELECT * FROM " + t.TableName() + " WHERE ID = ?")
	rows, err := stmt.Query(nextID)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(rows)

	//table := t.New()
	t.MapFromSqlRows(rows)

	return t
}

func (s *SqliteStorage) GetTableType(tableType string) types.Table {
	switch tableType {
	case "users":
		return &types.Users{}
		//case "orders":
		//	return types.Orders{}
	}
	return &types.Users{}
}

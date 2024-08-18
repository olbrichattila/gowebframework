package db

// Todo add sync, to avoid database locks as multiple go routines can use it the same time
import (
	"database/sql"
	"fmt"
	"framework/internal/app/env"
	"runtime"

	// This needs to be blank imported as not directly referenced, but required

	_ "github.com/mattn/go-sqlite3" // Import the SQLite driver
)

type DBer interface {
	Construct(env.Enver)
	Open()
	Close()
	QueryAll(string, ...any) <-chan map[string]interface{}
	QueryOne(string, ...any) (map[string]interface{}, error)
	Execute(string, ...any) (int64, error)
	GetLastError() error
}

type DB struct {
	env       env.Enver
	db        *sql.DB
	lastError error
}

func New() DBer {
	db := &DB{}
	runtime.SetFinalizer(db, func(db *DB) {
		db.Cleanup()
	})
	return db
}

func (d *DB) Construct(env env.Enver) {
	d.env = env
	d.Open()
}

func (d *DB) Cleanup() {
	d.Close()
}

func (d *DB) Open() {
	conn := d.env.Get("DB_CONNECTION")
	var err error

	if conn == "sqlite" {
		d.db, err = sql.Open("sqlite3", d.env.Get("DB_DATABASE"))
		if err != nil {
			panic(err)
		}
		return
	}

	panic("Connection type " + conn + " does not implemented")
}

func (d *DB) Close() {
	if d.db != nil {
		d.db.Close()
	}
}

func (d *DB) QueryAll(sql string, pars ...any) <-chan map[string]interface{} {
	ch := make(chan map[string]interface{}, 1)
	d.lastError = nil
	if d.db == nil {
		d.lastError = fmt.Errorf("db not open")
		close(ch)
		return ch
	}

	stmt, err := d.db.Prepare(sql)
	if err != nil {
		d.lastError = err
		close(ch)
		return ch
	}

	rows, err := stmt.Query(pars...)
	if err != nil {
		d.lastError = err
		stmt.Close()
		close(ch)
		return ch
	}

	cols, err := rows.Columns()
	if err != nil {
		d.lastError = err
		rows.Close()
		stmt.Close()
		close(ch)
		return ch
	}

	colCount := len(cols)

	row := make([]interface{}, colCount)
	for i := range row {
		row[i] = new(interface{})
	}

	go func() {
		for rows.Next() {
			err := rows.Scan(row...)
			if err != nil {
				d.lastError = err
				break
			}
			result := make(map[string]interface{}, colCount)
			for i, colName := range cols {
				result[colName] = *(row[i].(*interface{}))
			}
			ch <- result
			result = nil
		}
		rows.Close()
		stmt.Close()
		close(ch)
	}()

	return ch
}

func (d *DB) QueryOne(sql string, pars ...any) (map[string]interface{}, error) {
	if d.db == nil {
		return nil, fmt.Errorf("db not open")
	}

	stmt, err := d.db.Prepare(sql)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(pars...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	cols, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	colCount := len(cols)
	row := make([]interface{}, colCount)
	for i := range row {
		row[i] = new(interface{})
	}

	if !rows.Next() {
		return nil, fmt.Errorf("row cannot be found")
	}

	err = rows.Scan(row...)
	if err != nil {
		return nil, err
	}

	result := make(map[string]interface{}, colCount)
	for i, colName := range cols {
		result[colName] = *(row[i].(*interface{}))
	}

	return result, nil
}

func (d *DB) Execute(sql string, pars ...any) (int64, error) {
	if d.db == nil {
		return 0, fmt.Errorf("db not open")
	}

	stmt, err := d.db.Prepare(sql)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(pars...)
	if err != nil {
		return 0, err
	}

	return res.LastInsertId()
}

func (d *DB) GetLastError() error {
	return d.lastError
}

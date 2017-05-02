package sx

import (
	"database/sql"
	"errors"
	// Oracle SQL specific Driver
	_ "github.com/mattn/go-oci8"
)

// TaskFinder осуществляет поиск атрибутов
// Дела с ЕПГУ
type TaskFinder struct {
	DB *sql.DB
}

// Task содержит информацию о Заявке с ПГУ
type Task struct {
	Number            string
	MessageID         string
	ReasonServiceCode string
	// Comment    string
}

// NewCaseFinder create new instance
// connection string format "user/passw@host:port/sid"
func NewCaseFinder(connString string) (*TaskFinder, error) {
	db, err := sql.Open("oci8", connString)
	if err != nil {
		return nil, err
	}

	// // Set timeout (Go 1.8)
	// ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	// // Set prefetch count (Go 1.8)
	// ctx = ora.WithStmtCfg(ctx, ora.Cfg().StmtCfg.SetPrefetchCount(50000))

	return &TaskFinder{db}, nil
}

// Find deal in DB
func (finder TaskFinder) Find(ummsID string) (Task, error) {
	SQL := `select REASONCASENUMBER, EXTNUMBER, REASONSERVICECODE, FOIVCODE from INTASK where REASONCASENUMBER = :1 and Rownum < 2`

	rows, err := finder.DB.Query(SQL, ummsID)
	if err != nil {
		return Task{}, err
	}
	defer rows.Close()

	if !rows.Next() {
		return Task{}, errors.New("Can't find Case by CaseID")
	}

	var caseID string
	var extNumber string
	var reasonServiceCode string
	var foivCode string
	rows.Scan(&caseID, &extNumber, &reasonServiceCode, &foivCode)
	return Task{Number: caseID, MessageID: extNumber, ReasonServiceCode: reasonServiceCode}, nil
}

// Ping test DB connection
func (finder TaskFinder) Ping() (err error) {
	_, err = finder.Find("175426039")
	return err
}

// Close test DB connection
func (finder TaskFinder) Close() {
	if finder.DB != nil {
		finder.DB.Close()
	}
}

package phone

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// Rough schema
// id | raw_number | normalized

type Number struct {
	Id   int            `db:"id"`
	Raw  string         `db:"raw"`
	Norm sql.NullString `db:"norm"`
}

var ctx = context.Background()
var dbOld *sql.DB

func OpenOld() *sql.DB {
	connStr := "user=aethemba dbname=phone sslmode=disable"
	dbOld, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal(err)
	}

	return dbOld
}

func AllNumbers() []string {
	db := OpenOld()
	rows, err := db.Query("SELECT raw from number;")
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()
	defer db.Close()

	numbers := make([]string, 0)
	for rows.Next() {
		var raw string
		if err := rows.Scan(&raw); err != nil {
			log.Fatal(err)
		}
		numbers = append(numbers, raw)
	}

	return numbers
}

func InsertNumber(number string) {
	db := OpenOld()

	tx, err := db.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		log.Fatal(err)
	}

	q := fmt.Sprintf("INSERT INTO NUMBER (RAW) VALUES ('%s');", number)
	_, execErr := tx.Exec(q)

	if execErr != nil {
		_ = tx.Rollback()
		log.Fatal("EXEC ", execErr)
	}

	if err := tx.Commit(); err != nil {
		log.Fatal(err)
	}
}

func Delete() {
	db := OpenOld()

	tx, err := db.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		log.Fatal(err)
	}

	q := fmt.Sprintf("DELETE FROM NUMBER")
	_, execErr := tx.Exec(q)

	if execErr != nil {
		_ = tx.Rollback()
		log.Fatal("EXEC ", execErr)
	}

	if err := tx.Commit(); err != nil {
		log.Fatal(err)
	}
}

func Normalize(number string) string {

	n := strings.TrimSpace(number)
	n = strings.Replace(n, "(", "", -1)
	n = strings.Replace(n, ")", "", -1)
	n = strings.Replace(n, " ", "", -1)
	n = strings.Replace(n, "-", "", -1)

	return n
}

func OpenX() *sqlx.DB {
	db, err := sqlx.Connect("postgres", "user=aethemba dbname=phone sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func InsertX(number string) {
	db := OpenX()
	tx := db.MustBegin()

	tx.MustExec("INSERT INTO number (raw) VALUES ($1)", number)
	err := tx.Commit()
	if err != nil {
		log.Fatal(err)
	}

}

func AllNumbersX() []Number {
	db := OpenX()
	numbers := []Number{}

	err := db.Select(&numbers, "SELECT * from number;")
	if err != nil {
		fmt.Println("err", err)
	}
	return numbers
}

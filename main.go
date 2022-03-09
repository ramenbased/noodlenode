package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

//-- Div

func Er(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

//-- Postgres

func dbPing(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Connected to DB! Connections open:", db.Stats().OpenConnections)
	}
}

func dbHeight(db *sql.DB) int {
	var (
		height    int
		blockhash string
		time      int
		medianfee int
		rv        int
	)

	rows, _ := db.Query("select * from blockstats")
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&height, &blockhash, &time, &medianfee)
		Er(err)
		if height > rv {
			rv = height
		}
	}
	return rv
}

func routine(db *sql.DB, heightInDB int) {

	count := GetBlockCount()
	fmt.Println("Current Block Count:", count.Result)

	for b := heightInDB + 1; b <= count.Result; b++ {

		stats := GetBlockStats(b).Result
		conf := GetBlock(stats.Blockhash)

		if conf.Result.Confirmations >= 6 {
			tx, err := db.Begin()
			Er(err)
			defer tx.Rollback()
			stmt, err := tx.Prepare("INSERT INTO blockstats VALUES ($1,$2,$3,$4)")
			Er(err)
			defer stmt.Close()
			_, err = stmt.Exec(
				stats.Height,    //int
				stats.Blockhash, //string
				stats.Time,      //int
				stats.Medianfee, //int
			)
			Er(err)
			err = tx.Commit()
			fmt.Println(conf)
			Er(err)
			fmt.Println("Commited Height to DB:", stats.Height)
		} else {
			fmt.Println("Block", stats.Height, "discarded due to", conf.Result.Confirmations, "confirmations")
		}

	}
}

func main() {
	connStr := "host=localhost user=postgres password=postgres port=5432 dbname=noodledb"
	db, err := sql.Open("postgres", connStr)
	Er(err)
	defer db.Close()
	dbPing(db)
	cycle := 0
	for {
		fmt.Println("#########################")
		heightInDB := dbHeight(db)
		fmt.Println("Height DB:", heightInDB)
		routine(db, heightInDB)
		time.Sleep(time.Second * 10)
		cycle++
		fmt.Println("Cycles:", cycle)
	}
}

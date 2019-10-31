package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	// Connect to database
	db, err := sql.Open("postgres", "postgres://root@localhost:26257/test?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	// Start Changefeed (streaming stmt)
	rows, err := db.Query("EXPERIMENTAL CHANGEFEED FOR movies WITH UPDATED;")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		_ = rows.Close()
	}()
	type changeFeed struct {
		table string
		key   string
		value []byte
	}
	// This blocks forever and triggers when a new result comes
	for rows.Next() {
		changeFeed := &changeFeed{}
		err := rows.Scan(&changeFeed.table, &changeFeed.key, &changeFeed.value)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(changeFeed.value))
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}

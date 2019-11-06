package main

import (
	"database/sql"
	"encoding/json"
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
		Table string `json:"table"`
		Key   string `json:"key"`
		Value []byte `json:"value"`
	}
	// This blocks forever and triggers when a new result comes
	for rows.Next() {
		changeFeed := &changeFeed{}
		err := rows.Scan(&changeFeed.Table, &changeFeed.Key, &changeFeed.Value)
		if err != nil {
			log.Fatal(err)
		}
		jsonChangeFeed, err := json.Marshal(changeFeed)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(jsonChangeFeed))
		fmt.Println("")
		fmt.Println("value decoded:\n", string(changeFeed.Value))
		fmt.Println("")
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}

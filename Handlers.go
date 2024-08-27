package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"

	_ "github.com/lib/pq"

	"github.com/gorilla/mux"
)

const (
	host     = "postgres"
	port     = 5432
	user     = "qotd"
	password = "reallylongpassword99!"
	dbname   = "quotesdb"
)

func AllQuotes(w http.ResponseWriter, r *http.Request) {
	USE_DB := os.Getenv("USE_DB")
	if USE_DB == "true" {
		sqlCommand := "SELECT quoteId as id, quotation, author FROM quotes ORDER BY quoteId;"
		quotes := GetFromDatabase((sqlCommand))
		json.NewEncoder(w).Encode(quotes)
	} else {
		json.NewEncoder(w).Encode(quotes)
	}
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "qotd Oct 11 2023 10:41 am")
}

func OneQuote(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["quoteId"])
	if err == nil {
		USE_DB := os.Getenv("USE_DB")
		if USE_DB == "true" {
			sqlCommand := fmt.Sprintf("SELECT quoteId as id, quotation, author FROM quotes WHERE quoteId = %d", id)
			quotes := GetFromDatabase((sqlCommand))
			json.NewEncoder(w).Encode(quotes[0])
		} else {
			json.NewEncoder(w).Encode(quotes[id])
		}
	}
}

func RandomQuote(w http.ResponseWriter, r *http.Request) {
	USE_DB := os.Getenv("USE_DB")
	if USE_DB == "true" {
		sqlCommand := "SELECT quoteId as id, quotation, author FROM quotes ORDER BY random() LIMIT 1;"
		quotes := GetFromDatabase((sqlCommand))
		json.NewEncoder(w).Encode(quotes[0])
	} else {
		json.NewEncoder(w).Encode(quotes[rand.Intn(7)])
	}
}

func Version(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "3.0.0")
}

func WrittenIn(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "GO")
}

func GetFromDatabase(sqlCommand string) []Quote {
	// Connection string
	host := os.Getenv("DBHOST")
	port := 5432
	user := os.Getenv("DBUSER")
	password := os.Getenv("DBPASSWORD")
	dbname := os.Getenv("DBNAME")

	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	// Open a connection to the database
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	// Verify the connection
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query(sqlCommand)
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()
	// Parse the results
	var quotes []Quote
	for rows.Next() {
		var quote Quote
		err := rows.Scan(&quote.Id, &quote.Quotation, &quote.Author)
		if err != nil {
			log.Fatal(err)
		}
		quotes = append(quotes, quote)
	}

	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	// Print the results
	for _, quote := range quotes {
		fmt.Printf("ID: %d, Quotation: %s, Author: %s\n", quote.Id, quote.Quotation, quote.Author)
	}
	return quotes
}

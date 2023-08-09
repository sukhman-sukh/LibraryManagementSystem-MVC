package models

import (
	"database/sql"

	"context"
	"fmt"
	"log"
	"os"
	"time"
	"gopkg.in/yaml.v3"
	"lib-manager/pkg/types"
	_ "github.com/go-sql-driver/mysql"

)

func dsn() string {  

	// dataFile, err := os.Open("data.yaml")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// defer dataFile.Close()

	// var config types.ConfigSet
	// decoder := yaml.NewDecoder(dataFile)
	// err = decoder.Decode(&config)
	// if err != nil {
	// 	log.Fatalf("failed to decode config: %v", err)
	// }
	configFile, err := os.Open("data.yaml")
	if err != nil {
		log.Fatalf("failed to open config file: %v", err)
	}
	defer configFile.Close()

	var config types.ConfigSet
	decoder := yaml.NewDecoder(configFile)
	err = decoder.Decode(&config)
	if err != nil {
		log.Fatalf("failed to decode config: %v", err)
	}

	fmt.Printf("%+v\n", config)

	fmt.Println("======================================")
	fmt.Printf( config.DB_USERNAME, config.DB_PASSWORD, config.DB_HOST, config.DB_NAME)
	fmt.Println("======================================")
	return fmt.Sprintf("%s:%s@tcp(%s)/%s", config.DB_USERNAME, config.DB_PASSWORD, config.DB_HOST, config.DB_NAME)

    }

func Connection() (*sql.DB, error) {  
	db, err := sql.Open("mysql", dsn())  
	if err != nil {  
		log.Printf("Error: %s when opening DB", err)
		return nil, err
	}

    db.SetMaxOpenConns(20)
    db.SetMaxIdleConns(20)
    db.SetConnMaxLifetime(time.Minute * 5)

    ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancelfunc()
    err = db.PingContext(ctx)
    if err != nil {
        log.Printf("Errors %s pinging DB", err)
        return nil, err
    }
    log.Printf("Connected to DB successfully\n")
	return db, err
}


// Check If DB is empty 
//  Return True for empty DB

// func IsDbEmpty(tableName string , db *sql.DB , condition string) bool{

// 	query := fmt.Sprintf("SELECT COUNT(*) FROM %s WHERE "+condition, tableName)

// 	var rowCount int
// 	err := db.QueryRow(query).Scan(&rowCount)
// 	if err != nil {
// 		log.Fatal("Error querying the database:", err)
// 	}

// 	if rowCount == 0 {
// 		fmt.Println("Database is empty")
// 		return true 
// 	}
// 	return false
// }
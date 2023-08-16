package models

import (
	"context"
	"database/sql"
	"fmt"
	"lib-manager/pkg/types"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/yaml.v3"
)

func dsn() string {

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
		fmt.Printf("Errors %s pinging DB", err)
		return nil, err
	}
	log.Printf("Connected to DB successfully\n")
	return db, err
}

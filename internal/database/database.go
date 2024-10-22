package database

import (
	"context"
	"log"
	"os"
	"path/filepath"

	"entgo.io/ent/dialect"
	_ "github.com/mattn/go-sqlite3"
	"github.com/nikonhub/ai-learn-cli/ent"
)

type DB struct {
	Client *ent.Client
}

func NewDB() *DB {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("failed to get user home directory: %v", err)
	}

	dbDir := filepath.Join(homeDir, ".ai-learn")
	if err := os.MkdirAll(dbDir, 0755); err != nil {
		log.Fatalf("failed to create database directory: %v", err)
	}

	dbPath := filepath.Join(dbDir, "ai-learn.db")
	client, err := ent.Open(dialect.SQLite, "file:"+dbPath+"?_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	return &DB{Client: client}
}

func (db *DB) InitSchema() error {
	return db.Client.Schema.Create(context.Background())
}

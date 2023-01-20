package dbc

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/sheason2019/spoved/ent"

	_ "github.com/lib/pq"
)

var ins *ent.Client

func initClient() {
	createDatabase(dbname)
	dsn := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable", host, port, user, dbname, password)
	client, err := ent.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	ins = client
}

func createDatabase(dbname string) {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s sslmode=disable", host, port, user, password)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}

	db.Exec("create database " + dbname)
}

// 懒汉单例避免重复创建
func GetClient() *ent.Client {
	if ins == nil {
		initClient()
	}

	return ins
}

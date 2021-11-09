package utils

import (
	"context"
	"database/sql"
	"fmt"
	"log"
)

type TableDec struct {
	Field   string
	Type    string
	Null    string
	Key     string
	Default interface{}
	Extra   string
}

func TableInfo(table string, db *sql.DB, ctx context.Context) ([]TableDec, error) {
	rows, err := db.QueryContext(ctx, fmt.Sprintf("DESCRIBE %v", table))
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = rows.Close()
	}()
	tableDecList := make([]TableDec, 0)
	for rows.Next() {
		tableDec := TableDec{}
		if err := rows.Scan(&tableDec.Field, &tableDec.Type, &tableDec.Null, &tableDec.Key, &tableDec.Default, &tableDec.Key); err != nil {
			log.Fatal(err)
		}
		tableDecList = append(tableDecList, tableDec)
	}
	return tableDecList, nil
}

func Tables(db *sql.DB, ctx context.Context) ([]string, error) {
	rows, err := db.QueryContext(ctx, "SHOW TABLES")
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = rows.Close()
	}()
	tableList := make([]string, 0)
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			log.Fatal(err)
		}
		tableList = append(tableList, name)
	}
	return tableList, nil
}

package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var url = flag.String("url", "dtd_db_test:5k2VUnDfqLxGnu@tcp(rm-bp1j6307pxpihh49090130.mysql.rds.aliyuncs.com:3306)/tch_cms", "数据库链接")

type TableDec struct {
	Field string
	Type  string
	Null  string
	Key   string
	Default interface{}
	Extra string
}


func main()  {
	flag.Parse()

	if len(*url) == 0 {
		log.Fatalln("必须输入数据库链接")
		return
	}
	db, err := sql.Open("mysql", *url)
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = db.Close()
	}()
	ctx := context.Background()

	tableList, err := tables(db, ctx)
	if err != nil {
		return 
	}
	if len(tableList) > 0 {
		for _,v := range tableList {
			err = builder(v, db, ctx)
			if err != nil {
				log.Fatalf("生成错误 %v", err)
			}
		}
	}
}

func builder(table string, db *sql.DB, ctx context.Context) error {
	tableDecList, err := tableInfo(table, db, ctx)
	if err != nil {
		return err
	}
	fmt.Println(tableDecList)
	return err
}


func tableInfo(table string, db *sql.DB, ctx context.Context) ([]TableDec, error) {
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
		if err := rows.Scan(&tableDec.Field, &tableDec.Type, &tableDec.Null, &tableDec.Key, &tableDec.Default , &tableDec.Key); err != nil {
			log.Fatal(err)
		}
		tableDecList = append(tableDecList, tableDec)
	}
	return tableDecList, nil
}

func tables(db *sql.DB, ctx context.Context) ([]string, error) {
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



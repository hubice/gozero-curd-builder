package main

import (
	"bytes"
	"context"
	"database/sql"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"text/template"

	_ "github.com/go-sql-driver/mysql"
)

var url = flag.String("url", "dtd_db_test:5k2VUnDfqLxGnu@tcp(rm-bp1j6307pxpihh49090130.mysql.rds.aliyuncs.com:3306)/tch_cms", "数据库链接")
var service = flag.String("service", "service", "服务名称")
var files = []string{
	"api",
	"logic",
	"model",
}

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
	for _, v := range files {
		temp := bytes.Buffer{}
		files, err := template.ParseFiles(fmt.Sprintf("./tmpl/%v.tmpl", v))
		if err != nil {
			return err
		}
		_ = files.ExecuteTemplate(&temp, "table", table)
		_ = files.ExecuteTemplate(&temp, "service", *service)
		_ = files.ExecuteTemplate(&temp, "columns", tableDecList)
		_ = os.MkdirAll(fmt.Sprintf("./builder/%v/", table), 0666)
		_ = ioutil.WriteFile(fmt.Sprintf("./builder/%v/%v.api", table, table), temp.Bytes(), 0666)
	}
	return err
}

func capitalize(str string) string {
	var upperStr string
	vv := []rune(str)
	for i := 0; i < len(vv); i++ {
		if i == 0 {
			if vv[i] >= 97 && vv[i] <= 122 {
				vv[i] -= 32
				upperStr += string(vv[i])
			} else {
				return str
			}
		} else {
			upperStr += string(vv[i])
		}
	}
	return upperStr
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



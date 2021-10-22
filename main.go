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
	"strings"
	"text/template"

	_ "github.com/go-sql-driver/mysql"
)

var url = flag.String("url", "dtd_db_test:5k2VUnDfqLxGnu@tcp(rm-bp1j6307pxpihh49090130.mysql.rds.aliyuncs.com:3306)/tch_cms", "数据库链接")
var service = flag.String("service", "account", "服务名称")

type TmplRecipient struct {
	Table        string
	Service      string
	TableDecList []TableDec
	TableDecListExpectAutoSet []TableDec
}

type TableDec struct {
	Field   string
	Type    string
	Null    string
	Key     string
	Default interface{}
	Extra   string
}

func main() {
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
		for _, v := range tableList {
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
	tableDecListExpectAutoSet :=make([]TableDec, 0, len(tableDecList))
	for _, v := range tableDecList {
		if v.Field != "id" && v.Field != "create_time" && v.Field != "update_time" && v.Field != "delete_time" {
			tableDecListExpectAutoSet = append(tableDecListExpectAutoSet, v)
		}
	}
	fi, err := ioutil.ReadDir("./tmpl")
	if err != nil {
		return err
	}
	for _, v := range fi {
		if !v.IsDir() {
			buff := bytes.Buffer{}
			data, err := ioutil.ReadFile(fmt.Sprintf("./tmpl/%v", v.Name()))
			if err != nil {
				return err
			}
			tmpl, err := template.New(fmt.Sprintf("./tmpl/%v", v.Name())).Funcs(template.FuncMap{
				"Case2Camel": Case2Camel,
				"Case2Mid": Case2Mid,
				"DbType2Type": DbType2Type,
				"Case2CamelFirst": Case2CamelFirst,
			}).Parse(string(data))
			if err != nil {
				return err
			}
			err = tmpl.Execute(&buff, TmplRecipient{
				Table:        table,
				Service:      *service,
				TableDecList: tableDecList,
				TableDecListExpectAutoSet: tableDecListExpectAutoSet,
			})
			if err != nil {
				return err
			}
			_ = os.MkdirAll(fmt.Sprintf("./builder/%v/", Case2Empty(table)), 0666)
			suffix := ".go"
			if strings.Contains(v.Name(), "api") {
				suffix = ".api"
			}
			_ = ioutil.WriteFile(fmt.Sprintf("./builder/%v/%v", Case2Empty(table), Case2Empty(fmt.Sprintf("%v%v%v",table, strings.Replace(v.Name(), ".", "", -1), suffix))), buff.Bytes(), 0666)
		}
	}
	return nil
}

func DbType2Type(name string) string {
	list := map[string]string {
		"int": "int64",
		"datetime": "int64",
		"varchar": "string",
		"char": "string",
	}
	for k, v := range list {
		if strings.Contains(name, k) {
			return v
		}
	}
	return name
}

func Case2CamelFirst(name string) string {
	name = strings.Replace(name, "_", " ", -1)
	name = strings.Title(name)
	name = strings.Replace(name, " ", "", -1)
	return strings.ToLower(name[0:1]) + name[1:]
}

func Case2Camel(name string) string {
	name = strings.Replace(name, "_", " ", -1)
	name = strings.Title(name)
	return strings.Replace(name, " ", "", -1)
}

func Case2Empty(name string) string {
	 return strings.Replace(name, "_", "", -1)
}

func Case2Mid(name string) string {
	return strings.Replace(name, "_", "-", -1)
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
		if err := rows.Scan(&tableDec.Field, &tableDec.Type, &tableDec.Null, &tableDec.Key, &tableDec.Default, &tableDec.Key); err != nil {
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

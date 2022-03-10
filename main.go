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

	"gozero-curd-builder/utils"
)

var database = flag.String("url", "", "数据库链接")

var baseDir = "./tmpl/"

type TmplData struct {
	Table                    string
	TableDecList             []utils.TableDec
}

func main() {
	flag.Parse()

	if len(*database) == 0 {
		log.Fatalln("输入错误:必须输入数据库链接")
		return
	}
	db, err := sql.Open("mysql", *database)
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = db.Close()
	}()
	ctx := context.Background()

	tableList, err := utils.Tables(db, ctx)
	if err != nil {
		return
	}
	if len(tableList) > 0 {
		for _, v := range tableList {
			err = builder(v, db, ctx)
			if err != nil {
				log.Fatalf("生成错误:%v", err)
			}
		}
	}
}

func builder(table string, db *sql.DB, ctx context.Context) error {
	tableDecList, err := utils.TableInfo(table, db, ctx)
	if err != nil {
		return err
	}
	tmplData := TmplData{
		Table:        table,
		TableDecList: tableDecList,
	}
	// 方法
	tmplFunc := template.FuncMap{
		"Case2CamelLower": utils.Case2CamelLower,
		"Case2CamelUpper": utils.Case2CamelUpper,
		"Case2Mid":        utils.Case2Mid,
		"DbType2Type":     utils.DbType2Type,
	}
	// 创建
	fi, err := ioutil.ReadDir(baseDir)
	if err != nil {
		return err
	}
	for _, v := range fi {
		if !v.IsDir() {
			buff := bytes.Buffer{}
			data, err := ioutil.ReadFile(baseDir + v.Name())
			if err != nil {
				return err
			}
			tmpl, err := template.New(baseDir + v.Name()).Funcs(tmplFunc).Parse(string(data))
			if err != nil {
				return err
			}
			err = tmpl.Execute(&buff, tmplData)
			if err != nil {
				return err
			}
			err = os.MkdirAll(fmt.Sprintf("%v/%v/", baseDir+"builder", utils.Case2Empty(table)), 0666)
			if err != nil {
				return err
			}
			err = ioutil.WriteFile(fmt.Sprintf("%v/%v/%v", baseDir+"builder", utils.Case2Empty(table), v.Name()), buff.Bytes(), 0666)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

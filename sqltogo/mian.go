package main

import (
	"database/sql"
	"flag"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/logrusorgru/aurora"
	"os"
	"strings"
	"work/sqltogo/sql/config"
	"work/sqltogo/sql/gen"
)

func main() {
	var (
		outputPath string
		tableStr   string
		tables     []string
	)
	flag.StringVar(&outputPath, "outputPath", "", "")
	flag.StringVar(&tableStr, "tableStr", "", "")
	flag.Parse()
	tables = strings.Split(tableStr, ",")
	Run(outputPath, tables...)
}

func Run(outputPath string, tables ...string) {
	fileName := GenFile(tables...)
	SqlToGo(fileName, outputPath)
	defer os.Remove(fileName) //删除文件
}

func GenFile(tables ...string) string {
	db, err := sql.Open("mysql", "xxxx")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	var (
		ddlStatements []string
	)
	for _, tableName := range tables {
		rows, err := db.Query(fmt.Sprintf("SHOW CREATE TABLE %s", tableName))
		if err != nil {
			panic(err.Error())
		}
		defer rows.Close()

		for rows.Next() {
			var ddlStatement string
			err = rows.Scan(&tableName, &ddlStatement)
			if err != nil {
				panic(err.Error())
			}
			ddlStatements = append(ddlStatements, ddlStatement)
		}
	}

	// 将DDL语句保存到临时文件
	tmpFile, err := os.CreateTemp("", "ddl-*.sql")
	if err != nil {
		panic(err.Error())
	}
	defer tmpFile.Close()

	for _, ddlStatement := range ddlStatements {
		_, err = tmpFile.WriteString(ddlStatement + ";\n")
		if err != nil {
			panic(err.Error())
		}
	}
	return tmpFile.Name()
}

func SqlToGo(sqlFileName, outputPath string) {
	_ = gen.Clean()

	fmt.Println(aurora.BgRed(fmt.Sprintf("sql input directory------------------->%s", sqlFileName)))

	g, err := gen.NewDefaultGenerator(outputPath, &config.Config{
		NamingFormat: namingFormat,
	})
	if err != nil {
		panic(err)
	}
	err = g.StartFromDDL(sqlFileName, false, dataBase)
	if err != nil {
		panic(err)
	}
}

const (
	namingFormat = "GoZero" //命名规则,驼峰命名
	dataBase     = "slots2021"
) //命名规则,驼峰命名

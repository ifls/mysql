package mysql

import (
	"context"
	"database/sql/driver"
	"log"
	"testing"
)

func TestMysqlConn_Begin(t *testing.T) {
	dsn := "root@tcp(127.0.0.1)/simple?timeout=1s"
	md := &MySQLDriver{}
	conn, err := md.Open(dsn)
	if err != nil {
		t.Fatalf("error connecting: %s", err.Error())
	}
	defer conn.Close()
	my := conn.(*mysqlConn)

	err = my.Ping(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	trs, err := my.Query("select * from runoob_tbl2", nil)
	if err != nil {
		t.Fatal(err)
	}
	log.Printf("%#v\n", trs)
	result := make([]driver.Value, len(trs.Columns())-2)
	for trs.Next(result) == nil {
		for _, res := range result {
			log.Printf("%s\t", res)
		}
		log.Println()
	}
	log.Printf("%+v\n", trs.Columns())

	res, err := my.Exec(`insert into runoob_tbl2 (runoob_title, runoob_author) VALUES("学习 MySQL", "菜鸟教程");`, nil)
	if err != nil {
		t.Fatalf("%+v", err)
	}

	log.Printf("%+v\n", res)

	tx, err := my.Begin()
	if err != nil {
		t.Fatal(err)
	}

	trstx, err := my.Query("select * from runoob_tbl2", nil)
	if err != nil {
		t.Fatal(err)
	}
	for trstx.Next(make([]driver.Value, 3)) == nil {

	}
	log.Printf("%+v\n", trstx.Columns())

	restx, err := my.Exec(`insert into runoob_tbl2 (runoob_title, runoob_author) VALUES("学习 MySQL", "菜鸟教程");`, nil)
	if err != nil {
		t.Fatalf("%+v", err)
	}

	log.Printf("%+v\n", restx)

	err = tx.Commit()
	if err != nil {
		t.Fatal(err)
	}

	stmt, err := conn.Prepare("SELECT runoob_title FROM runoob_tbl2 WHERE runoob_id=?")
	if err != nil {
		t.Fatalf("error preparing statement: %s", err.Error())
	}
	defer stmt.Close()

	val, _ := converter{}.ConvertValue(string("1"))

	rows, err := stmt.Query([]driver.Value{val})
	if err != nil {
		t.Fatalf("error executing statement: %s", err.Error())
	}

	log.Println("prepare select", rows.Columns())
	for rows.Next(make([]driver.Value, 1)) == nil {

	}

	stmt, err = conn.Prepare(`insert into runoob_tbl2 (runoob_title, runoob_author, sub) VALUES("学习 MySQL", "菜鸟教程", ?);`)
	if err != nil {
		t.Fatalf("error preparing statement: %s", err.Error())
	}
	defer stmt.Close()

	// http://huanyouchen.github.io/2018/05/22/mysql-error-1406-Data-too-long-for-column/
	// SET @@global.sql_mode= ''; 让 mysql 不使用严格模式, 解决 超出长度的问题
	// set global max_allowed_packet = 524*1024*1024 修改服务器端最大包大小
	val, _ = converter{}.ConvertValue(string(make([]byte, 0xfefffe)))
	_, err = stmt.Exec([]driver.Value{val})
	if err != nil {
		t.Fatal(err)
	}
}

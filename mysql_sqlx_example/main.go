package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type User struct {
	Id   int64          `db:"id"`
	Name sql.NullString `db:"name"`
	Age  int            `db:"age"`
}

var DB *sqlx.DB

func initDB() error {
	var err error
	dsn := "root:123456@(mysql)/golang"
	DB, err = sqlx.Open("mysql", dsn)
	if err != nil {
		return err
	}

	//设置最大连接数
	DB.SetMaxOpenConns(100)
	//设置空闲时连接数
	DB.SetMaxIdleConns(16)
	return nil
}

//单条记录查询
func testQuery() {
	var user User
	sqlstr := "select `id`, `name`, `age` from `user` where id = ?"
	err := DB.Get(&user, sqlstr, 1)
	if err != nil {
		fmt.Printf("get data failed, err :%v\n", err)
		return
	}
	fmt.Printf("user id = 1: %v\n", user)
}

//多条记录查询
func testQueryMulti() {
	var user []User
	sqlstr := "select `id`, `name`, `age` from `user` where id > ?"
	err := DB.Select(&user, sqlstr, 0)
	if err != nil {
		fmt.Printf("select data failed, err :%v\n", err)
		return
	}
	fmt.Printf("user id > 0: %v\n", user)
}

//插入记录
func testInsert() {
	sqlstr := "insert into `user` (`name`, `age`) values (?, ?)"
	count, err := DB.Exec(sqlstr, "bob", 33)
	if err != nil {
		fmt.Printf("insert data failed, err :%v\n", err)
		return
	}
	id, err := count.LastInsertId()
	if err != nil {
		fmt.Printf("get insert id failed, err :%v\n", err)
		return
	}
	fmt.Printf("insert success, insert id : %d\n", id)
}

//更新记录
func testUpdate() {
	sqlstr := "update `user` set `name` = ? where id = ?"
	count, err := DB.Exec(sqlstr, "steven", 8)
	if err != nil {
		fmt.Printf("update data failed, err :%v\n", err)
		return
	}
	affect, err := count.RowsAffected()
	if err != nil {
		fmt.Printf("get affected rows failed, err :%v\n", err)
		return
	}
	fmt.Printf("update success, afftected rows : %d\n", affect)
}

//更新记录
func testDel() {
	sqlstr := "delete from `user` where id = ?"
	count, err := DB.Exec(sqlstr, 7)
	if err != nil {
		fmt.Printf("delete data failed, err :%v\n", err)
		return
	}
	affect, err := count.RowsAffected()
	if err != nil {
		fmt.Printf("get affected rows failed, err :%v\n", err)
		return
	}
	fmt.Printf("delete success, afftected rows : %d\n", affect)
}

//预处理 单条记录查询
func testPrepareGet() {
	var user User
	sqlstr := "select `id`, `name`, `age` from `user` where id = ?"
	stmt, err := DB.Prepare(sqlstr)
	if err != nil {
		fmt.Printf("prepare failed, err :%v\n", err)
		return
	}
	defer func() {
		if stmt != nil {
			stmt.Close()
		}
	}()
	row := stmt.QueryRow(1)
	err = row.Scan(&user.Id, &user.Name, &user.Age)
	if err != nil {
		fmt.Printf("scan failed, err :%v\n", err)
		return
	}
	fmt.Printf("user id = 1: %v\n", user)
}

//事务
func testTrans() {
	//模拟转账   即A扣钱 B加钱   此处使用年龄代替钱
	conn, err := DB.Begin()
	if err != nil {
		fmt.Printf("transaction failed, err:%v\n", err)
		return
	}

	sqlstr := "update `user` set `age` = `age` - ? where id = ?"
	_, err = conn.Exec(sqlstr, 5, 1)
	if err != nil {
		fmt.Printf("transaction first sql failed, err:%v\n", err)
		conn.Rollback()
		return
	}
	sqlstr = "update `user` set `age` = `age` + ? where id = ?"
	_, err = conn.Exec(sqlstr, 5, 2)
	if err != nil {
		fmt.Printf(" transaction second sql failed, err:%v\n", err)
		conn.Rollback()
		return
	}
	err = conn.Commit()
	if err != nil {
		conn.Rollback()
		fmt.Printf("transaction commit failed, err : %v \n", err)
		return
	}
}

func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("db init failed,err:%v\n", err)
		return
	}

	//单条查询
	//testQuery()
	//多条记录查询
	//testQueryMulti()
	//插入记录
	//testInsert()
	//更新记录
	//testUpdate()
	//删除记录
	//testDel()

	/*预处理 同mysql_sample*/
	//预处理获取一条
	//testPrepareGet()

	/*事务 同mysql_sample*/
	//testTrans()

}

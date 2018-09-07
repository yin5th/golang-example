package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql" //没有直接引用使用_   必须
)

var DB *sql.DB

func initDB() error {
	var err error
	dsn := "root:123456@tcp(mysql)/golang"
	//上面需单独声明err 否则DB会报错：定义的变量未使用
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	//设置最大连接数
	DB.SetMaxOpenConns(100)
	//设置空闲时连接数
	DB.SetMaxIdleConns(16)
	return nil
}

type User struct {
	Id   int64          `db:"id"`
	Name sql.NullString `db:"name"` //字段允许为空的时候需要使用
	Age  int            `db:"age"`
}

//查询一条记录
func testQueryRow() {
	var user User
	// 使用? 防止sql注入
	sqlStr := "select * from `user` where id=?"
	row := DB.QueryRow(sqlStr, 1)
	err := row.Scan(&user.Id, &user.Name, &user.Age)
	if err != nil {
		fmt.Printf("scan failed, err:%v\n", err)
		return
	}
	fmt.Printf("Id: %d  Name: %v Age:%d\n", user.Id, user.Name, user.Age)
}

//查询多条记录
func testQueryMultiRows() {
	var user User
	// 使用? 防止sql注入
	sqlStr := "select * from `user` where `id` > ?"
	rows, err := DB.Query(sqlStr, 0)
	//必须close 否则连接数量将很快使用完
	defer func() {
		if rows != nil {
			rows.Close()
		}
	}()

	if err != nil {
		fmt.Printf("db query failed, err:%v\n", err)
		return
	}
	for rows.Next() {
		err = rows.Scan(&user.Id, &user.Name, &user.Age)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return
		}
		fmt.Printf("Id:%d  Name: %v   Age:%d\n", user.Id, user.Name, user.Age)
	}

}

//插入数据
func testInsertData() {
	sqlStr := "insert into `user` (`name`, `age`) values (?, ?)"
	result, err := DB.Exec(sqlStr, "tom", 30)
	if err != nil {
		fmt.Printf("db exec failed, err: %v\n", err)
		return
	}
	id, err := result.LastInsertId()
	if err != nil {
		fmt.Printf("db insert failed, err: %v\n", err)
		return
	}
	fmt.Printf("db insert success, id is %d\n", id)
}

//修改数据
func testUpdateData() {
	sqlStr := "update `user` set name = ? where id = ? "
	result, err := DB.Exec(sqlStr, "jim", 2)
	if err != nil {
		fmt.Printf("db exec failed, err: %v\n", err)
		return
	}
	affectNum, err := result.RowsAffected()
	if err != nil {
		fmt.Printf("db update failed, err: %v\n", err)
		return
	}
	fmt.Printf("db update success, affected rows count: %d\n", affectNum)
}

func testDelData() {
	sqlStr := "delete  from `user` where id = ?"
	result, err := DB.Exec(sqlStr, 3)
	if err != nil {
		fmt.Printf("db exec failed, err: %v\n", err)
		return
	}
	affectNum, err := result.RowsAffected()
	if err != nil {
		fmt.Printf("db delete failed, err: %v\n", err)
		return
	}
	fmt.Printf("db delete success, affected rows count: %d\n", affectNum)
}

//预查询 查询一条记录
func testPrepareQueryRow() {
	var user User
	// 使用? 防止sql注入
	sqlStr := "select `id`, `name`, `age` from `user` where id=?"
	stmt, err := DB.Prepare(sqlStr)
	//stmt必须关闭
	defer func() {
		if stmt != nil {
			stmt.Close()
		}
	}()
	if err != nil {
		fmt.Printf("db prepare failed, err:%v\n", err)
		return
	}
	//查询ID=1的记录
	row := stmt.QueryRow(1)
	err = row.Scan(&user.Id, &user.Name, &user.Age)
	if err != nil {
		fmt.Printf("scan failed, err:%v\n", err)
		return
	}
	fmt.Printf("Id: %d  Name: %v Age:%d\n", user.Id, user.Name, user.Age)
}

//预查询 查询多条记录
func testPrepareQueryMultiRows() {
	var user User
	// 使用? 防止sql注入
	sqlStr := "select `id`, `name`, `age` from `user` where id>?"
	stmt, err := DB.Prepare(sqlStr)
	//stmt必须关闭
	defer func() {
		if stmt != nil {
			stmt.Close()
		}
	}()
	if err != nil {
		fmt.Printf("db prepare failed, err:%v\n", err)
		return
	}
	//查询ID>0的记录
	rows, err := stmt.Query(0)
	if err != nil {
		fmt.Printf("db query failed, err:%v\n", err)
		return
	}
	//rows必须关闭
	defer func() {
		if rows != nil {
			rows.Close()
		}
	}()
	for rows.Next() {
		err = rows.Scan(&user.Id, &user.Name, &user.Age)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return
		}
		fmt.Printf("Id: %d  Name: %v Age:%d\n", user.Id, user.Name, user.Age)
	}

}

//预处理 插入记录
func testPrepareInsertData() {
	sqlstr := "insert into `user` (`name`, `age`) values (?, ?)"
	stmt, err := DB.Prepare(sqlstr)
	if err != nil {
		fmt.Printf("prepare failed, err: %v\n", err)
		return
	}
	defer func() {
		if stmt != nil {
			stmt.Close()
		}
	}()

	res, err := stmt.Exec("bob", 35)
	if err != nil {
		fmt.Printf("insert failed, err: %v\n", err)
		return
	}

	insertID, err := res.LastInsertId()
	if err != nil {
		fmt.Printf("get insert ID failed, err: %v\n", err)
		return
	}
	fmt.Printf("insert success, id is %d\n", insertID)
}

//预处理 更新记录
func testPrepareUpdateData() {
	sqlstr := "update `user` set age = ? where id = ?"
	stmt, err := DB.Prepare(sqlstr)
	if err != nil {
		fmt.Printf("prepare failed, err: %v\n", err)
		return
	}
	defer func() {
		if stmt != nil {
			stmt.Close()
		}
	}()

	res, err := stmt.Exec(28, 1)
	if err != nil {
		fmt.Printf("update failed, err: %v\n", err)
		return
	}

	affected, err := res.RowsAffected()
	if err != nil {
		fmt.Printf("get insert ID failed, err: %v\n", err)
		return
	}
	fmt.Printf("update success,affected rows: %d\n", affected)
}

//预处理 删除记录
func testPrepareDelData() {
	sqlstr := "delete from `user` where id = ?"
	stmt, err := DB.Prepare(sqlstr)
	if err != nil {
		fmt.Printf("prepare failed, err: %v\n", err)
		return
	}
	defer func() {
		if stmt != nil {
			stmt.Close()
		}
	}()

	res, err := stmt.Exec(4)
	if err != nil {
		fmt.Printf("delete failed, err: %v\n", err)
		return
	}

	affected, err := res.RowsAffected()
	if err != nil {
		fmt.Printf("get deleted rows number failed, err: %v\n", err)
		return
	}
	fmt.Printf("delete success,deleted rows: %d\n", affected)
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
	_, err = conn.Exec(sqlstr, 20, 1)
	if err != nil {
		fmt.Printf("transaction first sql failed, err:%v\n", err)
		conn.Rollback()
		return
	}
	sqlstr = "update `user` set `age` = `age` + ? where id = ?"
	_, err = conn.Exec(sqlstr, 20, 2)
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
		fmt.Printf("init db failed, err:%v\n", err)
		return
	}
	defer DB.Close()

	/*普通sql查询*/
	//查询一条记录
	//testQueryRow()
	//查询多条记录
	//testQueryMultiRows()
	//插入记录
	//testInsertData()
	//修改记录
	//testUpdateData()
	//删除记录
	//testDelData()

	/*预查询*/
	//使用预查询 查询一条记录
	//testPrepareQueryRow()
	//使用预查询 查询多条记录
	//testPrepareQueryMultiRows()
	//使用预查询 插入记录
	//testPrepareInsertData()
	//使用预查询 修改记录
	//testPrepareUpdateData()
	//使用预查询 删除记录
	//testPrepareDelData()

	/*事务*/
	testTrans()
}

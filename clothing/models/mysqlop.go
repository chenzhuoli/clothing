package models

import (
	"database/sql"
	"fmt"
	"log"
	"time"
	"strings"
	"clothing/common"

	_ "github.com/go-sql-driver/mysql"
)

type ClothingData struct {
	ID int
	Title string
	Thumbnail string
	Likes int
	UserID string
	CreateTime time.Time
	UserName string
	UserHead string
}

func showMysqlVersion() {
	db, err := sql.Open("mysql", "debian-sys-maint:dzub1lB8YJaluX9N@tcp(127.0.0.1:3306)/clothing?charset=utf8mb4")
	db.Ping()
	defer db.Close()

	if err != nil {
		fmt.Println("数据库连接失败！")
		log.Fatalln(err)
	}

	var version string

	err2 := db.QueryRow("SELECT VERSION()").Scan(&version)

	if err2 != nil {
		log.Fatal(err2)
	}

	fmt.Println(version)
}

func createTable() {
	db, err := sql.Open("mysql", "debian-sys-maint:dzub1lB8YJaluX9N@tcp(127.0.0.1:3306)/clothing")
	db.Ping()
	defer db.Close()

	if err != nil {
		fmt.Println("connect DB error !")
		log.Fatalln(err)
	}

	_, err2 := db.Exec("CREATE TABLE user(id INT NOT NULL , name VARCHAR(20), PRIMARY KEY(ID));")
	if err2 != nil {
		log.Fatal(err2)
	}

	fmt.Println("successfully create table")
}

func insertItem() {
	db, err := sql.Open("mysql", "debian-sys-maint:dzub1lB8YJaluX9N@tcp(127.0.0.1:3306)/clothing?charset=utf8mb4")
	db.Ping()
	defer db.Close()

	if err != nil {
		fmt.Println("connect DB error !")
		log.Fatalln(err)
	}

	_, err2 := db.Query("INSERT INTO user VALUES(1, 'zhangsan')")
	if err2 != nil {
		log.Fatal(err2)
	}

	fmt.Println("successfully insert item")
}

func deleteItem() {
	db, err := sql.Open("mysql", "debian-sys-maint:dzub1lB8YJaluX9N@tcp(127.0.0.1:3306)/clothing?charset=utf8mb4")
	db.Ping()
	defer db.Close()

	if err != nil {
		fmt.Println("connect DB error !")
		log.Fatalln(err)
	}

	sql := "DELETE FROM user WHERE id = 1"
	res, err2 := db.Exec(sql)

	if err2 != nil {
		panic(err2.Error())
	}

	affectedRows, err := res.RowsAffected()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("delete item success, statement affected %d rows\n", affectedRows)
}

func alterItem() {
	db, err := sql.Open("mysql", "debian-sys-maint:dzub1lB8YJaluX9N@tcp(127.0.0.1:3306)/clothing?charset=utf8mb4")
	db.Ping()
	defer db.Close()

	if err != nil {
		fmt.Println("connect DB error !")
		log.Fatalln(err)
	}

	sql := "update user set name = ? WHERE id = ?"
	res, err2 := db.Exec(sql, "lisi", 1)

	if err2 != nil {
		panic(err2.Error())
	}

	affectedRows, err := res.RowsAffected()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("alter item success, statement affected %d rows\n", affectedRows)
}

func queryItem() {
	db, err := sql.Open("mysql", "debian-sys-maint:dzub1lB8YJaluX9N@tcp(127.0.0.1:3306)/clothing?charset=utf8mb4")
	db.Ping()
	defer db.Close()

	if err != nil {
		fmt.Println("connect DB error !")
		log.Fatalln(err)
	}

	 var mid int = 1

	 result, err2 := db.Query("SELECT * FROM user WHERE id = ?", mid)
	if err2 != nil {
		log.Fatal(err2)
	}

	for result.Next() {

		var id int
		var name string

		err = result.Scan(&id, &name)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("id: %d, name: %s\n", id, name)
	}
}

func dropTable() {
	db, err := sql.Open("mysql", "debian-sys-maint:dzub1lB8YJaluX9N@tcp(127.0.0.1:3306)/clothing")
	db.Ping()
	defer db.Close()

	if err != nil {
		fmt.Println("connect DB error !")
		log.Fatalln(err)
	}

	_, err2 := db.Exec("DROP TABLE user;")
	if err2 != nil {
		log.Fatal(err2)
	}

	fmt.Println("successfully drop table")
}

func QueryClothingRecommendItems() ([]ClothingData, error) {
	db, err := sql.Open("mysql", "debian-sys-maint:dzub1lB8YJaluX9N@tcp(127.0.0.1:3306)/clothing?charset=utf8mb4&parseTime=true")
	db.Ping()
	defer db.Close()

	if err != nil {
		fmt.Println("connect DB error !")
		log.Fatalln(err)
		return nil, err
	}

	 var user_id string = "lily"

	result, err2 := db.Query("SELECT clothing_recommend.id, clothing_recommend.title, clothing_recommend.thumbnail, clothing_recommend.likes, clothing_recommend.userid, clothing_recommend.create_time, users.name, users.head FROM clothing_recommend JOIN users ON clothing_recommend.userid = users.userid WHERE users.userid = ? ORDER BY create_time DESC", user_id)
	if err2 != nil {
		log.Fatal(err2)
		return nil, err2
	}


	var ClothingRecommendData []ClothingData

	for result.Next() {

		var item ClothingData

		err = result.Scan(&item.ID, &item.Title, &item.Thumbnail, &item.Likes, &item.UserID, &item.CreateTime, &item.UserName, &item.UserHead)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(item)
		ClothingRecommendData = append(ClothingRecommendData, item)
	}
	return ClothingRecommendData, nil
}

func InsertDesignerCertificationData(cert common.DesignerCertificationStruct) error {
	db, err := sql.Open("mysql", "debian-sys-maint:dzub1lB8YJaluX9N@tcp(127.0.0.1:3306)/clothing?charset=utf8mb4")
	db.Ping()
	defer db.Close()

	if err != nil {
		fmt.Println("connect DB error !")
		log.Fatalln(err)
		return err
	}
	works := strings.Join(cert.ClothingWorks, ",")	
	sql_str := fmt.Sprintf("INSERT INTO designers (userid,phone_number,email,clothing_works,create_time) VALUES('%s','%s','%s','%s',%s)", cert.UserID,cert.PhoneNumber,cert.Email,works,"NOW()")
	fmt.Println(sql_str)

	_, err2 := db.Query(sql_str)
	if err2 != nil {
		log.Fatal(err2)
		return err2
	}

	fmt.Println("successfully insert item")
	return nil
}


func main() {
	clothingInfo, err := QueryClothingRecommendItems() 
	fmt.Println(clothingInfo, err)
}

/*
func main() {
	showMysqlVersion()
	createTable()
	insertItem()
	queryItem()
	alterItem()
	queryItem()
	deleteItem()
	dropTable()
}
*/

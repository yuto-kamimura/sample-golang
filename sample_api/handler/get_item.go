package handler

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetItem(c echo.Context) error {

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true&loc=Local", "mariadb", "mariadb", "db", "mariadb")

	// データベースに接続
	db, err := sql.Open("mysql", "mariadb:mariadb@tcp(127.0.0.1:3307)/mariadb")
	if err != nil {
		panic(err)
	}

	// データベースにPingを送信して接続を確認
	err := db.Ping()
	if err != nil {
		log.Panicln(err)
	}

	// データベースからデータを取得するクエリ
	rows, err := db.Query("SELECT Name FROM Item WEHE Name = 'smart phone'")
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()

	return c.String(http.StatusOK, "")

}

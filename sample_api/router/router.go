package router

import (
	"net/http"
	"sampleApi/handler"

	"github.com/labstack/echo/v4"
)

// 新しいリクエスト先が必要な場合はここに追加する
func Route(e *echo.Echo) {
	/*
		httpメソッドは主にGET,POST,PUT,DELETEが使用されます。
		GET: DBからデータを取得する
		POST: DBのデータを更新する
		PUT: DBにデータを新規作成する
		DELTE: DBからデータを削除する

		golangのechoではe.〇〇の〇〇にGET～DELETEを入れることで
		クライアント（ブラウザからの操作）からDBに対して
		どのような処理を実行させたいか決定します。
	*/
	e.GET("", func(c echo.Context) error {
		return c.String(http.StatusOK, "Accessible")
	})

	// クライアントからhttp://{バックエンドのURL}:{バックエンドのポート}/helloが
	// リクエストされたときに、handler/hello.goのHello関数を実行する
	e.GET("/hello", handler.Hello)

	// 課題１ handlerディレクトリにget_item.goを作成し、ブラウザに文字列"smart phone"を表示させる

	// 課題２ handlerディレクトリにcreate_item.goを作成し、ターミナルから以下のコマンドを実行し、itemを登録する

}

# echo課題

## Challenge1
echoサーバーを実行してみましょう
1. vscodeのデバッグでLaunch Packageを実行しましょう。

＊APIの内容を変更する度に、echoサーバーの再起動を実行しましょう

## Challenge2
PUTメソッドを実行してみましょう
1. echoサーバーを起動したら、postmanのPUTメソッドを実行しましょう。(postmanがない場合はインストールしましょう)
    1. httpメソッドをPUTに設定して、`localhost:1324/items`を指定しましょう。
    1. postmanのbodyに01_challenge.jsonの内容を貼り付けましょう。bodyのタイプはjsonに変更しましょう。
1. ブラウザのURLにlocalhost:8081を入力し、curlコマンドから入力したデータを確認しましょう\
ログイン情報は以下 \
user: mariadb\
password: mariadb\
databasename: mariadb\


## Challenge3
GETメソッドで入力したデータを取得しましょう
1. postmanからGETメソッドを実行しましょう
    1. httpメソッドをGETに設定して、`localhost:1324/items`を指定しましょう。
    1. postmanのbodyに02_challenge.jsonの内容を貼り付けましょう。
1. 01_challenge.jsonの内容を取得できたか確認しましょう。


## Challenge4
DELETEで入力したデータを削除しましょう。
1. postmanからDELETEメソッドを実行しましょう
    1. httpメソッドをDELETEに設定して、`localhost:1423/items`を指定しましょう。
    1. postmanのbodyに02_challenge.jsonの内容を貼り付けましょう。
1. 02_challenge.jsonの内容が削除されたか確認するため、Challenge3をもう一度実行して、GETメソッドで取得できなことを確認しましょう。


## Challenge5
新しいGETリクエストのハンドラーを実装しましょう
1. localhost:1324/calculationを実行する関数`GetCalculationNum`をhandlerに作成しましょう
    1. 関数GetCalculationNum内では、for文を利用して1～10の値を足し合わせる処理を実行しましょう。
    1. ブラウザから`localhost:1324/calculation-num`をURLに指定されたとき、画面に計算結果を表示させるように、router.goを修正しましょう。


## Challenge6
新しいテーブルを作成しましょう。
1. entityに新しくuser.goを作成しましょう。
1. entity/item.goの内容を参考に、構造体Userを作成しましょう。
    1. Userの要素はID(int), Name(string), Password(string), Email(string), UpdatedAt(time.Time)としましょう。
1. main.goのinitDB()のdbInstance.AutoMigrate{}に新しく作成したUserを追加しましょう。
1. echoサーバーを起動しましょう。自動的にUserテーブルが作成されます。 
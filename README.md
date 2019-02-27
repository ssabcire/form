# features
GoでformをつかってDBとやりとりすることを考えてます。  
jsonでデータをやりとりできるといいなー。  

# form  
GoでformをつかってDBとやりとりするためのレポジトリです。練習用。  
いまのところ、クライアントでフォームに名前を入力すると、いろいろあってMySQLに追加されます。  

# Dependency
go 1.11.1

# Usage
mySQLのドライバをダウンロードします。  
```
go get https://github.com/go-sql-driver/mysql.git
```

テーブルのセットアップをします。  
table.sql

ビルドします。
```
go build
```

バイナリを実行します。
```
./form
```


ブラウザでlocalhostを起動します。
localhost:8080

すると、やりとりができると思います。

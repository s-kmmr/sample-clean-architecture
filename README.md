# sample-clean-architecture

```
何かの会員情報の管理に例えて、クリーンアーキテクチャの円を表現してみました。
今のところ、会員情報をすべて取得する機能だけ作ってます。

・使い方

[mysqlコンテナ起動（dockerが入っていること）]
cd docker/
docker-compose up -d
※サンプル用の初期データが入ります

[ルートディレクトリでサーバ起動]
go run main.go

※ginでサーバが起動します


```


***
### - Frameworks&Drivers
 - infrasturecture/
   - client
   - injector
   - persistence
   - router
### - Interface Adapters
- interfaces/
  - controllers
  - gateway/
    - database
- domain/
  - repository
### - Application Business Rules
- usecases/
### - Enterprise Business Rules
- domain/
  - model



***

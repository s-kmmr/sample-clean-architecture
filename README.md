# sample-clean-architecture

```
何かの会員情報の管理に例えて、クリーンアーキテクチャの円を表現してみました。
```

 ```
・使い方

[mysqlコンテナ起動（dockerが入っていること）]
cd docker/
docker-compose up -d
※サンプル用の初期データが入ります

[ルートディレクトリでサーバ起動]
go run main.go

※ginでサーバが起動します


```

- 今のところ作った機能
  - 会員情報をすべて取得する機能
  - 会員情報を新規登録する機能


***
### - Frameworks&Drivers
 - infrastructure/
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
 - 補足
   -  トランザクション管理はTransactionRepositoryで行うようにしてます
# caspian-serverside
caspianのバックエンドプロジェクト

GolangでAPIを開発

## 構成

- Golang
- MySQL
- Nginx

APIへのアクセスはNginxコンテナからGoのコンテナにフォワード

## 開発環境構築
```bash
git clone https://github.com/bryutus/caspian-serverside.git

docker-compose build

docker-compose up -d
```

## GoのAPIサーバ起動

```bash
# Goのコンテナに入る
docker-compose exec app /bin/bash

# migrate実行
go run /go/src/github.com/bryutus/caspian-serverside/app/migrate/migrate.go

# 外部APIよりデータ取得
go run /go/src/github.com/bryutus/caspian-serverside/app/main.go

# APIサーバ起動
go run /go/src/github.com/bryutus/caspian-serverside/app/server.go
```

## APIの疎通確認
```bash
curl http://localhost:8080/api/albums
```

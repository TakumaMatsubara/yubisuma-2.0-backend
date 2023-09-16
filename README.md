# 概要

<p>
指スマ2.0のバックエンドプロジェクトです。
</p>

## API 仕様

API 仕様は SwaggerUI を利用して閲覧します。

```
$ docker compose up swagger-ui
```

を実行することでローカルの Docker 上に SwaggerUI サーバが起動します。<br>
<br>
SwaggerUI サーバ起動後以下の URL から SwaggerUI へアクセスすることができます。

SwaggerUI: <http://localhost:8000/> <br>
定義ファイル: `./api-document.yaml`<br>

# 事前準備

プロジェクトを立ち上げるには

```
$ docker compose up -d
```

をして、

```
$ docker exec -it <container id(golang-appという名前)> bash
```

コンテナに入ったら

```
$ go run main.go
```

を実行してくれれば ok。<br>
<http://localhost:8080/>に接続してみてください。

### MySQL

MySQL はリレーショナルデータベースの 1 つです。

```
$ docker compose up mysql
```

を実行することでローカルの Docker 上に MySQL サーバが起動します。<br>
```
port:3307
```
で mysql に接続できます。

<br>
初回起動時に db/init ディレクトリ内の DDL, DML を読み込みデータベースの初期化を行います。<br>
(DDL(DataDefinitionLanguage)とはデータベースの構造や構成を定義するための SQL 文)<br>
(DML(DataManipulationLanguage)とはデータの管理・操作を定義するための SQL 文)

#### phpMyAdmin

MySQL データベースのテーブルやレコードの閲覧、変更するためのツールとして phpMyAdmin を用意しています。

```
$ docker compose up phpmyadmin
```

を実行することでローカルの Docker 上に phpMyAdmin サーバが起動します。<br>
phpMyAdmin サーバ起動後以下の URL からアクセスすることができます。

phpMyAdmin: <http://localhost:4000/>

# Go言語製WAF GinでWebアプリ

```
go get github.com/go-sql-driver/mysql
go get github.com/go-xorm/xorm
```

## MySQL DB準備

```
CREATE DATABASE bookshelf;
use bookshelf;
 
CREATE TABLE user (
id bigint(20) unsigned NOT NULL AUTO_INCREMENT,
nickname varchar(190) NOT NULL COMMENT 'ニックネーム',
PRIMARY KEY (id)
) ENGINE=InnoDB AUTO_INCREMENT=1002 DEFAULT CHARSET=utf8mb4 COMMENT='テスト';
```

## MySQL パスワード設定

```
CREATE USER hoge IDENTIFIED BY 'passwordPASSWORD+123';
GRANT ALL ON bookshelf.* to test@localhost IDENTIFIED BY 'passwordPASSWORD+123';
INSERT INTO USER VALUES(1, 'aho')
```

## 起動

```
go run main.go
```

[![golangci-lint](https://github.com/firdavsich/surl/actions/workflows/lint.yml/badge.svg)](https://github.com/firdavsich/surl/actions/workflows/lint.yml)

# surl
URL shortener api


### DB init

```
CREATE USER surl WITH ENCRYPTED PASSWORD 'surl';
CREATE DATABASE surl_db;
GRANT ALL PRIVILEGES ON DATABASE surl_db TO surl;
CREATE TABLE surl (
    id SERIAL PRIMARY KEY,
    url VARCHAR(2048) NOT NULL,
    hash VARCHAR(10) NOT NULL
);
```
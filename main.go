package main

import (
	"fmt"
	"database/sql"
	_ "github.com/lib/pq"
)

const (
	baseURL = "http://www.surl.com/"
)

func main(){
// open connect to postgresql
db, err:= sql.Open("postgres", "postgres://surl:surl@localhost/surl_db?sslmode=disable")
if err != nil {
	panic(err)
}

	fmt.Println(addUrl(db, "http://ya.ru"))
	fmt.Println(addUrl(db, "http://mail.ru"))
	fmt.Println(getUrl(db, "eZqHtqzXMp"))
	fmt.Println(getUrl(db, "1234567890"))
	fmt.Println(getUrl(db, "FIZKwhFXeE"))


}

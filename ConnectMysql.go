package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

const (
	host    string ="localhost"
	user    string ="wqq"
	passwd  string ="123456"
	charset string ="utf8"
	db      string ="mytest"
	port    int    =3306
)

type PrivateData struct {
	Id int
	Name string
}


func initDB()*sql.DB{
	driver:=fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s",user,passwd,host,port,db,charset)
	db,err:=sql.Open("mysql",driver)
	if err!=nil{
		panic(err)
	}
	if err:=db.Ping();err!=nil{
		fmt.Println("connect mysql fail")
	}
	return db
}


func SelectInfo(db *sql.DB,lang string) []PrivateData {
	myResult:=make([]PrivateData,0)
	rows,err:= db.Query(lang)
	if err !=nil{
		return myResult
	}
	for rows.Next(){
		var singData PrivateData
		err = rows.Scan(&singData.Id,&singData.Name)
		if err ==nil{
			myResult=append(myResult,singData)
			}else{
				continue
		}

	}
	return myResult
}

func main(){
	db:=initDB()
	queryLang:=`select id,name from infor`
	totalInfos := SelectInfo(db,queryLang)
	for i,infor:=range totalInfos {
		fmt.Printf("%d:%+v\n",i,infor)
	}
}
package main

import (
	"fmt"
)

type User struct {
	ID int64
	Email string
	AuthToken string
	UserName string
	PasswordDigest string
}

func main() {
	data := &Hoge{fuga: "fuga"}
	fmt.Println(hoge(data))
}

type Hoge struct {
	fuga string
}
func hoge(data interface{}) bool {
	fmt.Println(data.(*Hoge))
	fmt.Println(data.(*Hoge))
	// 権限のあるユーザーかどうか判別
	v, ok := data.(*Hoge)
	if ok {
		fmt.Println(v)
		return true
	}
	if "True" == "True" {
		return false
	}
	// falseの場合は403を返却
	return false
}
package test1

import "github.com/zhengweiye/test2/test2"

type Test1 struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func GetTest1() Test1 {
	test2 := test2.GetTest2()

	return Test1{
		Id:   test2.Id + 10,
		Name: "hello " + test2.Name,
	}
}

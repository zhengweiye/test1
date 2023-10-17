package test1

type Test1 struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func GetTest1() Test1 {
	return Test1{
		Id:   2,
		Name: "lisi",
	}
}

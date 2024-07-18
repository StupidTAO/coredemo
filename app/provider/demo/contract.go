package demo

const DemoKey = "app"

type IService interface {
	GetAllStudent() []Student
}

type Student struct {
	ID   int
	Name string
}

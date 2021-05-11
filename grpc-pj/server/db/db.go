package db

//var 会初始化吗
var db map[int]*Person

type Person struct {
	Id int
	Name string 
	Age int
}

func NewPerpson(id int, name string, age int) *Person {
	return &Person{Id: id, Name: name, Age: age}
}
func(p *Person) GetId() int {
	return p.Id
}

func(p *Person) GetName() string {
	return p.Name
}
func(p *Person) GetAge() int {
	return p.Age
}


func Get(id int ) (*Person,bool){
	if p,ok := db[id];ok{
		return p,true
	}
	return nil,false
}
func Add(p *Person) bool{
	if _,ok := db[p.Id];ok{
		return false
	}
	db[p.Id] = p
	return true
}

func init()  {
	db = make(map[int]*Person)
}



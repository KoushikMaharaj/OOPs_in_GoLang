// if you want to achieve encapsulation in GoLang
// create struct in different package and create field as private
// to set those fields use setter function as shown below line no. 18,23
// to get value of private fields use getter as shown at line 29
package structs

type Person struct {
	//Name string //public field
	name string //private field
	age  int
}

//constructor
func NewPerson(name string, age int) Person {
	return Person{name: name, age: age}
}

//reciver function
func (p Person) Walk() string {
	return p.name + " likes walking"
}

// reciver pointer function
// setter
func (p *Person) SetFirstName(name string) {
	p.name = name
}

func (p *Person) SetAgeName(age int) {
	p.age = age
}

//getter
func (p Person) Name() string {
	return p.name
}

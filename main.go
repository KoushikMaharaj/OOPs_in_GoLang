package main

import (
	//"example/oops/polymorphism"
	//"example/oops/structs"
	"example/oops/composition"
	"fmt"
)

func main() {

	// Encapsulation
	/* p := structs.Person{
		Name: "Tony",
		Age:  40,
	}

	p1 := structs.Person{
		Name: "Banner",
		Age:  40,
	}
	fmt.Println(p, "\n", p1)
	fmt.Println(p.Walk(), "\n", p1.Walk()) */
	/* p := structs.Person{
		Name: "Tony",
	} */
	/* p := structs.Person{}
	fmt.Println(p)
	p.SetFirstName("koushik")
	fmt.Println(p)
	p.SetAgeName(24)
	fmt.Println(p)
	fmt.Println(p.Name()) */

	//Polymorphism

	/* var c polymorphism.Shape = polymorphism.Circle{}
	var s polymorphism.Shape = polymorphism.Square{}
	fmt.Println(c, s)
	c.Render()
	s.Render() */

	car:=composition.NewCar("abc",600,15)
	fmt.Println(car)
	fmt.Println(car.HP())

}

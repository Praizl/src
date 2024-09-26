package main

import "fmt"

//import "fmt"

/*func main() {
	//two ways to create structs

	type cars struct {
		brand string
		year  int
		price float64
	}
	//using field names
	car1 := cars{
		brand: "Ford",
		year:  2005,
		price: 21000,
	}
	car2 := cars{"Honda", 20100, 2021}
	fmt.Println(car1)
	fmt.Println(car2)

	//updating and retrieving structs

	car1.brand = "toyota"
	fmt.Println(car1)

	var car3 cars
	car3.brand = "Renault"
	car3.year = 2008

	fmt.Printf("%#v\n", car3)
	fmt.Println(car1 == car3)

	//anonymous structs &anonymous fields(no alias created)

	rides := struct {
		brandName string
		yearMade  int
		cost      float64
	}{
		brandName: "Lexus",
		yearMade:  2018,
		cost:      21580,
	}

	fmt.Printf("%#v\n", rides)
	fmt.Printf("%v\n", rides.brandName)

	type moto struct {
		string
		int
		float64
	}
	m1 := moto{
		"Mazda",
		2012,
		12999.99,
	}
	fmt.Println(m1)

	//embedded structs
	type contact struct {
		phone   int
		email   string
		address string
	}
	type employee struct {
		name        string
		salary      int
		contactInfo contact
	}

	Aro := employee{
		"Aro 10x",
		9000,
		contact{
			phone:   8063063240,
			email:   "orirepraiz98@gmail.com",
			address: "6,Chinda street,Radio Estate,PH.",
		},
	}
	fmt.Println(Aro, Aro.contactInfo.address)

	fmt.Printf("%+v\n %v\n %#v\n", Aro, Aro, Aro)


}*/

type author struct {
	name  string
	age   int
	title string
}

func (a author) output() {
	fmt.Println("Nerds name:", a.name)
	fmt.Println("age:", a.age)
	fmt.Println("title of book:", a.title)
}

func main() {
	theStruct := author{
		name:  "Lee child",
		age:   64,
		title: "Trip Line",
	}

	theStruct.output()

	type person struct {
		name  string
		score int
	}

	person1 := person{
		"Aro",
		75,
	}
	fmt.Println(person1)
	person2 := &person1
	person1.name = "WTP"
	fmt.Println(person1, person2)

}

package main

//without interfaces

/*type circle struct {
	radius float64
}

type rectangle struct {
	lenght  float64
	breadth float64
}

func (c circle) area() float64 {
	return math.Pi * (math.Pow(c.radius, 2))
}

func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func (r rectangle) area() float64 {
	return r.lenght * r.breadth
}

func (r rectangle) perimeter() float64 {
	return 2 * (r.lenght + r.breadth)
}

/*func circleStats(c circle) {
	fmt.Println("Shape Name:", "Circle")
	fmt.Println("Area:", c.area())
	fmt.Println("Circumference", c.perimeter())
}

func rectangleStats(r rectangle) {
	fmt.Println("Shape Name:", "Rectangle")
	fmt.Println("Area", r.area())
	fmt.Println("Perimeter", r.perimeter())
}

// using interfaces
/*type shape interface {
	area() float64
	perimeter() float64
}

func print(s shape) {
	fmt.Printf("shape:%#v\n", s)
	fmt.Printf("Area:%v\n", s.area())
	fmt.Printf("Perimeter:%v\n", s.perimeter())

}

func (c circle) volume() float64 {
	return 4 / 3 * math.Pi * math.Pow(c.radius, 3)
}

// interface embedding
type shape interface {
	area() float64
}

type object interface {
	volume() float64
}

type geometry interface {
	shape
	object
	getColour() string
}

type cube struct {
	side   float64
	colour string
}

func (c cube) area() float64 {
	return 6 * c.side * c.side
}

func (c cube) volume() float64 {
	return math.Pow(c.side, 3)
}

func measure(g geometry) (float64, float64) {
	a := g.area()
	v := g.volume()
	return a, v
}

func (c cube) getColour() string {
	return c.colour
}

//type emptyInterface interface {}

type person struct {
	info interface{}
}*/

//func main() {
/*c1 := circle{radius: 4}
	r1 := rectangle{lenght: 3, breadth: 4}
	//circleStats(c1)
	//rectangleStats(r1)
	print(c1)
	print(r1)

	var s shape
	fmt.Printf("%T\n", s)

	ball := circle{radius: 2.5}
	s = ball
	print(s)

	s.(circle).volume() //type assertion
	fmt.Println("Volume", ball.volume())
	ball, ok := s.(circle) //assertion with err
	if ok {
		fmt.Printf("%v\n", ball.volume())
	}

	switch value := s.(type) {
	case circle:
		fmt.Printf("%#v is of type circle\n", value)

	case rectangle:
		fmt.Printf("%#v is of type rectangle\n", value)

	}

	c := cube{side: 2}
	a, v := measure(c)
	fmt.Println("Area:", a, "Volume:", v, c.colour)

	var empty interface{}
	fmt.Println(empty)

	empty = "empty"
	fmt.Println(empty)

	empty = []string{"chukwudi", "collins"}
	fmt.Println(empty)

	//using type assertion to get dynamic value
	fmt.Println(len(empty.([]int)))

	you := person{}
	you.info = "balablu" //can take any value
	fmt.Println(you.info)
}*/

package main

/*import (
	"fmt"
	"math"
	"strings"
)

func testFunc() {
	fmt.Println("My first function")
}

func Area(length int, breadth int) {
	fmt.Println(length * breadth)
}

func surface(length, height int, shape string) {
	fmt.Println("the surface area of the", shape, 0.5*float64(length)*float64(height))
}

func f1(a float64) float64 {
	return math.Pow(a, a)
}

func f2(b, c int) (int, int) {
	return (b / c), (b * b * c)
}

func f3(f, g int) (s int) {
	s = f + g
	return s
}

// variadic functions
func variad(a ...int) {
	fmt.Println(a)
}

func variad2(a ...int) {
	a[2] = 69

}

func sumAndProduct(a ...float64) (float64, float64) {
	sum := 0.
	product := 1.

	for _, value := range a {
		sum += value
		product *= value
	}
	return sum, product

}

// function with one variadic param
func info(age int, name ...string) string {
	nameComplete := strings.Join(name, "")
	returnString := fmt.Sprintf("age:%d ,name:%s", age, nameComplete)
	return returnString

}

func fu() {
	fmt.Println("This is Fu")
}

func bara() {
	fmt.Println("This is bara")
}

func fubara() {
	fmt.Println("This is fubara")
}

//anonymous functions

func increment(x int) func() int {
	return func() int {
		x++
		return x
	}
}



func main() {
	testFunc()
	Area(3, 5)
	surface(3, 10, "triangle")
	raise := f1(6.9)
	fmt.Println(raise)

	d, e := f2(6, 9)
	fmt.Println(d, e)

	s := f3(2, 3)
	fmt.Println(s)

	//with param values it
	variad(1, 2, 3, 4, 5, 6)
	variad()
	//passing a slice to a variadic function
	slyce := []int{7, 8, 9}

	variad2(slyce...)
	fmt.Println(slyce)

	//calling function
	sum, product := sumAndProduct(1.0, 2.0, 3.0, 4.0)
	fmt.Println(sum, product)

	info := info(30, "King", "Sunny", "Ade")
	fmt.Println(info)

	//Deferring a function delays its execution LIFO

	defer fu()
	bara()
	fubara()
	fmt.Println("This will execute before the deferred part too")

	//anon func within main
	func(anon string) {
		fmt.Println(anon)
	}("string in anon func")

	incre := increment(5)
	fmt.Println(incre())


}*/

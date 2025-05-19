package main

import (
	"fmt"
	"math"
	"math/rand"
	"strings"
	"time"
)

type Point struct {
	X, Y int
}

func (p Point) Abs() float64 {
	return math.Sqrt(float64(p.X*p.X + p.Y*p.Y))
}

func (p *Point) Scale(factor int) { // Pointer receiver to modify the original struct
	p.X *= factor
	p.Y *= factor
}

func (p *Point) Print() string {
	return fmt.Sprintf("Point(%d, %d)", p.X, p.Y)
}

func ScaleAsFunc(p *Point, factor int) { // Pointer is needed to modify the original struct
	p.X *= factor
	p.Y *= factor
}

type Abser interface {
	Abs() float64
	Print() string
}

type Person struct {
	Name string
	Age  int
}

func (p *Person) Abs() float64 {
	return float64(p.Age)
}

func (p *Person) Print() string {
	if(p == nil) {
		return "<nil>"
	}
	return fmt.Sprintf("Name: %s, Age: %d", p.Name, p.Age)
}

func (p *Point) String() string {
	return fmt.Sprintf("CoolPoint(%d, %d)", p.X, p.Y)
}

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}

func funcWithError() error {
	switch zeroOrOne := rand.Intn(2); zeroOrOne {
	case 0:
		return nil
	case 1:
		return &MyError{
			time.Now(),
			"it didn't work",
		}
	}
	return nil
}

type MyReader struct{}

func (r *MyReader) Read(bytes []byte) (int, error) {
	for i := range bytes {
		bytes[i] = 'A'
	}
	return len(bytes), nil
}

func main() {
	// Methods
	fmt.Println("Methods")

	somePoint := Point{3, 4}
	fmt.Println(somePoint.Abs())

	somePoint.Scale(10)
	fmt.Println(somePoint)

	somePoint2 := Point{3, 4}
	ScaleAsFunc(&somePoint2, 10)
	fmt.Println(somePoint2)

	// Interfaces
	fmt.Println("\nInterfaces")

	somePoint3 := Point{3, 4}
	var a Abser
	a = &somePoint3
	fmt.Println(a.Abs())

	somePerson := Person{"John", 30}
	a = &somePerson
	// a = somePerson // Doesn't compile, because Person doesn't implement Abser
	fmt.Println(a.Abs())

	// Handling nil pointers
	fmt.Println("\nHandling nil pointers")
	var p *Person = nil
	a = p
	fmt.Println(a.Print()) // This will not panic

	//a = nil
	//fmt.Println(a.Print()) // This will panic if not handled

	// Type assertion
	fmt.Println("\nType assertion")
	var unknown interface{} = "Hello, World!"
	// var unknown interface{} = 10
	str, strOk := unknown.(string)
	fmt.Println("Is string? ", str, strOk)
	intVal, intOk := unknown.(int)
	fmt.Println("Is int? ", intVal, intOk)

	// Type switch
	fmt.Println("\nType switch")
	switch switchCheck := unknown.(type) {
		case string:
			fmt.Println("String: ", strings.ToUpper(switchCheck))
		case int:
			fmt.Printf("Int: %d\n", switchCheck)
		default:
			fmt.Println("Unknown type:", switchCheck)
	}

	// Stringer interface
	fmt.Println("\nStringer interface")
	fmt.Println(&somePoint)

	// Error handling
	fmt.Println("\nError handling")
	if err := funcWithError(); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("No error")
	}

	// Readers 
	fmt.Println("\nReaders")
	r := &MyReader{}
	b := make([]byte, 10)
	n, _ := r.Read(b)
	fmt.Println(n, string(b))
}
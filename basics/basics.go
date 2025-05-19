package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"runtime"
	"strings"
)

func add(x, y int) int {
	return x + y
}

func swap(x, y int) (int, int) {
	return y, x
}

func toUpperLower(s string) (upper, lower string) {
	upper = strings.ToUpper(s)
	lower = strings.ToLower(s)
	return
}

var python, java bool

func Sqrt(x float64) float64 {
	lower := 0.0
	upper := x
	var mid float64
	for i := 0; i < 10; i++ {
		mid = lower + (upper - lower)/2
		//fmt.Println("Mid:", mid)
		if mid*mid < x {
			lower = mid
		} else if mid*mid > x {
			upper = mid
		} else {
			break
		}
	}
	return mid
}

type Point struct {
	X, Y int
}

func compose(outerFunc func(int) int, innerFunc func(int, int) int) func(int, int) int {
	return func(x, y int) int {
		return outerFunc(innerFunc(x, y))
	}
}

func logger() func(string) string {
	log := make([]string, 0)
	return func(s string) string {
		log = append(log, s)
		return strings.Join(log, "...")
	}
}

func main() {
	// Basic constructs
	fmt.Println("Basic constructs")
	fmt.Println(math.Pi)
	fmt.Println(add(42, 13))
	fmt.Println(swap(42, 13))
	fmt.Println(toUpperLower("Hello, World!"))
	var index = 42
	fmt.Println(index, python, java)
	differentVarDeclaration := 42
	fmt.Println(differentVarDeclaration)

	// Types
	fmt.Println("\nTypes")
	var (
		ToBe   bool       = false
		MaxInt uint64     = 1<<64 - 1
		complex      complex128 = cmplx.Sqrt(-5 + 12i)
	)
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", complex, complex)

	a := 42
	b := float64(a)
	var c = uint(b)
	fmt.Println(a, b, c)

	// Constants
	fmt.Println("\nConstants")
	const WORLD = "World"
	fmt.Println("Hello", WORLD)

	// Loops
	fmt.Println("\nLoops")
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println("Sum:", sum)

	// If statements
	fmt.Println("\nIf statements")
	if num := 0; num > 0 {
		fmt.Println("Num is positive:", num)
	} else {
		fmt.Println("Num isn't positive:", num)
	}

	// Test Sqrt function
	fmt.Println("\nTest Sqrt function")
	fmt.Println("Sqrt(2) =", Sqrt(2))
	fmt.Println("Sqrt(4) =", Sqrt(4))
	fmt.Println("Sqrt(25) =", Sqrt(25))
	fmt.Println("Sqrt(95) =", Sqrt(95))

	// Switch statement
	fmt.Println("\nSwitch")
	switch os := runtime.GOOS; os {
		case "darwin":
			fmt.Println("OS X.")
		case "linux":
			fmt.Println("Linux.")
		case "windows":
			fmt.Println("Windows.")
		default:
			fmt.Printf("%s.\n", os)
	}

	// Defer statement (uncomment to see the effect)
	changingVar := 1
	//defer fmt.Printf("World%d\n", changingVar)
	changingVar = 2
	//defer fmt.Printf("World%d\n", changingVar)
	changingVar = 40
	fmt.Println("Hello", changingVar)

	// Pointers
	fmt.Println("\nPointers")
	num1, num2 := 42, 2701
	var p1 *int = &num1
	fmt.Println(*p1)
	*p1 = 10
	fmt.Println(num1)

	p2 := &num2
	*p2 = *p2 / 37
	fmt.Println(num2)

	// Structs
	fmt.Println("\nStructs")
	point := Point{1, 2}
	point.X = 4
	fmt.Println(point.X, point.Y)

	pointerToPoint := &point
	pointerToPoint.X = 1e9
	fmt.Println(point.X, point.Y)

	var(
		point1 = Point{1, 2}
		point2 = Point{Y: 1}
		point3 = Point{}
		pointerToPoint2 = &Point{1, 2}
	)
	fmt.Println(point1, point2, point3, pointerToPoint2)

	// Arrays
	fmt.Println("\nArrays")
	var arr1 [2]string
	arr1[0] = "Hello"
	arr1[1] = "World"
	fmt.Println(arr1[0], arr1[1])
	arr2 := [2]string{"Hello", "World"}
	fmt.Println(arr2)

	// Slices
	fmt.Println("\nSlices")
	primes := [6]int{2, 3, 5, 7, 11, 13}
	var primesSlice []int = primes[1:4]
	fmt.Println(primesSlice)

	primesSlice2 := primesSlice[0:2]
	fmt.Println(primesSlice2)
	primesSlice[0] = 999
	fmt.Println(primes)
	fmt.Println(primesSlice2)

	structsSlice := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(structsSlice)

	sliceWithDefault := primes[:4]
	fmt.Println(sliceWithDefault)

	fmt.Println(primesSlice2)
	fmt.Println(len(primesSlice2), cap(primesSlice2))
	primesSlice2 = primesSlice2[:3]
	fmt.Println(primesSlice2)
	fmt.Println(len(primesSlice2), cap(primesSlice2))

	dynamicSlice := make([]int, 5)
	fmt.Println(dynamicSlice)

	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println(matrix)

	fmt.Println(primes)
	primesSlice3 := primes[1:4]
	fmt.Println(primesSlice3)
	primesSlice3 = append(primesSlice3, 1000)
	fmt.Println(primesSlice3)
	fmt.Println(primes)
	primesSlice3 = append(primesSlice3, 1001, 1002)
	fmt.Println(primesSlice3)
	fmt.Println(primes)

	// Ranges
	fmt.Println("\nRanges")
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
	for _, v := range pow {
		fmt.Printf("Value only = %d\n", v)
	}

	// Maps
	fmt.Println("\nMaps")
	pointMap := make(map[string]Point)
	pointMap["A"] = Point{1, 2}
	pointMap["B"] = Point{3, 4}
	fmt.Println(pointMap)

	pointMap2 := map[string]Point{
		"A": {2, 3},
		"B": {4, 5},
	}
	fmt.Println(pointMap2)
	pointMap2["C"] = Point{6, 7}
	fmt.Println(pointMap2["C"])
	delete(pointMap2, "B")
	fmt.Println(pointMap2)
	elem, ok := pointMap2["B"]
	fmt.Println(elem, ok)

	// Function Values
	fmt.Println("\nFunction Values")

	outerFunc := func(x int) int {
		return x * x
	}
	innerFunc := func(x, y int) int {
		return x + y
	}
	composedFunc := compose(outerFunc, innerFunc)
	fmt.Println(composedFunc(2, 3))
	fmt.Println(composedFunc(0, 3))

	loggerFunc := logger()
	fmt.Println(loggerFunc("Hello"))
	fmt.Println(loggerFunc("World"))
	fmt.Println(loggerFunc("!"))
}
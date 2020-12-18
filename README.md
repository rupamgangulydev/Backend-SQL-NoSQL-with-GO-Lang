

## Section 1 : Why use GO Lang as Backend :

|  NodeJs v/s Go | Reddit Post  | Reddit Post  | Valorent's Backend  |
| :---:   | :-: | :-: | :-: |
| <img src="/ASSETS/1.png" width=850> | <img src="/ASSETS/2.png"> | <img src="/ASSETS/3.png"> |  <img src="/ASSETS/4.png"> |


## Section 2 : GO Lang Fundamentals:
<details>
	<summary> Click here for expanding GO Lang Fundamentals</summary>
	
### VARIABLE
Inside a function, the := short assignment statement can be used in place of a var declaration with implicit type.

~~var int i  // syntax error~~
```go
var i int // // SYNTAX IS CORRECT
//Here i is the name of the variable. Type of the variable is int. 

package main

func main() {
	k := 3 
	m, n, p, q := 1, 34.67, true, "i am string"
	// Outside a function, every statement begins with a keyword (var, func, and so on)
	// and so the := construct is not available.
}

```
### FUNCTION
SIGNATURE of FUNCTION:
```go
func function_name (argument1 type, argument2 type) (return_type1, return_type2){
	return a,b
}
```
```go
package main

import "fmt"

func factorial(num int) int {
	if num > 1 {
		return num * factorial(num-1)
	}
	return 1
}
func anotherFunc() {
	fmt.Println("I am Another Function")
}

// Function as value - Anonymous function
var result = func(a, b int) int {
	return a + b
}

func main() {
	fmt.Println("Factorial of 4", factorial(4)) //Factorial of 4 24
	defer anotherFunc() //defer keword in GO makes a function execute at the end of the
	//execution (or when hits return statement) of parent function from where it is called.
	fmt.Println("HI I am MAIN and I execute first instead of defer function")
	fmt.Println("HI I am MAIN and I execute first instead of defer function")
	fmt.Println("HI I am MAIN and I execute first instead of defer function")
	fmt.Println("Calling Function as value - Anonymous function", result(6, 9))
	subtra := func(a, b int) int {
		return a - b
	}(89, 20)
	fmt.Println("Subtra is caling: ", subtra)
}
```
#### OUTPUT
```shell
Factorial of 4 24
HI I am MAIN and I execute first instead of defer function
HI I am MAIN and I execute first instead of defer function
HI I am MAIN and I execute first instead of defer function
Calling Function as value - Anonymous function 15
Subtra is caling:  69
I am Another Function
```



### ComandLine Argument Passing
```go
package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		// If you’re wondering why we expect 2 arguments,
		// it’s because the first argument – at index 0 – is always
		// the path of the currently running executable.
		os.Exit(1)
	}
	fmt.Println("It's over", os.Args[1])
}
```
#### OUTPUT
```shell
PS E:\PROJECTS\BACKEND\gobackendcrud> go run demo.go
exit status 1
PS E:\PROJECTS\BACKEND\gobackendcrud> go run demo.go 1234
It's over 1234
PS E:\PROJECTS\BACKEND\gobackendcrud> 
```
### POINTER
```go
package main

import (
	"fmt"
)

func passedasReferencee(a *int) { // d is passed
	fmt.Println("*a is : ", *a)
	fmt.Println("a is : ", a)
	fmt.Println("address of a variable is: ", &a)
	fmt.Println("Incrementing:")
	*a++ // incrementing value - *a was 7 , now *a will be 8
	fmt.Println("*a is : ", *a)
	fmt.Println("a is : ", a)

}

func passedasvariable(a int) { // a is copy of c
	fmt.Println("a is : ", a)
	fmt.Println("address of a variable is: ", &a)
	fmt.Println("Incrementing:")
	a++ //a is incremented from 7 to 8
	fmt.Println("a is : ", a)
}
func main() {
	c := 7
	fmt.Println("in main function address of c is : ", &c)
	d := &c // d is memory address of c
	fmt.Println("in main function d is : ", d)
	fmt.Println("in main function *d is : ", *d)
	fmt.Println("passedasReferencee called:--->")
	passedasReferencee(d)
	fmt.Println("after executing function :- ")
	fmt.Println("c is : ", c)
	fmt.Println("d is : ", d)
	fmt.Println("*d is ", *d)
	fmt.Println("passedasVariable called:--->")
	passedasvariable(c) // here we passed c
	fmt.Println("after executing function :- ")
	fmt.Println("c is : ", c) // no change
	fmt.Println("d is : ", d)
}

```
#### 		OUTPUT:
```shell
	in main function address of c is :  0xc000012090
	in main function d is :  0xc000012090
	in main function *d is :  7
	passedasReferencee called:--->       
	*a is :  7
	a is :  0xc000012090
	address of a variable is:  0xc000006030
	Incrementing:
	*a is :  8
	a is :  0xc000012090
	after executing function :-
	c is :  8
	d is :  0xc000012090
	*d is  8
	passedasVariable called:--->
	a is :  8
	address of a variable is:  0xc000012098
	Incrementing:
	a is :  9
	after executing function :-
	c is :  8
	d is :  0xc000012090
```

### Go lang tries to distance itself from object-oriented terminology, at the same time it can give you all object-oriented programming flavors and benefits. For instance, C++ or Java classes are called a struct, A class methods are called receivers. Inheritance is called embedding
### STRUCT
```go
package main

import (
	"fmt"
)

type DataStruct struct {
	string
	bool
}
type Employee struct {
	firstname string
	lastname  string
	salary    int
	fulltime  bool
}

func main() {

	var ross Employee
	fmt.Println(ross) // {"" 0 false}
	ross.firstname = "Rupam"
	ross.lastname = "Ganguly"
	ross.salary = 45000
	ross.fulltime = true
	fmt.Println(ross) // {Rupam Ganguly 45000 true}
	bos := Employee{
		firstname: "Rintu",
		lastname:  "Ganguly",
		salary:    58000,
		fulltime:  true,
	}
	fmt.Println(bos) // {Rintu Ganguly 58000 true}
	mos := Employee{"my mosh", "ganguly", 90000, true}
	fmt.Println(mos) // {my mosh ganguly 90000 true}
	// Anonymous struct
	monica := struct {
		age     int
		salary  int
		teacher bool
	}{
		age: 12, salary: 12345, teacher: true,
	}
	fmt.Println(monica) // {12 12345 true}
	// Pointer Struct
	pointerRoss := &Employee{
		firstname: "Rintu",
		lastname:  "Ganguly",
		salary:    58000,
		fulltime:  true,
	}
	fmt.Println(pointerRoss)                               // &{Rintu Ganguly 58000 true}
	fmt.Println(*pointerRoss)                              // {Rintu Ganguly 58000 true}
	fmt.Println("first name : ", (*pointerRoss).firstname) // first name :  Rintu
	fmt.Println(pointerRoss.firstname)                     // Rintu
	// DataStruct
	samp1 := DataStruct{"Monday time", true}
	fmt.Println(samp1) // {Monday time true}
	samp1.bool = false
	fmt.Println(samp1) // {Monday time false}
	samp2 := DataStruct{"Sunday", false}
	fmt.Println(samp2) // {Sunday false}
	samp3 := samp1
	fmt.Println(samp3) // {Monday time false}

	// Nested Struct
	type UpdatedEmployee struct {
		firstname string
		lastname  string
		salary    int
		fulltime  bool
		data      DataStruct
	}
	nestedstructMy := UpdatedEmployee{
		firstname: "Rintu",
		lastname:  "Ganguly",
		salary:    58000,
		fulltime:  true,
		data:      DataStruct{"new data inserted", true},
	}
	fmt.Println(nestedstructMy)             // {Rintu Ganguly 58000 true {new data inserted true}}
	fmt.Println(nestedstructMy.salary)      // 58000
	fmt.Println(nestedstructMy.data.string) // new data inserted
}
```
#### OUTPUT
```shell
{  0 false}
{Rupam Ganguly 45000 true}
{Rintu Ganguly 58000 true}
{my mosh ganguly 90000 true}
{12 12345 true}
&{Rintu Ganguly 58000 true}
{Rintu Ganguly 58000 true}
first name :  Rintu
Rintu
{Monday time true}
{Monday time false}
{Sunday false}
{Monday time false}
{Rintu Ganguly 58000 true {new data inserted true}}
58000
new data inserted
```

### INTERFACE

An interface is a collection of method signatures that a Type(struct) can implement (using methods). Hence interface defines (not declares) the behavior of the object
An interface is declared using the type keyword, followed by the name of the interface and the keyword interface. Then, we specify a set of method signatures inside curly braces.
Go Interface doesn’t have fields and also it doesn’t allow the definition of methods inside it. Any type needs to implements all methods of interface to become of that interface type.
To implement an interface, you just need to implement all the methods declared in the interface.

```go
package main

import (
	"fmt"
	"math"
)

// in order to successfully implemented an interface,
// we need to implement all the methods declared by the interface with exact signature
type Shape interface {
	Area() float64
	Perimeter() float64
}
type AnotherShape interface {
	NewArea() string
	NewPerimeter() float64
}
type Rect struct {
	width, height float64
}
type Circle struct {
	radius float64
}

func (r Rect) Area() float64 {
	fmt.Println(r.width * r.height)
	return 7
}
func (r Rect) Perimeter() float64 {
	fmt.Println(2 * (r.width + r.height))
	return 7
}
func (c Circle) Perimeter() float64 {
	fmt.Println(2 * (math.Pi * c.radius))
	return 7
}
func (c Circle) Area() float64 {
	fmt.Println(math.Pi * c.radius * 2)
	return 7
}
func (r Rect) NewArea() string {
	return "Implements multiple interfaces"
}
func (r Rect) NewPerimeter() float64 {
	return 9874321
}
func main() {
	var s Shape = Rect{10, 3}
	s.Area() // 30
	// multiple interface of Rect
	var ns AnotherShape = Rect{5, 8}
	fmt.Println(ns.NewArea())      // Implements multiple interfaces
	fmt.Println(ns.NewPerimeter()) // 9.874321e+06
	s = Circle{10}
	s.Perimeter() // 62.83185307179586

	//whrn Area() didnt implemented by c Circle then we get error:
	//cannot use Circle literal (type Circle) as type Shape in assignment:
	//Circle does not implement Shape (missing Area method)

}
```
#### OUTPUT
```shell
30
Implements multiple interfaces
9.874321e+06
62.83185307179586

```


### Type Assertion

Findout the Type of Dynamic Value of an interface

```go
package main

import "fmt"

func explain(i interface{}) {
	switch i.(type) {
	case string:
		fmt.Println(" Interface has a String")
	case int:
		fmt.Println("Interface has a Int")
	default:
		fmt.Println("Interface has other Type")
	}
}
func main() {
	explain("Rupam Ganguly")
	explain(45.987)
	explain(43)
}
```
#### OUTPUT
```shell
Interface has a String
Interface has other Type
Interface has a Int
```
### MIMIC OF INHERITANCE

```go
// In GO we can use struct for inheritance, we can compose using structs to form other objects
package main

import "fmt"

type ParentClass struct {
	color      string
	background string
}

func (ParentClass) basFuncOne() string {
	return "I am Base Function ONE"
}
func (ParentClass) basFuncTwo() string {
	return "I am Base Function TWO"
}

type ChildClass struct {
	ParentClass
	style    string
	fontsize int
}

func (c ChildClass) childFunc() string {
	return "I am Child Function "
}
func main() {
	// Patent object creation
	baseObj := ParentClass{
		color:      "Black",
		background: "White",
	}
	// Child object creation
	childObj := ChildClass{
		ParentClass: baseObj,
		style:       "bold",
		fontsize:    32,
	}
	fmt.Println(childObj)    // {{Black White} bold 32}
	childObj.color = "Green" // childObj can access color backgroud fields of Base STRUCT
	childObj.background = "Red"
	childObj.fontsize = 16
	fmt.Println(childObj)              // {{Green Red} bold 16}
	fmt.Println(childObj.basFuncOne()) // childObj can access basFunc Function of Base STRUCT
	fmt.Println(childObj.basFuncTwo()) // I am Base Function ONE // I am Base Function TWO
	fmt.Println(childObj.childFunc())
}
```
#### OUTPUT
```shell
{{Black White} bold 32}
{{Green Red} bold 16}
I am Base Function ONE
I am Base Function TWO
I am Child Function 
```
### Allocation

Go has two allocation primitives, the built-in functions new and make. new is a built-in function that allocates memory, but unlike its namesakes in some other languages it does not initialize the memory, it only zeros it. That is, new(T) allocates zeroed storage for a new item of type T and returns its address, a value of type *T. 

```go
package main

import "fmt"

type myBox struct {
	index int
	data  int
}

func main() {
	obj1 := new(myBox)
	fmt.Println(obj1)  // &{0 0}
	fmt.Println(*obj1) // {0 0}
	var obj2 myBox
	fmt.Println(obj2) // {0 0}
	//fmt.Println(*obj2) // invalid indirect of obj2 (type myBox)
}
```
#### OUTPUT
```shell
&{0 0}
{0 0}
{0 0}
```
The built-in function make(T, args) serves a purpose different from new(T). It creates slices, maps, and channels only, and it returns an initialized (not zeroed) value of type T (not *T). The reason for the distinction is that these three types represent, under the covers, references to data structures that must be initialized before use. Example of make function is present in Sliice section...

### Array

Basically, an array is a number of elements of same type stored in sequential order.

Once you declare an array with its size you are not allowed to change it.

If you try to insert more elements than array size, compiler will give you an error.

By default array size is 0 (zero)

Array index starts from 0th index

We can set value directly to array at particular index array_name[index]=value

The inbuilt len returns length of an array
```go
package main

import "fmt"

func main() {
	x := [12]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(x) // [1 2 3 4 5 6 7 8 9 0 0 0]
	var y [12]int
	fmt.Println(y) // [0 0 0 0 0 0 0 0 0 0 0 0]
	y[5] = 45
	y[2] = 12
	fmt.Println(y) // [0 0 12 0 0 45 0 0 0 0 0 0]
	y[5] = 123
	fmt.Println(y) // [0 0 12 0 0 123 0 0 0 0 0 0]
	fmt.Println("length of y: ", len(y)) // length of y:  12
	for i := 0; i < len(y); i++ {
		fmt.Print(y[i], " - ") // 0 - 0 - 12 - 0 - 0 - 123 - 0 - 0 - 0 - 0 - 0 - 0 -
	}
	fmt.Println()
	for i := 0; i < len(y); i++ {
		fmt.Print(y[i]+x[i], " - ") // 1 - 2 - 15 - 4 - 5 - 129 - 7 - 8 - 9 - 0 - 0 - 0 -
	}
}
```
#### OUTPUT
```shell
[1 2 3 4 5 6 7 8 9 0 0 0]
[0 0 0 0 0 0 0 0 0 0 0 0]
[0 0 12 0 0 45 0 0 0 0 0 0]
[0 0 12 0 0 123 0 0 0 0 0 0]
length of y:  12
0 - 0 - 12 - 0 - 0 - 123 - 0 - 0 - 0 - 0 - 0 - 0 -
1 - 2 - 15 - 4 - 5 - 129 - 7 - 8 - 9 - 0 - 0 - 0 - 
```
### SLICE

Slice is the same as an array but it has a variable length so we don’t need to specify the length to it. It will grow whenever it exceeds its size. Like an array, slice also has index and length but its length can be changed.

Slice also has continuous segments of memory locations

The default value of uninitialized slice is nil

Slices does not store the data. It just provides reference to an array

As we change the elements of slice, it will modify corresponding elements of that array
```go
package main

import "fmt"

func main() {
	var x []int
	fmt.Println(x) // []
	y := []int{2, 3, 4, 5, 6, 7}
	fmt.Println(y) // [2 3 4 5 6 7]
	z := make([]int, 12)
	fmt.Println(z)      // [0 0 0 0 0 0 0 0 0 0 0 0]
	fmt.Println(y[3:5]) // [5 6]
	fmt.Println(y[1:5]) // [3 4 5 6]
	fmt.Println(y[:4])  // [2 3 4 5]
	sl := make([][]int, 12)
	fmt.Println(sl) // [[] [] [] [] [] [] [] [] [] [] [] []]
	for i := range sl {
		fmt.Print(sl[i], " -> ") // [] -> [] -> [] -> [] -> [] -> [] -> [] -> [] -> [] -> [] -> [] -> [] ->
	}
	fmt.Println()
	slicef := [5][3]int{}
	fmt.Println(slicef) // [[0 0 0] [0 0 0] [0 0 0] [0 0 0] [0 0 0]]
	//slicef[1][1] = j
	for i := range slicef {
		fmt.Print(" row ", i, " -> ")
		for j := 0; j < len(slicef[0]); j++ {
			fmt.Print(" col ", j)
		}
		fmt.Println()

	}
	// row 0 ->  col 0 col 1 col 2
	// row 1 ->  col 0 col 1 col 2
	// row 2 ->  col 0 col 1 col 2
	// row 3 ->  col 0 col 1 col 2
	// row 4 ->  col 0 col 1 col 2
	k := 0
	for i := range slicef {
		for j := 0; j < len(slicef[0]); j++ {
			slicef[i][j] = k
			k++
		}
	}
	fmt.Println(slicef) // [[0 1 2] [3 4 5] [6 7 8] [9 10 11] [12 13 14]]
	fmt.Println(x)      // []
	x = append(x, 34)
	fmt.Println(x) // [34]
	x = append(x, 12, 7, 9, 1)
	fmt.Println(x) // [34 12 7 9 1]
	for i, v := range x {
		fmt.Printf("INDEX: %d Value: %d \n ", i, v)
	}
		//  INDEX: 0 Value: 34
		//  INDEX: 1 Value: 12
		//  INDEX: 2 Value: 7
		//  INDEX: 3 Value: 9 
		//  INDEX: 4 Value: 1
}
```
#### OUTPUT
```shell
[]
[2 3 4 5 6 7]
[0 0 0 0 0 0 0 0 0 0 0 0]
PS E:\PROJECTS\BACKEND\gobackendcrud> go run datastructure.go
[]
[2 3 4 5 6 7]
[0 0 0 0 0 0 0 0 0 0 0 0]
[5 6]
[3 4 5 6]
[2 3 4 5]
[[] [] [] [] [] [] [] [] [] [] [] []]
[] -> [] -> [] -> [] -> [] -> [] -> [] -> [] -> [] -> [] -> [] -> [] -> 
[[0 0 0] [0 0 0] [0 0 0] [0 0 0] [0 0 0]]
 row 0 ->  col 0 col 1 col 2
 row 1 ->  col 0 col 1 col 2
 row 2 ->  col 0 col 1 col 2
 row 3 ->  col 0 col 1 col 2
 row 4 ->  col 0 col 1 col 2
[[0 1 2] [3 4 5] [6 7 8] [9 10 11] [12 13 14]]
[]
[34]
[34 12 7 9 1]
INDEX: 0 Value: 34
 INDEX: 1 Value: 12
 INDEX: 2 Value: 7
 INDEX: 3 Value: 9 
 INDEX: 4 Value: 1
```
#### Deleting Element From Slice
```go
package main

import "fmt"

func main() {
	slicel := []string{"A", "B", "C", "D", "E", "F", "G"}
	i := 3
	fmt.Println(slicel) // [A B C D E F G]
	slicel[i] = slicel[len(slicel)-1]
	slicel[len(slicel)-1] = ""
	slicel = slicel[:len(slicel)-1]
	fmt.Println(slicel) // [A B C G E F]
	slicel = []string{"A", "B", "C", "D", "E", "F", "G"}
	fmt.Println(slicel) // [A B C D E F G]
	copy(slicel[i:], slicel[i+1:])
	fmt.Println(slicel) // [A B C E F G G]
	slicel[len(slicel)-1] = ""
	fmt.Println(slicel) // [A B C E F G ]
	slicel = slicel[:len(slicel)-1]
	fmt.Println(slicel) // [A B C E F G]
}
```
#### OUTPUT
```shell
[A B C D E F G]
[A B C G E F]
[A B C D E F G]
[A B C E F G G]
[A B C E F G ]
[A B C E F G]
```
### MAP
Maps are a convenient and powerful built-in data structure that associate values of one type (the key) with values of another type (the element or value). The key can be of any type for which the equality operator is defined, such as integers, floating point and complex numbers, strings, pointers, interfaces (as long as the dynamic type supports equality), structs and arrays. Slices cannot be used as map keys, because equality is not defined on them. Like slices, maps hold references to an underlying data structure. If you pass a map to a function that changes the contents of the map, the changes will be visible in the caller.
```go
package main

import (
	"fmt"
	"sort"
)

var myMap map[int]string

func main() {
	fmt.Println(myMap) // map[]
	//myMap[12] = "I am value of key 12 "
	//fmt.Println(myMap) // panic: assignment to entry in nil map

	// Map types are reference types, like pointers or slices,
	//  and so the value of myMap above is nil; it doesn't point to an initialized map.
	// A nil map behaves like an empty map when reading,
	// but attempts to write to a nil map will cause a runtime panic; don't do that.
	// To initialize a map, use the built in make function:
	myMap = make(map[int]string)
	fmt.Println(myMap) //map[]
	myMapOne := map[int]int{}
	fmt.Println(myMapOne) // map[]
	myMap[16] = "I am value of key 16"
	fmt.Println(myMap) // map[16:I am value of key 16]
	myMap[17] = "I am value of key 17"
	fmt.Println(myMap) // map[16:I am value of key 16 17:I am value of key 17]
	myMap[21] = "I am value of key 21"
	fmt.Println(myMap) // map[16:I am value of key 16 17:I am value of key 17 21:I am value of key 21]
	val := myMap[17]
	fmt.Println(val) // I am value of key 17
	anotherVal, isPresent := myMap[12345]
	fmt.Println(anotherVal, "->", isPresent) // -> false
	//The built in delete function removes an entry from the map:
	delete(myMap, 16)
	fmt.Println(myMap)      // map[17:I am value of key 17 21:I am value of key 21]
	fmt.Println(len(myMap)) // 2
	anotherMap := map[string]string{
		"one":   "i am value of 1",
		"two":   "i am value of 2",
		"three": "i am value of 3",
		"four":  "i am value of 4",
		"five":  "i am value of 5",
		"six":   "i am value of 6",
		"seven": "i am value of 7",
	}
	for key, val := range anotherMap {
		fmt.Println("Key is : ", key, " Value is : ", val, " - > ")
	}
	// Key is :  one  Value is :  i am value of 1  - >
	// Key is :  two  Value is :  i am value of 2  - >
	// Key is :  three  Value is :  i am value of 3  - >
	// Key is :  four  Value is :  i am value of 4  - >
	// Key is :  five  Value is :  i am value of 5  - >
	// Key is :  six  Value is :  i am value of 6  - >
	// Key is :  seven  Value is :  i am value of 7  - >

	//see when we run same code twice then the output will be different - not ordered

	// Key is :  six  Value is :  i am value of 6  - >
	// Key is :  seven  Value is :  i am value of 7  - >
	// Key is :  one  Value is :  i am value of 1  - >
	// Key is :  two  Value is :  i am value of 2  - >
	// Key is :  three  Value is :  i am value of 3  - >
	// Key is :  four  Value is :  i am value of 4  - >
	// Key is :  five  Value is :  i am value of 5  - >

	// When iterating over a map with a range loop, the iteration order is not specified
	// and is not guaranteed to be the same from one iteration to the next.
	// If you require a stable iteration order
	// you must maintain a separate data structure that specifies that order.
	myMap[34] = "I am value of key 34"
	myMap[41] = "I am value of key 41"
	myMap[4] = "I am value of key 4"
	var sliceOfKeys []int
	for k := range myMap {
		sliceOfKeys = append(sliceOfKeys, k)
	}
	fmt.Println(sliceOfKeys) // [41 4 34 17 21]
	sort.Ints(sliceOfKeys)

	fmt.Println(sliceOfKeys) // [4 17 21 34 41]

	for _, ki := range sliceOfKeys {
		fmt.Println("key: ", ki, " Value: ", myMap[ki])
	}
	// key:  4  Value:  I am value of key 4
	// key:  17  Value:  I am value of key 17
	// key:  21  Value:  I am value of key 21
	// key:  34  Value:  I am value of key 34
	// key:  41  Value:  I am value of key 41

}
```
#### OUTPUT
```shell
map[]
map[]
map[]
map[16:I am value of key 16]
map[16:I am value of key 16 17:I am value of key 17]
map[16:I am value of key 16 17:I am value of key 17 21:I am value of key 21]
I am value of key 17
 -> false
map[17:I am value of key 17 21:I am value of key 21]
2
Key is :  two  Value is :  i am value of 2  - >
Key is :  three  Value is :  i am value of 3  - >
Key is :  four  Value is :  i am value of 4  - >
Key is :  five  Value is :  i am value of 5  - >
Key is :  six  Value is :  i am value of 6  - >
Key is :  seven  Value is :  i am value of 7  - >
Key is :  one  Value is :  i am value of 1  - >
[41 4 34 17 21]
[4 17 21 34 41]
key:  4  Value:  I am value of key 4
key:  17  Value:  I am value of key 17
key:  21  Value:  I am value of key 21
key:  34  Value:  I am value of key 34
key:  41  Value:  I am value of key 41
```



</details>

## Section 3 : Data STructure and Algorithm:

<details>
	<summary> Click here for expanding Data STructure and Algorithm</summary>
	
### Why use Pointer in code-:
```go
package main

import "fmt"

type Stack []string

var index int

func push() {
	index = 3
	fmt.Println("I am Push function")
}
func newPush() {
	index = 1234
	fmt.Println("I am newPush")
}
func normalPush0() {
	index = 9
	fmt.Println("I am normalPush0")
}
func (Stack) normalPush1() {
	fmt.Println("I am normalPush1")

}
func (s Stack) abnormalPush0() {
	fmt.Println("I am abnormalPush0")
	s = append(s, "one")
	fmt.Println(s) //[one]
	s = append(s, "two")
	fmt.Println(s) //[one two]
	s = append(s, "three")
	fmt.Println(s) //[one two three]
}
func (s Stack) abnormalPush1() {
	fmt.Println("I am abnormalPush1")
	s = append(s, "zero")
	fmt.Println(s) // [zero]
	s = append(s, "minus one")
	fmt.Println(s) // [zero minus one]
	s = append(s, "minus two")
	fmt.Println(s) // [zero minus one minus two]
}
func (s *Stack) pointerPush0() {
	fmt.Println("pointerPush0")
	*s = append(*s, "four")
	fmt.Println(s) // &[four]
	*s = append(*s, "five")
	fmt.Println(s) // &[four five]
	*s = append(*s, "six")
	fmt.Println(s) // &[four five six]
	*s = append(*s, "seven")
	fmt.Println(s) //&[four five six seven]
	*s = append(*s, "eight")
	fmt.Println(s) // &[four five six seven eight]
	*s = append(*s, "nine")
	fmt.Println(s) //&[four five six seven eight nine]
}
func (s *Stack) pointerPush1() { // invalid receiver type *Stack (Stack is an interface type)
	fmt.Println("pointerPush1")
	*s = append(*s, "ten")
	fmt.Println(s) // &[four five six seven eight nine ten]
	*s = append(*s, "eleven")
	fmt.Println(s) // &[four five six seven eight nine ten eleven]
	*s = append(*s, "twelve")
	fmt.Println(s) // &[four five six seven eight nine ten eleven twelve]
	*s = append(*s, "thirteen")
	fmt.Println(s) // &[four five six seven eight nine ten eleven twelve thirteen]
	*s = append(*s, "fourteen")
	fmt.Println(s) // &[four five six seven eight nine ten eleven twelve thirteen fourteen]
	*s = append(*s, "fifteen")
	fmt.Println(s) // &[four five six seven eight nine ten eleven twelve thirteen fourteen fifteen]
}
func main() {
	fmt.Println("index is : ", index) // index is :  0
	push()
	fmt.Println("index is : ", index) // index is :  3
	newPush()
	fmt.Println("index is : ", index) // index is :  1234
	normalPush0()
	fmt.Println("index is : ", index) // index is :  9
	index = 80009
	fmt.Println("index is : ", index) // index is :  80009
	// According to this index is global variable
	var st Stack
	//normalPush1() // undefined: normalPush1
	st.normalPush1() // I am normalPush1
	st.abnormalPush0()
	// I am abnormalPush0
	// [one]
	// [one two]
	// [one two three]
	st.abnormalPush1()
	// I am abnormalPush1
	// [zero]
	// [zero minus one]
	// [zero minus one minus two]
	st.pointerPush0()
	// pointerPush0
	// &[four]
	// &[four five]
	// &[four five six]
	// &[four five six seven]
	// &[four five six seven eight]
	// &[four five six seven eight nine]
	st.pointerPush1()
	// pointerPush1
	// &[four five six seven eight nine ten]
	// &[four five six seven eight nine ten eleven]
	// &[four five six seven eight nine ten eleven twelve]
	// &[four five six seven eight nine ten eleven twelve thirteen]
	// &[four five six seven eight nine ten eleven twelve thirteen fourteen]
	// &[four five six seven eight nine ten eleven twelve thirteen fourteen fifteen]
	fmt.Println()
}


```
#### OUTPUT
```shell
index is :  0
I am Push function
index is :  3     
I am newPush      
index is :  1234
I am normalPush0
index is :  9
index is :  80009
I am normalPush1
I am abnormalPush0
[one]
[one two]
[one two three]
I am abnormalPush1
[zero]
[zero minus one]
[zero minus one minus two]
pointerPush0
&[four]
&[four five]
&[four five six]
&[four five six seven]
&[four five six seven eight]
&[four five six seven eight nine]
pointerPush1
&[four five six seven eight nine ten]
&[four five six seven eight nine ten eleven]
&[four five six seven eight nine ten eleven twelve]
&[four five six seven eight nine ten eleven twelve thirteen]
&[four five six seven eight nine ten eleven twelve thirteen fourteen]
&[four five six seven eight nine ten eleven twelve thirteen fourteen fifteen]
```
### STACK
```go
// A stack is an ordered data structure that follows the Last-In-First-Out (LIFO) principle.
//  Stacks are most easily implemented in Golang using slices:

// An element is pushed to the stack with the built-in append function.
// The element is popped from the stack by slicing off the top element.

package main

import "fmt"

type Stack []interface{}
//type Stack []string
func (s *Stack) Push(item interface{}) {
// func (s *Stack) Push(item string) {
	*s = append(*s, item)
}
func (s *Stack) Pop() bool {
	if len(*s) != 0 {
		index := len(*s) - 1
		*s = (*s)[:index]
		return true
	}
	return false
}
func main() {
	var st Stack
	st.Push("one")
	st.Push("two")
	st.Push("three")
	st.Push("four")
	st.Push(34)
	st.Push(123.567)
	fmt.Println(st)
	st.Pop()
	fmt.Println(st)
	st.Pop()
	fmt.Println(st)
	st.Pop()
	fmt.Println(st)
	st.Pop()
	fmt.Println(st)
	st.Pop()
	fmt.Println(st)
	st.Pop()
	fmt.Println(st)
	st.Pop()
	fmt.Println(st)
}

```
#### OUTPUT
```shell
[one two three four 34 123.567]
[one two three four 34]
[one two three four]
[one two three]
[one two]
[one]
[]
[]
```
| Store String |  Store Any Type | 
| :---:   | :-: |
| <img src="/ASSETS/stackgo1.PNG"> | <img src="/ASSETS/stackgo2.PNG"> | 

### QUEUE

Queue follows a FIFO (First-In-First-Out) structure, the dequeue and enqueue operations can be performed as follows:
Use the built-in append function to enqueue.
Slice off the first element to dequeue.

```go
package main

import "fmt"

type Queue []interface{}

func (q *Queue) push(item interface{}) {
	*q = append(*q, item)
}
func (q *Queue) pop() {
	if len(*q) != 0 {
		*q = (*q)[1:]
	}
}
func main() {
	var qu Queue
	qu.push("hi")
	qu.push(12)
	qu.push(34.678)
	qu.push(true)
	fmt.Println(qu)
	qu.pop()
	fmt.Println(qu)
	qu.pop()
	fmt.Println(qu)
	qu.pop()
	fmt.Println(qu)
	qu.pop()
	fmt.Println(qu)
	qu.pop()
	fmt.Println(qu)
	qu.pop()
	fmt.Println(qu)
}
```
#### OUTPUT
```shell
[hi 12 34.678 true]
[12 34.678 true]
[34.678 true]
[true]
[]
[]
[]
```
### Doubly Linked List
```go
// In simple words, we can say, linked list is a collection of nodes. Node consists of two parts:
// Data
// Pointers- next previous
package main

import "fmt"

type Node struct {
	data interface{}
	next *Node
	prev *Node
}
type LinkedList struct {
	head *Node
	size int
}

func (l *LinkedList) show() {
	list := l.head
	for list != nil {
		//fmt.Println(list.data)
		fmt.Printf("%v -> ", list.data)
		list = list.next
	}
	fmt.Println("-------------$$$$$$$$$$$$----------------")
}
func (l *LinkedList) insertAt(item interface{}, index int) {
	// IF YOU WANT EXPLANATION THEN PLEASE GO TO -
	// https://github.com/rupamgangulydev/DataStructures_and_Algorithms/blob/master/ADS_004_DoublyLinkedList/DoublyLinkedList.java
	node := &Node{
		data: item,
	}
	if index == 0 {

		if l.head != nil {
			node.next = l.head
			l.head.prev = node
		}
		l.head = node
	} else {
		n := l.head
		for i := 0; i < index-1; i++ {
			n = n.next
		}
		if n.next != nil {
			node.next = n.next
			n.next.prev = node
		}
		node.prev = n
		n.next = node
	}
	l.size++
}
func (l *LinkedList) deleteAt(index int) {
	if index < l.size {
		if index == 0 {
			if l.size == 1 {
				l.head = nil
				l.size--
			} else {
				l.head = l.head.next
				l.head.prev = nil
				l.size--
			}
		} else {
			n := l.head
			for i := 0; i < index-1; i++ {
				n = n.next
			}
			temp := n.next
			if temp.next != nil {
				n.next = temp.next
				temp.next.prev = n
				l.size--
			} else {
				n.next = nil
				l.size--
			}
		}
	}
}
func main() {
	var list LinkedList
	fmt.Println("12 Insert at 0 ")
	list.insertAt(12, 0)
	list.show()
	fmt.Println("13 Insert at 0 ")
	list.insertAt(13, 0)
	list.show()
	fmt.Println("15 Insert at 0 ")
	list.insertAt(15, 0)
	list.show()
	fmt.Println("rupam Insert at 0 ")
	list.insertAt("rupam", 0)
	list.show()
	fmt.Println("11 Insert at 0 ")
	list.insertAt(11, 0)
	list.show()
	fmt.Println("17 Insert at 3 ")
	list.insertAt(17, 3)
	list.show()
	fmt.Println("100 Insert at 6 ")
	list.insertAt(100, 6)
	list.show()
	fmt.Println("45 Insert at 1 ")
	list.insertAt(45, 1)
	list.show()
	fmt.Println("delete at 0 ")
	list.deleteAt(0)
	list.show()
	fmt.Println("delete at 6 ")
	list.deleteAt(6)
	list.show()
	fmt.Println("delete at 3 ")
	list.deleteAt(3)
	list.show()
	fmt.Println("delete at 21 ")
	list.deleteAt(21)
	list.show()

}
```
#### OUTPUT

```shell
12 Insert at 0 
12 -> -------------$$$$$$$$$$$$----------------
13 Insert at 0
13 -> 12 -> -------------$$$$$$$$$$$$----------------
15 Insert at 0
15 -> 13 -> 12 -> -------------$$$$$$$$$$$$----------------
rupam Insert at 0
rupam -> 15 -> 13 -> 12 -> -------------$$$$$$$$$$$$----------------
11 Insert at 0
11 -> rupam -> 15 -> 13 -> 12 -> -------------$$$$$$$$$$$$----------------
17 Insert at 3
11 -> rupam -> 15 -> 17 -> 13 -> 12 -> -------------$$$$$$$$$$$$----------------
100 Insert at 6
11 -> rupam -> 15 -> 17 -> 13 -> 12 -> 100 -> -------------$$$$$$$$$$$$----------------
45 Insert at 1
11 -> 45 -> rupam -> 15 -> 17 -> 13 -> 12 -> 100 -> -------------$$$$$$$$$$$$----------------
delete at 0
45 -> rupam -> 15 -> 17 -> 13 -> 12 -> 100 -> -------------$$$$$$$$$$$$----------------
delete at 6
45 -> rupam -> 15 -> 17 -> 13 -> 12 -> -------------$$$$$$$$$$$$----------------
delete at 3
45 -> rupam -> 15 -> 13 -> 12 -> -------------$$$$$$$$$$$$----------------
delete at 21
45 -> rupam -> 15 -> 13 -> 12 -> -------------$$$$$$$$$$$$----------------
```





</details>

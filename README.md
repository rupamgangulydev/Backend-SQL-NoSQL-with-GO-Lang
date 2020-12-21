

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

#### IF YOU WANT EXPLANATION THEN PLEASE GO TO MY DataStructures_and_Algorithms REPOSITORY, through this link :- 
##### https://github.com/rupamgangulydev/DataStructures_and_Algorithms/blob/master/ADS_004_DoublyLinkedList/DoublyLinkedList.java
 ### There you will get  
 
 | ADS_000_DynamicArray    |
ADS_001_Stack   |
ADS_002_Queue   |
ADS_003_CircularQueue   |
ADS_004_DoublyLinkedList    |
ADS_005_BinarySearchTree    |
ADS_006_Heap    |
ADS_007_HashTable   |
ADS_008_HashMap |
ADS_009_SegmentTree |
ADS_010_Graphs  |
ADS_011_PythonicListImplementation  |
DP_000_0-1Knapsack  |
DP_001_SubsetSumANDEqualSumPartition    |
DP_002_CountSubsetSumORTargetSum    |
DP_003_MinimumSubsetSumDifference   |
DP_004_UnboundedKnapsac |
DP_005_CoinChange2  |
DP_006_LongestCommonSubSequence |
DP_007_PrintLongestSubsequence  |
DP_008_ShortestSuperSequence    |
DP_009_-MinNoOfInsDeltoConvertStrtoStr  |
DP_010_-PrintingShortestSuperseq    |
DP_011_PalindromeString |
DP_012_MinNofInserDeltoMekPal   |
DP_013_MatrixChainMultiplication    |
DP_014_BrustBallons |
DP_015_CutsInStic   |
DP_016_PalindromePartition  |
DP_017_-BooleanExpressionTrue   |
DP_018_ScrambleString   |
DP_019_EggDroping   |
DP_020_MaxPathSumofBTree    |
 

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
### Section 3 : BACKEND REST API
<details>
	<summary> Click here for expanding Package mongo provides a MongoDB Driver API for Go.</summary>
func BatchCursorFromCursor
func BatchCursorFromCursor(c *Cursor) *driver.BatchCursor
BatchCursorFromCursor returns a driver.BatchCursor for the given Cursor. If there is no underlying driver.BatchCursor, nil is returned.

Deprecated: This is an unstable function because the driver.BatchCursor type exists in the "x" package. Neither this function nor the driver.BatchCursor type should be used by applications and may be changed or removed in any release.

func WithSession
func WithSession(ctx context.Context, sess Session, fn func(SessionContext) error) error
WithSession creates a new SessionContext from the ctx and sess parameters and uses it to call the fn callback. The SessionContext must be used as the Context parameter for any operations in the fn callback that should be executed under the session.

If the ctx parameter already contains a Session, that Session will be replaced with the one provided.

Any error returned by the fn callback will be returned without any modifications.

Example
type BSONAppender
type BSONAppender interface {
    AppendBSON([]byte, interface{}) ([]byte, error)
}
BSONAppender is an interface implemented by types that can marshal a provided type into BSON bytes and append those bytes to the provided []byte. The AppendBSON can return a non-nil error and non-nil []byte. The AppendBSON method may also write incomplete BSON to the []byte.

type BSONAppenderFunc
type BSONAppenderFunc func([]byte, interface{}) ([]byte, error)
BSONAppenderFunc is an adapter function that allows any function that satisfies the AppendBSON method signature to be used where a BSONAppender is used.

func (BSONAppenderFunc) AppendBSON
func (baf BSONAppenderFunc) AppendBSON(dst []byte, val interface{}) ([]byte, error)
AppendBSON implements the BSONAppender interface

type BulkWriteError
type BulkWriteError struct {
    WriteError            // The WriteError that occurred.
    Request    WriteModel // The WriteModel that caused this error.
}
BulkWriteError is an error that occurred during execution of one operation in a BulkWrite. This error type is only returned as part of a BulkWriteException.

func (BulkWriteError) Error
func (bwe BulkWriteError) Error() string
Error implements the error interface.

type BulkWriteException
type BulkWriteException struct {
    // The write concern error that occurred, or nil if there was none.
    WriteConcernError *WriteConcernError

    // The write errors that occurred during operation execution.
    WriteErrors []BulkWriteError

    // The categories to which the exception belongs.
    Labels []string
}
BulkWriteException is the error type returned by BulkWrite and InsertMany operations.

func (BulkWriteException) Error
func (bwe BulkWriteException) Error() string
Error implements the error interface.

func (BulkWriteException) HasErrorCode
func (bwe BulkWriteException) HasErrorCode(code int) bool
HasErrorCode returns true if any of the errors have the specified code.

func (BulkWriteException) HasErrorCodeWithMessage
func (bwe BulkWriteException) HasErrorCodeWithMessage(code int, message string) bool
HasErrorCodeWithMessage returns true if any of the contained errors have the specified code and message.

func (BulkWriteException) HasErrorLabel
func (bwe BulkWriteException) HasErrorLabel(label string) bool
HasErrorLabel returns true if the error contains the specified label.

func (BulkWriteException) HasErrorMessage
func (bwe BulkWriteException) HasErrorMessage(message string) bool
HasErrorMessage returns true if the error contains the specified message.

type BulkWriteResult
type BulkWriteResult struct {
    // The number of documents inserted.
    InsertedCount int64

    // The number of documents matched by filters in update and replace operations.
    MatchedCount int64

    // The number of documents modified by update and replace operations.
    ModifiedCount int64

    // The number of documents deleted.
    DeletedCount int64

    // The number of documents upserted by update and replace operations.
    UpsertedCount int64

    // A map of operation index to the _id of each upserted document.
    UpsertedIDs map[int64]interface{}
}
BulkWriteResult is the result type returned by a BulkWrite operation.

type ChangeStream
type ChangeStream struct {
    // Current is the BSON bytes of the current event. This property is only valid until the next call to Next or
    // TryNext. If continued access is required, a copy must be made.
    Current bson.Raw
    // contains filtered or unexported fields
}
ChangeStream is used to iterate over a stream of events. Each event can be decoded into a Go type via the Decode method or accessed as raw BSON via the Current field. For more information about change streams, see https://docs.mongodb.com/manual/changeStreams/.

func (*ChangeStream) Close
func (cs *ChangeStream) Close(ctx context.Context) error
Close closes this change stream and the underlying cursor. Next and TryNext must not be called after Close has been called. Close is idempotent. After the first call, any subsequent calls will not change the state.

func (*ChangeStream) Decode
func (cs *ChangeStream) Decode(val interface{}) error
Decode will unmarshal the current event document into val and return any errors from the unmarshalling process without any modification. If val is nil or is a typed nil, an error will be returned.

func (*ChangeStream) Err
func (cs *ChangeStream) Err() error
Err returns the last error seen by the change stream, or nil if no errors has occurred.

func (*ChangeStream) ID
func (cs *ChangeStream) ID() int64
ID returns the ID for this change stream, or 0 if the cursor has been closed or exhausted.

func (*ChangeStream) Next
func (cs *ChangeStream) Next(ctx context.Context) bool
Next gets the next event for this change stream. It returns true if there were no errors and the next event document is available.

Next blocks until an event is available, an error occurs, or ctx expires. If ctx expires, the error will be set to ctx.Err(). In an error case, Next will return false.

If Next returns false, subsequent calls will also return false.

Example
func (*ChangeStream) ResumeToken
func (cs *ChangeStream) ResumeToken() bson.Raw
ResumeToken returns the last cached resume token for this change stream, or nil if a resume token has not been stored.

Example
func (*ChangeStream) TryNext
func (cs *ChangeStream) TryNext(ctx context.Context) bool
TryNext attempts to get the next event for this change stream. It returns true if there were no errors and the next event document is available.

TryNext returns false if the change stream is closed by the server, an error occurs when getting changes from the server, the next change is not yet available, or ctx expires. If ctx expires, the error will be set to ctx.Err().

If TryNext returns false and an error occurred or the change stream was closed (i.e. cs.Err() != nil || cs.ID() == 0), subsequent attempts will also return false. Otherwise, it is safe to call TryNext again until a change is available.

This method requires driver version >= 1.2.0.

Example
type Client
type Client struct {
    // contains filtered or unexported fields
}
Client is a handle representing a pool of connections to a MongoDB deployment. It is safe for concurrent use by multiple goroutines.

The Client type opens and closes connections automatically and maintains a pool of idle connections. For connection pool configuration options, see documentation for the ClientOptions type in the mongo/options package.

Example
func Connect
func Connect(ctx context.Context, opts ...*options.ClientOptions) (*Client, error)
Connect creates a new Client and then initializes it using the Connect method. This is equivalent to calling NewClient followed by Client.Connect.

When creating an options.ClientOptions, the order the methods are called matters. Later Set* methods will overwrite the values from previous Set* method invocations. This includes the ApplyURI method. This allows callers to determine the order of precedence for option application. For instance, if ApplyURI is called before SetAuth, the Credential from SetAuth will overwrite the values from the connection string. If ApplyURI is called after SetAuth, then its values will overwrite those from SetAuth.

The opts parameter is processed using options.MergeClientOptions, which will overwrite entire option fields of previous options, there is no partial overwriting. For example, if Username is set in the Auth field for the first option, and Password is set for the second but with no Username, after the merge the Username field will be empty.

The NewClient function does not do any I/O and returns an error if the given options are invalid. The Client.Connect method starts background goroutines to monitor the state of the deployment and does not do any I/O in the main goroutine to prevent the main goroutine from blocking. Therefore, it will not error if the deployment is down.

The Client.Ping method can be used to verify that the deployment is successfully connected and the Client was correctly configured.

Example (AWS)
Example (Direct)
Example (Kerberos)
Example (PLAIN)
Example (Ping)
Example (ReplicaSet)
Example (SCRAM)
Example (SRV)
Example (Sharded)
Example (X509)
func NewClient
func NewClient(opts ...*options.ClientOptions) (*Client, error)
NewClient creates a new client to connect to a deployment specified by the uri.

When creating an options.ClientOptions, the order the methods are called matters. Later Set* methods will overwrite the values from previous Set* method invocations. This includes the ApplyURI method. This allows callers to determine the order of precedence for option application. For instance, if ApplyURI is called before SetAuth, the Credential from SetAuth will overwrite the values from the connection string. If ApplyURI is called after SetAuth, then its values will overwrite those from SetAuth.

The opts parameter is processed using options.MergeClientOptions, which will overwrite entire option fields of previous options, there is no partial overwriting. For example, if Username is set in the Auth field for the first option, and Password is set for the second but with no Username, after the merge the Username field will be empty.

func (*Client) Connect
func (c *Client) Connect(ctx context.Context) error
Connect initializes the Client by starting background monitoring goroutines. If the Client was created using the NewClient function, this method must be called before a Client can be used.

Connect starts background goroutines to monitor the state of the deployment and does not do any I/O in the main goroutine. The Client.Ping method can be used to verify that the connection was created successfully.

func (*Client) Database
func (c *Client) Database(name string, opts ...*options.DatabaseOptions) *Database
Database returns a handle for a database with the given name configured with the given DatabaseOptions.

func (*Client) Disconnect
func (c *Client) Disconnect(ctx context.Context) error
Disconnect closes sockets to the topology referenced by this Client. It will shut down any monitoring goroutines, close the idle connection pool, and will wait until all the in use connections have been returned to the connection pool and closed before returning. If the context expires via cancellation, deadline, or timeout before the in use connections have returned, the in use connections will be closed, resulting in the failure of any in flight read or write operations. If this method returns with no errors, all connections associated with this Client have been closed.

func (*Client) ListDatabaseNames
func (c *Client) ListDatabaseNames(ctx context.Context, filter interface{}, opts ...*options.ListDatabasesOptions) ([]string, error)
ListDatabaseNames executes a listDatabases command and returns a slice containing the names of all of the databases on the server.

The filter parameter must be a document containing query operators and can be used to select which databases are included in the result. It cannot be nil. An empty document (e.g. bson.D{}) should be used to include all databases.

The opts parameter can be used to specify options for this operation (see the options.ListDatabasesOptions documentation.)

For more information about the command, see https://docs.mongodb.com/manual/reference/command/listDatabases/.

Example
func (*Client) ListDatabases
func (c *Client) ListDatabases(ctx context.Context, filter interface{}, opts ...*options.ListDatabasesOptions) (ListDatabasesResult, error)
ListDatabases executes a listDatabases command and returns the result.

The filter parameter must be a document containing query operators and can be used to select which databases are included in the result. It cannot be nil. An empty document (e.g. bson.D{}) should be used to include all databases.

The opts paramter can be used to specify options for this operation (see the options.ListDatabasesOptions documentation).

For more information about the command, see https://docs.mongodb.com/manual/reference/command/listDatabases/.

func (*Client) NumberSessionsInProgress
func (c *Client) NumberSessionsInProgress() int
NumberSessionsInProgress returns the number of sessions that have been started for this client but have not been closed (i.e. EndSession has not been called).

func (*Client) Ping
func (c *Client) Ping(ctx context.Context, rp *readpref.ReadPref) error
Ping sends a ping command to verify that the client can connect to the deployment.

The rp paramter is used to determine which server is selected for the operation. If it is nil, the client's read preference is used.

If the server is down, Ping will try to select a server until the client's server selection timeout expires. This can be configured through the ClientOptions.SetServerSelectionTimeout option when creating a new Client. After the timeout expires, a server selection error is returned.

Using Ping reduces application resilience because applications starting up will error if the server is temporarily unavailable or is failing over (e.g. during autoscaling due to a load spike).

func (*Client) StartSession
func (c *Client) StartSession(opts ...*options.SessionOptions) (Session, error)
StartSession starts a new session configured with the given options.

If the DefaultReadConcern, DefaultWriteConcern, or DefaultReadPreference options are not set, the client's read concern, write concern, or read preference will be used, respectively.

Example (WithTransaction)
func (*Client) UseSession
func (c *Client) UseSession(ctx context.Context, fn func(SessionContext) error) error
UseSession creates a new Session and uses it to create a new SessionContext, which is used to call the fn callback. The SessionContext parameter must be used as the Context parameter for any operations in the fn callback that should be executed under a session. After the callback returns, the created Session is ended, meaning that any in-progress transactions started by fn will be aborted even if fn returns an error.

If the ctx parameter already contains a Session, that Session will be replaced with the newly created one.

Any error returned by the fn callback will be returned without any modifications.

func (*Client) UseSessionWithOptions
func (c *Client) UseSessionWithOptions(ctx context.Context, opts *options.SessionOptions, fn func(SessionContext) error) error
UseSessionWithOptions operates like UseSession but uses the given SessionOptions to create the Session.

Example
func (*Client) Watch
func (c *Client) Watch(ctx context.Context, pipeline interface{},
    opts ...*options.ChangeStreamOptions) (*ChangeStream, error)
Watch returns a change stream for all changes on the deployment. See https://docs.mongodb.com/manual/changeStreams/ for more information about change streams.

The client must be configured with read concern majority or no read concern for a change stream to be created successfully.

The pipeline parameter must be an array of documents, each representing a pipeline stage. The pipeline cannot be nil or empty. The stage documents must all be non-nil. See https://docs.mongodb.com/manual/changeStreams/ for a list of pipeline stages that can be used with change streams. For a pipeline of bson.D documents, the mongo.Pipeline{} type can be used.

The opts parameter can be used to specify options for change stream creation (see the options.ChangeStreamOptions documentation).

Example
type ClientEncryption
type ClientEncryption struct {
    // contains filtered or unexported fields
}
ClientEncryption is used to create data keys and explicitly encrypt and decrypt BSON values.

func NewClientEncryption
func NewClientEncryption(keyVaultClient *Client, opts ...*options.ClientEncryptionOptions) (*ClientEncryption, error)
NewClientEncryption creates a new ClientEncryption instance configured with the given options.

func (*ClientEncryption) Close
func (ce *ClientEncryption) Close(ctx context.Context) error
Close cleans up any resources associated with the ClientEncryption instance. This includes disconnecting the key-vault Client instance.

func (*ClientEncryption) CreateDataKey
func (ce *ClientEncryption) CreateDataKey(ctx context.Context, kmsProvider string, opts ...*options.DataKeyOptions) (primitive.Binary, error)
CreateDataKey creates a new key document and inserts it into the key vault collection. Returns the _id of the created document.

func (*ClientEncryption) Decrypt
func (ce *ClientEncryption) Decrypt(ctx context.Context, val primitive.Binary) (bson.RawValue, error)
Decrypt decrypts an encrypted value (BSON binary of subtype 6) and returns the original BSON value.

func (*ClientEncryption) Encrypt
func (ce *ClientEncryption) Encrypt(ctx context.Context, val bson.RawValue, opts ...*options.EncryptOptions) (primitive.Binary, error)
Encrypt encrypts a BSON value with the given key and algorithm. Returns an encrypted value (BSON binary of subtype 6).

type Collection
type Collection struct {
    // contains filtered or unexported fields
}
Collection is a handle to a MongoDB collection. It is safe for concurrent use by multiple goroutines.

func (*Collection) Aggregate
func (coll *Collection) Aggregate(ctx context.Context, pipeline interface{},
    opts ...*options.AggregateOptions) (*Cursor, error)
Aggregate executes an aggregate command against the collection and returns a cursor over the resulting documents.

The pipeline parameter must be an array of documents, each representing an aggregation stage. The pipeline cannot be nil but can be empty. The stage documents must all be non-nil. For a pipeline of bson.D documents, the mongo.Pipeline type can be used. See https://docs.mongodb.com/manual/reference/operator/aggregation-pipeline/#db-collection-aggregate-stages for a list of valid stages in aggregations.

The opts parameter can be used to specify options for the operation (see the options.AggregateOptions documentation.)

For more information about the command, see https://docs.mongodb.com/manual/reference/command/aggregate/.

Example
func (*Collection) BulkWrite
func (coll *Collection) BulkWrite(ctx context.Context, models []WriteModel,
    opts ...*options.BulkWriteOptions) (*BulkWriteResult, error)
BulkWrite performs a bulk write operation (https://docs.mongodb.com/manual/core/bulk-write-operations/).

The models parameter must be a slice of operations to be executed in this bulk write. It cannot be nil or empty. All of the models must be non-nil. See the mongo.WriteModel documentation for a list of valid model types and examples of how they should be used.

The opts parameter can be used to specify options for the operation (see the options.BulkWriteOptions documentation.)

Example
func (*Collection) Clone
func (coll *Collection) Clone(opts ...*options.CollectionOptions) (*Collection, error)
Clone creates a copy of the Collection configured with the given CollectionOptions. The specified options are merged with the existing options on the collection, with the specified options taking precedence.

func (*Collection) CountDocuments
func (coll *Collection) CountDocuments(ctx context.Context, filter interface{},
    opts ...*options.CountOptions) (int64, error)
CountDocuments returns the number of documents in the collection. For a fast count of the documents in the collection, see the EstimatedDocumentCount method.

The filter parameter must be a document and can be used to select which documents contribute to the count. It cannot be nil. An empty document (e.g. bson.D{}) should be used to count all documents in the collection. This will result in a full collection scan.

The opts parameter can be used to specify options for the operation (see the options.CountOptions documentation).

Example
func (*Collection) Database
func (coll *Collection) Database() *Database
Database returns the Database that was used to create the Collection.

func (*Collection) DeleteMany
func (coll *Collection) DeleteMany(ctx context.Context, filter interface{},
    opts ...*options.DeleteOptions) (*DeleteResult, error)
DeleteMany executes a delete command to delete documents from the collection.

The filter parameter must be a document containing query operators and can be used to select the documents to be deleted. It cannot be nil. An empty document (e.g. bson.D{}) should be used to delete all documents in the collection. If the filter does not match any documents, the operation will succeed and a DeleteResult with a DeletedCount of 0 will be returned.

The opts parameter can be used to specify options for the operation (see the options.DeleteOptions documentation).

For more information about the command, see https://docs.mongodb.com/manual/reference/command/delete/.

Example
func (*Collection) DeleteOne
func (coll *Collection) DeleteOne(ctx context.Context, filter interface{},
    opts ...*options.DeleteOptions) (*DeleteResult, error)
DeleteOne executes a delete command to delete at most one document from the collection.

The filter parameter must be a document containing query operators and can be used to select the document to be deleted. It cannot be nil. If the filter does not match any documents, the operation will succeed and a DeleteResult with a DeletedCount of 0 will be returned. If the filter matches multiple documents, one will be selected from the matched set.

The opts parameter can be used to specify options for the operation (see the options.DeleteOptions documentation).

For more information about the command, see https://docs.mongodb.com/manual/reference/command/delete/.

Example
func (*Collection) Distinct
func (coll *Collection) Distinct(ctx context.Context, fieldName string, filter interface{},
    opts ...*options.DistinctOptions) ([]interface{}, error)
Distinct executes a distinct command to find the unique values for a specified field in the collection.

The fieldName parameter specifies the field name for which distinct values should be returned.

The filter parameter must be a document containing query operators and can be used to select which documents are considered. It cannot be nil. An empty document (e.g. bson.D{}) should be used to select all documents.

The opts parameter can be used to specify options for the operation (see the options.DistinctOptions documentation).

For more information about the command, see https://docs.mongodb.com/manual/reference/command/distinct/.

Example
func (*Collection) Drop
func (coll *Collection) Drop(ctx context.Context) error
Drop drops the collection on the server. This method ignores "namespace not found" errors so it is safe to drop a collection that does not exist on the server.

func (*Collection) EstimatedDocumentCount
func (coll *Collection) EstimatedDocumentCount(ctx context.Context,
    opts ...*options.EstimatedDocumentCountOptions) (int64, error)
EstimatedDocumentCount executes a count command and returns an estimate of the number of documents in the collection using collection metadata.

The opts parameter can be used to specify options for the operation (see the options.EstimatedDocumentCountOptions documentation).

For more information about the command, see https://docs.mongodb.com/manual/reference/command/count/.

Example
func (*Collection) Find
func (coll *Collection) Find(ctx context.Context, filter interface{},
    opts ...*options.FindOptions) (*Cursor, error)
Find executes a find command and returns a Cursor over the matching documents in the collection.

The filter parameter must be a document containing query operators and can be used to select which documents are included in the result. It cannot be nil. An empty document (e.g. bson.D{}) should be used to include all documents.

The opts parameter can be used to specify options for the operation (see the options.FindOptions documentation).

For more information about the command, see https://docs.mongodb.com/manual/reference/command/find/.

Example
func (*Collection) FindOne
func (coll *Collection) FindOne(ctx context.Context, filter interface{},
    opts ...*options.FindOneOptions) *SingleResult
FindOne executes a find command and returns a SingleResult for one document in the collection.

The filter parameter must be a document containing query operators and can be used to select the document to be returned. It cannot be nil. If the filter does not match any documents, a SingleResult with an error set to ErrNoDocuments will be returned. If the filter matches multiple documents, one will be selected from the matched set.

The opts parameter can be used to specify options for this operation (see the options.FindOneOptions documentation).

For more information about the command, see https://docs.mongodb.com/manual/reference/command/find/.

Example
func (*Collection) FindOneAndDelete
func (coll *Collection) FindOneAndDelete(ctx context.Context, filter interface{},
    opts ...*options.FindOneAndDeleteOptions) *SingleResult
FindOneAndDelete executes a findAndModify command to delete at most one document in the collection. and returns the document as it appeared before deletion.

The filter parameter must be a document containing query operators and can be used to select the document to be deleted. It cannot be nil. If the filter does not match any documents, a SingleResult with an error set to ErrNoDocuments wil be returned. If the filter matches multiple documents, one will be selected from the matched set.

The opts parameter can be used to specify options for the operation (see the options.FindOneAndDeleteOptions documentation).

For more information about the command, see https://docs.mongodb.com/manual/reference/command/findAndModify/.

Example
func (*Collection) FindOneAndReplace
func (coll *Collection) FindOneAndReplace(ctx context.Context, filter interface{},
    replacement interface{}, opts ...*options.FindOneAndReplaceOptions) *SingleResult
FindOneAndReplace executes a findAndModify command to replace at most one document in the collection and returns the document as it appeared before replacement.

The filter parameter must be a document containing query operators and can be used to select the document to be replaced. It cannot be nil. If the filter does not match any documents, a SingleResult with an error set to ErrNoDocuments wil be returned. If the filter matches multiple documents, one will be selected from the matched set.

The replacement parameter must be a document that will be used to replace the selected document. It cannot be nil and cannot contain any update operators (https://docs.mongodb.com/manual/reference/operator/update/).

The opts parameter can be used to specify options for the operation (see the options.FindOneAndReplaceOptions documentation).

For more information about the command, see https://docs.mongodb.com/manual/reference/command/findAndModify/.

Example
func (*Collection) FindOneAndUpdate
func (coll *Collection) FindOneAndUpdate(ctx context.Context, filter interface{},
    update interface{}, opts ...*options.FindOneAndUpdateOptions) *SingleResult
FindOneAndUpdate executes a findAndModify command to update at most one document in the collection and returns the document as it appeared before updating.

The filter parameter must be a document containing query operators and can be used to select the document to be updated. It cannot be nil. If the filter does not match any documents, a SingleResult with an error set to ErrNoDocuments wil be returned. If the filter matches multiple documents, one will be selected from the matched set.

The update parameter must be a document containing update operators (https://docs.mongodb.com/manual/reference/operator/update/) and can be used to specify the modifications to be made to the selected document. It cannot be nil or empty.

The opts parameter can be used to specify options for the operation (see the options.FindOneAndUpdateOptions documentation).

For more information about the command, see https://docs.mongodb.com/manual/reference/command/findAndModify/.

Example
func (*Collection) Indexes
func (coll *Collection) Indexes() IndexView
Indexes returns an IndexView instance that can be used to perform operations on the indexes for the collection.

func (*Collection) InsertMany
func (coll *Collection) InsertMany(ctx context.Context, documents []interface{},
    opts ...*options.InsertManyOptions) (*InsertManyResult, error)
InsertMany executes an insert command to insert multiple documents into the collection. If write errors occur during the operation (e.g. duplicate key error), this method returns a BulkWriteException error.

The documents parameter must be a slice of documents to insert. The slice cannot be nil or empty. The elements must all be non-nil. For any document that does not have an _id field when transformed into BSON, one will be added automatically to the marshalled document. The original document will not be modified. The _id values for the inserted documents can be retrieved from the InsertedIDs field of the returnd InsertManyResult.

The opts parameter can be used to specify options for the operation (see the options.InsertManyOptions documentation.)

For more information about the command, see https://docs.mongodb.com/manual/reference/command/insert/.

Example
func (*Collection) InsertOne
func (coll *Collection) InsertOne(ctx context.Context, document interface{},
    opts ...*options.InsertOneOptions) (*InsertOneResult, error)
InsertOne executes an insert command to insert a single document into the collection.

The document parameter must be the document to be inserted. It cannot be nil. If the document does not have an _id field when transformed into BSON, one will be added automatically to the marshalled document. The original document will not be modified. The _id can be retrieved from the InsertedID field of the returned InsertOneResult.

The opts parameter can be used to specify options for the operation (see the options.InsertOneOptions documentation.)

For more information about the command, see https://docs.mongodb.com/manual/reference/command/insert/.

Example
func (*Collection) Name
func (coll *Collection) Name() string
Name returns the name of the collection.

func (*Collection) ReplaceOne
func (coll *Collection) ReplaceOne(ctx context.Context, filter interface{},
    replacement interface{}, opts ...*options.ReplaceOptions) (*UpdateResult, error)
ReplaceOne executes an update command to replace at most one document in the collection.

The filter parameter must be a document containing query operators and can be used to select the document to be replaced. It cannot be nil. If the filter does not match any documents, the operation will succeed and an UpdateResult with a MatchedCount of 0 will be returned. If the filter matches multiple documents, one will be selected from the matched set and MatchedCount will equal 1.

The replacement parameter must be a document that will be used to replace the selected document. It cannot be nil and cannot contain any update operators (https://docs.mongodb.com/manual/reference/operator/update/).

The opts parameter can be used to specify options for the operation (see the options.ReplaceOptions documentation).

For more information about the command, see https://docs.mongodb.com/manual/reference/command/update/.

Example
func (*Collection) UpdateByID
func (coll *Collection) UpdateByID(ctx context.Context, id interface{}, update interface{},
    opts ...*options.UpdateOptions) (*UpdateResult, error)
UpdateByID executes an update command to update the document whose _id value matches the provided ID in the collection. This is equivalent to running UpdateOne(ctx, bson.D{{"_id", id}}, update, opts...).

The id parameter is the _id of the document to be updated. It cannot be nil. If the ID does not match any documents, the operation will succeed and an UpdateResult with a MatchedCount of 0 will be returned.

The update parameter must be a document containing update operators (https://docs.mongodb.com/manual/reference/operator/update/) and can be used to specify the modifications to be made to the selected document. It cannot be nil or empty.

The opts parameter can be used to specify options for the operation (see the options.UpdateOptions documentation).

For more information about the command, see https://docs.mongodb.com/manual/reference/command/update/.

func (*Collection) UpdateMany
func (coll *Collection) UpdateMany(ctx context.Context, filter interface{}, update interface{},
    opts ...*options.UpdateOptions) (*UpdateResult, error)
UpdateMany executes an update command to update documents in the collection.

The filter parameter must be a document containing query operators and can be used to select the documents to be updated. It cannot be nil. If the filter does not match any documents, the operation will succeed and an UpdateResult with a MatchedCount of 0 will be returned.

The update parameter must be a document containing update operators (https://docs.mongodb.com/manual/reference/operator/update/) and can be used to specify the modifications to be made to the selected documents. It cannot be nil or empty.

The opts parameter can be used to specify options for the operation (see the options.UpdateOptions documentation).

For more information about the command, see https://docs.mongodb.com/manual/reference/command/update/.

Example
func (*Collection) UpdateOne
func (coll *Collection) UpdateOne(ctx context.Context, filter interface{}, update interface{},
    opts ...*options.UpdateOptions) (*UpdateResult, error)
UpdateOne executes an update command to update at most one document in the collection.

The filter parameter must be a document containing query operators and can be used to select the document to be updated. It cannot be nil. If the filter does not match any documents, the operation will succeed and an UpdateResult with a MatchedCount of 0 will be returned. If the filter matches multiple documents, one will be selected from the matched set and MatchedCount will equal 1.

The update parameter must be a document containing update operators (https://docs.mongodb.com/manual/reference/operator/update/) and can be used to specify the modifications to be made to the selected document. It cannot be nil or empty.

The opts parameter can be used to specify options for the operation (see the options.UpdateOptions documentation).

For more information about the command, see https://docs.mongodb.com/manual/reference/command/update/.

Example
func (*Collection) Watch
func (coll *Collection) Watch(ctx context.Context, pipeline interface{},
    opts ...*options.ChangeStreamOptions) (*ChangeStream, error)
Watch returns a change stream for all changes on the corresponding collection. See https://docs.mongodb.com/manual/changeStreams/ for more information about change streams.

The Collection must be configured with read concern majority or no read concern for a change stream to be created successfully.

The pipeline parameter must be an array of documents, each representing a pipeline stage. The pipeline cannot be nil but can be empty. The stage documents must all be non-nil. See https://docs.mongodb.com/manual/changeStreams/ for a list of pipeline stages that can be used with change streams. For a pipeline of bson.D documents, the mongo.Pipeline{} type can be used.

The opts parameter can be used to specify options for change stream creation (see the options.ChangeStreamOptions documentation).

Example
type CollectionSpecification
type CollectionSpecification struct {
    // The collection name.
    Name string

    // The type of the collection. This will either be "collection" or "view".
    Type string

    // Whether or not the collection is readOnly. This will be false for MongoDB versions < 3.4.
    ReadOnly bool

    // The collection UUID. This field will be nil for MongoDB versions < 3.6. For versions 3.6 and higher, this will
    // be a primitive.Binary with Subtype 4.
    UUID *primitive.Binary

    // A document containing the options used to construct the collection.
    Options bson.Raw

    // An IndexSpecification instance with details about the collection's _id index. This will be nil if the NameOnly
    // option is used and for MongoDB versions < 3.4.
    IDIndex *IndexSpecification
}
CollectionSpecification represents a collection in a database. This type is returned by the Database.ListCollectionSpecifications function.

func (*CollectionSpecification) UnmarshalBSON
func (cs *CollectionSpecification) UnmarshalBSON(data []byte) error
UnmarshalBSON implements the bson.Unmarshaler interface.

type CommandError
type CommandError struct {
    Code    int32
    Message string
    Labels  []string // Categories to which the error belongs
    Name    string   // A human-readable name corresponding to the error code
    Wrapped error    // The underlying error, if one exists.
}
CommandError represents a server error during execution of a command. This can be returned by any operation.

func (CommandError) Error
func (e CommandError) Error() string
Error implements the error interface.

func (CommandError) HasErrorCode
func (e CommandError) HasErrorCode(code int) bool
HasErrorCode returns true if the error has the specified code.

func (CommandError) HasErrorCodeWithMessage
func (e CommandError) HasErrorCodeWithMessage(code int, message string) bool
HasErrorCodeWithMessage returns true if the error has the specified code and Message contains the specified message.

func (CommandError) HasErrorLabel
func (e CommandError) HasErrorLabel(label string) bool
HasErrorLabel returns true if the error contains the specified label.

func (CommandError) HasErrorMessage
func (e CommandError) HasErrorMessage(message string) bool
HasErrorMessage returns true if the error contains the specified message.

func (CommandError) IsMaxTimeMSExpiredError
func (e CommandError) IsMaxTimeMSExpiredError() bool
IsMaxTimeMSExpiredError returns true if the error is a MaxTimeMSExpired error.

func (CommandError) Unwrap
func (e CommandError) Unwrap() error
Unwrap returns the underlying error.

type Cursor
type Cursor struct {
    // Current contains the BSON bytes of the current change document. This property is only valid until the next call
    // to Next or TryNext. If continued access is required, a copy must be made.
    Current bson.Raw
    // contains filtered or unexported fields
}
Cursor is used to iterate over a stream of documents. Each document can be decoded into a Go type via the Decode method or accessed as raw BSON via the Current field.

func (*Cursor) All
func (c *Cursor) All(ctx context.Context, results interface{}) error
All iterates the cursor and decodes each document into results. The results parameter must be a pointer to a slice. The slice pointed to by results will be completely overwritten. This method will close the cursor after retrieving all documents. If the cursor has been iterated, any previously iterated documents will not be included in results.

This method requires driver version >= 1.1.0.

Example
func (*Cursor) Close
func (c *Cursor) Close(ctx context.Context) error
Close closes this cursor. Next and TryNext must not be called after Close has been called. Close is idempotent. After the first call, any subsequent calls will not change the state.

func (*Cursor) Decode
func (c *Cursor) Decode(val interface{}) error
Decode will unmarshal the current document into val and return any errors from the unmarshalling process without any modification. If val is nil or is a typed nil, an error will be returned.

func (*Cursor) Err
func (c *Cursor) Err() error
Err returns the last error seen by the Cursor, or nil if no error has occurred.

func (*Cursor) ID
func (c *Cursor) ID() int64
ID returns the ID of this cursor, or 0 if the cursor has been closed or exhausted.

func (*Cursor) Next
func (c *Cursor) Next(ctx context.Context) bool
Next gets the next document for this cursor. It returns true if there were no errors and the cursor has not been exhausted.

Next blocks until a document is available, an error occurs, or ctx expires. If ctx expires, the error will be set to ctx.Err(). In an error case, Next will return false.

If Next returns false, subsequent calls will also return false.

Example
func (*Cursor) RemainingBatchLength
func (c *Cursor) RemainingBatchLength() int
RemainingBatchLength returns the number of documents left in the current batch. If this returns zero, the subsequent call to Next or TryNext will do a network request to fetch the next batch.

Example
func (*Cursor) TryNext
func (c *Cursor) TryNext(ctx context.Context) bool
TryNext attempts to get the next document for this cursor. It returns true if there were no errors and the next document is available. This is only recommended for use with tailable cursors as a non-blocking alternative to Next. See https://docs.mongodb.com/manual/core/tailable-cursors/ for more information about tailable cursors.

TryNext returns false if the cursor is exhausted, an error occurs when getting results from the server, the next document is not yet available, or ctx expires. If ctx expires, the error will be set to ctx.Err().

If TryNext returns false and an error occurred or the cursor has been exhausted (i.e. c.Err() != nil || c.ID() == 0), subsequent attempts will also return false. Otherwise, it is safe to call TryNext again until a document is available.

This method requires driver version >= 1.2.0.

Example
type Database
type Database struct {
    // contains filtered or unexported fields
}
Database is a handle to a MongoDB database. It is safe for concurrent use by multiple goroutines.

func (*Database) Aggregate
func (db *Database) Aggregate(ctx context.Context, pipeline interface{},
    opts ...*options.AggregateOptions) (*Cursor, error)
Aggregate executes an aggregate command the database. This requires MongoDB version >= 3.6 and driver version >= 1.1.0.

The pipeline parameter must be a slice of documents, each representing an aggregation stage. The pipeline cannot be nil but can be empty. The stage documents must all be non-nil. For a pipeline of bson.D documents, the mongo.Pipeline type can be used. See https://docs.mongodb.com/manual/reference/operator/aggregation-pipeline/#db-aggregate-stages for a list of valid stages in database-level aggregations.

The opts parameter can be used to specify options for this operation (see the options.AggregateOptions documentation).

For more information about the command, see https://docs.mongodb.com/manual/reference/command/aggregate/.

func (*Database) Client
func (db *Database) Client() *Client
Client returns the Client the Database was created from.

func (*Database) Collection
func (db *Database) Collection(name string, opts ...*options.CollectionOptions) *Collection
Collection gets a handle for a collection with the given name configured with the given CollectionOptions.

func (*Database) CreateCollection
func (db *Database) CreateCollection(ctx context.Context, name string, opts ...*options.CreateCollectionOptions) error
CreateCollection executes a create command to explicitly create a new collection with the specified name on the server. If the collection being created already exists, this method will return a mongo.CommandError. This method requires driver version 1.4.0 or higher.

The opts parameter can be used to specify options for the operation (see the options.CreateCollectionOptions documentation).

Example
func (*Database) CreateView
func (db *Database) CreateView(ctx context.Context, viewName, viewOn string, pipeline interface{},
    opts ...*options.CreateViewOptions) error
CreateView executes a create command to explicitly create a view on the server. See https://docs.mongodb.com/manual/core/views/ for more information about views. This method requires driver version >= 1.4.0 and MongoDB version >= 3.4.

The viewName parameter specifies the name of the view to create.

The viewOn parameter specifies the name of the collection or view on which this view will be created
The pipeline parameter specifies an aggregation pipeline that will be exececuted against the source collection or view to create this view.

The opts parameter can be used to specify options for the operation (see the options.CreateViewOptions documentation).

Example
func (*Database) Drop
func (db *Database) Drop(ctx context.Context) error
Drop drops the database on the server. This method ignores "namespace not found" errors so it is safe to drop a database that does not exist on the server.

func (*Database) ListCollectionNames
func (db *Database) ListCollectionNames(ctx context.Context, filter interface{}, opts ...*options.ListCollectionsOptions) ([]string, error)
ListCollectionNames executes a listCollections command and returns a slice containing the names of the collections in the database. This method requires driver version >= 1.1.0.

The filter parameter must be a document containing query operators and can be used to select which collections are included in the result. It cannot be nil. An empty document (e.g. bson.D{}) should be used to include all collections.

The opts parameter can be used to specify options for the operation (see the options.ListCollectionsOptions documentation).

For more information about the command, see https://docs.mongodb.com/manual/reference/command/listCollections/.

Example
func (*Database) ListCollectionSpecifications
func (db *Database) ListCollectionSpecifications(ctx context.Context, filter interface{},
    opts ...*options.ListCollectionsOptions) ([]*CollectionSpecification, error)
ListCollectionSpecifications executes a listCollections command and returns a slice of CollectionSpecification instances representing the collections in the database.

The filter parameter must be a document containing query operators and can be used to select which collections are included in the result. It cannot be nil. An empty document (e.g. bson.D{}) should be used to include all collections.

The opts parameter can be used to specify options for the operation (see the options.ListCollectionsOptions documentation).

For more information about the command, see https://docs.mongodb.com/manual/reference/command/listCollections/.

func (*Database) ListCollections
func (db *Database) ListCollections(ctx context.Context, filter interface{}, opts ...*options.ListCollectionsOptions) (*Cursor, error)
ListCollections executes a listCollections command and returns a cursor over the collections in the database.

The filter parameter must be a document containing query operators and can be used to select which collections are included in the result. It cannot be nil. An empty document (e.g. bson.D{}) should be used to include all collections.

The opts parameter can be used to specify options for the operation (see the options.ListCollectionsOptions documentation).

For more information about the command, see https://docs.mongodb.com/manual/reference/command/listCollections/.

func (*Database) Name
func (db *Database) Name() string
Name returns the name of the database.

func (*Database) ReadConcern
func (db *Database) ReadConcern() *readconcern.ReadConcern
ReadConcern returns the read concern used to configure the Database object.

func (*Database) ReadPreference
func (db *Database) ReadPreference() *readpref.ReadPref
ReadPreference returns the read preference used to configure the Database object.

func (*Database) RunCommand
func (db *Database) RunCommand(ctx context.Context, runCommand interface{}, opts ...*options.RunCmdOptions) *SingleResult
RunCommand executes the given command against the database. This function does not obey the Database's read preference. To specify a read preference, the RunCmdOptions.ReadPreference option must be used.

The runCommand parameter must be a document for the command to be executed. It cannot be nil. This must be an order-preserving type such as bson.D. Map types such as bson.M are not valid. If the command document contains a session ID or any transaction-specific fields, the behavior is undefined.

The opts parameter can be used to specify options for this operation (see the options.RunCmdOptions documentation).

Example
func (*Database) RunCommandCursor
func (db *Database) RunCommandCursor(ctx context.Context, runCommand interface{}, opts ...*options.RunCmdOptions) (*Cursor, error)
RunCommandCursor executes the given command against the database and parses the response as a cursor. If the command being executed does not return a cursor (e.g. insert), the command will be executed on the server and an error will be returned because the server response cannot be parsed as a cursor. This function does not obey the Database's read preference. To specify a read preference, the RunCmdOptions.ReadPreference option must be used.

The runCommand parameter must be a document for the command to be executed. It cannot be nil. This must be an order-preserving type such as bson.D. Map types such as bson.M are not valid. If the command document contains a session ID or any transaction-specific fields, the behavior is undefined.

The opts parameter can be used to specify options for this operation (see the options.RunCmdOptions documentation).

func (*Database) Watch
func (db *Database) Watch(ctx context.Context, pipeline interface{},
    opts ...*options.ChangeStreamOptions) (*ChangeStream, error)
Watch returns a change stream for all changes to the corresponding database. See https://docs.mongodb.com/manual/changeStreams/ for more information about change streams.

The Database must be configured with read concern majority or no read concern for a change stream to be created successfully.

The pipeline parameter must be a slice of documents, each representing a pipeline stage. The pipeline cannot be nil but can be empty. The stage documents must all be non-nil. See https://docs.mongodb.com/manual/changeStreams/ for a list of pipeline stages that can be used with change streams. For a pipeline of bson.D documents, the mongo.Pipeline{} type can be used.

The opts parameter can be used to specify options for change stream creation (see the options.ChangeStreamOptions documentation).

Example
func (*Database) WriteConcern
func (db *Database) WriteConcern() *writeconcern.WriteConcern
WriteConcern returns the write concern used to configure the Database object.

type DatabaseSpecification
type DatabaseSpecification struct {
    Name       string // The name of the database.
    SizeOnDisk int64  // The total size of the database files on disk in bytes.
    Empty      bool   // Specfies whether or not the database is empty.
}
DatabaseSpecification contains information for a database. This type is returned as part of ListDatabasesResult.

type DeleteManyModel
type DeleteManyModel struct {
    Filter    interface{}
    Collation *options.Collation
    Hint      interface{}
}
DeleteManyModel is used to delete multiple documents in a BulkWrite operation.

func NewDeleteManyModel
func NewDeleteManyModel() *DeleteManyModel
NewDeleteManyModel creates a new DeleteManyModel.

func (*DeleteManyModel) SetCollation
func (dmm *DeleteManyModel) SetCollation(collation *options.Collation) *DeleteManyModel
SetCollation specifies a collation to use for string comparisons. The default is nil, meaning no collation will be used.

func (*DeleteManyModel) SetFilter
func (dmm *DeleteManyModel) SetFilter(filter interface{}) *DeleteManyModel
SetFilter specifies a filter to use to select documents to delete. The filter must be a document containing query operators. It cannot be nil.

func (*DeleteManyModel) SetHint
func (dmm *DeleteManyModel) SetHint(hint interface{}) *DeleteManyModel
SetHint specifies the index to use for the operation. This should either be the index name as a string or the index specification as a document. This option is only valid for MongoDB versions >= 4.4. Server versions >= 3.4 will return an error if this option is specified. For server versions < 3.4, the driver will return a client-side error if this option is specified. The driver will return an error if this option is specified during an unacknowledged write operation. The default value is nil, which means that no hint will be sent.

type DeleteOneModel
type DeleteOneModel struct {
    Filter    interface{}
    Collation *options.Collation
    Hint      interface{}
}
DeleteOneModel is used to delete at most one document in a BulkWriteOperation.

func NewDeleteOneModel
func NewDeleteOneModel() *DeleteOneModel
NewDeleteOneModel creates a new DeleteOneModel.

func (*DeleteOneModel) SetCollation
func (dom *DeleteOneModel) SetCollation(collation *options.Collation) *DeleteOneModel
SetCollation specifies a collation to use for string comparisons. The default is nil, meaning no collation will be used.

func (*DeleteOneModel) SetFilter
func (dom *DeleteOneModel) SetFilter(filter interface{}) *DeleteOneModel
SetFilter specifies a filter to use to select the document to delete. The filter must be a document containing query operators. It cannot be nil. If the filter matches multiple documents, one will be selected from the matching documents.

func (*DeleteOneModel) SetHint
func (dom *DeleteOneModel) SetHint(hint interface{}) *DeleteOneModel
SetHint specifies the index to use for the operation. This should either be the index name as a string or the index specification as a document. This option is only valid for MongoDB versions >= 4.4. Server versions >= 3.4 will return an error if this option is specified. For server versions < 3.4, the driver will return a client-side error if this option is specified. The driver will return an error if this option is specified during an unacknowledged write operation. The default value is nil, which means that no hint will be sent.

type DeleteResult
type DeleteResult struct {
    DeletedCount int64 `bson:"n"` // The number of documents deleted.
}
DeleteResult is the result type returned by DeleteOne and DeleteMany operations.

type Dialer
type Dialer interface {
    DialContext(ctx context.Context, network, address string) (net.Conn, error)
}
Dialer is used to make network connections.

type EncryptionKeyVaultError
type EncryptionKeyVaultError struct {
    Wrapped error
}
EncryptionKeyVaultError represents an error while communicating with the key vault collection during client-side encryption.

func (EncryptionKeyVaultError) Error
func (ekve EncryptionKeyVaultError) Error() string
Error implements the error interface.

func (EncryptionKeyVaultError) Unwrap
func (ekve EncryptionKeyVaultError) Unwrap() error
Unwrap returns the underlying error.

type IndexModel
type IndexModel struct {
    // A document describing which keys should be used for the index. It cannot be nil. This must be an order-preserving
    // type such as bson.D. Map types such as bson.M are not valid. See https://docs.mongodb.com/manual/indexes/#indexes
    // for examples of valid documents.
    Keys interface{}

    // The options to use to create the index.
    Options *options.IndexOptions
}
IndexModel represents a new index to be created.

type IndexOptionsBuilder
type IndexOptionsBuilder struct {
    // contains filtered or unexported fields
}
IndexOptionsBuilder specifies options for a new index.

Deprecated: Use the IndexOptions type in the mongo/options package instead.

func NewIndexOptionsBuilder
func NewIndexOptionsBuilder() *IndexOptionsBuilder
NewIndexOptionsBuilder creates a new IndexOptionsBuilder.

Deprecated: Use the Index function in mongo/options instead.

func (*IndexOptionsBuilder) Background
func (iob *IndexOptionsBuilder) Background(background bool) *IndexOptionsBuilder
Background specifies a value for the background option.

Deprecated: Use the IndexOptions.SetBackground function in mongo/options instead.

func (*IndexOptionsBuilder) Bits
func (iob *IndexOptionsBuilder) Bits(bits int32) *IndexOptionsBuilder
Bits specifies a value for the bits option.

Deprecated: Use the IndexOptions.SetBits function in mongo/options instead.

func (*IndexOptionsBuilder) BucketSize
func (iob *IndexOptionsBuilder) BucketSize(bucketSize int32) *IndexOptionsBuilder
BucketSize specifies a value for the bucketSize option.

Deprecated: Use the IndexOptions.SetBucketSize function in mongo/options instead.

func (*IndexOptionsBuilder) Build
func (iob *IndexOptionsBuilder) Build() bson.D
Build finishes constructing an the builder.

Deprecated: Use the IndexOptions type in the mongo/options package instead.

func (*IndexOptionsBuilder) Collation
func (iob *IndexOptionsBuilder) Collation(collation interface{}) *IndexOptionsBuilder
Collation specifies a value for the collation option.

Deprecated: Use the IndexOptions.SetCollation function in mongo/options instead.

func (*IndexOptionsBuilder) DefaultLanguage
func (iob *IndexOptionsBuilder) DefaultLanguage(defaultLanguage string) *IndexOptionsBuilder
DefaultLanguage specifies a value for the default_language option.

Deprecated: Use the IndexOptions.SetDefaultLanguage function in mongo/options instead.

func (*IndexOptionsBuilder) ExpireAfterSeconds
func (iob *IndexOptionsBuilder) ExpireAfterSeconds(expireAfterSeconds int32) *IndexOptionsBuilder
ExpireAfterSeconds specifies a value for the expireAfterSeconds option.

Deprecated: Use the IndexOptions.SetExpireAfterSeconds function in mongo/options instead.

func (*IndexOptionsBuilder) LanguageOverride
func (iob *IndexOptionsBuilder) LanguageOverride(languageOverride string) *IndexOptionsBuilder
LanguageOverride specifies a value for the language_override option.

Deprecated: Use the IndexOptions.SetLanguageOverride function in mongo/options instead.

func (*IndexOptionsBuilder) Max
func (iob *IndexOptionsBuilder) Max(max float64) *IndexOptionsBuilder
Max specifies a value for the max option.

Deprecated: Use the IndexOptions.SetMax function in mongo/options instead.

func (*IndexOptionsBuilder) Min
func (iob *IndexOptionsBuilder) Min(min float64) *IndexOptionsBuilder
Min specifies a value for the min option.

Deprecated: Use the IndexOptions.SetMin function in mongo/options instead.

func (*IndexOptionsBuilder) Name
func (iob *IndexOptionsBuilder) Name(name string) *IndexOptionsBuilder
Name specifies a value for the name option.

Deprecated: Use the IndexOptions.SetName function in mongo/options instead.

func (*IndexOptionsBuilder) PartialFilterExpression
func (iob *IndexOptionsBuilder) PartialFilterExpression(partialFilterExpression interface{}) *IndexOptionsBuilder
PartialFilterExpression specifies a value for the partialFilterExpression option.

Deprecated: Use the IndexOptions.SetPartialFilterExpression function in mongo/options instead.

func (*IndexOptionsBuilder) Sparse
func (iob *IndexOptionsBuilder) Sparse(sparse bool) *IndexOptionsBuilder
Sparse specifies a value for the sparse option.

Deprecated: Use the IndexOptions.SetSparse function in mongo/options instead.

func (*IndexOptionsBuilder) SphereVersion
func (iob *IndexOptionsBuilder) SphereVersion(sphereVersion int32) *IndexOptionsBuilder
SphereVersion specifies a value for the 2dsphereIndexVersion option.

Deprecated: Use the IndexOptions.SetSphereVersion function in mongo/options instead.

func (*IndexOptionsBuilder) StorageEngine
func (iob *IndexOptionsBuilder) StorageEngine(storageEngine interface{}) *IndexOptionsBuilder
StorageEngine specifies a value for the storageEngine option.

Deprecated: Use the IndexOptions.SetStorageEngine function in mongo/options instead.

func (*IndexOptionsBuilder) TextVersion
func (iob *IndexOptionsBuilder) TextVersion(textVersion int32) *IndexOptionsBuilder
TextVersion specifies a value for the textIndexVersion option.

Deprecated: Use the IndexOptions.SetTextVersion function in mongo/options instead.

func (*IndexOptionsBuilder) Unique
func (iob *IndexOptionsBuilder) Unique(unique bool) *IndexOptionsBuilder
Unique specifies a value for the unique option.

Deprecated: Use the IndexOptions.SetUnique function in mongo/options instead.

func (*IndexOptionsBuilder) Version
func (iob *IndexOptionsBuilder) Version(version int32) *IndexOptionsBuilder
Version specifies a value for the version option.

Deprecated: Use the IndexOptions.SetVersion function in mongo/options instead.

func (*IndexOptionsBuilder) Weights
func (iob *IndexOptionsBuilder) Weights(weights interface{}) *IndexOptionsBuilder
Weights specifies a value for the weights option.

Deprecated: Use the IndexOptions.SetWeights function in mongo/options instead.

type IndexSpecification
type IndexSpecification struct {
    // The index name.
    Name string

    // The namespace for the index. This is a string in the format "databaseName.collectionName".
    Namespace string

    // The keys specification document for the index.
    KeysDocument bson.Raw

    // The index version.
    Version int32
}
IndexSpecification represents an index in a database. This type is returned by the IndexView.ListSpecifications function and is also used in the CollectionSpecification type.

func (*IndexSpecification) UnmarshalBSON
func (i *IndexSpecification) UnmarshalBSON(data []byte) error
UnmarshalBSON implements the bson.Unmarshaler interface.

type IndexView
type IndexView struct {
    // contains filtered or unexported fields
}
IndexView is a type that can be used to create, drop, and list indexes on a collection. An IndexView for a collection can be created by a call to Collection.Indexes().

func (IndexView) CreateMany
func (iv IndexView) CreateMany(ctx context.Context, models []IndexModel, opts ...*options.CreateIndexesOptions) ([]string, error)
CreateMany executes a createIndexes command to create multiple indexes on the collection and returns the names of the new indexes.

For each IndexModel in the models parameter, the index name can be specified via the Options field. If a name is not given, it will be generated from the Keys document.

The opts parameter can be used to specify options for this operation (see the options.CreateIndexesOptions documentation).

For more information about the command, see https://docs.mongodb.com/manual/reference/command/createIndexes/.

Example
func (IndexView) CreateOne
func (iv IndexView) CreateOne(ctx context.Context, model IndexModel, opts ...*options.CreateIndexesOptions) (string, error)
CreateOne executes a createIndexes command to create an index on the collection and returns the name of the new index. See the IndexView.CreateMany documentation for more information and an example.

func (IndexView) DropAll
func (iv IndexView) DropAll(ctx context.Context, opts ...*options.DropIndexesOptions) (bson.Raw, error)
DropAll executes a dropIndexes operation to drop all indexes on the collection. If the operation succeeds, this returns a BSON document in the form {nIndexesWas: <int32>}. The "nIndexesWas" field in the response contains the number of indexes that existed prior to the drop.

The opts parameter can be used to specify options for this operation (see the options.DropIndexesOptions documentation).

For more information about the command, see https://docs.mongodb.com/manual/reference/command/dropIndexes/.

func (IndexView) DropOne
func (iv IndexView) DropOne(ctx context.Context, name string, opts ...*options.DropIndexesOptions) (bson.Raw, error)
DropOne executes a dropIndexes operation to drop an index on the collection. If the operation succeeds, this returns a BSON document in the form {nIndexesWas: <int32>}. The "nIndexesWas" field in the response contains the number of indexes that existed prior to the drop.

The name parameter should be the name of the index to drop. If the name is "*", ErrMultipleIndexDrop will be returned without running the command because doing so would drop all indexes.

The opts parameter can be used to specify options for this operation (see the options.DropIndexesOptions documentation).

For more information about the command, see https://docs.mongodb.com/manual/reference/command/dropIndexes/.

func (IndexView) List
func (iv IndexView) List(ctx context.Context, opts ...*options.ListIndexesOptions) (*Cursor, error)
List executes a listIndexes command and returns a cursor over the indexes in the collection.

The opts parameter can be used to specify options for this operation (see the options.ListIndexesOptions documentation).

For more information about the command, see https://docs.mongodb.com/manual/reference/command/listIndexes/.

Example
func (IndexView) ListSpecifications
func (iv IndexView) ListSpecifications(ctx context.Context, opts ...*options.ListIndexesOptions) ([]*IndexSpecification, error)
ListSpecifications executes a List command and returns a slice of returned IndexSpecifications

type InsertManyResult
type InsertManyResult struct {
    // The _id values of the inserted documents. Values generated by the driver will be of type primitive.ObjectID.
    InsertedIDs []interface{}
}
InsertManyResult is a result type returned by an InsertMany operation.

type InsertOneModel
type InsertOneModel struct {
    Document interface{}
}
InsertOneModel is used to insert a single document in a BulkWrite operation.

func NewInsertOneModel
func NewInsertOneModel() *InsertOneModel
NewInsertOneModel creates a new InsertOneModel.

func (*InsertOneModel) SetDocument
func (iom *InsertOneModel) SetDocument(doc interface{}) *InsertOneModel
SetDocument specifies the document to be inserted. The document cannot be nil. If it does not have an _id field when transformed into BSON, one will be added automatically to the marshalled document. The original document will not be modified.

type InsertOneResult
type InsertOneResult struct {
    // The _id of the inserted document. A value generated by the driver will be of type primitive.ObjectID.
    InsertedID interface{}
}
InsertOneResult is the result type returned by an InsertOne operation.

type ListDatabasesResult
type ListDatabasesResult struct {
    // A slice containing one DatabaseSpecification for each database matched by the operation's filter.
    Databases []DatabaseSpecification

    // The total size of the database files of the returned databases in bytes.
    // This will be the sum of the SizeOnDisk field for each specification in Databases.
    TotalSize int64
}
ListDatabasesResult is a result of a ListDatabases operation.

type MarshalError
type MarshalError struct {
    Value interface{}
    Err   error
}
MarshalError is returned when attempting to transform a value into a document results in an error.

func (MarshalError) Error
func (me MarshalError) Error() string
Error implements the error interface.

type MongocryptError
type MongocryptError struct {
    Code    int32
    Message string
}
MongocryptError represents an libmongocrypt error during client-side encryption.

func (MongocryptError) Error
func (m MongocryptError) Error() string
Error implements the error interface.

type MongocryptdError
type MongocryptdError struct {
    Wrapped error
}
MongocryptdError represents an error while communicating with mongocryptd during client-side encryption.

func (MongocryptdError) Error
func (e MongocryptdError) Error() string
Error implements the error interface.

func (MongocryptdError) Unwrap
func (e MongocryptdError) Unwrap() error
Unwrap returns the underlying error.

type Pipeline
type Pipeline []bson.D
Pipeline is a type that makes creating aggregation pipelines easier. It is a helper and is intended for serializing to BSON.

Example usage:

mongo.Pipeline{
	{{"$group", bson.D{{"_id", "$state"}, {"totalPop", bson.D{{"$sum", "$pop"}}}}}},
	{{"$match", bson.D{{"totalPop", bson.D{{"$gte", 10*1000*1000}}}}}},
}
type ReplaceOneModel
type ReplaceOneModel struct {
    Collation   *options.Collation
    Upsert      *bool
    Filter      interface{}
    Replacement interface{}
    Hint        interface{}
}
ReplaceOneModel is used to replace at most one document in a BulkWrite operation.

func NewReplaceOneModel
func NewReplaceOneModel() *ReplaceOneModel
NewReplaceOneModel creates a new ReplaceOneModel.

func (*ReplaceOneModel) SetCollation
func (rom *ReplaceOneModel) SetCollation(collation *options.Collation) *ReplaceOneModel
SetCollation specifies a collation to use for string comparisons. The default is nil, meaning no collation will be used.

func (*ReplaceOneModel) SetFilter
func (rom *ReplaceOneModel) SetFilter(filter interface{}) *ReplaceOneModel
SetFilter specifies a filter to use to select the document to replace. The filter must be a document containing query operators. It cannot be nil. If the filter matches multiple documents, one will be selected from the matching documents.

func (*ReplaceOneModel) SetHint
func (rom *ReplaceOneModel) SetHint(hint interface{}) *ReplaceOneModel
SetHint specifies the index to use for the operation. This should either be the index name as a string or the index specification as a document. This option is only valid for MongoDB versions >= 4.2. Server versions >= 3.4 will return an error if this option is specified. For server versions < 3.4, the driver will return a client-side error if this option is specified. The driver will return an error if this option is specified during an unacknowledged write operation. The default value is nil, which means that no hint will be sent.

func (*ReplaceOneModel) SetReplacement
func (rom *ReplaceOneModel) SetReplacement(rep interface{}) *ReplaceOneModel
SetReplacement specifies a document that will be used to replace the selected document. It cannot be nil and cannot contain any update operators (https://docs.mongodb.com/manual/reference/operator/update/).

func (*ReplaceOneModel) SetUpsert
func (rom *ReplaceOneModel) SetUpsert(upsert bool) *ReplaceOneModel
SetUpsert specifies whether or not the replacement document should be inserted if no document matching the filter is found. If an upsert is performed, the _id of the upserted document can be retrieved from the UpsertedIDs field of the BulkWriteResult.

type ServerError
type ServerError interface {
    error
    // HasErrorCode returns true if the error has the specified code.
    HasErrorCode(int) bool
    // HasErrorLabel returns true if the error contains the specified label.
    HasErrorLabel(string) bool
    // HasErrorMessage returns true if the error contains the specified message.
    HasErrorMessage(string) bool
    // HasErrorCodeWithMessage returns true if any of the contained errors have the specified code and message.
    HasErrorCodeWithMessage(int, string) bool
    // contains filtered or unexported methods
}
ServerError is the interface implemented by errors returned from the server. Custom implementations of this interface should not be used in production.

type Session
type Session interface {
    // Functions to modify session state.
    StartTransaction(...*options.TransactionOptions) error
    AbortTransaction(context.Context) error
    CommitTransaction(context.Context) error
    WithTransaction(ctx context.Context, fn func(sessCtx SessionContext) (interface{}, error),
        opts ...*options.TransactionOptions) (interface{}, error)
    EndSession(context.Context)

    // Functions to retrieve session properties.
    ClusterTime() bson.Raw
    OperationTime() *primitive.Timestamp
    Client() *Client
    ID() bson.Raw

    // Functions to modify mutable session properties.
    AdvanceClusterTime(bson.Raw) error
    AdvanceOperationTime(*primitive.Timestamp) error
    // contains filtered or unexported methods
}
Session is an interface that represents a MongoDB logical session. Sessions can be used to enable causal consistency for a group of operations or to execute operations in an ACID transaction. A new Session can be created from a Client instance. A Session created from a Client must only be used to execute operations using that Client or a Database or Collection created from that Client. Custom implementations of this interface should not be used in production. For more information about sessions, and their use cases, see https://docs.mongodb.com/manual/reference/server-sessions/, https://docs.mongodb.com/manual/core/read-isolation-consistency-recency/#causal-consistency, and https://docs.mongodb.com/manual/core/transactions/.

StartTransaction starts a new transaction, configured with the given options, on this session. This method will return an error if there is already a transaction in-progress for this session.

CommitTransaction commits the active transaction for this session. This method will return an error if there is no active transaction for this session or the transaction has been aborted.

AbortTransaction aborts the active transaction for this session. This method will return an error if there is no active transaction for this session or the transaction has been committed or aborted.

WithTransaction starts a transaction on this session and runs the fn callback. Errors with the TransientTransactionError and UnknownTransactionCommitResult labels are retried for up to 120 seconds. Inside the callback, sessCtx must be used as the Context parameter for any operations that should be part of the transaction. If the ctx parameter already has a Session attached to it, it will be replaced by this session. The fn callback may be run multiple times during WithTransaction due to retry attempts, so it must be idempotent. Non-retryable operation errors or any operation errors that occur after the timeout expires will be returned without retrying. If the callback fails, the driver will call AbortTransaction. Because this method must succeed to ensure that server-side resources are properly cleaned up, context deadlines and cancellations will not be respected during this call. For a usage example, see the Client.StartSession method documentation.

ClusterTime, OperationTime, Client, and ID return the session's current operation time, the session's current cluster time, the Client associated with the session, and the ID document associated with the session, respectively. The ID document for a session is in the form {"id": <BSON binary value>}.

EndSession method should abort any existing transactions and close the session.

AdvanceClusterTime and AdvanceOperationTime are for internal use only and must not be called.

func SessionFromContext
func SessionFromContext(ctx context.Context) Session
SessionFromContext extracts the mongo.Session object stored in a Context. This can be used on a SessionContext that was created implicitly through one of the callback-based session APIs or explicitly by calling NewSessionContext. If there is no Session stored in the provided Context, nil is returned.

type SessionContext
type SessionContext interface {
    context.Context
    Session
}
SessionContext combines the context.Context and mongo.Session interfaces. It should be used as the Context arguments to operations that should be executed in a session. This type is not goroutine safe and must not be used concurrently by multiple goroutines.

There are two ways to create a SessionContext and use it in a session/transaction. The first is to use one of the callback-based functions such as WithSession and UseSession. These functions create a SessionContext and pass it to the provided callback. The other is to use NewSessionContext to explicitly create a SessionContext.

func NewSessionContext
func NewSessionContext(ctx context.Context, sess Session) SessionContext
NewSessionContext creates a new SessionContext associated with the given Context and Session parameters.

Example
type SingleResult
type SingleResult struct {
    // contains filtered or unexported fields
}
SingleResult represents a single document returned from an operation. If the operation resulted in an error, all SingleResult methods will return that error. If the operation did not return any documents, all SingleResult methods will return ErrNoDocuments.

func (*SingleResult) Decode
func (sr *SingleResult) Decode(v interface{}) error
Decode will unmarshal the document represented by this SingleResult into v. If there was an error from the operation that created this SingleResult, that error will be returned. If the operation returned no documents, Decode will return ErrNoDocuments.

If the operation was successful and returned a document, Decode will return any errors from the unmarshalling process without any modification. If v is nil or is a typed nil, an error will be returned.

func (*SingleResult) DecodeBytes
func (sr *SingleResult) DecodeBytes() (bson.Raw, error)
DecodeBytes will return the document represented by this SingleResult as a bson.Raw. If there was an error from the operation that created this SingleResult, both the result and that error will be returned. If the operation returned no documents, this will return (nil, ErrNoDocuments).

func (*SingleResult) Err
func (sr *SingleResult) Err() error
Err returns the error from the operation that created this SingleResult. If the operation was successful but did not return any documents, Err will return ErrNoDocuments. If the operation was successful and returned a document, Err will return nil.

type StreamType
type StreamType uint8
StreamType represents the cluster type against which a ChangeStream was created.

const (
    CollectionStream StreamType = iota
    DatabaseStream
    ClientStream
)
These constants represent valid change stream types. A change stream can be initialized over a collection, all collections in a database, or over a cluster.

type UpdateManyModel
type UpdateManyModel struct {
    Collation    *options.Collation
    Upsert       *bool
    Filter       interface{}
    Update       interface{}
    ArrayFilters *options.ArrayFilters
    Hint         interface{}
}
UpdateManyModel is used to update multiple documents in a BulkWrite operation.

func NewUpdateManyModel
func NewUpdateManyModel() *UpdateManyModel
NewUpdateManyModel creates a new UpdateManyModel.

func (*UpdateManyModel) SetArrayFilters
func (umm *UpdateManyModel) SetArrayFilters(filters options.ArrayFilters) *UpdateManyModel
SetArrayFilters specifies a set of filters to determine which elements should be modified when updating an array field.

func (*UpdateManyModel) SetCollation
func (umm *UpdateManyModel) SetCollation(collation *options.Collation) *UpdateManyModel
SetCollation specifies a collation to use for string comparisons. The default is nil, meaning no collation will be used.

func (*UpdateManyModel) SetFilter
func (umm *UpdateManyModel) SetFilter(filter interface{}) *UpdateManyModel
SetFilter specifies a filter to use to select documents to update. The filter must be a document containing query operators. It cannot be nil.

func (*UpdateManyModel) SetHint
func (umm *UpdateManyModel) SetHint(hint interface{}) *UpdateManyModel
SetHint specifies the index to use for the operation. This should either be the index name as a string or the index specification as a document. This option is only valid for MongoDB versions >= 4.2. Server versions >= 3.4 will return an error if this option is specified. For server versions < 3.4, the driver will return a client-side error if this option is specified. The driver will return an error if this option is specified during an unacknowledged write operation. The default value is nil, which means that no hint will be sent.

func (*UpdateManyModel) SetUpdate
func (umm *UpdateManyModel) SetUpdate(update interface{}) *UpdateManyModel
SetUpdate specifies the modifications to be made to the selected documents. The value must be a document containing update operators (https://docs.mongodb.com/manual/reference/operator/update/). It cannot be nil or empty.

func (*UpdateManyModel) SetUpsert
func (umm *UpdateManyModel) SetUpsert(upsert bool) *UpdateManyModel
SetUpsert specifies whether or not a new document should be inserted if no document matching the filter is found. If an upsert is performed, the _id of the upserted document can be retrieved from the UpsertedIDs field of the BulkWriteResult.

type UpdateOneModel
type UpdateOneModel struct {
    Collation    *options.Collation
    Upsert       *bool
    Filter       interface{}
    Update       interface{}
    ArrayFilters *options.ArrayFilters
    Hint         interface{}
}
UpdateOneModel is used to update at most one document in a BulkWrite operation.

func NewUpdateOneModel
func NewUpdateOneModel() *UpdateOneModel
NewUpdateOneModel creates a new UpdateOneModel.

func (*UpdateOneModel) SetArrayFilters
func (uom *UpdateOneModel) SetArrayFilters(filters options.ArrayFilters) *UpdateOneModel
SetArrayFilters specifies a set of filters to determine which elements should be modified when updating an array field.

func (*UpdateOneModel) SetCollation
func (uom *UpdateOneModel) SetCollation(collation *options.Collation) *UpdateOneModel
SetCollation specifies a collation to use for string comparisons. The default is nil, meaning no collation will be used.

func (*UpdateOneModel) SetFilter
func (uom *UpdateOneModel) SetFilter(filter interface{}) *UpdateOneModel
SetFilter specifies a filter to use to select the document to update. The filter must be a document containing query operators. It cannot be nil. If the filter matches multiple documents, one will be selected from the matching documents.

func (*UpdateOneModel) SetHint
func (uom *UpdateOneModel) SetHint(hint interface{}) *UpdateOneModel
SetHint specifies the index to use for the operation. This should either be the index name as a string or the index specification as a document. This option is only valid for MongoDB versions >= 4.2. Server versions >= 3.4 will return an error if this option is specified. For server versions < 3.4, the driver will return a client-side error if this option is specified. The driver will return an error if this option is specified during an unacknowledged write operation. The default value is nil, which means that no hint will be sent.

func (*UpdateOneModel) SetUpdate
func (uom *UpdateOneModel) SetUpdate(update interface{}) *UpdateOneModel
SetUpdate specifies the modifications to be made to the selected document. The value must be a document containing update operators (https://docs.mongodb.com/manual/reference/operator/update/). It cannot be nil or empty.

func (*UpdateOneModel) SetUpsert
func (uom *UpdateOneModel) SetUpsert(upsert bool) *UpdateOneModel
SetUpsert specifies whether or not a new document should be inserted if no document matching the filter is found. If an upsert is performed, the _id of the upserted document can be retrieved from the UpsertedIDs field of the BulkWriteResult.

type UpdateResult
type UpdateResult struct {
    MatchedCount  int64       // The number of documents matched by the filter.
    ModifiedCount int64       // The number of documents modified by the operation.
    UpsertedCount int64       // The number of documents upserted by the operation.
    UpsertedID    interface{} // The _id field of the upserted document, or nil if no upsert was done.
}
UpdateResult is the result type returned from UpdateOne, UpdateMany, and ReplaceOne operations.

func (*UpdateResult) UnmarshalBSON
func (result *UpdateResult) UnmarshalBSON(b []byte) error
UnmarshalBSON implements the bson.Unmarshaler interface.

type WriteConcernError
type WriteConcernError struct {
    Name    string
    Code    int
    Message string
    Details bson.Raw
}
WriteConcernError represents a write concern failure during execution of a write operation. This error type is only returned as part of a WriteException or a BulkWriteException.

func (WriteConcernError) Error
func (wce WriteConcernError) Error() string
Error implements the error interface.

type WriteError
type WriteError struct {
    // The index of the write in the slice passed to an InsertMany or BulkWrite operation that caused this error.
    Index int

    Code    int
    Message string
}
WriteError is an error that occurred during execution of a write operation. This error type is only returned as part of a WriteException or BulkWriteException.

func (WriteError) Error
func (we WriteError) Error() string
type WriteErrors
type WriteErrors []WriteError
WriteErrors is a group of write errors that occurred during execution of a write operation.

func (WriteErrors) Error
func (we WriteErrors) Error() string
Error implements the error interface.

type WriteException
type WriteException struct {
    // The write concern error that occurred, or nil if there was none.
    WriteConcernError *WriteConcernError

    // The write errors that occurred during operation execution.
    WriteErrors WriteErrors

    // The categories to which the exception belongs.
    Labels []string
}
WriteException is the error type returned by the InsertOne, DeleteOne, DeleteMany, UpdateOne, UpdateMany, and ReplaceOne operations.

func (WriteException) Error
func (mwe WriteException) Error() string
Error implements the error interface.

func (WriteException) HasErrorCode
func (mwe WriteException) HasErrorCode(code int) bool
HasErrorCode returns true if the error has the specified code.

func (WriteException) HasErrorCodeWithMessage
func (mwe WriteException) HasErrorCodeWithMessage(code int, message string) bool
HasErrorCodeWithMessage returns true if any of the contained errors have the specified code and message.

func (WriteException) HasErrorLabel
func (mwe WriteException) HasErrorLabel(label string) bool
HasErrorLabel returns true if the error contains the specified label.

func (WriteException) HasErrorMessage
func (mwe WriteException) HasErrorMessage(message string) bool
HasErrorMessage returns true if the error contains the specified message.

type WriteModel
type WriteModel interface {
    // contains filtered or unexported methods
}
WriteModel is an interface implemented by models that can be used in a BulkWrite operation. Each WriteModel represents a write.

This interface is implemented by InsertOneModel, DeleteOneModel, DeleteManyModel, ReplaceOneModel, UpdateOneModel, and UpdateManyModel. Custom implementations of this interface must not be used.

type XSession
type XSession interface {
    ClientSession() *session.Client
}
XSession is an unstable interface for internal use only.

Deprecated: This interface is unstable because it provides access to a session.Client object, which exists in the "x" package. It should not be used by applications and may be changed or removed in any release.
</details>

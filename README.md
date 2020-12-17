

## Section 1 : Why use GO Lang as Backend :

| Attempt 1 | Attempt 2  | Attempt 3  | Attempt 4  |
| :---:   | :-: | :-: | :-: |
| <img src="/ASSETS/1.png" width=850> | <img src="/ASSETS/2.png"> | <img src="/ASSETS/3.png"> |  <img src="/ASSETS/4.png"> |


## Section 2 : GO Lang Fundamentals:
### FUNCTION
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
	defer anotherFunc()
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
defer keword in GO makes a function execute at the end of the execution (or when hits return statement) of parent function from where it is called.

SIGNATURE of FUNCTION:
```go
func function_name (argument1 type, argument2 type) (return_type1, return_type2){
	return a,b
}
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
P
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

package main

import (
	"fmt"
	"os"
	"sort"
	"time"
)

var fiboCache map[int]int

func init() {
	fiboCache = make(map[int]int)
}

func main() {
	/*
		link types in Go:

		array
		slice
		map
		pointer
		channel
	*/

	// constants

	const name string = "Alex"
	const age int8 = 30
	const weight float32 = 75.5
	fmt.Println("name:", name, "age:", age, "weight:", weight)

	// variables

	var a int = 10
	var b int = 2

	fmt.Printf("type: %T, value: %v\n", a, a)

	var result int = a + b
	// var result = a - b
	// var result = a * b
	// var result = a / b
	// var result = a % b
	fmt.Println("result:", result)

	// strings

	var greeting string = "Hello"
	fmt.Println(greeting+" "+name+"!", len(greeting))

	// pointers

	// new returns pointer
	link1 := new(int) // 0
	fmt.Printf("link1: address - %v, value - %v\n", link1, *link1)

	// get pointer from simple variable
	var link2 int
	fmt.Printf("link2: address - %v, value - %v\n", &link2, link2)

	// format

	var isDone bool = true
	fmt.Printf("%s %s! Your weight is %.2f kilo.\n", greeting, name, weight)
	fmt.Printf("%T %t\n", isDone, isDone)

	// conditions

	// if-else
	if age < 18 {
		fmt.Println("age < 18 - You are underage")
	} else if age > 70 && isDone {
		fmt.Println("You are a pensioner")
	} else {
		fmt.Println("You are an adult")
	}

	if _, err := os.Stat("./hello.go"); err != nil {
		fmt.Println("os.Stat error:", err)
	}

	// switch-case
	switch age { // switch age := 30; age {
	case 30:
		fmt.Println("You are 30")
	case 40, 50, 60:
		if age == 40 {
			fmt.Println("You are 40")
			break // exit from case
		}

		fmt.Println("You are older than 30")
		fallthrough // go to the next case
	default:
		fmt.Println("You are not 30")
	}

	// it looks like if-else
	switch {
	case age == 30:
		fmt.Println("You are 30")
	default:
		fmt.Println("You are not 30")
	}

	// cycles

	var i int8 = 0

	for i < 3 { // it looks like while
		fmt.Println("i:", i)
		i++
	}

	for j := 0; j < 3; j++ {
		fmt.Println("j:", j)
	}

	// arrays

	var ages [3]int8
	ages[0] = 10
	ages[1] = 20
	ages[2] = 30
	fmt.Println("ages[1]:", ages[1])

	ages2 := [...]int8{10, 20}
	fmt.Println("ages2 len 2:", len(ages2))

	// var weights = [4]float32{55.5, 65.5, 75.5, 85.5}
	weights := [4]float32{55.5, 65.5, 75.5, 85.5}

	for i, value := range weights {
		fmt.Println("i:", i, "value:", value)
	}

	// slices

	var slice1 []int
	// slice1 := make([]int) // len 1 by default, cap = len
	// slice1 := make([]int, 2) // len 2, cap = len
	// slice1 := make([]int, 2, 3) // len 2, cap 3
	slice1 = append(slice1, 10)
	slice1 = append(slice1, 30)
	slice1 = append(slice1, 20)
	fmt.Println("slice 1:", slice1, len(slice1), cap(slice1))

	// sort slice
	sort.Ints(slice1)
	fmt.Println("sorted slice 1:", slice1)

	slice2 := []int{100, 200, 300, 400, 500, 600}
	fmt.Println("slice 2:", slice2, len(slice2), cap(slice2), slice2[:2])
	fmt.Println("slice2[:2]", slice2[:2])
	fmt.Println("slice2[2:]", slice2[2:])
	fmt.Println("slice2[1:5]", slice2[1:5])

	// use ... operator
	slice1 = append(slice1, slice2...)
	fmt.Println("slice 1 after ...", slice1, len(slice1), cap(slice1))

	// create slice with defined cap - use append without re-creation array
	slice3 := make([]int, 0, 8) // len 0, cap 8
	fmt.Println("slice 3", slice3, len(slice3), cap(slice3))

	// create slice from array
	agesSlice := ages[:]
	fmt.Println("agesSlice", agesSlice, len(agesSlice), cap(agesSlice))

	// maps

	// sites := map[string]int{}
	sites := make(map[string]int) // len 1 by default
	// sites := make(map[string]int, 2) // len 2
	sites["google"] = 100
	sites["yandex"] = 80
	sites["duckduckgo"] = 60

	// delete key
	delete(sites, "yandex")

	/*
		// check existing key - option 1
		yandex, exists := sites["yandex"]
		if exists {
			fmt.Println("yandex:", yandex)
		} else {
			fmt.Println("yandex does not exist")
		}
	*/

	// check existing key - option 2
	if yandex, exists := sites["yandex"]; exists {
		fmt.Println("yandex:", yandex)
	} else {
		fmt.Println("yandex does not exist")
	}

	fmt.Println("sites:", sites)
	fmt.Println("google:", sites["google"])

	// iterate by map
	for k, v := range sites {
		fmt.Printf("sites Map: key - %v, value - %v\n", k, v)
	}

	// it creates new link to the same map
	// sites2 := sites

	// functions

	fmt.Println("Sum result:", Sum(a, b))
	fmt.Println("Sum all args result:", SumAll(10, 20, 30, 40, 50))

	var SumRes int
	var mulRes int
	SumRes, mulRes = SumAndMult(a, b)
	fmt.Println("Sum and mult result:", SumRes, mulRes)

	// anonymous
	func(res int) {
		fmt.Println("anonymous result:", res)
	}(SumRes)

	// closures

	mutipleB := func() int {
		return b * 1000
	}

	fmt.Println("mutiple b result:", mutipleB())

	increment := GenerateCounter()
	fmt.Println("increment result:", increment())
	fmt.Println("increment result:", increment())
	fmt.Println("increment result:", increment())

	fmt.Println("ReduceSlice with Sum:", ReduceSlice(slice2, Sum, 0))

	// deferment

	/*
		defer PrintTwo()
		PrintOne()

		timer := GetTimer()
		defer timer()
		time.Sleep(time.Second)
	*/

	// pointers (links in memory)

	var c int = 123
	ModifyInt(&c)
	fmt.Println("c was modified from 123 to", c)

	// structures

	// key : value init
	bob := Human{
		Name:   "Bob",
		Age:    40,
		Height: 175.5,
		Weight: 80.5,
	}

	fmt.Println("Bob's age is", bob.Age)
	fmt.Println("Bob's height to weight ratio is", bob.CalcWeightToHeight())

	// init with new - new returns link (pointer), a := new(int)
	alice := new(Human)
	alice.Name = "Alice"
	alice.Age = 30
	alice.Height = 165.5
	alice.Weight = 55.5

	fmt.Println("Human Alice:", alice)

	// init with inline args
	andrew := Android{Human{"Andrew", 1, 200, 200}, "Release 1"}
	andrew.TakeWorld()
	fmt.Println("Andrew is Android:", andrew)

	// type assertion

	CastWarriorToAndroid(andrew)
}

// functions

// Sum makes sum
func Sum(a int, b int) int {
	return a + b
}

// SumAll sum all args
func SumAll(args ...int) int {
	acc := 0

	for _, val := range args {
		acc += val
	}

	return acc
}

// recursion

// Fibo is recursive calculation with memoization
func Fibo(n int) int {
	if n < 2 {
		return n
	}

	if cached, exists := fiboCache[n]; exists {
		return cached
	}

	sum := Fibo(n-1) + Fibo(n-2)
	fiboCache[n] = sum

	return sum
}

// closures

// GenerateCounter generates incremental counter
func GenerateCounter() func() int {
	n := 0
	return func() int {
		n++
		return n
	}
}

// GetTimer prints time from start to end
func GetTimer() func() {
	start := time.Now()
	return func() {
		fmt.Printf("Time from start %v\n", time.Since(start))
	}
}

// SumAndMult returns sum and multiply
func SumAndMult(a int, b int) (int, int) {
	return (a + b), (a * b)
}

// PrintOne prints one
func PrintOne() {
	fmt.Println("one")
}

// PrintTwo prints two
func PrintTwo() {
	fmt.Println("two")
}

// ModifyInt modifies int by link in memory
func ModifyInt(x *int) {
	*x = 321
}

// ReduceSlice reduces integers from slice into one value
func ReduceSlice(numbers []int, callback func(int, int) int, initial int) int {
	result := initial

	for _, v := range numbers {
		result = callback(result, v)
	}

	return result
}

// structures

// Human is simple human type
type Human struct {
	Name   string
	Age    int
	Height float32
	Weight float32
}

// CalcWeightToHeight is Human method
func (h *Human) CalcWeightToHeight() float32 {
	return h.Weight / (h.Height - 100)
}

// Android use embedding instead of inheritance
type Android struct {
	Human // Human struct, Android extends Human
	Model string
}

// interfaces

// Warrior is warrior interface
type Warrior interface {
	TakeWorld()
}

// TakeWorld prints phrase
func (a Android) TakeWorld() {
	fmt.Println("I am the king of the world")
}

// type assertion

// CastWarriorToAndroid casting Warrior interface to Android struct
func CastWarriorToAndroid(warrior Warrior) {
	warrior.TakeWorld()

	// do not do this - may raise panic
	// android := warrior.(Android)

	// try cast with Go type assertion
	if android, ok := warrior.(Android); ok {
		fmt.Println("cast success", android.Model)
	}

	// switch type assertion
	switch any := warrior.(type) {
	case Android:
		fmt.Println("This is type Android", any.Model)
	default:
		fmt.Println("This is unknown type")
	}
}

// binary shift

// PowerOfTwo is binary shift
func PowerOfTwo(v int) int {
	v--
	v |= v >> 1
	v |= v >> 2
	v |= v >> 4
	v |= v >> 8
	v |= v >> 16
	v++

	return v
}

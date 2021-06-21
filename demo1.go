package main

import "fmt"

func main() {

	// IF example
	// statePopulations := map[string]int{
	// 	"Gujarat": 10,
	// 	"mp":      20,
	// 	"up":      30,
	// }
	// if pop, ok := statePopulations["mp"]; ok {
	// 	fmt.Println(pop)
	// }

	// Switch example
	// switch i := 1 + 1; i { //initilizer
	// // switch 3 {
	// case 1:
	// 	fmt.Println("first case")
	// case 2:
	// 	fmt.Println("second case")
	// default:
	// 	fmt.Println("default case")
	// }

	// i := 10
	// switch {
	// case i <= 10:
	// 	fmt.Println("less then or equal to 10")
	// 	// fallthrough
	// fallthrough will execute next case no matter what
	// case i >= 20:
	// 	fmt.Println("greater than 20")
	// default:
	// 	fmt.Println("default case")
	// }

	// var i interface{} = []int{}

	// switch i.(type) {
	// case int:
	// 	fmt.Println("i is an int")
	// case float64:
	// 	fmt.Println("i is an float64")
	// case string:
	// 	fmt.Println("i is an string")
	// case []int:
	// 	fmt.Println("i is an int slice")
	// default:
	// 	fmt.Println("default")
	// }

	// Loop Example
	// for i, j := 1, 1; i <= 5; i, j = i+1, j+2 {
	// 	fmt.Println(i, j)
	// }

	// i := 1
	// for i <= 5 {
	// 	fmt.Println(i)
	// 	i++
	// }

	// for i := 1; i <= 10; i++ {
	// 	if i%2 != 0 {
	// 		continue
	// 	}
	// 	fmt.Println(i)
	// }

	// s := []int{1, 2, 3}
	// for k, v := range s {
	// 	fmt.Printf("key = %v, value = %v\n", k, v)
	// }

	// Defer Example
	// defer takes value at a time defer called not take a value at last
	// a := "A"
	// defer fmt.Println(a)
	// a = "D"

	// Panic, recover and defer example
	// panic occurs because our application can't continue to function
	// it's starting to panic it can't figure out what to do.
	// panic happen after deferred statement are executed.
	// execution sequence main function > defer > panic > return

	// fmt.Println("start")
	// // defer fmt.Println("this was deferred")
	// defer func() {
	// 	if err := recover(); err != nil {
	// 		log.Println(err)
	// 	}
	// }()
	// panic("something bad happened")
	// fmt.Println("end")

	// Pointer Example
	// type myStruct struct {
	// 	foo int
	// }

	// var ms *myStruct

	// ms = new(myStruct)
	// ms.foo = 42
	// fmt.Println(ms.foo)

	// ms = &myStruct{foo: 10}
	// fmt.Println(*ms)

	// slice and maps are self referenced
	// a := []int{1, 2, 3}
	// b := a
	// fmt.Println(a, b)
	// a[1] = 5
	// fmt.Println(a, b)

	// astrisk(*)

	// a := map[string]int{
	// 	"a": 1,
	// 	"b": 2,
	// }
	// b := a
	// fmt.Println(a, b)
	// b["a"] = 3
	// fmt.Println(a, b)

	// variadic parameters
	// result := sum(1, 2, 3, 4, 5)
	// fmt.Println("total is", *result)

	// result, err := divide(5.0, 0.0)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println(result)

	// anonymous function example
	// func() {
	// 	fmt.Println("hello go!")
	// }()

	// var f func() = func() {
	// 	fmt.Println("hello")
	// }
	// f()

	divide := func(a, b float64) (float64, error) {
		if b == 0.0 {
			return 0.0, fmt.Errorf("cannot divide by 0")
		}
		return a / b, nil
	}
	d, err := divide(5, 0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(d)

}

// func divide(a, b float64) (float64, error) {
// 	if b == 0.0 {
// 		return 0.0, fmt.Errorf("cannot divide by 0")
// 	}
// 	return a / b, nil
// }

// func sum(values ...int) *int {
// 	// fmt.Printf("%T\n", values)
// 	// fmt.Println(values)
// 	result := 0
// 	for _, v := range values {
// 		result += v
// 	}
// 	return &result
// }

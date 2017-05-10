//A Tour of Go
//Author: Ellysa Stanton
//Version: May 2017

package main

import (
	//Imports
	"fmt"
	"image"
	"io"
	"math"
	"math/cmplx"
	"runtime"
	"strings"
	"sync"
	"time"
)

const (
	//Constants
	GoldenRatio = 1.618034
	Big = 1 << 100
	Small = Big >> 99
)

var (
	//Basic Types
	ToBe bool = false
	MaxInt uint64 = 1<<64 - 1
	z complex128 = cmplx.Sqrt(-5 + 12i)

	//Struct Literals
	vert1 = Vertex{1,2}
	vert2 = Vertex{X: 1}
	vert3 = Vertex{}
	vertPtr = &Vertex{1,2}

	//Maps
	mmap map[string]Vert

	//Map Literals
	lmap = map[string]Vert{
		"Bell Labs": Vert{
			40.68433, -74.39967,
		},
		"Google": Vert{
			37.42202, -122.08408,
		},
	}
)

type Vertex struct {
	//Structs
	X int
	Y int
}

type Vert struct {
	//Maps
	Lat, Long float64
}


type mVert struct {
	//Methods
	X, Y float64
}

type myFloat float64
	//Methods Continued

type Abser interface {
	//Interfaces
	Abs() float64
}

type I interface {
	//Interfaces Implemented Implicitly
	M()
}

type T struct {
	//Interfaces Implemented Implicitly
	S string
}

type Person struct {
	//Stringers
	Name string
	Age int
}

type myError struct {
	//Errors
	When time.Time
	What string
}

type SafeCounter struct {
	//sync.Mutex
	v map[string]int
	mux sync.Mutex
}

func main() {

	//Imports
	fmt.Println("\n===Imports===")
	fmt.Printf("Now you have %g problems\n", math.Sqrt(7))

	//Exported Names
	fmt.Println("\n===Exported Names===")
	fmt.Printf("Pi = %f\n", math.Pi)

	//Functions
	fmt.Println("\n===Functions===")
	fmt.Printf("42 + 13 = %d\n", add(42, 13))

	//Multiple Results
	fmt.Println("\n===Multiple Results===")
	a, b := swap("HELLO", "GOODBYE")
	fmt.Printf("swap(HELLO, GOODBYE) returns ")
	fmt.Println(a, b)

	//Named Return Values
	fmt.Println("\n===Named Return Values===")
	fmt.Printf("split(17) returns ")
	fmt.Println(split(17))

	//Variables
	//... With Initializers
	//... Short Declarations
	fmt.Println("\n===Variables===")
	var i, j int = 1, 2
	k := 3
	c, python, java := true, false, "NO"
	fmt.Printf("i j k c python java = ")
	fmt.Println(i, j, k, c, python, java)

	//Basic Types
	fmt.Println("\n===Basic Types===")
	fmt.Printf("Type: %T, Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T, Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T, Value: %v\n", z, z)

	//Zero Values
	fmt.Println("\n===Zero Values===")
	var i2 int
	var f float64
	var b2 bool
	var s string
	fmt.Printf("i2 f b2 s = ")
	fmt.Printf("%v %v %v %q\n", i2, f, b2, s)

	//Type Conversions
	fmt.Println("\n===Type Conversions===")
	var x2, y2 int = 3, 4
	var f2 float64 = math.Sqrt(float64(x2*x2 + y2*y2))
	var z2 uint = uint(f2)
	fmt.Printf("x2 y2 z2 = ")
	fmt.Println(x2, y2, z2)

	//Type Inference
	fmt.Println("\n===Type Inference===")
	v := 42
	fmt.Printf("v is of type %T\n", v)

	//Constants
	fmt.Println("\n===Constants===")
	fmt.Printf("Golden Ratio: %f\n", GoldenRatio)
	const Truth = true
	fmt.Println("Go rules?", Truth)

	//Numeric Constants
	fmt.Println("\n===Numeric Constants===")
	fmt.Printf("needInt(Small) returns ")
	fmt.Println(needInt(Small))
	fmt.Printf("needFloat(Small) returns ")
	fmt.Println(needFloat(Small))
	fmt.Printf("needFloat(Big) returns ")
	fmt.Println(needFloat(Big))

	//For
	fmt.Println("\n===For===")
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Printf("sum = %d\n", sum)

	//For as "while"
	fmt.Println("\n===For as 'while'===")
	sum2 := 1
	for ; sum2 < 1000; {
		sum2 += sum2
	}
	fmt.Printf("sum2 = %d\n", sum2)

	//Forever
	/*for {
	}
	*/

	//If
	fmt.Println("\n===If===")
	fmt.Printf("sqrt(2) = %d\n", mySqrt(2))
	fmt.Printf("sqrt(-4) = %s\n", mySqrt(-4))

	//If With a Short Statement
	fmt.Println("\n===If With a Short Statement===")
	fmt.Printf("pow(3, 2, 10) = %f\n", pow(3, 2, 10))
	fmt.Printf("pow(3, 3, 20) = %f\n", pow(3, 3, 20))

	//If and Else
	fmt.Println("\n===If and Else===")
	fmt.Printf("pow2(3, 2, 10) = %f\n", pow2(3, 2, 10))
	fmt.Printf("pow2(3, 3, 20) = %f\n", pow2(3, 3, 20))

	//Switch
	fmt.Println("\n===Switch===")
	fmt.Print("Go runs on... ")
	switch os := runtime.GOOS; os {
		case "darwin":
			fmt.Println("OS X.")
		case "linux":
			fmt.Println("Linux.")
		default:
			fmt.Printf("%s.\n", os)
	}

	//Switch Evaluation Order
	fmt.Println("\n===Switch Evaluation Order===")
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
		case today + 0:
			fmt.Println("Today.")
		case today + 1:
			fmt.Println("Tomorrow.")
		case today + 2:
			fmt.Println("In two days.")
		default:
			fmt.Println("Too far away.")
	}

	//Switch With No Condition
	fmt.Println("\n===Switch With No Condition===")
	t := time.Now()
	switch {
		case t.Hour() < 12:
			fmt.Println("Good morning!")
		case t.Hour() < 17:
			fmt.Println("Good afternoon!")
		default:
			fmt.Println("Good evening!")
	}

	//Defer
	/*defer fmt.Println("world.")
	fmt.Println("Hello")
	*/

	//Stacking Defer
	for i := 0; i < 10; i ++ {
		defer fmt.Printf("defer #%d\n", i)
		if i == 9 {
			defer fmt.Println("\n===Defers===")
		}
	}

	//Pointers
	fmt.Println("\n===Pointers===")
	ii, jj := 42, 2701
	pp := &ii
	fmt.Printf("*pp = ")
	fmt.Println(*pp)
	*pp = 21
	fmt.Printf("ii = ")
	fmt.Println(ii)
	pp = &jj
	*pp = *pp / 37
	fmt.Printf("jj = ")
	fmt.Println(jj)

	//Structs
	fmt.Println("\n===Structs===")
	fmt.Println("X = 1, Y = 2")
	fmt.Printf("Vertex = ")
	fmt.Println(Vertex{1,2})

	//Struct Fields
	fmt.Println("\n===Struct Fields===")
	vert := Vertex{1,2}
	vert.X = 4
	fmt.Println("X changed to 4")
	fmt.Printf("Vertex = ")
	fmt.Println(vert)

	//Pointers to Structs
	fmt.Println("\n===Pointers to Structs===")
	vertP := &vert
	vertP.Y = 1e9
	fmt.Println("Y changed to 1000000000")
	fmt.Printf("Vertex = ")
	fmt.Println(vert)

	//Struct Literals
	fmt.Println("\n===Struct Literals===")
	fmt.Printf("vert1 vertPtr vert2 vert3 = ")
	fmt.Println(vert1, vertPtr, vert2, vert3)

	//Arrays
	fmt.Println("\n===Arrays===")
	var arr [2]string
	arr[0] = "Hello"
	arr[1] = "World"
	fmt.Printf("arr[0] arr[1] = ")
	fmt.Println(arr[0], arr[1])
	fmt.Printf("arr = ")
	fmt.Println(arr)
	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Printf("primes = ")
	fmt.Println(primes)

	//Slices
	fmt.Println("\n===Slices===")
	var slc []int = primes[1:4]
	fmt.Printf("primes[1:4] = ")
	fmt.Println(slc)

	//Slices Are Like References to Arrays
	fmt.Println("\n===Slices Are Like References to Arrays===")
	names  := [4]string {
		"Jim",
		"Pam",
		"Dwight",
		"Michael",
	}
	fmt.Printf("names = ")
	fmt.Println(names)
	slcA := names[0:2]
	slcB := names[1:3]
	fmt.Printf("slcA slcB = ")
	fmt.Println(slcA, slcB)
	fmt.Printf("slcB[1] = %s\n", slcB[1])
	slcB[1] = "XXX"
	fmt.Println("slcB[1] set to XXX")
	fmt.Printf("slcB[1] = %s\n", slcB[1])
	fmt.Printf("slcA, slcB = ")
	fmt.Println(slcA, slcB)
	fmt.Printf("names = ")
	fmt.Println(names)

	//Slice Literals
	fmt.Println("\n===Slice Literals===")
	fmt.Printf("primes = ")
	fmt.Println(primes)
	primesR := []bool{true, false, true, true, false, true}
	fmt.Printf("primesR = ")
	fmt.Println(primesR)
	primesS := []struct {
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
	fmt.Printf("primesS = ")
	fmt.Println(primesS)

	//Slice Defaults
	fmt.Println("\n===Slice Defaults===")
	fmt.Printf("primes = ")
	fmt.Println(primes)
	slc = primes[1:4]
	fmt.Printf("slc = primes[1:4] = ")
	fmt.Println(slc)
	slc = primes[:2]
	fmt.Printf("slc = primes[:2] = ")
	fmt.Println(slc)
	slc = primes[1:]
	fmt.Printf("slc = primes[1:] = ")
	fmt.Println(slc)

	//Slice Length and Capacity
	fmt.Println("\n===Slice Length and Capacity===")
	slc2 := []int{2, 3, 5, 7, 11, 13}
	printSlice("slc2", slc2)
	slc2 = slc2[:0]
	printSlice("slc2[:0]", slc2)
	slc2 = slc2[:4]
	printSlice("slc2[:4]", slc2)
	slc2 = slc2[2:]
	printSlice("slc2[2:]", slc2)

	//Nil Slices
	fmt.Println("\n===Nil Slices===")
	var slc3 []int
	printSlice("slc3", slc3)
	if slc3 == nil {
		fmt.Println("NIL")
	}
	slc4 := []int{9, 233, -4}
	printSlice("slc4", slc4)
	if slc4 == nil {
		fmt.Println("NIL")

	}

	//Creating a Slice With Make
	fmt.Println("\n===Creating a Slice With Make===")
	slcC := make([]int, 5)
	printSlice("slcC", slcC)
	slcD := make([]int, 0, 5)
	printSlice("slcD", slcD)
	slcE := slcD[:2]
	printSlice("slcE = slcD[:2]", slcE)
	slcF := slcE[2:5]
	printSlice("sclF = slcE[2:5]", slcF)

	//Slices of Slices
	fmt.Println("\n===Slices of Slices===")
	board := [][]string {
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"
	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}

	//Appending to a Slice
	fmt.Println("\n===Appending to a Slice===")
	var slcG []int
	printSlice("slcG", slcG)
	slcG = append(slcG, 0)
	printSlice("slcG + 0", slcG)
	slcG = append(slcG, 1)
	printSlice("slcG + 1", slcG)
	slcG = append(slcG, 2, 3, 4)
	printSlice("slcG + 2, 3, 4", slcG)

	//Range
	fmt.Println("\n===Range===")
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
	fmt.Println("2^pow:")
	for i, v := range pow {
		fmt.Printf("2^%d = %d\n", i, v)
	}
	pow2 := make([]int, 10)
	fmt.Println("2^pow:")
	for i := range pow2 {
		pow2[i] = 1 << uint(i) // == 2^i
	}
	for _, value := range pow2 {
		fmt.Printf("%d\n", value)
	}

	//Maps
	fmt.Println("\n===Maps===")
	mmap = make(map[string]Vert)
	mmap["Bell Labs"] = Vert{
		40.68433, -74.39967,
	}
	fmt.Printf("Bell Labs is at ")
	fmt.Println(mmap["Bell Labs"])

	//Map Literals
	fmt.Println("\n===Map Literals===")
	fmt.Printf("lmap = ")
	fmt.Println(lmap)

	//Mutating Maps
	fmt.Println("\n===Mutating Maps===")
	umap := make(map[string]int)
	umap["Answer"] = 42
	fmt.Println("Answer = ", umap["Answer"])
	umap["Answer"] = 82
	fmt.Println("Answer changed to 82")
	fmt.Println("Answer = ", umap["Answer"])
	delete(umap, "Answer")
	fmt.Println("Answer = ", umap["Answer"])
	v, ok := umap["Answer"]
	fmt.Println("Answer = ", v, "Present?", ok)

	//Function Values
	fmt.Println("\n===Function Values===")
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Printf("hypot = %f\n", hypot(5, 12))
	fmt.Printf("compute(hypot) = %f\n", compute(hypot))
	fmt.Printf("compute(math.Pow) = %f\n", compute(math.Pow))

	//Function Closures
	fmt.Println("\n===Function Closures===")
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}

	//Methods
	fmt.Println("\n===Methods===")
	vertM := mVert{3, 4}
	fmt.Printf("mVert{3,4} = ")
	fmt.Println(vertM.Abs())

	//Methods are Functions
	fmt.Println("\n===Methods are Functions===")
	fmt.Printf("mVert{3,4} = ")
	fmt.Println(Abs(vertM))

	//Methods Continued
	fmt.Println("\n===Methods Continued===")
	fl := myFloat(-math.Sqrt2)
	fmt.Printf("fl = ")
	fmt.Println(fl.Abs())

	//Pointer Receivers
	fmt.Println("\n===Pointer Receivers===")
	vertM.Scale(10)
	fmt.Printf("mVert{3,4} scaled by 10 = ")
	fmt.Println(vertM.Abs())

	//Interfaces
	fmt.Println("\n===Interfaces===")
	var ab Abser
	ab = fl		//a myFloat implements Abser
	ab = &vertM	//a *mVert implements Abser
	ab = vertM	//a mVert doesn't implement Abser
	fmt.Printf("vertM = ")
	fmt.Println(ab.Abs())

	//Interfaces Implemented Implicitly
	fmt.Println("\n===Interfaces Implemented Implicitly===")
	var iface I = T{"hello"}
	fmt.Printf("interface message = ")
	iface.M()

	//Nil Interface Values
	fmt.Println("\n===Nil Interface Values===")
	fmt.Printf("interface = ")
	describe(iface)
	fmt.Printf("interface message = ")
	iface.M()

	//Empty Interface
	fmt.Println("\n===Empty Interface===")
	var iface2 interface{}
	fmt.Printf("interface = ")
	describe2(iface2)
	iface2 = 42
	fmt.Printf("interface = ")
	describe2(iface2)
	iface2 = "hello"
	fmt.Printf("interface = ")
	describe2(iface2)

	//Type Assertions
	fmt.Println("\n===Type Assertions===")
	var iface3 interface{} = "hello"
	str := iface3.(string)
	fmt.Printf("interface = ")
	fmt.Println(str)
	str, okay := iface3.(string)
	fmt.Printf("interface, string? = ")
	fmt.Println(str, okay)
	flt, okay := iface3.(float64)
	fmt.Printf("interface, float64? = ")
	fmt.Println(flt, okay)

	//Type Switches
	fmt.Println("\n===Type Switches===")
	fmt.Printf("do(21) = ")
	do(21)
	fmt.Printf("do('hello') = ")
	do("hello")
	fmt.Printf("do(true) = ")
	do(true)

	//Stringers
	fmt.Println("\n===Stringers===")
	jim := Person{"Jim Halpert", 29}
	pam := Person{"Pam Halpert", 28}
	fmt.Println(jim, pam)

	//Errors
	fmt.Println("\n===Errors===")
	if err := run(); err != nil {
		fmt.Println(err)
	}

	//Readers
	fmt.Println("\n===Readers===")
	rd := strings.NewReader("Hiya, buddy!")
	byt := make([]byte, 8)
	for {
		n, err := rd.Read(byt)
		fmt.Printf("n = %v err = %v byt = %v\n", n, err, byt)
		fmt.Printf("byt[:n] = %q\n", byt[:n])
		if err == io.EOF {
			break
		}
	}

	//Images
	fmt.Println("\n===Images===")
	img := image.NewRGBA(image.Rect(0, 0, 100, 100))
	fmt.Printf("image bounds = ")
	fmt.Println(img.Bounds())
	fmt.Printf("colors at (0,0) = ")
	fmt.Println(img.At(0,0).RGBA())

	//Goroutines
	fmt.Println("\n===Goroutines===")
	go say("world")
	say("hello")

	//Channels
	fmt.Println("\n===Channels===")
	stuff := []int{7, 2, 8, -9, 4, 0}
	ch := make(chan int)
	go cSum(stuff[:len(stuff)/2], ch)
	go cSum(stuff[len(stuff)/2:], ch)
	xx, yy := <-ch, <-ch	//receive xx and yy from ch
	fmt.Printf("xx yy xx+yy = ")
	fmt.Println(xx, yy, xx+yy)

	//Buffered Channels
	fmt.Println("\n===Buffered Channels===")
	ch2 := make(chan int, 2)
	ch2 <- 1
	ch2 <- 2
	fmt.Printf("1 = ")
	fmt.Println(<-ch2)
	fmt.Printf("2 = ")
	fmt.Println(<-ch2)

	//Range and Close
	fmt.Println("\n===Range and Close===")
	ch3 := make(chan int, 10)
	go fibonacci(cap(ch3), ch3)
	for ifib := range ch3 {
		fmt.Println(ifib)
	}

	//sync.Mutex
	fmt.Println("\n===sync.Mutex===")
	cMut := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go cMut.Inc("somekey")
	}
	time.Sleep(time.Second)
	fmt.Println(cMut.Value("somekey"))
	
	//Select
	fmt.Println("\n===Select===")
	ch4 := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-ch4)
		}
		quit <- 0
	}()
	fibonacci2(ch4, quit)

	//Default Selection
	fmt.Println("\n===Default Selection===")
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			return
		default:
			fmt.Println(". . .")
			time.Sleep(50 * time.Millisecond)
		}
	}

}

//Functions
func add(x, y int) int {
	return x + y
}

//Multiple Results
func swap(x, y string) (string, string) {
	return y, x
}

//Named Return Values
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

//Numeric Constants
func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x*0.1
}

//If
func mySqrt(x float64) string {
	if x < 0 {
		return mySqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

//If With a Short Statement
func pow(x, n, lim float64) float64 {
	if v  := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

//If and Else
func pow2(x, n, lim float64) float64 {
	if v  := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("pow2 failed: %g >= %g\n", v, lim)
	}
	return lim
}

//Slice Length and Capacity
func printSlice(s string, x []int) {
	fmt.Printf("%s length = %d | capacity = %d | value = %v\n", s, len(x), cap(x), x)
}

//Function Values
func compute(fn func(float64, float64) float64) float64 {
	return fn(3,4)
}

//Function Closures
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

//Methods
//Interfaces
func (v mVert) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

//Methods are Functions
func Abs(v mVert) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}


//Methods Continued
func (f myFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

//Pointer Receivers
func (v *mVert) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}


//Interfaces Implemented Implicitly
func (t T) M() {
	fmt.Println(t.S)
}

//Nil Interface Values
func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

//Empty Interface
func describe2(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

//Type Switches
func do(i interface{}) {
	switch sw := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", sw, sw*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", sw, len(sw))
	default:
		fmt.Printf("I don't know about type %T\n", sw)
	}
}

//Stringers
func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

//Errors
func (e *myError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

func run() error {
	return &myError {
		time.Now(),
		"it didn't work",
	}
}

//Goroutines
func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

//Channels
func cSum(s []int, c chan int) {
	cSum := 0
	for _, v := range s {
		cSum += v
	}
	c <- cSum	//sends sum to channel c
}

//Range and Close
func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

//Select
func fibonacci2(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

//sync.Mutex
func (c *SafeCounter) Inc(key string) {
	c.mux.Lock()
	c.v[key]++
	c.mux.Unlock()
}

func (c *SafeCounter) Value(key string) int {
	c.mux.Lock()
	defer c.mux.Unlock()
	return c.v[key]
}

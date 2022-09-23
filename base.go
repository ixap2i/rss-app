package main

import (
	"fmt"
	"time"
	"math"
)
const (
	// Create a huge number by shifting a 1 bit left 100 places.
	// In other words, the binary number that is 1 followed by 100 zeroes.
	Big = 1 << 100
	// Shift it right again 99 places, so we end up with 1<<1, or 2.
	Small = Big >> 99
)
func WordCount(s string) map[string]int {
	return map[string]int{"x":1}
}
func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

type Vertex struct {
	X,Y int
}

type MapSample struct {
	Lat, Long float64
}

type ErrNegativeSqrt struct {
	When time.Time
	What string
}
var mmap map[string]MapSample
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}
func fibonacci() func() int {
    f, g := 0, 1
    return func() int {
        f, g = g, f+g // 前の値と新しい値を扱う定形
        return f
    }
}

func (e *ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}

func run(vv float64) error {
	if math.Signbit(vv) {
		return &ErrNegativeSqrt{
			time.Now(),
			"it didn't work",
		}
	} else {
		return nil
	}
}

func Sqrt(x float64) (float64, error) {
	if error := run(x); error != nil {
		return x, error
	} else {
		return x, nil
	}
}

// type MyReader struct{}

// func (t MyReader) Read([]byte) (int, error) {

// }
// TODO: Add a Read([]byte) (int, error) method to MyReader.

func main() {
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))

	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	var p *int

	i := 42
	p = &i

	fmt.Println(*p)
	*p = 21

	fmt.Println(Vertex{1, 2})
	var v = Vertex{1, 2}
	v.X = 2

	fmt.Println(v.X)
	fmt.Println(v.Y)

	var test [2]string

	test[0] = "Hello"
	test[1] = "World"

	fmt.Println(test[0], test[1])

	var primes = [6]int{1, 2, 3, 4, 5, 6}
	fmt.Println(primes)
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	a := names[0:2]
	b := names[1:3]

	fmt.Println(a, b)
	fmt.Println(len(a), cap(b))

	aaa := make([]int, 5)
	fmt.Println(aaa)

	board := [][]string{
		[]string{"_","_","_"},
		[]string{"_","_","_"},
		[]string{"_","_","_"},
	}
	var tttt = append(board, []string{"_","_","_"},)
	fmt.Println(tttt)
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

	for _, v := range pow {
		fmt.Printf("test %d %d", i, v)
	}

	mmap = make(map[string]MapSample)
	mmap["Bell Labs"] = MapSample {
		40.68433, -74.39967,
	}

	var mmap2 = map[string]MapSample {
		"Facebook": { 40.68433, -74.39967, },
		"Google": { 37.42202, -122.08408, },
	}
	fmt.Println(mmap["Bell Labs"])
	for i, v := range mmap2 {
		fmt.Printf("ssss %s %f", i, v)
	}
	// delete(mmap, "Bell Labs")
	for i, v := range mmap {
		fmt.Printf("ssss %s %f", i, v)
	}

	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}

	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}

	// negativeな値を渡すとエラーする
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}


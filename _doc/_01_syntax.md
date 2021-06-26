# SYNTAX

## SLICE

###[Slice literal](https://tour.golang.org/moretypes/9)
```shell script
s := []struct {
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
```

###[len, cap](https://tour.golang.org/moretypes/13)
```shell script
s := make([]int, 5) 
len(s)
cap(s)
```

###[Slices of slices](https://tour.golang.org/moretypes/14)
```shell script
board := [][]string{
	[]string{"_", "_", "_"},
	[]string{"_", "_", "_"},
	[]string{"_", "_", "_"},
}
```

```shell script
s = append(s, 2, 3, 4)
```

###[Traverse](https://tour.golang.org/moretypes/17)
```shell script
pow := make([]int, 10)
for i, _ := range pow
for _, value := range pow
```

## MAP
### [maps](https://tour.golang.org/moretypes/19)
```shell script
type Vertex struct {
	Lat, Long float64
}
var firstMap map[string]Vertex

var secondMap = map[string]Vertex{
	"Bell Labs": Vertex{
		40.68433, -74.39967,
	},
	"Google": Vertex{
		37.42202, -122.08408,
	},
}
secondMap[key] = value
delete(secondMap, key)
```

### [pass func as a param](https://tour.golang.org/moretypes/24)
```shell script
func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}
```

### [Function closure]()
#### <i>A closure is a function value that references variables from outside its body</i>
```shell script

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}
```

## METHOD
#### Methods for type
#### A method is just a function with a receiver argument.
### [Methods](https://tour.golang.org/methods/1)
```shell script
type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	vertex := Vertex{3, 4}
	fmt.Println(vertex.Abs())
}
```

### [Pointer receivers](https://tour.golang.org/methods/4)
#### Methods with pointer receivers can modify the value to which the receiver points
```shell script
type Vertex struct {
	X, Y float64
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}
// With a value receiver, the Scale method operates on a copy of the original Vertex value. 
// (This is the same behavior as for any other function argument.) 
func (v Vertex) Scale1(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	
	v.Scale1(10)
	fmt.Println(v.Abs()) # 5
	
	v.Scale(10)
	fmt.Println(v.Abs()) # 50
}

```

### [Interfaces](https://tour.golang.org/methods/9)
```shell script
type I interface {
	M()
}

type T struct {
	S string
}

func (t T) M() {
	fmt.Println(t.S)
}

func main() {
	var i I = T{"hello"}
	i.M()
}

```

## GOROUTINE
### [goroutines](https://tour.golang.org/concurrency/1)
#### A goroutine is a lightweight thread managed by the Go runtime.
#### Goroutines run in the same address space, so access to shared memory must be synchronized.
```shell script
go f(x, y, z)
```

### [Channels](https://tour.golang.org/concurrency/2)
#### Channels are a typed conduit through which you can send and receive values with the channel operator, <-
#### By default, sends and receives block until the other side is ready. This allows goroutines to synchronize without explicit locks or condition variables.
```shell script
ch := make(chan int)
ch <- v    // Send v to channel ch.
v := <-ch  // Receive from ch, and assign value to v.
```
### [Buffered Channels](https://tour.golang.org/concurrency/3)
#### Sends to a buffered channel block only when the buffer is full.
#### Receives block when the buffer is empty.
```shell script
ch := make(chan int, 100)
```

### [Range and Close] (https://tour.golang.org/concurrency/4)
#### A sender can close a channel to indicate that no more values will be sent.
```shell script
v, ok := <-ch
```
#### ok is false if there are no more values to receive and the channel is closed.

### [Select](https://tour.golang.org/concurrency/5)
#### The select statement lets a goroutine wait on multiple communication operations.
#### A select blocks until one of its cases can run, then it executes that case.
#### It chooses one at RANDOM if multiple are ready.
### The default case in a select is run if no other case is ready.
```shell script
func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for { // loop infinite
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit: // signal để thoát loop infinite
			fmt.Println("quit")
			return // break loop
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci(c, quit)
}
```

### [sync.Mutex](https://tour.golang.org/concurrency/9)


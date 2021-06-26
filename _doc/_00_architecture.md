#SCALABLE SYSTEM
###[scaling-pinterest](https://www.slideshare.net/InfoQ/scaling-pinterest)
###[sharding-mysql](https://medium.com/pinterest-engineering/sharding-pinterest-how-we-scaled-our-mysql-fleet-3f341e96ca6f)

# [PATTERN](https://ravindraelicherla.medium.com/10-design-patterns-every-software-architect-must-know-b33237bc01c2)

## [GOLANG] channel pattern
### [Streaming - read from a channel](https://play.golang.org/p/9C-u721el8p)
```shell script
package main

import (
	"fmt"
	"time"
)

func startSender(name string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 1; i <= 5; i++ {
			c <- (name + " hello")
			time.Sleep(time.Second)
		}
	}()
	return c
}

func main() {
	sender := startSender("Ti")
	for i := 1; i <= 5; i++ {
		fmt.Println(<-sender)
	}
}

```

### [Streaming - multiple worker](https://play.golang.org/p/df7WGblymJB)
```shell script
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func startSender(name string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 1; i <= 5; i++ {
			c <- (name + " hello")
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c
}

func main() {
	ti := startSender("Ti")
	teo := startSender("Teo")

	for {
		select {
		case msgTi := <-ti:
			fmt.Println(msgTi)
		case msgTeo := <-teo:
			fmt.Println(msgTeo)
		}
	}
}

```

### [call -- APIs aggregation](https://play.golang.org/p/RTMRgC2ddf4)
```shell script
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func fetchAPI(model string) string {
	time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	return model
}

func main() {
	responseChan := make(chan string)
	var results []string

	go func() { responseChan <- fetchAPI("users") }()
	go func() { responseChan <- fetchAPI("categories") }()
	go func() { responseChan <- fetchAPI("products") }()

	for i := 1; i <= 3; i++ {
		results = append(results, <-responseChan)
	}
	fmt.Println(results)
}

```
### [query-first](https://play.golang.org/p/QMqTcX-doWy)
```shell script
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func query(url string) string {
	time.Sleep(time.Duration(rand.Intn(2e3)) * time.Millisecond)
	return url
}

func queryFirst(servers ...string) <-chan string {
	c := make(chan string)
	for _, serv := range servers {
		go func(s string) { c <- query(s) }(serv)
	}
	return c
}

func main() {
	result := queryFirst("server 1", "server 2", "server 3")
	fmt.Println(<-result)
}

```
// lets work  go routine
// A goroutine is a theread of execution that's managed  by go's runtine      
// part of Go's concurrency model and enables you to run functions or methods simultaneously with other functions or methods.

package main
import  (
    "fmt"
    "time"
    "sync"
)

func sayHello() {
    fmt.Println("Hello goroutine")
}

// func  worker(ch chan int) {
//     time.Sleep(time.Second)
//     ch <- 42
// }

func worker(id int, wg *sync.WaitGroup) {
    defer wg.Done()
    fmt.Printf("Worker %d starting\n", id)
    time.Sleep(time.Second)
    fmt.Printf("Worker %d done\n", id)
}

var (
    counter int
    mu sync.Mutex
)

func increment(wg *sync.WaitGroup) {
   defer wg.Done()
   mu.Lock()
   counter++
   mu.Unlock()
}
func main() {
    go sayHello()
    // time.Sleep(time.Second)
    fmt.Println("hello main function")
    // In Golang,Channels is a powerful feature that allows goroutines to communicate with each other and synchronize their execution.
    // They provide a way to send and receive values between goroutines, making it easy to build concurrent programs. It is created using the make function.
    // ch := make(chan int)
    // // ch := make(chan int, 100)
    // go worker(ch)
    // value := <- ch
    // fmt.Println("Recived", value)
    // select statement goroutine to wait on multiple communication operations.
    //  In this example, the program waits for messages from ch1 and ch2 and display them as they arrive. The select statement ensures that the program can handle whichever message comes first without blocking.
    ch1 := make(chan string)
    ch2 := make(chan string)
    go func() {
        ch1 <- "get massage from channel 1"    
    }()
    go func() {
        ch2 <- "get massage from channel 2"
    }()
    for i := 0; i<2; i++ {
        select {
        case msg1 := <- ch1:
            fmt.Println(msg1)
        case msg2 :=  <- ch2:
            fmt.Println(msg2)
        }
    }
    // A WaitGroup is a synchronization primitive provided by the sync package that allows you to wait for a collection of goroutines to finish executing.
    // In this example, the program waits for all worker goroutines to complete their tasks before printing "All workers done". 
    var  wg sync.WaitGroup
    for i :=1; i <=  3; i++  {
        wg.Add(1)
        go worker(i, &wg)
        go increment(&wg)

    }
    wg.Wait()
    fmt.Println("all  worker are done")
    fmt.Println()
    //  Mutex
    // In Golang, the mutexesis provided by the sync package, are synchronization primitives used to protect shared resources and prevent race conditions when multiple goroutines access and modify those resources concurrently.
    // unbuffred   sender and receiver must meet at the same time.
    // Buffered channel = channel has a queue/storage, so sender can continue until the buffer is full.
    q := make(chan int, 2)
    q <- 10
    q <- 20
    q <- 30 // BLOCKS
    // recive  only feature and 
    func worker(ch <-chan int) {
    value := <-ch
    fmt.Println(value)
    //  send  only feature
    func worker(ch chan<- int) {
        value := <-ch
        fmt.Println(value)
    }
}
}   



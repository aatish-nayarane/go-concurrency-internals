package main
import "fmt"
import "sync"

// we going to use  even and oddd chan.
func even(evench chan bool, oddch chan bool, wg *sync.WaitGroup) {
    // this will  be manage  even number using chan 
    defer  wg.Done()
   
    for i:= 0; i<10; i +=2 {
        <-evench
        fmt.Println("EvenchNumber")
        fmt.Println("even: ", i)
        oddch <- true
    }
}

// we going to use  even and oddd chan.
func odd(evench chan bool, oddch chan bool, wg *sync.WaitGroup) {
    // this will  be manage  even number using chan
    defer wg.Done()
    
    for i := 1; i < 10; i += 2 {
        <-oddch
        fmt.Println("oddch calilng")
        fmt.Println("odd", i)
        if i < 9  {
            fmt.Println("here you are")
            evench <- true 
        }

    }
}

func main() {
    var  wg sync.WaitGroup

    evench := make(chan bool)
    oddch :=  make(chan bool)
    
    wg.Add(2)
    go even(evench, oddch, &wg)
    go odd(evench, oddch, &wg)
    evench <- true
    wg.Wait()
}

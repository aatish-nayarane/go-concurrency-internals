package main
import "fmt"
import "sync"
//  we have to write program for sender and  reciver which  will five vm create and recive as an  five vm creted
func sender(ch chan string, wg *sync.WaitGroup) {
    defer wg.Done()
    input := <- ch
    for i:= 0; i<5; i++ {
        ch <- input
        fmt.Println("vm creating", i)
    }    

}

func recive(ch chan string, wg *sync.WaitGroup) {
    defer wg.Done()
    for i:= 0; i<5; i++ {
        value := <- ch
        ch <- "VM created"
        fmt.Println("vm recive", value)
    }
}

func main() {
    // var wg sync.WaitGroup
    // ch := make(chan string)

    // // add  wait group
    // wg.Add(2)
    // go sender(ch, &wg)
    // go recive(ch, &wg)
    // ch <- "VM created"
    // There is no receiver ready to receive from the unbuffered channel, so the send blocks forever. Since main is the only goroutine, the Go runtime detects that all goroutines are blocked and reports a deadlock
    // ch := make(chan string)

    // ch <- "VM created"

    // fmt.Println("Done")
    // wg.Wait()
    

    // lets  practice  on buffred channel
    // var wg sync.WaitGroup
    // ch := make(chan string, 1)
    // ch <-  "vm created"
    // fmt.Println("done")

    //  predict the output will work or not.
    ch := make(chan string, 2)
    // ch <- "VM-1"
    // ch <- "VM-2"
    // fmt.Println("adfasdA")
    // ch <- "VM-3"
    // fmt.Println("B")

    // lets  practice with one go channel
    ch <-  "vm1"
    ch <- "vm2"
    fmt.Println("a")
    close(ch)
    // go  func() {
    //     ch <- "vm3"  
    //     // fmt.Println(<-ch)
    // }()
    // ch <- "vm3"
    // fmt.Println("b")
    value, ok := <-ch
    if !ok {
    fmt.Println("Channel is closed")
    }
    fmt.Println(<-ch, value)
    fmt.Println(<-ch)
    fmt.Println(<-ch)
    fmt.Println("main done")
    // channel closing.
    
    // Receiving from a closed channel is safe. Sending to a closed channel causes a panic.


}

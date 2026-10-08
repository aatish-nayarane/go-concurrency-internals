                 ┌── Worker 1 ──┐
                 │              │
Jobs ────────────┼── Worker 2 ──┼──→ Results
                 │              │
                 └── Worker 3 ──┘


package main
import "fmt"
// import "time"

// lets practice in fanout and  fain architecture using go lang.

func  worker(fanoutch <-chan string, faninch chan string, id int ) {
    for job := range  fanoutch {
        fmt.Println("dfsdf", id, job)
        faninch <-  job
    }
}



func main() {
    fanoutch := make(chan string, 10)
    faninch := make(chan string, 10)
    go  worker(fanoutch, faninch, 1)
    go  worker(fanoutch, faninch, 2)
    go  worker(fanoutch, faninch, 3)
    fanoutch <- "vm1"
    fanoutch <- "vm2"
    fanoutch <- "vm3"
    fanoutch <- "vm4"
    fanoutch <- "vm5"
    fanoutch <- "vm6"
    fanoutch <- "vm7"
    fanoutch <- "vm8"
    fanoutch <- "vm9"
    fanoutch <- "vm10"
    // time.Sleep(time.Second)

    for i := 0;  i<10;  i++ {
        fmt.Println(<-faninch )
    } 


}

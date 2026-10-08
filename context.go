package main
import "fmt"
import "context"
import  "time"
import "errors"

// create vm

// main
//  │
//  ├── create context
//  │
//  ├── start createvm goroutine
//  │
//  │       ┌─────────────────────┐
//  │       │ vm running 0         │
//  │       │ vm running 1         │
//  │       └─────────────────────┘
//  │
//  ├── wait 2 seconds
//  │
//  ├── cancel()
//  │       │
//  │       ▼
//  │   ctx.Done()
//  │       │
//  │       ▼
//  │   "ctx is done"
//  │       │
//  │       ▼
//  │     return
//  │
//  └── worker stops

func createvm(ctx  context.Context) error {
    for i:= 0; i<10; i++ {
        select {
            case  <-ctx.Done():
                fmt.Println("VM creation cancelled")
                return errors.New("VM creation cancelled")
            case  <-time.After(time.Second):
                fmt.Println("vm running", i)
                // time.Sleep(time.Second)
        }   
    }
    return nil
}

func main() {
    //  here we have to manage context 
    // ctx, cancel := context.WithCancel(context.Background())
    // timeout perpose we to use  TimeOut
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()
    go createvm(ctx)
    // time.Sleep(time.Second*4)
    // cancel()
    time.Sleep(time.Second*6)
    
}

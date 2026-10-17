package main
import "fmt"
import  "errors"

//which  is  represent  method signatures 
// The struct data type implements these interfaces to have method.
type VmEngine interface {
    create() string
}


type Kvm  struct {
    payload string
}

type  Virtual struct {
    payload string
}

func (vm Kvm) create() string {
    return vm.payload
}
func (vm  Virtual) create() string {
    return  vm.payload
}

func  create_Vm(vmc VmEngine)  string {
    return  vmc.create()
}

// We can type assertion for   
// error Handler 

func  Sqrt(value float64)(float64,  error) {
    if value < 0{
        return  0, errors.New("somethig erorr is therr") 
    } else {
        return value, nil
    }  
}

func main() {
    create_vm := Kvm{payload: "create vm"}
    create_virtual := Virtual{payload: "virtual vm"}
    fmt.Println(create_Vm(create_virtual))
    fmt.Println(create_Vm(create_vm))
    //  assertion  test 
    var i interface{} = "hello"
    // var  data map[string]interface{}
    // data["aatish"] = "aatish"
    // fmt.Println(data)
     
    s :=  i.(string)
    fmt.Println("output of go lang", s)
    
    // The Ok-idiom is a common pattern for handling errors.
    if s, ok := i.(string); ok {
        fmt.Println("string data: ", s)
    } else {
        fmt.Println("Not a string")
    }
    // handling failed type Assertion.
    result, err :=  Sqrt(-1)
    if  err != nil {
        fmt.Println(err)
    }  else  {
        fmt.Println(result)
    }

}

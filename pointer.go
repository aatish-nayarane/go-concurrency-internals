package main
import "fmt"

type  Student struct {
    name string 
    age  int
}
// lets take an example   interface  as well 


func  add(a int, b int) int  {
    return a+b
}

func main() {
    // fmt.Println("Hello, World!")
    // // lets take and  example  for  go pointer 
    
    // var i int = 10
    // var ptr *int
    // ptr = &i
    // fmt.Println("this is ptr value",  *ptr,  &ptr)
    // var j int =  10
    // ptr =  &j
    // *ptr = j
    // fmt.Println("this after value",  *ptr,  &ptr)
    //  lets work on struct based  pointer
    // p :=&Student{"Indu", 18}
    // fmt.Println("Name =", p.name)
    // fmt.Println("GET AGE", p.age)
    // (*p).age = 17
    // p.age = 16
    // fmt.Println("take all value",  p.age)
    // // Dereferencing Pointers in Functions.
    // var FuncPtr func(int, int) int
    // FuncPtr = add
    // result := FuncPtr(46, 64)
    // fmt.Println("Result =", result)
    // //  array of pointer.
    // const MAX int = 3
    // a :=  []int{10, 100, 200}
    // var i int
    // var ptr [MAX]*int;
    // for i=0; i < MAX; i++ {
    //     ptr[i] = &a[i] // assign  the address  of integer
    // }
    // // var j int
    // for j := 0; j < MAX; j++ {
    //     fmt.Println("this is value", *ptr[j])
    // }
    // fmt.Println("output value for port", ptr)
    //  double pointer
    // pointer to pointer us a form of chain pointer
    var b int
    var ptr *int
    var pptr **int
    b = 3000
    ptr = &b
    pptr =  &ptr
   /* take the value using pptr */
   fmt.Printf("Value of a = %d\n", b )
   fmt.Printf("Value available at *ptr = %d\n", *ptr )
   fmt.Printf("Value available at **pptr = %d\n", **pptr)
   var  x int =  100
   var  y int =  200

   fmt.Printf("Before swap, value of a : %d\n", x )
   fmt.Printf("Before swap, value of b : %d\n", y )

   swap(&x, &y)

   fmt.Printf("after swap, value of a : %d\n", x )
   fmt.Printf("after swap, value of b : %d\n", y )


}
func swap(x *int, y *int) {
   var temp int
   temp = *x    /* save the value at address x */
   *x = *y      /* put y into x */
   *y = temp    /* put temp into y */
}


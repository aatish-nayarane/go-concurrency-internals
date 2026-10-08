package main

import "fmt"

type VM struct {
	Name string
	CPU  int
	RAM  int
}

// Case 1
func case1() {
	vm := VM{
		Name: "vm-1",
		CPU:  4,
		RAM:  8192,
	}

	fmt.Println(vm.Name)
}

// Case 2
func case2() *VM {
	vm := VM{
		Name: "vm-2",
		CPU:  4,
		RAM:  8192,
	}

	return &vm
}

// Case 3
func updateVM(vm *VM) {
	vm.CPU = 8
}

func case3() {
	vm := VM{
		Name: "vm-3",
		CPU:  4,
		RAM:  8192,
	}

	updateVM(&vm)

	fmt.Println(vm.CPU)
}

// Case 4
func case4() *VM {
	return &VM{
		Name: "vm-4",
		CPU:  4,
		RAM:  8192,
	}
}

// Case 5
func case5() {
	vm := &VM{
		Name: "vm-5",
		CPU:  4,
		RAM:  8192,
	}

	fmt.Println(vm.Name)
}

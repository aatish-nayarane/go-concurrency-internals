package main
import "fmt"
import "errors"


// 
type VMDetails struct {
    Name string
    CPU int
    RAM int
}

// define Interface for VMPROVIDER
type VMProvider interface  {
    CreateVM(vm VMDetails) error
    DeleteVM(vm VMDetails)
}

// snapshot provider 
type SnapshotProvider interface  {
    CreateSnapshot(vm VMDetails)
}

// Nerwork Provider 
type NetworkProvider interface {
    AttachNetwork(vm VMDetails)
}

// cloud provider
type CloudProvider interface {
    VMProvider
    SnapshotProvider
    NetworkProvider
}

// custom  error  handlening 
type  error  interface  {
    Error() string
}

//Create struct for
type OpenStackProvider struct{}
// type OpenStackProvider struct{}
// custom error struct
type CPUQuotaError struct {
    VMName  string
    RequestedCPU int
    AvailableCPU int

}

func (e CPUQuotaError) Error() string {
    return fmt.Sprintf(
        "CPU quota exceeded for VM %s: requested=%d, available=%d",
        e.VMName,
        e.RequestedCPU,
        e.AvailableCPU,
    )
}

var ErrCPUQuota = errors.New("CPU quota exceeded")

func (p OpenStackProvider) CreateVM(vm VMDetails) error {
    if vm.CPU > 64 {
        return CPUQuotaError{
                    VMName:       vm.Name,
                    RequestedCPU: vm.CPU,
                    AvailableCPU: 64,
                }
            }

    return nil
}

func (vm  OpenStackProvider )  DeleteVM(details VMDetails) {
    fmt.Println(details)
}

func (vm  OpenStackProvider )  CreateSnapshot(details VMDetails) {
    fmt.Println(details)
}

func (vm  OpenStackProvider)  AttachNetwork(details VMDetails) {
    fmt.Println(details)
}


func deployApplication(provider CloudProvider, vm VMDetails) error {
    err := provider.CreateVM(vm)
    if err != nil {
        return err
    }
    provider.AttachNetwork(vm)
    provider.CreateSnapshot(vm)
    return nil
}


func main() {

    vm := VMDetails{
        Name: "test_vm",
        CPU: 85,
        RAM: 83,
    }
    //  
    var data CloudProvider
    data = OpenStackProvider{}
    err := deployApplication(data, vm)
    if err != nil {
        fmt.Println("Deployment failed:", err)
    }
}

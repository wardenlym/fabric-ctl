package vehiclerecord

import (
	"fmt"

	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-protos-go/peer"
)

type VehicleRecord struct {
}

func main() {
	err := shim.Start(new(VehicleRecord))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}

func (t *VehicleRecord) Init(stub shim.ChaincodeStubInterface) peer.Response {
	return shim.Success(nil)
}

func (t *VehicleRecord) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	function, args := stub.GetFunctionAndParameters()
	fmt.Println("invoke is running " + function)
	// Handle different functions
	if function == "funcA" {
		return t.funcA(stub, args)
	} else if function == "funcB" {
		return t.funcB(stub, args)
	} else if function == "funcC" {
		return t.funcC(stub, args)
	}

	fmt.Println("invoke did not find func: " + function) //error
	return shim.Error("Received unknown function invocation")
}

func (t *VehicleRecord) funcA(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	return shim.Success(nil)
}

func (t *VehicleRecord) funcB(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	return shim.Success(nil)
}

func (t *VehicleRecord) funcC(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	return shim.Success(nil)
}

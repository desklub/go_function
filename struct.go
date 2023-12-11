package main

import "fmt"

// Define a struct
type MyStruct struct {
    Name  string
    Value int
}

// Method to modify the struct
func (ms *MyStruct) ModifyStruct(newName string, newValue int) {
    ms.Name = newName
    ms.Value = newValue
}

// Method to access the struct without modifying it
func (ms MyStruct) AccessStruct() {
    fmt.Printf("Name: %s, Value: %d\n", ms.Name, ms.Value)
}

func main() {
    // Create an instance of the struct
    myInstance := MyStruct{Name: "Keneth", Value: 50}

    // Print the initial values
    fmt.Println("First values:")
    myInstance.AccessStruct()

    // Modify the struct using the method
    myInstance.ModifyStruct("Updated", 99)

    // Print the modified values
    fmt.Println("Modified values:")
    myInstance.AccessStruct()
}
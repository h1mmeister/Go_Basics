package main

import "fmt"

var number int = 5          

func main(){
    var decision bool =  true 
    fmt.Println("Original Value of number: ",number)
    number = 10               
    fmt.Println("New Value of number: ",number)
    fmt.Println("Value of decision: " ,decision)
}
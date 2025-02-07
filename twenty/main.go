package main

import (
    "fmt"
    "github.com/hermetoproject/integration-tests/twentyone"
)

func main() {
    fmt.Println("The gomod/twenty module requires minimum go version 1.20")
    twentyone.PrintTwentyOne()
}

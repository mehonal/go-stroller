package main

import (
    "fmt"
    "os"
)


func main() {
    fmt.Printf("Environment variable: %s\n", os.Getenv("TEST"))
}

package main

import (
    "fmt"
    "net/mail"
)

func main() {
    email := "example@invozone.us"
    _, err := mail.ParseAddress(email)
    if err != nil {
        fmt.Println("Invalid email:", err)
		return
    } else {
        fmt.Println("Valid email:", email)
    }
}

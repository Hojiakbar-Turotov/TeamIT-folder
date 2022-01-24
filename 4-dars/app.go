//go 1.10.4

package main  
import (
    "fmt"
    "os"
) 

func main() { 
    if os.Args[1] == "Hojiakbar" || os.Args[1] == "Shaxriyor" {
        fmt.Printf("sizning ismingiz: %s", os.Args[1])
    } else if os.Args[1] == "Turotov" || os.Args[1] == "" {
        fmt.Printf("sizning familyangiz: %s", os.Args[1])
    } else if os.Args[1] == 18 || os.Args[1] == 22 {
        fmt.Printf("sizning yoshingiz: %d", os.Args[1])
    } else {
        fmt.Printf("kiritilgan ma'lumotlar topilmadi")
    }
}
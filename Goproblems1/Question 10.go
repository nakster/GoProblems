package main

import "fmt"

func reverse(value string) string {
    // Convert string to rune slice.
    // ... This method works on the level of runes, not bytes.
    data := []rune(value)
    result := []rune{}

    // Add runes in reverse order.
    for i := len(data) - 1; i >= 0; i-- {
        result = append(result, data[i])
    }

    // Return new string.
    return string(result)
}

func main() {
    // Test our method.
    value1 := "cat"
    reversed1 := reverse(value1)
    fmt.Println(value1)
    fmt.Println(reversed1)

    value2 := "abcde"
    reversed2 := reverse(value2)
    fmt.Println(value2)
    fmt.Println(reversed2)
}

package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
    "unicode"
)

func isPalindrome(text string) bool {
    runes := []rune(strings.ToLower(text))
    i, j := 0, len(runes)-1

    for i < j {
        if !unicode.IsLetter(runes[i]) {
            i++
            continue
        }
        if !unicode.IsLetter(runes[j]) {
            j--
            continue
        }
        if runes[i] != runes[j] {
            return false
        }
        i++
        j--
    }
    return true
}

func main() {
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Enter text: ")
    input, _ := reader.ReadString('\n')
    input = strings.TrimSpace(input)

    if isPalindrome(input) {
        fmt.Println("Yes, it is a palindrome!")
    } else {
        fmt.Println("No, it is not a palindrome.")
    }
}

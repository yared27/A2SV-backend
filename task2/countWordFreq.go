package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func count_freq(st string)map[string]int{
	st = strings.ToLower(st)
	cleaned := ""
	for _, r := range st{
		if unicode.IsLetter(r) || unicode.IsSpace(r){
			cleaned += string(r)
		}
	}

	words := strings.Fields(cleaned)

	 freq := make(map[string]int)

	 for _, ch := range words{
			freq[(ch)]++
		
	 }
	 return freq
}

func main(){
	fmt.Print("Enter text: ")
	reader := bufio.NewReader(os.Stdin)
	st, _ := reader.ReadString('\n')
	dict := count_freq(st)
	fmt.Println(dict)
}
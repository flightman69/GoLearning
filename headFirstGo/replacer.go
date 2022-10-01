package main

import (
	"fmt"
	"strings"
)

func main() {
	broken := "s#nj#i is #wesome"
	fmt.Println("broken string:",broken)
	replacer := strings.NewReplacer("#", "a")
	fixed := replacer.Replace(broken)
	fmt.Println("fixed string:",fixed)
	againBroken := strings.ReplaceAll(fixed,"a", "$")
	fmt.Println("Breaking again using diff method for fun:",againBroken)

}

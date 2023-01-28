package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

var words = [...]string{
	"auto", "double", "int", "struct",
	"break", "else", "long", "switch",
	"case", "enum", "register", "typedef",
	"char", "extern", "return", "union",
	"continue", "for", "signed", "void",
	"do", "if", "static", "while", "default",
	"goto", "sizeof", "volatile", "const",
	"float", "short", "unsinged",
	"class", "new", "this", "delete",
	"operator", "throw", "except",
	"private", "public", "finally",
	"protected", "virtual", "friend",
	"try", "is", "with", "assert",
	"range", "import", "failthrough", "defer",
	"func", "select", "interface", "package",
	"true", "false", "print", "var",
	"as", "get", "part", "set",
}

func main() {
	kbReader := bufio.NewReader(os.Stdin)
	for {
		for i := 3; i > 0; i-- {
			fmt.Printf("\x1bc")
			fmt.Print(i)
			fmt.Printf("\a")
			time.Sleep(1 * time.Second)
		}
		typo := 0
		fmt.Printf("\x1bc")
		start := time.Now()
		for i := len(words) - 1; i > 0; i-- {
			word := words[i]
			for {
				fmt.Printf("  %s %d\n  ", word, i)
				for i := 0; i < len(word); i++ {
					fmt.Printf("─")
				}
				fmt.Printf("\n  ")
				input, _ := kbReader.ReadString('\n')
				input = strings.TrimSpace(input)
				fmt.Printf("\x1bc")
				if input == word {
					break
				}
				if input != "" {
					typo++
					fmt.Printf("\a")
				}
			}
		}
		fmt.Printf("\x1bc")
		fmt.Printf("time: %.1f minutes", time.Since(start).Minutes())
		fmt.Printf("\ntypo: %d", typo)
		time.Sleep(5 * time.Second)
	}
}

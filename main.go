package main

import "fmt"

//Basicamente o desafio é verificar se o map[string] é uma expressão válida, sendo assim, temos que fazer com que no output da função main deve ser retornado apenas true.

func main() {
	items := map[string]bool{
		"()             ": true,
		"([)]           ": false,
		"({a/b})        ": true,
		"[({()a}d)aa]   ": true,
		"[(--})))]      ": false,
		"ajdkjdsjkdj    ": true,
		"(([[]}}}})     ": false,
		")()            ": false,
	}

	for k, v := range items {
		fmt.Printf("%s expected: %t, passed: %t\n", k, v, solution(k) == v)
	}
}

func solution(str string) bool {
	return true
}

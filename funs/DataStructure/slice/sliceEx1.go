// return a slice holding only the non-empty strings.
package main

import "fmt"

func main() {
	data := []string{"one", "", "three", "", "five"}
	fmt.Println(nonempty(data))
}

func nonempty(strings []string) []string {
	//out := []string{}
	var out []string
	for _, v := range strings {
		if v != "" {
			out = append(out, v)
		}
	}
	return out
}

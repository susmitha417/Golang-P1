package maps

import "fmt"

// PrintMap prints all key-value pairs in a given map.
func PrintMap(m map[string]string) {
	for key, value := range m {
		fmt.Println(key, "->", value)
	}
}

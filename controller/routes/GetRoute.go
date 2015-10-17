package routes

import (
	"fmt"
	"net/http"
)

// Test does stuff
func Test() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("hello")
	}
}

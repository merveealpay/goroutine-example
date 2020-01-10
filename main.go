//independent and concurrency: go-routine
package main

import (
	"fmt"
	"time"
)

func main() {
	names := []string{"Merve", "Kumsal", "Cakil", "Deniz", "Kumsal"}
	//add go before function for concurrency
	go findX(names)
	go findY(names)

}

func findX(array []string) {
	for i := 0; i < len(array); i++ {
		if array[i] == "Kumsal" {
			fmt.Println("findX found the " + array[i])
		}
		time.Sleep(time.Second)
	}
}
func findY(array []string) {
	for i := 0; i < len(array); i++ {
		if array[i] == "Kumsal" {
			fmt.Println("findY found the " + array[i])
		}
		time.Sleep(time.Second)
	}
}

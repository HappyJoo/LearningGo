// 方法 1
package main

import (
	"fmt"
	"sync"
	"strings"
)

func main() {
	letter, number := make(chan bool), make(chan bool)
	wait := sync.WaitGroup{}
	
	go func() {
		i := 1
		for {
			select {
			case <- number:
				fmt.Print(i)
				i++
				fmt.Print(i)
				i++
				letter <- true
				break
			default:
				break
			}
		}
	}()

	wait.Add(1)
	go func(wait *sync.WaitGroup) {
		str := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

		i := 0
		for {
			select {
			case <- letter:
				if i >= strings.Count(str, "") - 1{
					wait.Done()
					return 
				}
				fmt.Print(str[i:i+1])
				i++
				if i >= strings.Count(str, "") {
					i = 0
				}
				fmt.Print(str[i:i+1])
				i++
				number <- true
				break
			default:
				break
			}
		}
	}(&wait)
	number <- true
	wait.Wait()
}

// 方法 2
package main

import (
	"fmt"
)

func main() {
	chanNumber := make(chan bool)
	chanLetter := make(chan bool, 1)
	done := make(chan int)

	go func() {
		for i := 1; i < 26; i += 2 {
			<-chanLetter 
			fmt.Print(i)
			fmt.Print(i + 1)
			chanNumber <- true
		}
	}()

	go func() {
		letters := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", 
		"L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
		for i := 0; i < 26; i += 2 {
			<-chanNumber 
			fmt.Print(letters [i])
			fmt.Print(letters [i+1])
			chanLetter <- true
		}
		done <- 1
	}()

	chanLetter <- true
	<-done
}


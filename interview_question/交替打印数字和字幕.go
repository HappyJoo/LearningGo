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
				letter <- true // call letter func
				break
			default:
				break
			}
		}
	}()

	wait.Add(1)
	go func(wait *sync.WaitGroup) {
		str := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
		strLength := strings.Count(str, "") - 1
		i := 0
		for {
			select {
			case <- letter:
				if i >= strLength {
					wait.Done()
					return 
				}
				fmt.Print(str[i:i+1])
				i++
				if i >= strLength {
					i = 0
				}
				fmt.Print(str[i:i+1])
				i++
				number <- true // call number func
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
	chanNumber := make(chan bool, 1)
	chanLetter := make(chan bool)
	done := make(chan int)

	go func() {
		for i := 1; i < 26; i += 2 {
			<-chanNumber 
			fmt.Print(i)
			fmt.Print(i + 1)
			chanLetter <- true // chanLetter in to call letter func
		}
	}()

	go func() {
		letters := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", 
		"L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
		for i := 0; i < 26; i += 2 {
			<-chanLetter 
			fmt.Print(letters [i])
			fmt.Print(letters [i+1])
			chanNumber <- true // chanNumber in to call number func
		}
		done <- 1
	}()
	
	// chanNumber 传入之后才可以触发 number func
	chanNumber<- true
	// 全部跑完之后 done <- 1 然后触发 done
	<-done

}


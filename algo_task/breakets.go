// Нам известно число n 
// Необходимо вывести все правильные скобочные последовательности в лексикографическом порядке с n открывающимися скобками

package main

import (
	"fmt"
)

func genBracket(cur string, open int, closed int, n int) {
    if len(cur) == 2*n {
        fmt.Println(cur)
        return
    }
    if open < n {
        genBracket(cur + "(", open+1, closed, n)
    }
    if open > closed {
        genBracket(cur + ")", open, closed+1, n)
    }
}

func run(n int) {
    genBracket("", 0, 0, n)
}

func main() {
  run(10)
}

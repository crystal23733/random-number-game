package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

var stdin = bufio.NewReader(os.Stdin)

func InputIntValue() (int, error) {
	var n int
	_, err := fmt.Scanln(&n)
	if err != nil {
		stdin.ReadString('\n')
	}

	return n, err
}

func main() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source).Intn(100)

	cnt := 1
	for {
		fmt.Printf("숫자를 입력하세요:")
		n, err := InputIntValue()
		if err != nil {
			fmt.Println("숫자만 입력하세요.")
		} else {
			if n > r {
				fmt.Println("입력된 값이 더 큽니다.")
			} else if n < r {
				fmt.Println("입력된 값이 더 작습니다.")
			} else {
				fmt.Println("정답입니다! 시도한 횟수:", cnt)
				break
			}
			cnt++
		}
	}
}

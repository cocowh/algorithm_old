package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(getHint("00112233445566778899", "16872590340158679432"))
}

func getHint(secret string, guess string) string {
	bull, cow, m := 0, 0, make([]int, 10)
	for i := 0; i < len(secret); i++ {
		if secret[i] == guess[i] {
			bull++
			continue
		}
		m[secret[i]-'0']++
		m[guess[i]-'0']--
	}
	for i := 0; i < 10; i++ {
		if m[i] > 0 {
			cow += m[i]
		}
	}
	cow = len(secret) - cow - bull
	return strconv.Itoa(bull) + "A" + strconv.Itoa(cow) + "B"
}

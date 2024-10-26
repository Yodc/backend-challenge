package main

import (
	"fmt"
	"math"
)

func decode(encoded string) string {
	var minSum = math.MaxInt64
	var bestCombo string

	var generateCombinations func(current string, previous int)
	generateCombinations = func(current string, previous int) {
		if len(current) == len(encoded)+1 {
			sum := 0
			for _, char := range current {
				sum += int(char - '0')
			}
			if sum < minSum {
				minSum = sum
				bestCombo = current
			}
			return
		}

		for i := 0; i < 10; i++ {
			if len(current) > 0 {
				switch encoded[len(current)-1] {
				case 'L':
					if previous > i {
						generateCombinations(current+string(rune('0'+i)), i)
					}
				case 'R':
					if previous < i {
						generateCombinations(current+string(rune('0'+i)), i)
					}
				case '=':
					if previous == i {
						generateCombinations(current+string(rune('0'+i)), i)
					}
				}
			} else {
				generateCombinations(current+string(rune('0'+i)), i)
			}
		}
	}

	generateCombinations("", 0)

	return bestCombo
}

func main() {
	var encoded string
	fmt.Print("กรุณาใส่ข้อความที่เข้ารหัสแล้ว: ")
	fmt.Scanln(&encoded)
	decoded := decode(encoded)
	fmt.Printf("ตัวเลขชุดที่มีผลรวมของทุกตัวเลขมีค่าน้อยที่สุดคือ: %s\n", decoded)
}

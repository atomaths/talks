package main

import (
	"fmt"
	"log"
	"unicode"
)

func main() {
	var ga rune
	ga = '\uAC00' // '가'
	fmt.Printf("%q\n", ga)


	delta := 'δ' // delta has type rune.
	var DELTA rune

	DELTA = unicode.ToUpper(delta) // delta의 대문자 'Δ' 문자가 리턴
	epsilon := unicode.ToLower(DELTA + 1) // DELTA('Δ')에 1을 더하면 Epsilon(대문자)가 나오는데 여기다 ToLower
	if epsilon != 'δ'+1 {
		log.Fatal("inconsistent casing for Greek")
	}
}

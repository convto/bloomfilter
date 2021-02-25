package main

import (
	"crypto/md5"
	"fmt"
)

var bits uint8

func main() {
	add("test1")
	fmt.Printf("added bits: %b\n\n", bits)
	// 存在する
	fmt.Printf("is exists 'test1': %t\n\n", exists("test1"))
	// 存在しない
	fmt.Printf("is exists 'test2': %t\n\n", exists("test2"))
	// 存在しないのに存在すると判定している(偽陽性)
	fmt.Printf("is exists 'test12': %t\n\n", exists("test12"))
}

func add(s string) {
	idxes := hash(s)
	for _, idx := range idxes {
		fmt.Printf("bit on: %b\n", 1<<idx)
		// idx番目のbitを立てる
		bits |= 1 << idx
	}
}

func exists(s string) bool {
	idxes := hash(s)
	for _, idx := range idxes {
		fmt.Printf("bit check: %b\n", 1<<idx)
		// idx番目のbitが立っていないケースを判定
		if 1<<idx != 1<<idx&bits {
			return false
		}
	}
	return true
}

func hash(s string) (idxes [3]uint8) {
	// md5結果の終端byteの上位3bitの値(0~7)をbitsのindex値として返す
	// ことなるハッシュがとりたいので適当にsaltをふる
	idxes[0] = md5.Sum([]byte(s + "1"))[15] >> 5
	idxes[1] = md5.Sum([]byte(s + "2"))[15] >> 5
	idxes[2] = md5.Sum([]byte(s + "3"))[15] >> 5
	return idxes
}
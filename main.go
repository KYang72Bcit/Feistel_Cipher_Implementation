package main

import (
	"errors"
	"fmt"
	"strings"
)


const (
	ROUND int = 8;
	BLOCK_SIZE int = 128

)

var permutation = []int{
	12, 3, 24, 7, 18, 9, 21, 34,
	1, 13, 27, 8, 19, 5, 25, 20,
	31, 2, 14, 6, 22, 10, 23, 11,
	30, 4, 0, 26, 17, 15, 28, 32,
	33, 29, 16, 35,
}


type roundFunction func(int, []byte, []byte) []byte


func encryptionRounds(plainText []byte, key []byte, rd roundFunction, round int ) []byte {
	mid := len(plainText)/2

	left := plainText[:mid]
	right := plainText[mid:]

	//padding to make sure they are the same length
	if (len(left) < len(right)) {
		left = append(left, make([]byte, len(right) - len(left))...)
	} else if len(right) < len(left) {
		right = append(left, make([]byte, len(right) - len(left))...)
	}

	newLeft := right
	roundResult := rd(round, right, key)

	newRight := make([]byte, len(right))
	for i := 0; i < len(right); i++ {
		newRight[i] = left[i] ^ roundResult[i]
	}
	return append(newLeft, newRight...)

}

// key must be 32 bits
func keyPermutation(key string) (string, error) {
	if (len(key)!= len(permutation)){
		return "", errors.New("key must be 32 bits")
	}
	subkey := make([]byte, len(permutation))
	for i, pos := range permutation {
		subkey[i] = key[pos] // permutation array is 0-based index
	}
	return string(subkey), nil

}

func bytesTobit(bytes []byte ) string {
	var bits strings.Builder
	for _, b := range bytes {
		bits.WriteString(fmt.Sprintf("%08b", b))
	}
	return bits.String()
}

func bitsToBytes(bits string) ([]byte, error) {
	if len(bits)%8 != 0 {
		return nil, errors.New("Incorrect length of bit string")
	}

	bytes := make([]byte, len(bits)/8)
	for i := 0; i < len(bits); i += 8 {
		var b byte
		for j := 0; j < 8; j++ {
			if bits[i+j] == '1' {
				b |= 1 << (7 - j)
			}
		}
		bytes[i/8] = b
	}
	return bytes, nil
}

func rd(round int, text []byte, key []byte ) []byte {
	
	return nil 
}
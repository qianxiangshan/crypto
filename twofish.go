package main

import (
	"bytes"
	"crypto/ct"
	"fmt"
	"math/rand"

	"golang.org/x/crypto/twofish"
)

func randomdata() []byte {
	var randbyte = make([]byte, 16)
	for i := 0; i < 16; i++ {
		randbyte[i] = byte(rand.Uint32())
	}
	return randbyte
}

func main() {

	for i := 0; i < 100; i++ {
		key := randomdata()
		data := randomdata()
		var gto = make([]byte, 16)
		var cto = make([]byte, 16)
		dogotwofish(key, data, gto)
		ct.CtwoFish(key, data, cto)
		result := bytes.Compare(gto, cto)
		if result != 0 {
			panic("not equal")
		}
		fmt.Println(cto)
	}
}

func dogotwofish(key, in, out []byte) error {
	cipher, err := twofish.NewCipher(key)
	if err != nil {
		return nil
	}
	cipher.Encrypt(out, in)
	return nil
}

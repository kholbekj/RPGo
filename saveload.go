package main

import (
	/*"crypto/aes"
	"crypto/cipher"
	"crypto/rand"*/

	"encoding/json"
	"fmt"
	//"io"
)

func saveState(p Player) bool {
	save, err := json.Marshal(p)
	if err != nil {
		fmt.Println(err)
		return false
	}
	fmt.Println(string(save))
	return true
}

func loadState() {

}

/*func encryptSave(jsonData string) {
	key := []byte("this is an awesome key muddafukaaaa")
	cleartext := []byte(jsonData)

	if len(cleartext)%aes.BlockSize != 0 {
		panic("json data is not a multiple of the block size")
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	ciphertext := make([]byte, aes.BlockSize+len(cleartext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext[aes.BlockSize:], cleartext)

	fmt.Printf("%x\n", ciphertext)
}*/

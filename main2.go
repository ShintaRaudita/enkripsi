package main

import (
    "crypto/cipher"
    "crypto/des"
    "encoding/hex"
    "decrypt"
)

func main() {
    // Generate a new DES key 
    // Key and plaintext
    key := []byte("8bytekey")
    plaintext := []byte("My secret doc")

    // Pad plaintext to 8 byte multiple
  padding := 8 - (len(plaintext) % 8)
  padded := make([]byte, len(plaintext)+padding)
  copy(padded, plaintext)

    // Encrypt
  block, _ := des.NewCipher(key)
  encrypter := cipher.NewCBCEncrypter(block, key)
  ciphertext := make([]byte, len(padded))
  encrypter.CryptBlocks(ciphertext, padded)

    // Encode and print
  hexCiphertext := hex.EncodeToString(ciphertext)
  println(hexCiphertext)


    // To decrypt:
    // 1. Decode the hex string back to bytes
    // 2. Create a DES decrypter 
    // 3. Decrypt the ciphertext 
}
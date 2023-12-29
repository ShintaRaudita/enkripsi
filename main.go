package main

import (
  "crypto/cipher"
  "crypto/des"
  "encoding/hex"
)

func main() {

  // Hex ciphertext 
  hexCiphertext := "d44cae770b83817f12d05aefab2151f5"

  // Decode hex
  ciphertext, _ := hex.DecodeString(hexCiphertext)

  // DES key
  key := []byte("8bytekey")

  // Create DES decrypter
  block, _ := des.NewCipher(key)
  decrypter := cipher.NewCBCDecrypter(block, key)

  // Decrypt ciphertext
  plaintext := make([]byte, len(ciphertext))
  decrypter.CryptBlocks(plaintext, ciphertext)

  // Print decrypted plaintext
  println(string(plaintext))
}





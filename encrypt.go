package encrypt

import (
    "crypto/des"
    "crypto/cipher"
    "crypto/rand"
    "encoding/base64"
    "fmt"
)

// Encrypt mengenkripsi data dengan kunci enkripsi
func Encrypt(data []byte) []byte {
    // Generate a cipher
    cipher, err := des.NewCipher(Key)
    if err != nil {
        fmt.Println(err)
        return nil
    }

    // Encrypt the data
    ciphertext := make([]byte, len(data))
    cipher.Encrypt(ciphertext, data)

    // Encode the ciphertext to base64
    return base64.StdEncoding.EncodeToString(ciphertext)
}

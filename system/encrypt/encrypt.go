/*                                                                          
|---------------------------------------------------------------            
| Encrypt
|---------------------------------------------------------------            
|
| Encryption utility helper 
| To use a 44-character key, we'll need to hash it down
| to 32 bytes for AES-256 encryption. We'll use SHA-256 for this purpose.
| 
| @author: IgnitedCMS
| @license: MIT
| @version: 1.0
| @since: 1.0 
|
*/


/* --Usage example
import (
    "fmt"
    "system/gozen/encrypt"
)
e := encrypt.New()

encryptedValue := e.Encrypt("Hello, World!")
fmt.Println("Encrypted:", encryptedValue)

decryptedValue, err := e.Decrypt(encryptedValue)
if err != nil {
    log.Fatal(err)
}
fmt.Println("Decrypted:", decryptedValue)
*/

package encrypt

import (
    "crypto/aes"
    "crypto/cipher"
    "crypto/rand"
    "crypto/sha256"
    "encoding/base64"
    "errors"
    "io"
    "os"

    "github.com/joho/godotenv"
)

type Encryptor struct {
    key []byte
}

func New() *Encryptor {
    err := godotenv.Load()
    if err != nil {
        panic("Error loading .env file")
    }

    key := os.Getenv("APP_KEY")
    if key == "" {
        panic("APP_KEY not found in .env file")
    }

    if len(key) != 44 {
        panic("APP_KEY must be exactly 44 characters long")
    }

    // Hash the 44-character key to get a 32-byte key
    hasher := sha256.New()
    hasher.Write([]byte(key))
    hashedKey := hasher.Sum(nil)

    return &Encryptor{
        key: hashedKey,
    }
}

func (e *Encryptor) Encrypt(plaintext string) string {
    block, err := aes.NewCipher(e.key)
    if err != nil {
        panic(err)
    }

    gcm, err := cipher.NewGCM(block)
    if err != nil {
        panic(err)
    }

    nonce := make([]byte, gcm.NonceSize())
    if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
        panic(err)
    }

    ciphertext := gcm.Seal(nonce, nonce, []byte(plaintext), nil)
    return base64.StdEncoding.EncodeToString(ciphertext)
}

func (e *Encryptor) Decrypt(ciphertext string) (string, error) {
    data, err := base64.StdEncoding.DecodeString(ciphertext)
    if err != nil {
        return "", err
    }

    block, err := aes.NewCipher(e.key)
    if err != nil {
        return "", err
    }

    gcm, err := cipher.NewGCM(block)
    if err != nil {
        return "", err
    }

    nonceSize := gcm.NonceSize()
    if len(data) < nonceSize {
        return "", errors.New("ciphertext too short")
    }

    nonce, encryptedData := data[:nonceSize], data[nonceSize:]
    plaintext, err := gcm.Open(nil, nonce, encryptedData, nil)
    if err != nil {
        return "", err
    }

    return string(plaintext), nil
}

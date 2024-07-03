/*
|---------------------------------------------------------------
| Key
|---------------------------------------------------------------
|
| Generates a crypto secure app_key
| 32 bytes will give you a 44-character base64 string
|
| @author: IgnitedCMS
| @license: MIT
| @version: 1.0
| @since: 1.0
|
*/
package main

import (
	"fmt"
	"gozen/system/hash"
)

func main() {
	key, err := hash.GenerateKey(32)
	if err != nil {
		fmt.Print("erro")
	}
	fmt.Print(key)
}

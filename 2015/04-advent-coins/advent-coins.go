package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strings"
)

var secretKey = "yzbqklnj"

func main() {
	md5Hash := MD5Hash{secretKey}
	fmt.Println("Lowest number whose MD5 hash yields 00000 start:", md5Hash.GetLowest("00000"))
	fmt.Println("Lowest number whose MD5 hash yields 000000 start:", md5Hash.GetLowest("000000"))
}

func (md5Hash *MD5Hash) GetLowest(startsWith string) int {
	num := 1
	hasher := md5.New()

	for {
		hasher.Write([]byte(secretKey + fmt.Sprint(num)))
		hex := hex.EncodeToString(hasher.Sum(nil))
		index := strings.Index(hex, startsWith)
		if index == 0 {
			break
		}
		num++
		hasher.Reset()
	}

	return num
}

type MD5Hash struct {
	secretKey string
}

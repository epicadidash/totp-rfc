package main

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base32"
	"encoding/binary"
	"strings"
	"time"
)

func main() {
	text := "password"
	conter := 0

}
func TOTP(secret string) int {
	counter := time.Now().Unix()
	HOTP(secret, int(counter)/30)
}
func HOTP(secret string, counter int) int {
	decodedText, err := base32.StdEncoding.DecodeString(strings.ToUpper(secret))
	if err != nil {

	}
	bs := make([]byte, 8)
	binary.BigEndian.PutUint64(bs, uint64(counter))
	hash := hmac.New(sha1.New, []byte(decodedText))
	hash.Write(bs)
	h := hash.Sum(nil)
	code := dynamictruncation(h)
	return code
}
func dynamictruncation(h []byte) int {

	offset := int(low_order_4_bits(h[19]))
	p := h[offset : offset+4]
	otp := (last_31_bits(p)) % 1000000
	return otp
}

func low_order_4_bits(star byte) byte {
	hei := star & 15
	return hei
}

func last_31_bits(p []byte) int {
	var header uint32
	r := bytes.NewReader(p)
	var err = binary.Read(r, binary.BigEndian, &header)
	if err != nil {

	}
	return int(header) & 0x7fffffff
}

package main

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base32"
	"encoding/binary"
	"strings"
)

func main() {
	text := "password"
	conter := 0
	encodedText := base32.StdEncoding.EncodeToString([]byte(text))
	key := strings.ToUpper(encodedText)

	bs := make([]byte, 8)
	binary.BigEndian.PutUint64(bs, uint64(conter))

	hash := hmac.New(sha1.New, []byte(key))
	hash.Write(bs)
	h := hash.Sum(nil)
	dynamictruncation(h)

}
func dynamictruncation(h []byte) int {

	offset := int(low_order_4_bits(h[19]))
	p := h[offset : offset+4]
	otp := (last_31_bits(p)) % 1000000
	println(otp)
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

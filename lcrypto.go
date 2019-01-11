package lcrypto

import (
	"fmt"
	"strconv"
	"unicode/utf8"
	"strings"
)

var(

	key = []rune{'d', 'y', '!', 'U', '7', 'e', '0', 'z', 'B', '@', 'm', 'A', 'N', 'D', '^', 'o', '3',
				 'g', 'k', 'R', '5', 'z', '(', 'W', 'w', '$', 'Z', 'C', 'M', '2', 'v', '1', 'Z', 'd', '-', '2' }
)


func Decrypt(value string) (string, string){

	split := strings.Split(value, "=")

	checksum := split[1]
	encoded := split[0]

	decoded := ""

	rollingIndex := (len(encoded) / 4) / 2

	for i := 0; i < len(encoded); i+=4 {


		hey := []byte{encoded[i], encoded[i+1], encoded[i+2], encoded[i+3]}

		s := string(hey[:4])

		tmpCode, _ := strconv.ParseInt(s, 16, 64)

		index := (i/4) + rollingIndex

		code := tmpCode / int64(key[index])

		decoded += string(code)

	}


	return decoded, checksum

}

func Encrypt(value string) string{

	encoded := ""
	checksum := 0

	rollingIndex := len(value) / 2

	for i, w := 0, 0; i < len(value); i += w {

		runeValue, width := utf8.DecodeRuneInString(value[i:])

		ASCII := int(runeValue)

		index := i + rollingIndex
		KEYCODE := int(key[index])

		value := ASCII * KEYCODE

		checksum += value

		hex := strconv.FormatInt(int64(value), 16)

		switch hexsize := len(hex); hexsize {

		case 1:
			encoded += "000" + hex
		case 2:
			encoded += "00" + hex
		case 3:
			encoded += "0" + hex
		default:
			encoded += hex

		}

		w = width
	}

	return encoded + "=" + strconv.FormatInt(int64(checksum), 16)

}


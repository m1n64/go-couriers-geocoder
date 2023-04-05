package encoder

import (
	"hash/crc32"
	"strings"
)

var keychars = "7BJSL2T5KCW61QH9GYDVZXE3PU8R4AFNM"
var divider = 21
var noise = 23
var biasDivider = 13

func Encode(i int) string {
	if i == 0 {
		return ""
	}

	i = i + noise           // Другое число кодируем
	bias := i % biasDivider // Смещение по словарю - остаток от деления на 10

	var res strings.Builder

	for i > 0 {
		div := i % divider
		i = i / divider
		res.WriteByte(keychars[div+bias])
	}

	verChars := "VZKX913"
	crc := crc32.ChecksumIEEE([]byte(res.String()))
	res.WriteByte(verChars[int(crc)%len(verChars)]) // Предпоследний символ - это версия алгоритма. Если один из этих, то этот используется
	res.WriteByte(keychars[bias])                   // Сохраняем последним символов смещение

	return res.String()
}

func Decode(code string) int {
	if code == "" {
		return 0
	}

	biasC := string(code[len(code)-1])
	bias := strings.Index(keychars, biasC)
	code = code[:len(code)-2]
	codeRunes := []rune(code)
	for i, j := 0, len(codeRunes)-1; i < j; i, j = i+1, j-1 { // Развернуть строку
		codeRunes[i], codeRunes[j] = codeRunes[j], codeRunes[i]
	}

	val := 0

	for _, c := range codeRunes {
		val *= divider
		val += strings.Index(keychars, string(c)) - bias
	}

	return val - int(noise)
}

package util

import (
	"fmt"
	"math/rand"
	"regexp"
	"time"
)

var r *rand.Rand

func init() {
	r = rand.New(rand.NewSource(time.Now().Unix()))
}

// RandString 生成随机字符串
func RandString(len int) string {
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		b := r.Intn(26) + 65
		bytes[i] = byte(b)
	}
	return string(bytes)
}

// GetRandomString 生成随机字符串
func GetRandomString(n int) string {
	str := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	bytes := []byte(str)
	var result []byte
	for i := 0; i < n; i++ {
		result = append(result, bytes[rand.Intn(len(bytes))])
	}
	return string(result)
}

// GetRandomString2 生成随机字符串
func GetRandomString2(n int) string {
	randBytes := make([]byte, n/2)
	rand.Read(randBytes)
	return fmt.Sprintf("%x", randBytes)
}

// ToLower 单词转化为小写
func ToLower(text []byte) []byte {
	output := make([]byte, len(text))
	for i, t := range text {
		if t >= 'A' && t <= 'Z' {
			output[i] = t - 'A' + 'a'
		} else {
			output[i] = t
		}
	}
	return output
}

// IsAlphabet 是否是字母
func IsAlphabet(r rune) bool {
	return (r >= 65 && r <= 90) || (r >= 97 && r <= 122)
}

// 是否是无效字符
func IsInvaildSymbol(str string) bool {
	special := "[,\\[\\]@#$%￥%……&\\(\\)'\\|?；：【】‘；：”“’。，、？ /.]"
	b, _ := regexp.MatchString(special, str)
	return b
}

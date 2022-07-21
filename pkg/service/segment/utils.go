package segment

import (
	"bytes"
	"comma/pkg/library/util"
	"fmt"
)

// SegmentsToSlice
//
// 	有两种输出模式，以"中华人民共和国"为例
//  	普通模式（searchType=false）输出一个分词："[中华人民共和国]"
//  	搜索模式（searchType=true） 输出普通模式的再细致切分："[中华 人民 共和 共和国 人民共和国 中华人民共和国]"
//
// 搜索模式主要用于给搜索引擎提供尽可能多的关键字，详情请见Token结构体的注释。
func SegmentsToSlice(segSlice []Segment, searchType bool, threeParticiple bool) (output []string) {
	var sliceStr []string
	if searchType {
		for _, seg := range segSlice {
			r := []rune(seg.token.GetTextSlice())
			if len(r) == 3 && threeParticiple == true {
				// 三个字的词不在进行二次分词
				sliceStr = append(sliceStr, seg.token.GetTextSlice())
				continue
			}
			t := tokenToSlice(seg.token)
			for _, s := range t {
				sliceStr = append(sliceStr, s)
			}
		}
	} else {
		for _, seg := range segSlice {
			sliceStr = append(sliceStr, textSliceToString(seg.token.textSlice))
		}
	}
	// 剔除重复的词 剔除非法字符 剔除stop
	for _, s := range sliceStr {
		if util.IsInvaildSymbol(s) {
			continue
		}
		output = append(output, s)
	}
	return
}

func SegmentsToString(segSlice []Segment, searchType bool) (output string) {
	if searchType {
		for _, seg := range segSlice {
			output += tokenToString(seg.token)
		}
	} else {
		for _, seg := range segSlice {
			output += fmt.Sprintf("%s/%s ", textSliceToString(seg.token.textSlice), seg.token.classify)
		}
	}
	return
}

func tokenToSlice(token *Token) (output []string) {
	output = append(output, token.GetTextSlice())
	hasOnlyTerminalToken := true
	for _, s := range token.segmentSlice {
		if len(s.token.segmentSlice) > 1 {
			hasOnlyTerminalToken = false
		}
	}
	if !hasOnlyTerminalToken {
		for _, s := range token.segmentSlice {
			if s != nil {
				if s.token.GetFrequency() > 1 {
					output = append(output, s.token.GetTextSlice())
				}
			}
		}
	}
	return
}

func tokenToString(token *Token) (output string) {
	hasOnlyTerminalToken := true
	for _, s := range token.segmentSlice {
		if len(s.token.segmentSlice) > 1 {
			hasOnlyTerminalToken = false
		}
	}
	if !hasOnlyTerminalToken {
		for _, s := range token.segmentSlice {
			if s != nil {
				output += tokenToString(s.token)
			}
		}
	}
	output += fmt.Sprintf("%s", textSliceToString(token.textSlice))
	return
}

// SegmentsToSlice2
//
// 	有两种输出模式，以"中华人民共和国"为例
//  	普通模式（searchType=false）输出一个分词："[中华人民共和国]"
//  	搜索模式（searchType=true） 输出普通模式的再细致切分："[中华 人民 共和 共和国 人民共和国 中华人民共和国]"
//
// 搜索模式主要用于给搜索引擎提供尽可能多的关键字，详情请见Token结构体的注释。
func SegmentsToSlice2(segSlice []Segment, searchType bool) (output []string) {
	if searchType {
		for _, seg := range segSlice {
			output = append(output, tokenToSlice2(seg.token)...)
		}
	} else {
		for _, seg := range segSlice {
			output = append(output, seg.token.GetTextSlice())
		}
	}
	return
}

func tokenToSlice2(token *Token) (output []string) {
	hasOnlyTerminalToken := true
	for _, s := range token.segmentSlice {
		if len(s.token.segmentSlice) > 1 {
			hasOnlyTerminalToken = false
		}
	}
	if !hasOnlyTerminalToken {
		for _, s := range token.segmentSlice {
			output = append(output, tokenToSlice2(s.token)...)
		}
	}
	output = append(output, textSliceToString(token.textSlice))
	return
}

// 将多个字元拼接一个字符串输出
func textSliceToString(textSlice []Text) string {
	return join(textSlice)
}

func join(textSlice []Text) string {
	switch len(textSlice) {
	case 0:
		return ""
	case 1:
		return string(textSlice[0])
	case 2:
		// Special case for common small values.
		// Remove if golang.org/issue/6714 is fixed
		return string(textSlice[0]) + string(textSlice[1])
	case 3:
		// Special case for common small values.
		// Remove if golang.org/issue/6714 is fixed
		return string(textSlice[0]) + string(textSlice[1]) + string(textSlice[2])
	}
	n := 0
	for i := 0; i < len(textSlice); i++ {
		n += len(textSlice[i])
	}
	b := make([]byte, n)
	bp := copy(b, textSlice[0])
	for _, s := range textSlice[1:] {
		bp += copy(b[bp:], s)
	}
	return string(b)
}

// 返回多个字元的字节总长度
func textSliceByteLength(textSlice []Text) (length int) {
	for _, word := range textSlice {
		length += len(word)
	}
	return
}

func textSliceToBytes(textSlice []Text) []byte {
	var buf bytes.Buffer
	for _, word := range textSlice {
		buf.Write(word)
	}
	return buf.Bytes()
}

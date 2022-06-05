package segment

import (
	"regexp"
	"strings"
	"sync"
)

type service struct {
	Segmenter *SegmenterService
}

const titleWeight = 0.2

var (
	s    *service
	once sync.Once
)

func GetInstance() *service {
	once.Do(func() {
		ps := GetTaskServiceInstance()
		s = &service{
			Segmenter: ps.Segmenter,
		}
	})
	return s
}

func (s *service) GetSegmentData(title, tags string) (*map[string]float32, *map[string][]string, *map[string][]string, *[]string) {
	// fmt.Println(strings.ToLower(word))
	re := regexp.MustCompile(`[\p{P}\p{S}\n]`) // 去除换行符以及所有符号
	// re := regexp.MustCompile(`([#@])|[\p{P}\p{S}\d]`) // 替换除＃和@以外的所有符号
	// re := regexp.MustCompile(`[\p{P}\p{S}\s\n\d]`) // 去除空格、去除换行符以及所有符号
	title = re.ReplaceAllString(title, "$1")
	tags = re.ReplaceAllString(tags, "$1")
	words := title + tags
	wordMap, synMap, hypMap, orderArray := s.Segmenter.SegmentIndexMode(words, false, 0, false)
	// 提高标题词权重
	for kw := range *wordMap {
		if strings.Contains(title, kw) {
			(*wordMap)[kw] = (*wordMap)[kw] + float32(titleWeight)
		}
	}
	//{
	//	arr := p.SegmenterService.SegmentSearchMode(words)
	//	for _, val := range arr {
	//		fmt.Print(val + "|")
	//	}
	//}
	return wordMap, synMap, hypMap, orderArray
}

func (s *service) AddWordToDict(word, classify string, frequency int) {
	// s.Segmenter.AddWordTokenToDict(word, classify, frequency)
}

func (s service) DelWordToDict(word string) {
}

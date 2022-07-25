package segment

import (
	"comma/pkg/library/util"
	"comma/pkg/library/util/text/t2c"
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
	"sync"
	"unicode"
	"unicode/utf8"

	"github.com/caicaispace/gohelper/logx"
)

const (
	minTokenFrequency = 2 // 仅从字典文件中读取大于等于此频率的分词
	maxTokenFrequency = 20000
	maxIdf            = 1.2
	minIdf            = 1.0
)

// SegmenterService 分词器结构体
type SegmenterService struct {
	t2s               *t2c.OpenCC
	dict              *Dictionary
	synonymsDict      []*map[string][]string // 同义词
	stopDict          *map[string]bool       // 停词 本代码粗略使用停词功能 本应在分词中处理 现在就当过滤掉一些停词
	bannedDict        *map[string]bool       // 违禁词
	hypernymDict      *map[string][]string   // 上位词 例如 mv  视频,音乐
	hyponymDict       *map[string][]string   // 下位词 例如 音乐 mp3,mv
	highFrequencyDict []*map[string]bool     // 高频词
	maxFrequency      int                    // 最高频
	maxIdf            float32                // 最大逆向文档频率
	minIdf            float32                // 最小逆向文档频率
}

// 该结构体用于记录Viterbi算法中某字元处的向前分词跳转信息
type jumper struct {
	minDistance float32
	token       *Token
}

var (
	segService     *SegmenterService
	segServiceOnce sync.Once
)

func GetSegmenterServiceInstance() *SegmenterService {
	segServiceOnce.Do(func() {
		segService = &SegmenterService{
			dict:         NewDictionary(),
			maxFrequency: maxTokenFrequency,
			maxIdf:       maxIdf,
			minIdf:       minIdf,
		}
	})
	return segService
}

// LoadDict
func (seg *SegmenterService) LoadDict() {
	seg.LoadDataToT2cDict()
	seg.LoadDataToDict()
	seg.LoadDataToBannedDict()
	seg.LoadDataToStopDict()
	seg.LoadDataToSynonymsDict()
	seg.LoadDataToHighFrequencyDict()
	seg.LoadDataToHyponymDict()
	logx.Debug("-------- dictionary load success --------")
}

// GetDict 返回分词器使用的词典
func (seg *SegmenterService) GetDict() *Dictionary {
	return seg.dict
}

// SetMinMaxIdf 设置逆向文档频率
func (seg *SegmenterService) SetMinMaxIdf(minIdf float32, maxIdf float32) {
	seg.minIdf = minIdf
	seg.maxIdf = maxIdf
}

// LoadDataToT2cDict 繁体转简体
func (seg *SegmenterService) LoadDataToT2cDict() {
	t2s, err := t2c.New("t2s")
	if err != nil {
		panic("LoadDataToT2cDict err")
	}
	seg.t2s = t2s
}

func (seg *SegmenterService) reloadCore() {
	// 计算每个分词的路径值，路径值含义见Token结构体的注释
	logTotalFrequency := float32(math.Log2(float64(seg.dict.totalFrequency)))
	for i := range seg.dict.tokens {
		token := &seg.dict.tokens[i]
		token.distance = logTotalFrequency - float32(math.Log2(float64(token.frequency)))
	}
	// 对每个分词进行细致划分，用于搜索引擎模式，该模式用法见Token结构体的注释。
	for i := range seg.dict.tokens {
		token := &seg.dict.tokens[i]
		segments := seg.segmentBatchWord(token.textSlice, true)
		// 计算需要添加的子分词数目
		tokensToAddCount := 0
		for iToken := 0; iToken < len(segments); iToken++ {
			if len(segments[iToken].token.textSlice) > 0 {
				tokensToAddCount++
			}
		}
		token.segmentSlice = make([]*Segment, tokensToAddCount)
		// 添加子分词
		iSegmentsToAdd := 0
		for iToken := 0; iToken < len(segments); iToken++ {
			if len(segments[iToken].token.textSlice) > 0 {
				token.segmentSlice[iSegmentsToAdd] = &segments[iToken]
				iSegmentsToAdd++
			}
		}
	}
}

func (seg *SegmenterService) AddWordToDict(word, classify string, frequency int) {
	seg.AddWordTokenToDict(word, classify, frequency)
	seg.reloadCore()
}

func (seg *SegmenterService) LoadDataToDict() {
	words, _ := LoadDictFromDB()
	for _, word := range words {
		seg.AddWordTokenToDict(word.Word, word.Classify, word.Frequency)
	}
	seg.reloadCore()
}

func (seg *SegmenterService) AddWordToBannedDict(word string) {
	(*seg.bannedDict)[word] = true
}

func (seg *SegmenterService) LoadDataToBannedDict() {
	seg.bannedDict, _ = LoadBannedDictFromDB()
}

func (seg *SegmenterService) AddWordToStopDict(word string) {
	(*seg.stopDict)[word] = true
}

func (seg *SegmenterService) LoadDataToStopDict() {
	seg.stopDict, _ = LoadStopDictFromDB()
}

func (seg *SegmenterService) LoadDataToSynonymsDict() {
	//project := LoadProjectFromDB()
	//var synonymsListT []*map[string][]string
	//for i := 0; i < project.LastId+1; i++ {
	//	synonymsListT = append(synonymsListT, nil)
	//}
	//for _, project := range project.List {
	//	//if len(*synonymsList[proId]) == 0 {
	//	//	logger.Debug("-------- v2 empty use v3 --------")
	//	//	synonymsList[proId], _ = QuerySynonymsDicV3(proId)
	//	//}
	//	synonymsListT[project.ID], _ = LoadSynonymsDictFromDB(project.ID)
	//}
	//seg.synonymsDict = synonymsListT
	project := LoadProjectFromDB()
	seg.synonymsDict = make([]*map[string][]string, project.LastId+1)
	for _, project := range project.List {
		//if len(*synonymsList[proId]) == 0 {
		//	logger.Debug("-------- v2 empty use v3 --------")
		//	synonymsList[proId], _ = QuerySynonymsDicV3(proId)
		//}
		seg.synonymsDict[project.ID], _ = LoadSynonymsDictFromDB(project.ID)
	}
}

func (seg *SegmenterService) LoadDataToHighFrequencyDict() {
	//project := LoadProjectFromDB()
	//var highFrequencyListT []*map[string]bool
	//for i := 0; i < project.LastId+1; i++ {
	//	highFrequencyListT = append(highFrequencyListT, nil)
	//}
	//for _, project := range project.List {
	//	highFrequencyListT[project.ID], _ = LoadHighFrequencyDictFromDB(project.ID)
	//}
	//seg.highFrequencyDict = highFrequencyListT
	project := LoadProjectFromDB()
	seg.highFrequencyDict = make([]*map[string]bool, project.LastId+1)
	for _, project := range project.List {
		seg.highFrequencyDict[project.ID], _ = LoadHighFrequencyDictFromDB(project.ID)
	}
}

func (seg *SegmenterService) LoadDataToHyponymDict() {
	seg.hyponymDict, _ = LoadHyponymDictFromDB()
}

// AddWordTokenToDict 将分词添加到字典中
func (seg *SegmenterService) AddWordTokenToDict(word, classify string, frequency int) {
	// 过滤频率太小的词
	if frequency < minTokenFrequency {
		return
	}
	if frequency > seg.maxFrequency {
		seg.maxFrequency = frequency
	}
	seg.dict.addToken(Token{
		textSlice: splitTextToWords([]byte(word)),
		frequency: frequency,
		classify:  classify,
	})
}

// Segment 文本分词
func (seg *SegmenterService) Segment(wordByte []byte) []Segment {
	return seg.segmentInternal(wordByte, false)
}

// SegmentInternal 内部分词
func (seg *SegmenterService) SegmentInternal(wordByte []byte, searchMode bool) []Segment {
	return seg.segmentInternal(wordByte, searchMode)
}

// SegmentIndexMode 索引模式分词
// 增加 threeParticiple 三个汉字的是否分词
// 依次返回正常分词 同义词  下位词（逗号分隔）
// 增加项目 id 如果没有 默认填写 0
// 逆向文档频率 https://www.elastic.co/guide/cn/elasticsearch/guide/current/scoring-theory.html
func (seg *SegmenterService) SegmentIndexMode(
	str string,
	threeParticiple bool,
	projectId int,
	idf bool,
) (
	*map[string]float32,
	*map[string][]string,
	*map[string][]string,
	*[]string,
) {
	segments := seg.segmentInternal([]byte(seg.preDelString(str)), false)
	segmentSlice := SegmentsToSlice(segments, true, threeParticiple)
	wordMap := make(map[string]float32) // 带权重的词
	synMap := make(map[string][]string) // 同义词
	hypMap := make(map[string][]string) // 下位词、上位词
	orderSlice := make([]string, 0)     // 搜索词排序
	// 进行二次处理
	for _, v := range segmentSlice {
		if _, ok := wordMap[v]; ok {
			continue
		}
		if _, ok := (*seg.stopDict)[v]; ok {
			continue
		}
		wordMap[v] = 1
		if idf {
			wordMap[v] = seg.GetIdfValue(v)
		}
		orderSlice = append(orderSlice, v)
		// 同义词
		for sK, sV := range seg.synonymsDict {
			if sV == nil {
				continue
			}
			if projectId != 0 && sK == projectId {
				if sv, ok := (*sV)[v]; ok {
					synMap[v] = sv
				}
				break
			}
			if projectId == 0 {
				if sv, ok := (*sV)[v]; ok {
					synMap[v] = sv
					break
				}
			}
		}
		// 下位词、上位词
		if hv, ok := (*seg.hyponymDict)[v]; ok {
			hypMap[v] = hv
		}
	}
	return &wordMap, &synMap, &hypMap, &orderSlice
}

// SegmentSearchMode 搜索模式分词
func (seg *SegmenterService) SegmentSearchMode(str string) []string {
	out := seg.preDelString(str)
	segments := seg.segmentInternal([]byte(out), false)
	return SegmentsToSlice(segments, false, true)
}

// GetFreValue 获取词频值
func (seg *SegmenterService) GetFreValue(str string) int {
	id, err := seg.dict.trie.Get([]byte(str))
	if err != nil {
		return 0
	}
	return seg.dict.tokens[id].frequency
}

// GetIdfValue 获取逆向文档频率值
func (seg *SegmenterService) GetIdfValue(str string) float32 {
	value := seg.GetFreValue(str)
	idf := (1 - float32(value)/float32(seg.maxFrequency)) * (seg.maxIdf - seg.minIdf)
	idf2, _ := strconv.ParseFloat(fmt.Sprintf("%.3f", idf), 32)
	return seg.minIdf + float32(idf2)
}

// IsStop 是否是停词
func (seg *SegmenterService) IsStop(word string) bool {
	_, ok := (*seg.stopDict)[word]
	return ok
}

// IsHighFrequency 查询是否是高频词
// 入参 参数第一参数 查询的词
// 入参 参数第二参数 项目id (没有项目组可以填写为 0, 0 默认查询所有的)
func (seg *SegmenterService) IsHighFrequency(word string, projectId uint8) bool {
	for i, v := range seg.highFrequencyDict {
		if i == int(projectId) || projectId == 0 {
			if v == nil {
				return false
			}
			if _, ok := (*v)[word]; ok {
				return true
			}
			if projectId != 0 {
				return false
			}
		}
	}
	return false
}

// IsBanned 是否是违禁词
func (seg *SegmenterService) IsBanned(word string) bool {
	_, ok := (*seg.bannedDict)[word]
	return ok
}

func (seg *SegmenterService) segmentInternal(byteSlice []byte, searchMode bool) []Segment {
	// 处理特殊情况
	if len(byteSlice) == 0 {
		return []Segment{}
	}
	// 划分字元
	text := splitTextToWords(byteSlice)
	return seg.segmentBatchWord(text, searchMode)
}

// 批量分词
func (seg *SegmenterService) segmentBatchWord(textSlice []Text, searchMode bool) []Segment {
	// 搜索模式下该分词已无继续划分可能的情况
	if searchMode && len(textSlice) == 1 {
		return []Segment{}
	}
	// jumpers定义了每个字元处的向前跳转信息，包括这个跳转对应的分词，
	// 以及从文本段开始到该字元的最短路径值
	jumperArr := make([]jumper, len(textSlice))
	tokenArr := make([]*Token, seg.dict.maxTokenLength)
	for current := 0; current < len(textSlice); current++ {
		// 找到前一个字元处的最短路径，以便计算后续路径值
		var baseDistance float32
		if current == 0 {
			// 当本字元在文本首部时，基础距离应该是零
			baseDistance = 0
		} else {
			baseDistance = jumperArr[current-1].minDistance
		}
		// 寻找所有以当前字元开头的分词
		tokenCount := seg.dict.getTokenCount(textSlice[current:util.MinInt(current+seg.dict.maxTokenLength, len(textSlice))], tokenArr)
		// 对所有可能的分词，更新分词结束字元处的跳转信息
		for iToken := 0; iToken < tokenCount; iToken++ {
			location := current + len(tokenArr[iToken].textSlice) - 1
			if !searchMode || current != 0 || location != len(textSlice)-1 {
				updateJumper(&jumperArr[location], baseDistance, tokenArr[iToken])
			}
		}
		// 当前字元没有对应分词时补加一个伪分词
		if tokenCount == 0 || len(tokenArr[0].textSlice) > 1 {
			updateJumper(
				&jumperArr[current],
				baseDistance,
				&Token{
					textSlice: []Text{textSlice[current]},
					frequency: 1,
					distance:  32,
					classify:  "x",
				},
			)
		}
	}
	// 从后向前扫描第一遍得到需要添加的分词数目
	numSeg := 0
	for index := len(textSlice) - 1; index >= 0; {
		location := index - len(jumperArr[index].token.textSlice) + 1
		numSeg++
		index = location - 1
	}
	// 从后向前扫描第二遍添加分词到最终结果
	outData := make([]Segment, numSeg)
	for index := len(textSlice) - 1; index >= 0; {
		location := index - len(jumperArr[index].token.textSlice) + 1
		numSeg--
		outData[numSeg].token = jumperArr[index].token
		index = location - 1
	}
	// 计算各个分词的字节位置
	bytePosition := 0
	for iSeg := 0; iSeg < len(outData); iSeg++ {
		outData[iSeg].start = bytePosition
		bytePosition += textSliceByteLength(outData[iSeg].token.textSlice)
		outData[iSeg].end = bytePosition
	}
	return outData
}

// 	更新跳转信息:
// 	1. 当该位置从未被访问过时(jumper.minDistance为零的情况)，或者
//	2. 当该位置的当前最短路径大于新的最短路径时
// 	将当前位置的最短路径值更新为baseDistance加上新分词的概率
func updateJumper(jumper *jumper, baseDistance float32, token *Token) {
	newDistance := baseDistance + token.distance
	if jumper.minDistance == 0 || jumper.minDistance > newDistance {
		jumper.minDistance = newDistance
		jumper.token = token
	}
}

// 将文本划分成字元
func splitTextToWords(text Text) []Text {
	outData := make([]Text, 0, len(text)/3)
	current := 0
	inAlphanumeric := true
	alphanumericStart := 0
	for current < len(text) {
		r, size := utf8.DecodeRune(text[current:])
		if size <= 2 && (unicode.IsLetter(r) || unicode.IsNumber(r)) {
			// 当前是拉丁字母或数字（非中日韩文字）
			if !inAlphanumeric {
				alphanumericStart = current
				inAlphanumeric = true
			}
		} else {
			if inAlphanumeric {
				inAlphanumeric = false
				if current != 0 {
					outData = append(outData, util.ToLower(text[alphanumericStart:current]))
				}
			}
			outData = append(outData, text[current:current+size])
		}
		current += size
	}
	// 处理最后一个字元是英文的情况
	if inAlphanumeric {
		if current != 0 {
			outData = append(outData, util.ToLower(text[alphanumericStart:current]))
		}
	}
	return outData
}

func (seg *SegmenterService) preDelString(in string) string {
	// 进行简繁体转化
	out, err := seg.t2s.Convert(in)
	if err != nil {
		out = in
		log.Fatal(err)
	}
	r := []rune(out)
	// 中文英文之间加上空格
	if len(r) <= 1 {
		return out
	}
	preBool := false
	if unicode.IsLetter(r[0]) || unicode.IsNumber(r[0]) {
		preBool = true
	}
	var outData []rune
	var preAplNum string
	for i := 0; i < len(r); i++ {
		t := r[i]
		if util.IsAlphabet(t) || unicode.IsNumber(t) {
			// 这里功能就是为了让数字和字母分开 但是如果这是一个词 那就不能分开
			if preBool && i != 0 {
				if (util.IsAlphabet(t) && unicode.IsNumber(r[i-1])) || (util.IsAlphabet(r[i-1]) && unicode.IsNumber(t)) {
					ids := seg.dict.trie.PrefixPredict([]byte(strings.ToLower(preAplNum+string(t))), 0)
					if len(ids) > 0 {
						maxFindStr := ""
						for j := 0; j < len(ids); j++ {
							m, _ := seg.dict.trie.Key(ids[j])
							isCon := strings.Contains(strings.ToLower(in), string(m))
							if isCon {
								if len(string(m)) > len(maxFindStr) {
									// 最长匹配
									maxFindStr = string(m)
								}
							}
						}
						if maxFindStr == "" {
							outData = append(outData, ' ')
							preAplNum = ""
						} else if maxFindStr == (preAplNum + string(t)) {
							preAplNum = ""
						} else {
							// 继续最长匹配 数字字母的最长匹配
						}
					} else {
						outData = append(outData, ' ')
						preAplNum = ""
					}
				}
			}
			preAplNum = preAplNum + string(t)
			preBool = true
		} else {
			preBool = false
			preAplNum = ""
		}
		outData = append(outData, t)
	}
	return string(outData)
}

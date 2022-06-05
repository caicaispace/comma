package segment

// Text
//	字串类型，可以用来表达
//	1. 一个字元，比如"中"又如"国", 英文的一个字元是一个词
//	2. 一个分词，比如"中国"又如"人口"
//	3. 一段文字，比如"中国有十三亿人口"
type Text []byte

// Token 一个分词
type Token struct {
	// 分词的字串，这实际上是个字元数组
	textSlice []Text
	// 分词在语料库中的词频
	frequency int
	// log2(总词频/该分词词频)，这相当于log2(1/p(分词))，用作动态规划中
	// 该分词的路径长度。求解prod(p(分词))的最大值相当于求解
	// sum(distance(分词))的最小值，这就是“最短路径”的来历。
	distance float32
	// 词性标注
	classify string
	// 该分词文本的进一步分词划分，见Segments函数注释。
	segmentSlice []*Segment
}

// GetTextSlice 返回分词文本
func (token *Token) GetTextSlice() string {
	return textSliceToString(token.textSlice)
}

// GetFrequency 返回分词在语料库中的词频
func (token *Token) GetFrequency() int {
	return token.frequency
}

// GetClassify 返回分词词性标注
func (token *Token) GetClassify() string {
	return token.classify
}

// GetSegmentSlice 该分词文本的进一步分词划分，比如"中华人民共和国中央人民政府"这个分词
// 有两个子分词"中华人民共和国"和"中央人民政府"。子分词也可以进一步有子分词
// 形成一个树结构，遍历这个树就可以得到该分词的所有细致分词划分，这主要
// 用于搜索引擎对一段文本进行全文搜索。
func (token *Token) GetSegmentSlice() []*Segment {
	return token.segmentSlice
}

func (token *Token) TextEquals(str string) bool {
	tokenLen := 0
	for _, t := range token.textSlice {
		tokenLen += len(t)
	}
	if tokenLen != len(str) {
		return false
	}
	byteStr := []byte(str)
	index := 0
	for i := 0; i < len(token.textSlice); i++ {
		textByte := []byte(token.textSlice[i])
		for j := 0; j < len(textByte); j++ {
			if textByte[j] != byteStr[index] {
				index = index + 1
				return false
			}
			index = index + 1
		}
	}
	return true
}

package segment

import (
	tda "goaway/pkg/library/util/tree/trie-double-array"
)

// Dictionary 结构体实现了一个字串前缀树，一个分词可能出现在叶子节点也有可能出现在非叶节点
type Dictionary struct {
	trie           *tda.Cedar // Cedar 前缀树
	tokens         []Token    // 词典中所有的分词，方便遍历
	maxTokenLength int        // 词典中最长的分词
	totalFrequency int64      // 词典中所有分词的频率之和
}

func NewDictionary() *Dictionary {
	return &Dictionary{trie: tda.New()}
}

// GetMaxTokenLength 词典中最长的分词
func (dict *Dictionary) GetMaxTokenLength() int {
	return dict.maxTokenLength
}

// GetTokensLength 词典中分词数目
func (dict *Dictionary) GetTokensLength() int {
	return len(dict.tokens)
}

// GetTotalFrequency 词典中所有分词的频率之和
func (dict *Dictionary) GetTotalFrequency() int64 {
	return dict.totalFrequency
}

// GetTokenFrequency 词典中所对应分词的频率
func (dict *Dictionary) GetTokenFrequency(index int) int {
	return dict.tokens[index].frequency
}

// 向词典中加入一个分词
func (dict *Dictionary) addToken(token Token) {
	bytes := textSliceToBytes(token.textSlice)
	_, err := dict.trie.Get(bytes)
	if err == nil {
		return
	}
	dict.trie.Insert(bytes, dict.GetTokensLength())
	dict.tokens = append(dict.tokens, token)
	dict.totalFrequency += int64(token.frequency)
	if len(token.textSlice) > dict.maxTokenLength {
		dict.maxTokenLength = len(token.textSlice)
	}
}

// 在词典中查找和字元组 wordSlice 可以前缀匹配的所有分词,返回值为找到的分词数
func (dict *Dictionary) getTokenCount(wordSlice []Text, tokenSlice []*Token) (TokensCount int) {
	var id, value int
	var err error
	for _, word := range wordSlice {
		id, err = dict.trie.Jump(word, id)
		if err != nil {
			break
		}
		value, err = dict.trie.Value(id)
		if err == nil {
			tokenSlice[TokensCount] = &dict.tokens[value]
			TokensCount++
		}
	}
	return
}

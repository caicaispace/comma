package segment

// Segment 文本中的一个分词
type Segment struct {
	start int    // 分词在文本中的起始字节位置
	end   int    // 分词在文本中的结束字节位置（不包括该位置）
	token *Token // 分词信息
}

// Start 返回分词在文本中的起始字节位置
func (s *Segment) Start() int {
	return s.start
}

// End 返回分词在文本中的结束字节位置（不包括该位置）
func (s *Segment) End() int {
	return s.end
}

// Token 返回分词信息
func (s *Segment) Token() *Token {
	return s.token
}
package segment

import (
	"sync"

	service "comma/pkg/service/segment"
)

type Segment struct{}

type SegmentParams struct {
	Title string `json:"title"`
	Tags  string `json:"tags"`
}

type SegmentResult struct {
	Word  *map[string]float32  `json:"word"`
	Syn   *map[string][]string `json:"syn"`
	Hyp   *map[string][]string `json:"hyp"`
	Order *[]string            `json:"order"`
}

var (
	s    *Segment
	once sync.Once
)

func GetInstance() *Segment {
	once.Do(func() {
		service.GetInstance()
		s = &Segment{}
	})
	return s
}

// Segment 分词
func (p *Segment) Segment(params *SegmentParams, result *SegmentResult) error {
	ps := service.GetInstance()
	wordMap, synMap, hypMap, orderArray := ps.GetSegmentData(params.Title, params.Tags)
	rspData := SegmentResult{
		Word:  wordMap,
		Syn:   synMap,
		Hyp:   hypMap,
		Order: orderArray,
	}
	*result = interface{}(rspData).(SegmentResult)
	return nil
}

type BatchSegmentParams struct {
	List map[string]SegmentParams `json:"list"`
}

type BatchSegmentResult struct {
	List map[string]SegmentResult `json:"list"`
}

// BatchSegment 批量分词
func (p *Segment) BatchSegment(params *BatchSegmentParams, result *BatchSegmentResult) error {
	ps := service.GetInstance()
	listMap := make(map[string]SegmentResult)
	for id, item := range params.List {
		wordMap, synMap, hypMap, orderArray := ps.GetSegmentData(item.Title, item.Tags)
		listMap[id] = SegmentResult{
			Word:  wordMap,
			Syn:   synMap,
			Hyp:   hypMap,
			Order: orderArray,
		}
	}
	rspData := BatchSegmentResult{
		List: listMap,
	}
	*result = interface{}(rspData).(BatchSegmentResult)
	return nil
}

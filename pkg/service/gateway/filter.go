package gateway

import (
	"regexp"
	"strings"

	"comma/pkg/library/util/config"

	segmentService "comma/pkg/service/segment"
)

type SEARCH_TYPE int

const (
	PROXY_SEARCH_TYPE SEARCH_TYPE = iota
	PASS_SEARCH_TYPE
	STANDARD_SEARCH_TYPE
)

type FilterService struct {
	Segmenter *segmentService.SegmenterService
}

type PayLoadFilter struct {
	Segmenter *segmentService.SegmenterService
}

type IFilter interface {
	FilterRequest(er *EsRequest)
}

type EsRequest struct {
	SearchType    SEARCH_TYPE
	RequestUrl    string
	RequestMethod string
	IndexName     string
	TypeName      string
	BodyRaw       string
	Banned        string
	ProjectId     uint8
}

var (
	payLoadFilterTypeCommon  *PayLoadFilter
	payLoadFilterTypePayload *PayLoadFilter
)

// GetFilterInstance
// 0 common
// 1 payload
func (fs *FilterService) GetFilterInstance(indexName string, typeName string) IFilter {
	filterType := config.GetInstance().GetEsFilterType(indexName, typeName)
	switch filterType {
	case 0:
		payLoadFilterTypeCommon = &PayLoadFilter{
			Segmenter: fs.Segmenter,
		}
		return payLoadFilterTypeCommon
	case 1:
		payLoadFilterTypePayload = &PayLoadFilter{
			Segmenter: fs.Segmenter,
		}
		return payLoadFilterTypePayload
	default:
		return nil
	}
}

func (pf *PayLoadFilter) FilterRequest(er *EsRequest) {
	er.Banned = ""
	body := er.BodyRaw

	reg := regexp.MustCompile(`"query"\s*:\s*"([^"]*)"`)
	matches := reg.FindAllSubmatch([]byte(er.BodyRaw), -1)

	hasHighFre := false
	hanCount := 0
	letCount := 0
	numCount := 0

	var termSlice []string
	var bannedSlice []string
	for i := 0; i < len(matches); i++ {
		segWords := pf.Segmenter.SegmentSearchMode(string(matches[i][1])) // 搜索词分词
		if len(segWords) == 0 {
			continue
		}
		termSlice = make([]string, 0)
		bannedSlice = make([]string, 0)
		for _, word := range segWords {
			// stop
			if pf.Segmenter.IsStop(word) {
				continue
			}
			// banned
			if pf.Segmenter.IsBanned(word) {
				bannedSlice = append(bannedSlice, word)
				continue
			}
			termSlice = append(termSlice, word)
			// High frequency
			if pf.Segmenter.IsHighFrequency(word, er.ProjectId) {
				hasHighFre = true
			}
		}
		r := []rune(string(matches[i][1]))
		for i := 0; i < len(r); i++ {
			if r[i] <= 40869 && r[i] >= 19968 {
				hanCount++
			} else if (64 < r[i] && r[i] < 91) || (96 < r[i] && r[i] < 123) {
				letCount++
			} else if 47 < r[i] && r[i] < 58 {
				numCount++
			}
		}
		er.Banned = strings.Join(bannedSlice, " ")
		body = strings.Replace(body, string(matches[i][1]), strings.Join(termSlice, " "), -1)
	}
	if hasHighFre || (letCount > 0 || numCount > 0) || (hanCount >= 20) {
		body = strings.Replace(body, "50%", "100%", -1)
	}
	er.BodyRaw = body
}

package trie_double_array_dict

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	tda "goaway/pkg/library/util/tree/trie_double_array"
)

// BuildFromFile builds the da dict from fileName
func BuildFromFile(fileName string) (*Dict, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return Build(file)
}

// Build da dict from io.Reader
func Build(in io.Reader) (*Dict, error) {
	trie := tda.New()
	var values [][]string
	br := bufio.NewReader(in)
	for {
		line, err := br.ReadString('\n')
		if err == io.EOF {
			break
		}
		items := strings.Split(strings.TrimSpace(line), "\t")
		if len(items) < 2 {
			continue
		}
		err = trie.Insert([]byte(items[0]), len(values))
		if err != nil {
			return nil, err
		}

		if len(items) > 2 {
			values = append(values, items[1:])
		} else {
			values = append(values, strings.Fields(items[1]))
		}
	}
	return &Dict{Trie: trie, Values: values}, nil
}

// PrefixMatch str by Dict, returns the matched string and its according values
func (d *Dict) PrefixMatch(str string) (map[string][]string, error) {
	if d.Trie == nil {
		return nil, fmt.Errorf("Trie is nil")
	}
	ret := make(map[string][]string)
	for _, id := range d.Trie.PrefixMatch([]byte(str), 0) {
		key, err := d.Trie.Key(id)
		if err != nil {
			return nil, err
		}
		value, err := d.Trie.Value(id)
		if err != nil {
			return nil, err
		}
		ret[string(key)] = d.Values[value]
	}
	return ret, nil
}

// Get the values of str, like map
func (d *Dict) Get(str string) ([]string, error) {
	if d.Trie == nil {
		return nil, fmt.Errorf("trie is nil")
	}
	id, err := d.Trie.Get([]byte(str))
	if err != nil {
		return nil, err
	}
	value, err := d.Trie.Value(id)
	if err != nil {
		return nil, err
	}
	return d.Values[value], nil
}

package trie_double_array_dict

import (
	"bufio"
	"encoding/gob"
	"os"
	"path/filepath"

	tda "comma/pkg/library/util/tree/trie_double_array"
)

// Load gob serialized dict from dir
func Load(dir string) (*Dict, error) {
	trieFile := filepath.Join(dir, "trie")
	valueFile := filepath.Join(dir, "values")
	trie := tda.New()
	if err := trie.LoadFromFile(trieFile, "gob"); err != nil {
		return nil, err
	}
	file, err := os.OpenFile(valueFile, os.O_RDONLY, 0o600)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	in := bufio.NewReader(file)
	dataDecoder := gob.NewDecoder(in)
	var values [][]string
	if err = dataDecoder.Decode(&values); err != nil {
		return nil, err
	}
	return &Dict{Trie: trie, Values: values}, nil
}

// Save gob serialized dict to dir
func (d *Dict) Save(dir string) error {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		os.Mkdir(dir, 0o755)
	}
	trieFile := filepath.Join(dir, "trie")
	valueFile := filepath.Join(dir, "values")
	if err := d.Trie.SaveToFile(trieFile, "gob"); err != nil {
		return err
	}
	file, err := os.OpenFile(valueFile, os.O_CREATE|os.O_WRONLY, 0o666)
	if err != nil {
		return err
	}
	defer file.Close()
	out := bufio.NewWriter(file)
	defer out.Flush()
	dataEncoder := gob.NewEncoder(out)
	return dataEncoder.Encode(d.Values)
}

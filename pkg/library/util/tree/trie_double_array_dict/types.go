package trie_double_array_dict

import (
	tda "comma/pkg/library/util/tree/trie_double_array"
)

// Dict contains the Trie and dict values
type Dict struct {
	Trie   *tda.Cedar
	Values [][]string
}

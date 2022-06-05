package trie_double_array_dict

import (
	tda "goaway/pkg/library/util/tree/trie-double-array"
)

// Dict contains the Trie and dict values
type Dict struct {
	Trie   *tda.Cedar
	Values [][]string
}

package t

// Frequently-used map alias.
type (
	Map            = map[string]interface{}
	MapAnyAny      = map[interface{}]interface{}
	MapAnyStr      = map[interface{}]string
	MapAnyInt      = map[interface{}]int
	MapStrAny      = map[string]interface{}
	MapStrStr      = map[string]string
	MapStrInt      = map[string]int
	MapIntAny      = map[int]interface{}
	MapIntStr      = map[int]string
	MapIntInt      = map[int]int
	MapAnyBool     = map[interface{}]bool
	MapStrBool     = map[string]bool
	MapIntBool     = map[int]bool
	MapStrSliceStr = map[string][]string
)

// Frequently-used slice alias.
type (
	SliceMap        = []Map
	SliceMapAnyAny  = []MapAnyAny
	SliceMapAnyStr  = []MapAnyStr
	SliceMapAnyInt  = []MapAnyInt
	SliceMapStrAny  = []MapStrAny
	SliceMapStrStr  = []MapStrStr
	SliceMapStrInt  = []MapStrInt
	SliceMapIntAny  = []MapIntAny
	SliceMapIntStr  = []MapIntStr
	SliceMapIntInt  = []MapIntInt
	SliceMapAnyBool = []MapAnyBool
	SliceMapStrBool = []MapStrBool
	SliceMapIntBool = []MapIntBool
)

// Frequently-used slice alias.
type (
	Slice     = []interface{}
	SliceAny  = []interface{}
	SliceStr  = []string
	SliceInt  = []int
	SliceBool = []bool
)

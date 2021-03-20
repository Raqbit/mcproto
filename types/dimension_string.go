// Code generated by "stringer -type=Dimension -output dimension_string.go -linecomment"; DO NOT EDIT.

package types

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[DimensionNether - -1]
	_ = x[DimensionOverworld-0]
	_ = x[DimensionEnd-1]
}

const _Dimension_name = "NetherOverworldEnd"

var _Dimension_index = [...]uint8{0, 6, 15, 18}

func (i Dimension) String() string {
	i -= -1
	if i < 0 || i >= Dimension(len(_Dimension_index)-1) {
		return "Dimension(" + strconv.FormatInt(int64(i+-1), 10) + ")"
	}
	return _Dimension_name[_Dimension_index[i]:_Dimension_index[i+1]]
}
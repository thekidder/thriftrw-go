// Code generated by "stringer -type=BaseTypeID"; DO NOT EDIT

package ast

import "fmt"

const _BaseTypeID_name = "BoolTypeIDByteTypeIDI16TypeIDI32TypeIDI64TypeIDDoubleTypeIDStringTypeIDBinaryTypeID"

var _BaseTypeID_index = [...]uint8{0, 10, 20, 29, 38, 47, 59, 71, 83}

func (i BaseTypeID) String() string {
	i -= 1
	if i < 0 || i >= BaseTypeID(len(_BaseTypeID_index)-1) {
		return fmt.Sprintf("BaseTypeID(%d)", i+1)
	}
	return _BaseTypeID_name[_BaseTypeID_index[i]:_BaseTypeID_index[i+1]]
}

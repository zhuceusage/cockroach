// Code generated by "stringer -type=prepareType"; DO NOT EDIT

package pgwire

import "fmt"

const (
	_prepareType_name_0 = "preparePortal"
	_prepareType_name_1 = "prepareStatement"
)

var (
	_prepareType_index_0 = [...]uint8{0, 13}
	_prepareType_index_1 = [...]uint8{0, 16}
)

func (i prepareType) String() string {
	switch {
	case i == 80:
		return _prepareType_name_0
	case i == 83:
		return _prepareType_name_1
	default:
		return fmt.Sprintf("prepareType(%d)", i)
	}
}
// Code generated by "enumer -type=AddressFlag -linecomment -text"; DO NOT EDIT.

//
package nethelpers

import (
	"fmt"
)

const _AddressFlagName = "temporarynodadoptimisticdadfailedhomeaddressdeprecatedtentativepermanentmngmtmpaddrnoprefixroutemcautojoinstableprivacy"

var _AddressFlagMap = map[AddressFlag]string{
	1:    _AddressFlagName[0:9],
	2:    _AddressFlagName[9:14],
	4:    _AddressFlagName[14:24],
	8:    _AddressFlagName[24:33],
	16:   _AddressFlagName[33:44],
	32:   _AddressFlagName[44:54],
	64:   _AddressFlagName[54:63],
	128:  _AddressFlagName[63:72],
	256:  _AddressFlagName[72:83],
	512:  _AddressFlagName[83:96],
	1024: _AddressFlagName[96:106],
	2048: _AddressFlagName[106:119],
}

func (i AddressFlag) String() string {
	if str, ok := _AddressFlagMap[i]; ok {
		return str
	}
	return fmt.Sprintf("AddressFlag(%d)", i)
}

var _AddressFlagValues = []AddressFlag{1, 2, 4, 8, 16, 32, 64, 128, 256, 512, 1024, 2048}

var _AddressFlagNameToValueMap = map[string]AddressFlag{
	_AddressFlagName[0:9]:     1,
	_AddressFlagName[9:14]:    2,
	_AddressFlagName[14:24]:   4,
	_AddressFlagName[24:33]:   8,
	_AddressFlagName[33:44]:   16,
	_AddressFlagName[44:54]:   32,
	_AddressFlagName[54:63]:   64,
	_AddressFlagName[63:72]:   128,
	_AddressFlagName[72:83]:   256,
	_AddressFlagName[83:96]:   512,
	_AddressFlagName[96:106]:  1024,
	_AddressFlagName[106:119]: 2048,
}

// AddressFlagString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func AddressFlagString(s string) (AddressFlag, error) {
	if val, ok := _AddressFlagNameToValueMap[s]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to AddressFlag values", s)
}

// AddressFlagValues returns all values of the enum
func AddressFlagValues() []AddressFlag {
	return _AddressFlagValues
}

// IsAAddressFlag returns "true" if the value is listed in the enum definition. "false" otherwise
func (i AddressFlag) IsAAddressFlag() bool {
	_, ok := _AddressFlagMap[i]
	return ok
}

// MarshalText implements the encoding.TextMarshaler interface for AddressFlag
func (i AddressFlag) MarshalText() ([]byte, error) {
	return []byte(i.String()), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface for AddressFlag
func (i *AddressFlag) UnmarshalText(text []byte) error {
	var err error
	*i, err = AddressFlagString(string(text))
	return err
}

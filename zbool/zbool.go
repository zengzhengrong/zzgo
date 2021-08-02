package zbool

import "errors"

// BoolFlag is type of map[string]bool
type BoolFlag map[string]bool

// BoolFlagMap is mapping bool
var BoolFlagMap = BoolFlag{
	"1":     true,
	"true":  true,
	"True":  true,
	"0":     false,
	"false": false,
	"False": false,
}

// BoundCheck is return false and err if is not found
func (b BoolFlag) BoundCheck(flag string) (bool, error) {
	bl, ok := b[flag]
	if !ok {
		return false, errors.New("flag not mapping any key")
	}
	return bl, nil
}

// Check is not found return flase
func (b BoolFlag) Check(flag string) bool {
	bl, ok := b[flag]
	if !ok {
		return false
	}
	return bl
}

// Code generated by "stringer -type=TextDecorations"; DO NOT EDIT.

package gi

import (
	"errors"
	"strconv"
)

var _ = errors.New("dummy error")

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[DecoNone-0]
	_ = x[DecoUnderline-1]
	_ = x[DecoOverline-2]
	_ = x[DecoLineThrough-3]
	_ = x[DecoBlink-4]
	_ = x[DecoDottedUnderline-5]
	_ = x[DecoParaStart-6]
	_ = x[DecoSuper-7]
	_ = x[DecoSub-8]
	_ = x[DecoBgColor-9]
	_ = x[TextDecorationsN-10]
}

const _TextDecorations_name = "DecoNoneDecoUnderlineDecoOverlineDecoLineThroughDecoBlinkDecoDottedUnderlineDecoParaStartDecoSuperDecoSubDecoBgColorTextDecorationsN"

var _TextDecorations_index = [...]uint8{0, 8, 21, 33, 48, 57, 76, 89, 98, 105, 116, 132}

func (i TextDecorations) String() string {
	if i < 0 || i >= TextDecorations(len(_TextDecorations_index)-1) {
		return "TextDecorations(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _TextDecorations_name[_TextDecorations_index[i]:_TextDecorations_index[i+1]]
}

func (i *TextDecorations) FromString(s string) error {
	for j := 0; j < len(_TextDecorations_index)-1; j++ {
		if s == _TextDecorations_name[_TextDecorations_index[j]:_TextDecorations_index[j+1]] {
			*i = TextDecorations(j)
			return nil
		}
	}
	return errors.New("String: " + s + " is not a valid option for type: TextDecorations")
}

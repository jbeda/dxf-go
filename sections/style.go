package sections

import (
	"fmt"
	"github.com/rpaloschi/dxf-go/core"
)

const verticalTextBit = 0x4
const shapeBit = 0x1

const backwardsBit = 0x2
const upsideDownBit = 0x4

type Style struct {
	Name           string
	Height         float64
	Width          float64
	Oblique        float64
	IsBackwards    bool
	IsUpsideDown   bool
	IsShape        bool
	IsVerticalText bool
	Font           string
	BigFont        string
}

func NewStyle(tags core.TagSlice) *Style {
	style := new(Style)

	style.Height = 1.0
	style.Width = 1.0

	for _, tag := range tags.RegularTags() {
		switch tag.Code {
		case 2:
			style.Name, _ = core.AsString(tag.Value)

		case 3:
			style.Font, _ = core.AsString(tag.Value)

		case 4:
			style.BigFont, _ = core.AsString(tag.Value)

		case 40:
			style.Height, _ = core.AsFloat(tag.Value)

		case 41:
			style.Width, _ = core.AsFloat(tag.Value)

		case 50:
			style.Oblique, _ = core.AsFloat(tag.Value)

		case 70:
			flags, _ := core.AsInt(tag.Value)
			style.IsShape = flags&shapeBit != 0
			style.IsVerticalText = flags&verticalTextBit != 0

		case 71:
			flags, _ := core.AsInt(tag.Value)
			style.IsBackwards = flags&backwardsBit != 0
			style.IsUpsideDown = flags&upsideDownBit != 0

		default:
			fmt.Printf("Discarding tag for Style: %+v\n", tag.ToString())
		}
	}

	return style
}
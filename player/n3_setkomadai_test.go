package player

import (
	"image/color"
	"testing"

	"github.com/littletrainee/GunGiBoardGameGUI/font"
	"github.com/littletrainee/gunginotationgenerator/enum/level"
)

func TestSetKomaDai(t *testing.T) {
	P := Player{SelfColor: color.Black}
	// P.KomaDai = []pair.Pair[koma.Koma, int]{}
	P.SetKomaDai(level.ELEMENTARY, font.ConvertToFace())
	if len(P.KomaDai) == 11 {
		t.Log(true)
	} else {
		t.Error(false)
	}
}

package player

import (
	"image/color"
	"testing"

	"github.com/littletrainee/GunGiBoardGameGUI/font"
	"github.com/littletrainee/GunGiBoardGameGUI/koma"
	"github.com/littletrainee/gunginotationgenerator/enum/level"
	"github.com/littletrainee/pair"
)

func TestSetKomaDai(t *testing.T) {
	P := Player{SelfColor: color.Black}
	P.KomaTai = []pair.Pair[koma.Koma, int]{}
	P.SetKomaTai(level.ELEMENTARY, font.ConvertToFace())
	if len(P.KomaTai) == 11 {
		t.Log(true)
	} else {
		t.Error(false)
	}
}

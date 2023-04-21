package player

import (
	"image/color"
	"testing"

	"github.com/littletrainee/GunGiBoardGameGUI/koma"
	"github.com/littletrainee/gunginotationgenerator/enum/level"
	"github.com/littletrainee/pair"
)

func TestSetKomaDai(t *testing.T) {
	P := Player{Color: color.Black}
	P.KomaDai = []pair.Pair[koma.Koma, int]{}
	P.SetKomaTai(level.ELEMENTARY)
	if len(P.KomaDai) == 11 {
		t.Log(true)
	} else {
		t.Error(false)
	}
}

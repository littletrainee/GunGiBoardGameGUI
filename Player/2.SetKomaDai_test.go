package Player

import (
	"image/color"
	"testing"

	"github.com/littletrainee/GunGiBoardGameGUI/Const"
	"github.com/littletrainee/GunGiBoardGameGUI/Koma"
	"github.com/littletrainee/GunGiBoardGameGUI/Player/Pair"
)

func TestSetKomaDai(t *testing.T) {
	P := Player{Color: color.Black}
	P.KomaDai = []Pair.Pair[Koma.Koma, int]{}
	P.SetKomaDai(Const.Elementary)
	if len(P.KomaDai) == 11 {
		t.Log(true)
	} else {
		t.Error(false)
	}
}

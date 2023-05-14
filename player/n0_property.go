package player

import (
	"image"
	"image/color"

	"github.com/littletrainee/GunGiBoardGameGUI/enum/playerstate"
	"github.com/littletrainee/GunGiBoardGameGUI/koma"
	"github.com/littletrainee/pair"
)

type Player struct {
	KomaTai                 []pair.Pair[koma.Koma, int]
	KomaTaiIndex            int
	KomaTaiCoordinate       image.Point
	KomaOnKomaTaiCoordinate image.Point
	SelfColor               color.Gray16
	State                   playerstate.PlayerState
	IsOwn                   bool
	IsSuMi                  bool
	MaxRange                int
}

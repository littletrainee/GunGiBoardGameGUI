package player

import (
	"image"
	"image/color"

	"github.com/littletrainee/GunGiBoardGameGUI/koma"
	"github.com/littletrainee/pair"
)

// 玩家物件
type Player struct {
	KomaDai                 []pair.Pair[koma.Koma, int] // 駒台Item1為駒的資訊，Item2為駒的剩餘數量
	KomaDaiCoordinate       image.Point                 // 駒台的起始繪製座標
	KomaOnKomaDaiCoordinate image.Point                 // 駒在駒台上繪製的座標
	KomaDaiBackGroundColor  color.RGBA                  // 駒台的背景顏色
	SelfColor               color.Gray16                // 自家的顏色
	IsOwn                   bool                        // 是否為莊家
	IsSetSuI                bool                        // 是否已布置好帥
	IsSuMi                  bool                        // 是否宣告"布陣完畢"
	MaxRange                int                         // 可以"新"的最大範圍
}

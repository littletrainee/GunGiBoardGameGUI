// 遊戲中的玩家物件
package player

import (
	"image"
	"image/color"

	"github.com/littletrainee/GunGiBoardGameGUI/koma"
)

// 玩家物件
type Player struct {
	KomaDai                 [][]koma.Koma // 駒台Item1為駒的資訊，Item2為駒的剩餘數量
	KomaDaiCoordinate       image.Point   // 駒台的起始繪製座標
	KomaOnKomaDaiCoordinate image.Point   // 駒在駒台上繪製的座標
	KomaDaiBackground       color.RGBA    // 駒台的背景顏色
	SelfColor               color.Gray16  // 自家的顏色
	IsOwn                   bool          // 是否為莊家
	IsSetSuI                bool          // 是否已布置好帥
	IsSuMi                  bool          // 是否宣告"布陣完畢"
	MaxRange                int           // 可以"新"的最大範圍
	SelectBouShou           bool          // 謀是否被選取
	MutinyList              [][]int       // 叛變清單
}

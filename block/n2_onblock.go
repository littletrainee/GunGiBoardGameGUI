package block

import (
	"image"

	"github.com/littletrainee/GunGiBoardGameGUI/constant"
)

// OnBlock 檢查滑鼠是否有按在該區塊上。
//
// 若有在按鈕上則回傳true，否則false。
//
// 參數x與y為滑鼠的座標。
func (b *Block) OnBlock(x, y int) bool {
	return image.Point{x, y}.In(image.Rect(int(b.Coordinate.X),
		int(b.Coordinate.Y), int(b.Coordinate.X)+int(constant.BLOCK_SIZE),
		int(b.Coordinate.Y)+int(constant.BLOCK_SIZE)))
}

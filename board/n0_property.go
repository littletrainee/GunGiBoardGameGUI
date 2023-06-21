// 棋盤物件
package board

import (
	"image"

	"github.com/littletrainee/GunGiBoardGameGUI/block"
)

// 放置區塊的棋盤
type Board struct {
	Blocks map[image.Point]block.Block // 棋盤的區塊雜湊表，鍵為區塊的位置物件，值為區塊物件
}

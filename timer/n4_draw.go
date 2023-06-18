package timer

import (
	"image/color"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/littletrainee/GunGiBoardGameGUI/constant"
	_phase "github.com/littletrainee/GunGiBoardGameGUI/enum/phase"
	"github.com/littletrainee/gunginotationgenerator/enum/phase"
	"golang.org/x/image/font"
)

func (t *Timer) Draw(screen *ebiten.Image, font font.Face, phase phase.Phase) {
	if phase == _phase.BEGIN_COUNTDOWN_FOR_GAMING {
		// 重設畫布
		t.Img = ebiten.NewImage(constant.WINDOW_SIZE_WIDTH, constant.WINDOW_SIZE_HEIGHT)
		// 將文字繪製到畫布上
		text.Draw(t.Img, strconv.Itoa(t.RemainingTime), font,
			int(t.CurrentCoordinate.X)+100, int(t.CurrentCoordinate.Y)+75,
			color.Black)
	} else {
		// 計時器背景
		vector.DrawFilledRect(screen, float32(t.CurrentCoordinate.X), float32(t.CurrentCoordinate.Y),
			float32(constant.TIMER_SIZE), float32(constant.TIMER_SIZE),
			t.BackgroundColor, true)
		// 每次重新顯示文字必須要重製畫布，否則會有殘留影像
		t.Img = ebiten.NewImage(constant.TIMER_SIZE, constant.TIMER_SIZE)
		s := strconv.Itoa(t.RemainingTime / 60)
		s += ":"
		s += strconv.Itoa(t.RemainingTime % 60)
		text.Draw(t.Img, s, font, 0, 20, color.Black)

	}
	screen.DrawImage(t.Img, t.Op)
}

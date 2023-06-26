package mutiny

import (
	"github.com/littletrainee/GunGiBoardGameGUI/koma"
	"github.com/littletrainee/GunGiBoardGameGUI/player"
)

// MunityCheck 判斷是否有可以觸發叛變技能的核可，目標位置謀所控制的駒在自家駒台上有包含的話，則可以觸發叛變
func MunityCheck(targetKomaStack []koma.Koma, player *player.Player) {
	for i, v := range targetKomaStack {
		if v.Name == "謀" {
			break
		}
		for ki, V := range player.KomaDai {
			if len(V) > 0 && V[0].Name == v.Name {
				player.MutinyList = append(player.MutinyList, []int{ki, i})
			}
		}
	}
}

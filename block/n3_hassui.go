package block

// 確認目標位置是否有帥
func (b *Block) HasSuI() bool {
	if len(b.KomaStack) > 0 {
		for _, v := range b.KomaStack {
			if v.Name == "帥" {
				return true
			}
		}
	}
	return false
}

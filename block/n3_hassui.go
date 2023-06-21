package block

// HasSuI 檢查該區塊是否包含帥
//
// 若有在按鈕上則回傳true，否則false。
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

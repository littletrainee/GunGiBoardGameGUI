package capture

// Reset 重製顯示畫面與控制按鈕的屬性
func (c *Capture) Reset() {
	c.Show = false
	c.ControlBool = false
}

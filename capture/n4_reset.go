package capture

func (c *Capture) Reset() {
	c.Show = false
	c.CaptureBool = false
	c.CancelBool = false
	c.ControlBool = false
}

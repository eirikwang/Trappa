package drivers

import (
	"github.com/hybridgroup/gobot/platforms/i2c"
	"github.com/hybridgroup/gobot"
	"bytes"
	"encoding/binary"
)

type i2cDriver struct {
	gobot.Driver
}
func newI2cDriver(driver *gobot.Driver) i2cDriver{
	return i2cDriver{Driver: *driver}
}
func (h *i2cDriver) i2c() (i2c.I2cInterface) {
	return h.Adaptor().(i2c.I2cInterface)
}
func (h *i2cDriver) read16(command byte)([]byte){
	h.i2c().I2cWrite([]byte{command});
	return h.i2c().I2cRead(uint(2))
}
func (h *i2cDriver) readInt16(command byte)(in int16){
	binary.Read(bytes.NewReader(h.read16(command)), binary.BigEndian, &in)
	return
}

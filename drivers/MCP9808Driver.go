package drivers

import (
	"github.com/hybridgroup/gobot"
	"github.com/hybridgroup/gobot/platforms/i2c"
	"math/big"
	"time"
	"bytes"
)

var (
	MCP9808_COM_MANIFACTURER  []byte = []byte{0x06}
	MCP9808_COM_DEVICE       []byte  = []byte{0x07}
	MCP9808_COM_AMBIENT_TEMP []byte  = []byte{0x05}
)

type MCP9808Driver struct {
	i2cDriver
	Address byte
	Temperature *big.Rat
}

func createi2cDriver(i2cInterface i2c.I2cInterface, name string) *gobot.Driver {
	return gobot.NewDriver(
		name,
		"MCP9808Driver",
		i2cInterface.(gobot.AdaptorInterface),
		time.Duration(1*time.Second),
	)
}

func NewMCP9808Driver(i2cInterface i2c.I2cInterface, name string) (*MCP9808Driver) {
	return &MCP9808Driver{
		Address: 0x18,
		i2cDriver: newI2cDriver(createi2cDriver(i2cInterface, name)),
	}
}

func (h *MCP9808Driver) Start() bool {
	h.i2c().I2cStart(h.Address)
	if (!h.init()) {return false}

	gobot.Every(h.Interval(), func() {
			in := h.readInt16(MCP9808_COM_AMBIENT_TEMP)
			h.Temperature = big.NewRat(int64(in & 0x0fff), 16);
		})
	return true
}


func (h *MCP9808Driver) init() bool {
	mani := h.read16(MCP9808_COM_MANIFACTURER)
	dev := h.read16(MCP9808_COM_DEVICE)
	if (bytes.Equal(mani, []byte{0x00, 0x54}) && bytes.Equal(dev, []byte{0x04, 0x00})) {return true}
	return false
}
func (h *MCP9808Driver) Halt() bool { return true }

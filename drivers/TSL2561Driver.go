package drivers

import (
	"github.com/hybridgroup/gobot/platforms/i2c"
	"fmt"
	"github.com/hybridgroup/gobot"
	"encoding/binary"
)

var (
	TSL2561_DEVICE_ID []byte = []byte {0x0A};
)
type TSL2561Func interface {
	getLux() int64
}
type TSL2561Driver struct {
	i2cDriver
	Address byte
	Lux     int64

}
func (h TSL2561Driver) getLux() int64{
	return h.Lux
}

func NewTSL2561Driver(i2cInterface i2c.I2cInterface, name string) *TSL2561Driver {
	return &TSL2561Driver{
		//Address: 0x39 (uten jord, fungerer ikke) 0x48 med 3v
		Address: 0x29, //Addr pin med gnd
		i2cDriver: newI2cDriver(createi2cDriver(i2cInterface, name)),
	}
}

func (h *TSL2561Driver) Start() bool {
	h.i2c().I2cStart(h.Address)
	if (!h.init()) {return false}
	gobot.Every(h.Interval(), func() {
			fmt.Println("-----")
			fmt.Println(newCommand().word().address(data0low).cmd())
			fmt.Println(newCommand().word().address(data0low).command)
			fmt.Println(h.read16(newCommand().word().address(data0low).cmd()));
			fmt.Println(h.read16(newCommand().word().address(data1low).cmd()));
			fmt.Println("%%%%%")
		})

	return true
}


func (h *TSL2561Driver) init() bool {
	mani := h.read(TSL2561_DEVICE_ID, uint(2))
	fmt.Println(mani)
	if (mani[0] == byte(0x0A)) {return true}
	return false
}
func (h *TSL2561Driver) Halt() bool { return true }


type commandAddress byte
const(
	control commandAddress =  iota
	timing
	treshholdlowlow
	treshholdlowhigh
	treshholdhighhigh
	treshholdhighlow
	interrupt
	reserved1
	crc
	reserved2
	id
	reserved3
	data0low
	data0high
	data1low
	data1high
)

type command struct{
	command uint16
}
func newCommand() command {return command{0x80}} //Must be set
func (c command) clear() command { c.command |= 0x40; return c} //Clear register
func (c command) word() command { c.command |= 0x20; return c}  //Use weither word bit or byte bit. Not both
func (c command) byte() command { c.command |= 0x10; return c}
func (c command) address(adr commandAddress) command  { c.command |= uint16(adr); return c}
func (c command) cmd() (b[]byte ) {bt := make([]byte,2); binary.BigEndian.PutUint16(bt, c.command);return bt}



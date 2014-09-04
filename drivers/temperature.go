package drivers
import (
	"github.com/hybridgroup/gobot"
	"github.com/hybridgroup/gobot/platforms/i2c"
	"math/big"
)

const (
	IO_port string             = "/dev/i2c-2"
	IO_addr byte               = 0x18
    READ_SIZE uint      = 2
	COM_MANIFACTURER = 0x06
	COM_DEVICE       = 0x07
	COM_AMBIENT_TEMP = 0x05
)

type Sensor interface {
	Init() bool
}
type MCP9828Driver struct {
	gobot.Driver
	Temperature *big.Rat
}
func NewMCP9828Driver(a i2c.I2cInterface, name string)  (*MCP9828Driver){
	return &MCP9828Driver{
			Driver: *gobot.NewDriver(
				name,
				"MCP9828Driver",
				a.(gobot.AdaptorInterface),
			),
	}
}
func (h *MCP9828Driver) Start() bool {
	h.i2c().I2cStart(IO_addr)
	h.i2c().I2cWrite([]byte("A"))

	gobot.Every(h.Interval(), func() {
			i:= new(big.Int)
			i.SetBytes(h.i2c().I2cRead(READ_SIZE))
			i.And(i, big.NewInt(0x0fff))
			h.Temperature = big.NewRat(i.Int64(), 10);
		})
	return true
}
func (h *MCP9828Driver) i2c() (i2c.I2cInterface){
	return h.Adaptor().(i2c.I2cInterface)
}
func (h *MCP9828Driver) Init() bool {
	/**if(h.i2c().I2cRead(READ_SIZE) == COM_MANIFACTURER){

	}**/

	return true
}
func (h *MCP9828Driver) Halt() bool { return true }

/*
func (obj *MCP9828) Read(){
	gpio.NewDirectPinDriver(obj.Reader, "", "");

if(!b.i2cOpen( port, addr)) return false;
if(b.i2cReadBytes(port, COM_MANIFACTURER, 2).toString("hex") != "0054") return false;
if(b.i2cReadBytes(port, COM_DEVICE, 2) .toString("hex") != "0400") return false;
console.log("Temperature sensor running")
return true;
}
function checkTemperature(){
var t = b.i2cReadBytes(port, COM_AMBIENT_TEMP, 2).readInt16BE(0);
var temp = (t & 0x0fff) / 16;
console.log(temp);
console.log(t & 0x1000)
if (t & 0x1000 <0) temp -= 256;

console.log("temperature: " + temp);
}

if(setupTemperature()){
setInterval(checkTemperature, 1000);
}
*/

var b = require('bonescript');
var port =  '/dev/i2c-2';
var addr = 0x18;
var COM_MANIFACTURER = 0x06;
var COM_DEVICE = 0x07;
var COM_AMBIENT_TEMP = 0x05;

function setupTemperature(){
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

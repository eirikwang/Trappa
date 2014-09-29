var b = require('bonescript');
var port =  '/dev/i2c-2';
var addr = 0x18;
var ADDR = 0x80;
var ADDR_COM_REGISTER = 0x00;
var ADDR_COM_WORD = 0x20;
var ADDR_COM_TIMING = 0x01;
var ADDR_REGISTER_ID = 0x0A;

var COM_PWR_ON = 0x03
var COM_PWR_OFF = 0x00

var INT_13 = 0x00;
var INT_101 = 0x01;
var INT_402 = 0x02;

var COM_MANIFACTURER = 0x06;
var COM_DEVICE = 0x07;
var COM_AMBIENT_TEMP = 0x05;

var REGISTER_CHAN0_LOW        = 0x0C;
var REGISTER_CHAN0_HIGH       = 0x0D;
var REGISTER_CHAN1_LOW        = 0x0E;
var REGISTER_CHAN1_HIGH       = 0x0F;

var intTime = 0x00; //13.7ms
var gain = 0x00; //no gain
function enable(){
    console.log((ADDR | ADDR_COM_REGISTER).toString(16))
    b.i2cWriteBytes(port, 0x80, COM_PWR_ON)
}
function disable(){
    b.i2cWriteBytes(port, 0x80, COM_PWR_OFF)
}
function setIntegrationTime(newIntTime){
    enable();
    b.i2cWriteBytes(port, (ADDR| ADDR_COM_TIMING).toString(16), (newIntTime|gain).toString(16));
    intTime = newIntTime;
    disable();
}
function setGain(newGain){
    enable();
    b.i2cWriteBytes(port, (ADDR| ADDR_COM_TIMING).toString(16), (intTime|newGain).toString(16));
    gain = newGain;
    disable();
}
function getReading(){
    //var chan = (ADDR | ADDR_COM_WORD | REGISTER_CHAN0_LOW).toString(16) ;
    var chan = 0xac;
    var chan2=0xae;
    //var chan2 = (ADDR | ADDR_COM_WORD | REGISTER_CHAN1_LOW).toString(16) ;
    var broadband = b.i2cReadBytes(port,chan,2);
    console.log(chan)
    console.log(chan2)
    var ir = b.i2cReadBytes(port, chan2, 2);
    disable();
    console.log( {broadband: broadband, ir: ir});
}
function getData(){
    enable();
    var res;
    if(intTime == INT_13){
        setTimeout(getReading, 14);
    }
    return res;
}

function getLumox(){
    return getData();
}
function begin(){
    if(!b.i2cOpen( port, ADDR_REGISTER_ID)) return false;
    var rdy = b.i2cReadBytes(port, ADDR_REGISTER_ID, 1);
    console.log(rdy);
    if(!rdy) return false;
    setIntegrationTime(intTime);
    setGain(gain);
    disable();
    return true;
}

if(begin()){
    setInterval(getLumox, 1000);
}


package drivers

type integrationTime uint16
const(
    none integrationTime = 0x0214
	short integrationTime = 0x7517
	long integrationTime = 0x0fe7
)


const (
	LUX_SCALE uint16 = 14 // scale by 2^14

	RATIO_SCALE uint16 = 9 // scale ratio by 2^9
	//−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−
	// Integration time scaling factors
	//−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−
	CH_SCALE uint16      = 10     // scale channel values by 2^10
	CHSCALE_TINT0 uint16 = 0x7517 // 322/11 * 2^CH_SCALE
	CHSCALE_TINT1 uint16 = 0x0fe7 // 322/81 * 2^CH_SCALE
	//−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−
	// T, FN, and CL Package coefficients
	//−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−
	K1T uint16 = 0x0040 // 0.125 * 2^RATIO_SCALE
	B1T uint16 = 0x01f2 // 0.0304 * 2^LUX_SCALE
	M1T uint16 = 0x01be // 0.0272 * 2^LUX_SCALE
	K2T uint16 = 0x0080 // 0.250 * 2^RATIO_SCALE TSL2560, TSL256
	B2T uint16 = 0x0214 // 0.0325 * 2^LUX_SCALE
	M2T uint16 = 0x02d1 // 0.0440 * 2^LUX_SCALE
	K3T uint16 = 0x00c0 // 0.375 * 2^RATIO_SCALE
	B3T uint16 = 0x023f // 0.0351 * 2^LUX_SCALE
	M3T uint16 = 0x037b // 0.0544 * 2^LUX_SCALE
	K4T uint16 = 0x0100 // 0.50 * 2^RATIO_SCALE
	B4T uint16 = 0x0270 // 0.0381 * 2^LUX_SCALE
	M4T uint16 = 0x03fe // 0.0624 * 2^LUX_SCALE
	K5T uint16 = 0x0138 // 0.61 * 2^RATIO_SCALE
	B5T uint16 = 0x016f // 0.0224 * 2^LUX_SCALE
	M5T uint16 = 0x01fc // 0.0310 * 2^LUX_SCALE
	K6T uint16 = 0x019a // 0.80 * 2^RATIO_SCALE
	B6T uint16 = 0x00d2 // 0.0128 * 2^LUX_SCALE
	M6T uint16 = 0x00fb // 0.0153 * 2^LUX_SCALE
	K7T uint16 = 0x029a // 1.3 * 2^RATIO_SCALE
	B7T uint16 = 0x0018 // 0.00146 * 2^LUX_SCALE
	M7T uint16 = 0x0012 // 0.00112 * 2^LUX_SCALE
	K8T uint16 = 0x029a // 1.3 * 2^RATIO_SCALE
	B8T uint16 = 0x0000 // 0.000 * 2^LUX_SCALE
	M8T uint16 = 0x0000 // 0.000 * 2^LUX_SCALE
	//−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−
	// CS package coefficients
	//−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−
	0.130
	0.260
	0.390
	0.520
	K1C uint16 = 0x0043 // 0.130 * 2^RATIO_SCALE
	B1C uint16 = 0x0204 // 0.0315 * 2^LUX_SCALE
	M1C uint16 = 0x01ad // 0.0262 * 2^LUX_SCALE
	K2C uint16 = 0x0085 // 0.260 * 2^RATIO_SCALE
	B2C uint16 = 0x0228 // 0.0337 * 2^LUX_SCALE
	M2C uint16 = 0x02c1 // 0.0430 * 2^LUX_SCALE
	K3C uint16 = 0x00c8 // 0.390 * 2^RATIO_SCALE
	B3C uint16 = 0x0253 // 0.0363 * 2^LUX_SCALE
	M3C uint16 = 0x0363 // 0.0529 * 2^LUX_SCALE    TSL2560, TSL2561
	B4C uint16 = 0x0282 // 0.0392 * 2^LUX_SCALE
	M4C uint16 = 0x03df // 0.0605 * 2^LUX_SCALE
	K5C uint16 = 0x014d // 0.65 * 2^RATIO_SCALE
	B5C uint16 = 0x0177 // 0.0229 * 2^LUX_SCALE
	M5C uint16 = 0x01dd // 0.0291 * 2^LUX_SCALE
	K6C uint16 = 0x019a // 0.80 * 2^RATIO_SCALE
	B6C uint16 = 0x0101 // 0.0157 * 2^LUX_SCALE
	M6C uint16 = 0x0127 // 0.0180 * 2^LUX_SCALE
	K7C uint16 = 0x029a // 1.3 * 2^RATIO_SCALE
	B7C uint16 = 0x0037 // 0.00338 * 2^LUX_SCALE
	M7C uint16 = 0x002b // 0.00260 * 2^LUX_SCALE
	K8C uint16 = 0x029a // 1.3 * 2^RATIO_SCALE
	B8C uint16 = 0x0000 // 0.000 * 2^LUX_SCALE
	M8C uint16 = 0x0000 // 0.000 * 2^LUX_SCALE
)

// lux equation approximation without floating point calculations
//////////////////////////////////////////////////////////////////////////////
// Routine: unsigned int CalculateLux(unsigned int ch0, unsigned int ch0, int iType)
//
// Description: Calculate the approximate illuminance (lux) given the raw
// channel values of the TSL2560. The equation if implemented
// as a piece−wise linear approximation.
//
// Arguments: unsigned int iGain − gain, where 0:1X, 1:16X
// unsigned int tInt − integration time, where 0:13.7mS, 1:100mS, 2:402mS,
// 3:Manual
// unsigned int ch0 − raw channel value from channel 0 of TSL2560
// unsigned int ch1 − raw channel value from channel 1 of TSL2560
// unsigned int iType − package type (T or CS)
//
// Return: unsigned int − the approximate illuminance (lux)
//
//////////////////////////////////////////////////////////////////////////////
func CalculateLux(gain uint16, integrationTime integrationTime, ch0 uint16, ch1 uint16, iPackage uint16) uint32 {
//−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−
// first, scale the channel values depending on the gain and integration time
// 16X, 402mS is nominal.
// scale if integration time is NOT 402 msec
var chScale, channel1, channel0 uint16;
chScale = uint16(integrationTime)

// scale if gain is NOT 16X
if (!gain) {chScale <<= 4}; // scale 1X to 16X
// scale the channel values
channel0 = (ch0 * chScale) >> CH_SCALE;
channel1 = (ch1 * chScale) >> CH_SCALE;
//−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−−
// find the ratio of the channel values (Channel1/Channel0)
// protect against divide by zero
var ratio1 uint16 = 0;
if (channel0 != 0) {ratio1 = (channel1 << (RATIO_SCALE+1)) / channel0}
// round the ratio value
var ratio uint16 = (ratio1 + 1) >> 1;
// is ratio <= eachBreak ?
var b, m uint;
switch (iPackage)
{
case 0: // T, FN and CL package
if ((ratio >= 0) && (ratio <= K1T))
{b=B1T; m=M1T;}
else if (ratio <= K2T)
{b=B2T; m=M2T;}
else if (ratio <= K3T)
{b=B3T; m=M3T;}
else if (ratio <= K4T)
{b=B4T; m=M4T;}
else if (ratio <= K5T)
{b=B5T; m=M5T;}
else if (ratio <= K6T)
{b=B6T; m=M6T;}
else if (ratio <= K7T)
{b=B7T; m=M7T;}
else if (ratio > K8T)
{b=B8T; m=M8T;}
break;
case 1:// CS package
if ((ratio >= 0) && (ratio <= K1C))
{b=B1C; m=M1C;}
else if (ratio <= K2C)
{b=B2C; m=M2C;}
else if (ratio <= K3C)
{b=B3C; m=M3C;}
else if (ratio <= K4C)
{b=B4C; m=M4C;}
else if (ratio <= K5C)
{b=B5C; m=M5C;}
else if (ratio <= K6C)
{b=B6C; m=M6C;}
else if (ratio <= K7C)
{b=B7C; m=M7C;}    TSL2560, TSL2561
LIGHT-TO-DIGITAL CONVERTER
TAOS059N − MARCH 2009
28


Copyright  2009, TAOS Inc. The LUMENOLOGY  Company
www.taosinc.com
else if (ratio > K8C)
{b=B8C; m=M8C;}
break;
}
unsigned long temp;
temp = ((channel0 * b) − (channel1 * m));
// do not allow negative lux value
if (temp < 0) temp = 0;
// round lsb (2^(LUX_SCALE−1))
temp += (1 << (LUX_SCALE−1));
// strip off fractional portion
unsigned long lux = temp >> LUX_SCALE;
return(lux);
}    

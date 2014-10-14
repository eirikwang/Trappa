Trappa
======


Fungerende temperatursensor og bevegelsesensor. 

Temperatursensor må være koblet til t3 - P9-19(i2c2_scl), t4 - P9-20(i2c2_sda), t1 - p9-01(GND) t2 -  p9-03 (3.3V)

Bevegelsesensor må være koblet til rød - p9-02, sort  - p9-01, gul - p9-12)

Api: 
Api kan ses på: 
http://192.168.7.2:4040/robots/trappa

Temperatur leses ved: 
http://192.168.7.2:4040/robots/trappa/devices/temp/commands/Read

Sonar leses ved: 
http://192.168.7.2:4040/robots/trappa/devices/sonar/commands/Read

Bevegelse leses ved: 
http://192.168.7.2:4040/robots/trappa/devices/motion/commands/DigitalRead

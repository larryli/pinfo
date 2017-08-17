package main

// ====================== 26 pin =====================
// OLED VCC ==>   3.3V+  [1] (2)  5V+   ==> TTP223 VCC
// OLED SDA ==> IIC SDA  (3) (4)  5V+
// OLED SCL ==> IIC SCL  (5) (6)  GND
//                GPIO6  (7) (8)  UART TX
// OLED GND ==>     GND  (9) (10) UART RX
//                GPIO1 (11) (12) GPIO7 ==> TTP223 SIG
//                GPIO0 (13) (14) GND   ==> TTP223 GND
//                GPIO3 (15) (16) GPIO19
//                3.3V+ (17) (18) GPIO18
//               GPIO15 (19) (20) GND
//               GPIO16 (21) (22) GPIO2
//               GPIO14 (23) (24) GPIO13
//                  GND (25) (26) GPIO10

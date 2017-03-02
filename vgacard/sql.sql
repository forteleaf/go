CREATE TABLE manufacturer {

}

CREATE TABLE vga {
vid int(8),
  manufacturer varchar(100),
  model varchar(100), -- gtx 1060
  core varchar(10), -- Graphics Core
  diesize varchar(20),  -- Die Size
  transistors varchar(20),
  cuda int(10), -- cuda core 1152 CUDA Cores
  base int(10), -- BAse clock MHz
  boost int(10), -- Boost Clock = 133 MHz
  vram varchar(20),
  bus varchar(20),
  power varchar(20), -- power connect = singgle 6-pin power
  tdp varchar(5), -- TDP = 120W
  display varchar(100), -- display outputs port
  launch date
}
Graphics Core	GP106	GP106	GP104	GP104	GP102
Process Node	16nm FinFET	16nm FinFET	16nm FinFET	16nm FinFET	16nm FinFET
Die Size	200mm2	200mm2	314mm2	314mm2	471mm2
Transistors	4.4 Billion	4.4 Billion	7.2 Billion	7.2 Billion	12 Billion
CUDA Cores	1152 CUDA Cores	1280 CUDA Cores	1920 CUDA Cores	2560 CUDA Cores	3584 CUDA Cores
Base Clock	1518 MHz	1506 MHz	1506 MHz	1607 MHz	1417 MHz
Boost Clock	1733 MHz	1708 MHz	1683 MHz	1733 MHz	1530 MHz
FP32 Compute	4.0 TFLOPs	4.4 TFLOPs	6.5 TFLOPs	9.0 TFLOPs	11 TFLOPs
VRAM	3 GB GDDR5	6 GB GDDR5	8 GB GDDR5	8 GB GDDR5X	12 GB GDDR5X
Bus Interface	192-bit bus	192-bit bus	256-bit bus	256-bit bus	384-bit bus
Power Connector	Single 6-Pin Power	Single 6-Pin Power	Single 8-Pin Power	Single 8-Pin Power	8+6 Pin Power
TDP	120W	120W	150W	180W	250W
Display Outputs	3x Display Port 1.4
1x HDMI 2.0b
1x DVI	3x Display Port 1.4
1x HDMI 2.0b
1x DVI	3x Display Port 1.4
1x HDMI 2.0b
1x DVI	3x Display Port 1.4
1x HDMI 2.0b
1x DVI	3x Display Port 1.4
1x HDMI 2.0b
1x DVI
Launch Date

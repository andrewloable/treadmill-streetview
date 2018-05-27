# Treadmill StreeView PC component

This is the PC component of the project.
It basically receives the signal from the connected arduino device how many loops of the treadmill belt were detected using the treadmills tachometer.
This data is then used by the software to move the street view display forward depending on the configured number of loops detected.

## Prerequisites

You need to install the Google GO to be able to compile the code.
It is available here: https://golang.org/dl/

## Getting Started

### PC Connection to Arduino Device
1. Connect the arduino device to your pc via usb.
2. Check the designated COM port of your arduino device and modify the config.json file with the correct port. (The default port in the config.json file is COM4).

### Compiling and Run the code
1. Navigate to the pc folder.
2. Execute ```go build```
3. Run the generated executable file.

## Modifications

Open config.json and enter the correct COM port when the arduino device is connected.
The baud rate is set to 9600 by default. To change it, you may have to change the code in the arduino device.

# Treadmill StreeView Arduino component

This is the arduino component of the project.
The arduino device is connected to the treadmill's tachometer signal to detect its pulses.
The pulses are then converted into loops, where 1 loop is the complete cycle of the treadmill belt.

## Prerequisites

You need to install the Arduino IDE to be able to upload the code.
It is available here: https://www.arduino.cc/en/Main/Software

## Getting Started

### Treadmill Connection to Arduino Device
1. Identify the signal and ground wires of the tachometer sensor in your treadmill. Also, identify the voltage of high and low pulse of the signal using an oscilloscope.
2. Connect the signal wire to Analog Pin 3 and the ground wire to the Ground pin of the arduino device.

(Note: Your treadmill connection may vary, the treadmill available for me to test outputs a 1 volt high for each pulses of the tachometer.)

### Uploading Code to Device
1. Open Arduino IDE in your pc.
2. Open the INO file
3. Connect your arduino device to your pc's usb.
4. In Arduino IDE, select your device's model and connection port.
5. Click the upload button.

## Modifications

The code contains a segment where you can edit the values of the variables depending on your existing treadmill.

```
const int ANALOG_HIGH = 200
```
This variable sets the value of the analog input when the tachometer pulse is in a high state. 

```
const int PULSES_ONE_LOOP = 10000;
```
PULSES_ONE_LOOP is the number of pulses detected that is equivalent to one loop of the treadmill belt.

# Treadmill Street View
Street View walk around using your treadmill

## Overview
```
Treadmill -> Arduino -> PC -> App with maps installed
```

## Arduino Component
Please navigate to the arduino sub folder to view the details on how to upload and use the code to your arduino device.

## PC Component
Please navigate to the pc sub folder to fiew the details on how to compile and run the app.

## Disclaimer
This software is specifically designed to a specific treadmill where the tachometer signal generates pulses with a high signal at 1 volt.

## Using the app
Once the arduino and pc components are compiled and up and running, do the following:
1. Run the executable file in the pc component and it will automatically open a browser that has an embedded google map segment. 
2. Open the street view of a place you want to view then orient it to the direction of the road to follow.
3. Press the UP key in your keyboard to check if the street view move forward. If not, select another street view then do this step again.
4. If the street view moves forward, you are good to go. Just select the number of loops before the street view moves forward.
5. Click the START button.
6. Turn on your treadmill and walk. The street view should move forward once it detects the number of loops from the tachometer.

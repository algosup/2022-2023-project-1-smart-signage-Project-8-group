<h1 style="font-size: 2.5rem;">Technical specifications - Group 8</h1>

# Project Scope

The purpose of this project to create a product to function as a controller for outdoor LED signes.
The final product must be capable of:

- Turn the LEDs on and off.
- Dim the LEDs using PWM.
- Read the ambient light levels using a photoresistor and change the LEDs luminosity accordingly.
- Verify if the LEDs are powered.
- Verify if LEDs are functional.
- Count the amount of time the LEDs have been up.
- Connect to the LoRaWan network.
- Send status reports trough the LoRaWan network (such as estimated remaining LED liftime, error messages, etc.)
- Change device settings trough the LoRaWan network.
- Linking multiple devices together and shotting them all down in case of an error. 

The project must be realised using the TinyGo programming language. This is a limitation imposed by Algosup and we recommend C++ for the actual product as TinyGo has several limitations.
On the same note we also recommend a slightly different hardware, this will be detaild later on.

# Hardware

## Arduino Uno

Most everything in the project will be controlled using an Arduino Uno for the following reasons:

- It is easily accessible
- It is compatible with every other module used for the project (including the Lora-E5)
- It is compatible with TinyGo

## Lora-E5

We will control the final product trought the LoRaWan network.
We will be using a Lora-E5 for this purpose.

This device can be controlled using [AT Commands](https://files.seeedstudio.com/products/317990687/res/LoRa-E5%20AT%20Command%20Specification_V1.0%20.pdf).

## XY-MOS

Switch used to turn the LEDs on/off.
[Here](https://www.youtube.com/watch?v=tCJ2Q-CT6Q8) is an example of how to use it.

## ZMCT103C Current sensor

Sensor used to verify if the LEDs are powered.
[Here](https://electropeak.com/learn/interfacing-zmct103c-5a-ac-current-transformer-module-with-arduino/) is an example of how to get its output in miliAmpers.

## ACS712 current sensor

Sensore used to verify LED status by comaring the input current to the actual current present in the circuit.
[Here](https://www.engineersgarage.com/acs712-current-sensor-with-arduino/) and [here](https://www.electronicshub.org/interfacing-acs712-current-sensor-with-arduino/). are examples on how to use it.

ACS712 outputs are in mV (type float). The expected voltage depends on the input voltage and the power usage of the LEDs (hence why we should be able to verfy the leds using it).

The formula to calculate the output is ``Ln = I/Li`` where

```
Ln = Number of LEDs
I = Current (Low intensity)
Li = Nominal current of LEDs
```

## Photoresistor

Component used to measure ambient light.
To use its output, simply wire it up as shown [here](https://create.arduino.cc/projecthub/MisterBotBreak/how-to-use-a-photoresistor-46c5eb)

## Materials used for testing purposes

- 12V LED strip
- GPV-18-12 AC/DC converter

# Recommended hardware

## Arduino MKR WAN 1310

This Arduino has an inbuilt system to connect to the LoRaWan network.
This means it could replace the Arduino Uno as well as the Lora-E5.
It could reduce the electricity usage of the product as well as severly simplify it.
Peer-to-Peer communication is also simplified with this micro controller.

Unfortunatelly, we had found out about thes contoller too late to actually implement it.

## DS3231 Real-Time Clock Module

The Arduinos inbuilt clock is both unprecise and temperature-sesitive.
Therefore we recommend the use of an RTC module with an integrated temperature-compensated crystal oscillator.

## DHT-22 (also named as AM2302) Humidity sensor

Optionally, it could be useful to be able to tell humidity levels because in case of a strong fog the photosensor still might get enought light to lower the LEDs intensity even if in that situation, we might not want to do so.

# Electronical configuration

This is how our model must be set up.

![Schematic](./Images/Schematic.png)

# Naming conventions

We'll be following the naming conventions described [here](https://www.golangprograms.com/naming-conventions-for-golang-functions.html).

# Software architecture

The main loop will function on 15 minute intervals.
Each time the available data will be sent to the server. This will help avoid data loss during uplink communication.

Functions (inputs/outputs):
![Sofware_Architecture](./Images/Software_Architecture.png)

## readLoraDownLink()

Read data recieved from the server if available.

## readSensorData()

Gather all sensor data.

## changeLocalData()

Set data to match freshly gathered information (both from downlink and sensors).

## setLeds()

Turns leds on and off depending on available data.

## sendLoraUplink()

Send report to the server trough the LoRaWan network

## isConnectedToPower()

Read the input from the ZMCT103C module to learn if the LEDs are powered by an outlet or not.
ZMCT103C outputs are in terms of mA (type float). In our case the output should be around 0.25A under power.

## getAmbientLightLevel

Read the ambient light level using a photoresistor and convert it into a percentage.

## verifLeds()

Approximates LED status using output from a ACS712 low voltage sensor.

For details on how to use the module look at [this example](https://www.electronicshub.org/interfacing-acs712-current-sensor-with-arduino/).

ACS712 outputs are in mV (type float). The expected voltage depends on the input voltage and the power usage of the LEDs (hence why we should be able to verfy the leds using it).

# Testing plan

Dummy data must be created to replace downlink and sensory data.
The exact dummy data and its expected outputs are left to the discretion of the software engineer.

Tests must NOT be uploaded into the Arduino.

# Networking

The datastructure of the information packets sent and recieved by the Lora will need to be engineered to minimise the data packet size and therefore the data loss in transmission.
The data packets structure is as follows:

![Data structure](./Images/Data_Structure.png)

## Uplink (Lora-To-Server)

- LED brightness: Byte that contains a percentage value between 1-100.
- LED usage: Byte that contains the amount of time the LEDs have been up during the last 2 hours, each bit represing a 15 minute interval (Ex.: 11001111 = 30min up > 30min down > 60min up).
- Is powered: LEDs are powered.
- ZMCT Error: The ZMCT103C Current sensor does not give useful output.
- XY-MOS Error: The LEDs do not switch on even if they are powered.
- ACS Error: The ACS712 current sensor does not give useful output.
- PhotoResistor Error: The photoresistor does not give useful output.
- Request backup data: The arduino has been switched off for some reason and requires instructions on what it is supposed to do.

## Downlink (Server-To-Lora)

- Turn on/off: Turn LEDs on/off (Turn device on/off later if a more precise clock is available)
- Change min/max light level: Set the minimum and maximum operational light levels for the leds in increments of 10%. First 4 bits represent the minimum light level and second 4 bits represent the maximum light level (values between 0-10).
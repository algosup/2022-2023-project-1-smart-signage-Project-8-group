<h1 style="font-size: 2.5rem;">Technical specifications - Group 8</h1>

# Project Scope

This project was proposed by SignAll, a French company that has been manufacturing large luminous signage since 1962. They supply a large number of customers such as McDonald’s, Burger King, La Poste, Orange, AXA, Crédit Agricole, Total, etc. to name a few.

Their existing products are not connected therefore users must be on-site to know if the signage is on, functional, or out of order. Also, users cannot switch the signage on/off remotely (even when the law requires them to switch it off at a given time, when the shop closes for instance).

When the owner of the brand on the signage is not the same as the owner of the place (think of a Burger Kind’s restaurant for instance), the maintenance team from the brand does not know what is going on at the place where the signage is installed and must go on-site on a regular basis just to check if everything is working. This results in additional costs and damage to the brand when the signage if out of order for too long.

Lately, environmental concerns and cost of energy has increased the pressure on the manufacturer to produce more efficient solutions such as dimming the signage when it is getting dark or switch it off completely at a given time or when there is a shortage of electricity.

In other words, this project need to manage a set of L.E.Ds in order to control these remotly. That control shall include turning those said L.E.Ds ON or OFF and being able to dim the light therefore reducing the energy comsumption. The device should also be able to work on is own, dim or turn OFF L.E.Ds at a given time of the day or depending of the ambiant light for exemple.

# Hardware

## Arduino Uno

Most everything in the project will be controlled using an Arduino Uno for the following reasons:

- It is easily accessible
- It is compatible with every other module used for the project (Lora module is to be determined)
- It is compatible with TinyGo
- It contains an inbuilt crystal oscillator (which is theorically precise enought to be used as a clock)

## Lora-E5

We will attempt to control the final product trought the LoRaWan network.
If the LoRa-E5 is the best solution is yet to be determined.

## XY-MOS

Switch used to turn the LEDs on/off.

## ZMCT103C

Sensor used to verify if the LEDs are powered.

## Photoresistor

Component used to measure ambient light.

## Materials used for testing purposes

- 12V LED strip
- GPV-18-12 AC/DC converter

# Roadmap

- [x] Control LEDs with XY-MOS
- [x] Dim lights by changing on/off frequency of the LEDs

- [x] Read data from ZMCT103C
- [x] Read data from Photoresistor

- [ ] Etablish connection between the Arduino and the LoRa-E5 for AT commands
- [ ] Etablish connection with TheThingsNetwork
- [ ] Set up protocol(s) for long range communication

# Electronical configuration

This is how model of the final product must be set up.

![Schematic](./Images/Schematic.png)

# Naming conventions

We'll be following the naming conventions described [here](https://www.golangprograms.com/naming-conventions-for-golang-functions.html).

- Every variable and function must be writen in camelcase and must start with a lower case letter.


# Software architecture

## Global variables

- 

## Functions

<details>
    <summary>How to read</summary>
    -Description-

    | Title       |
    | ----------- |
    | Inputs      |
    | Outputs     |
</details>

### isConnectedToPower()

Read the input from the ZMCT103C module to learn if the LEDs are powered by an outlet or not.
ZMCT103C outputs are in terms of mA (type float). In our case the output should be around 0.25A under power.

| isConnectedToPower() |
| ----------- |
| - |
| - bool isConnected |

### getAmbientLightLevel

Read the input from the ACS712 sensor to read the ambient light level using a photoresistor.
For details on how to use the module look at [this example](https://www.electronicshub.org/interfacing-acs712-current-sensor-with-arduino/).

ACS712 outputs are in mV (type float). The expected voltage depends on the input voltage and the ambient light level. The output is to be normalised into a float between 0-1.

| getAmbientLightLevel() |
| ----------- |
| - |
| - float ambientLightLevel |

### switchLeds()

Turns leds on if imut is True and off if it is False.

| switchLeds() |
| ----------- |
| - bool turnOn |
| - |




# Networking

No idea about this one yet

<style>
    h3{
        font-size: 1.25rem;
        font-weight: larger;
        text-decoration: underline;
    }
</style>



---
###### Notes


XY-MOS:
- ~D
- GND /

ZMC: 
- 5V /
- A
- GND /

PR: /
- A /
- PW (5V?) /
- GND /

LORA:
- 3V3 /
- D RX
- ~D TX
- GND /

ACS712:
- 5V
- A
- GND

DS3: 
- 5V
- A
- A
- GND

Total:
- 5V    -   3(4)
- 3V3   -   1
- A     -   5
- D     -   1
- ~D    -   2
- GND   -   6
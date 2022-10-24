# **<ins>smart signage</ins>**

## **<ins>test report</ins>**

### test n°1

- name: Primary connection
- description: test the connection between the LoRa and the computer
- want: the LoRa and the pc connect
- get: the dot on the LoRa lighted up
- result: positive
- comment: went to a gateway 40 km from Algosup to try connecting the LoRa using TTN

### test n°2

- name: secondary connection
- description: test the connection between the LoRa and the LEDS
- want: the LoRa and the LED connect
- get: we can turn on and off the LEDs
- result: positive
- comment: we use a XY-MOS in order to turn on and off the LEDs

### test n°3

- name: additional connection
- description: test an additionnal connection between the LoRa and a arduino uno
- want: the LoRa and the arduino uno connect
- get: the LEDs on the arduino lights up
- result: positive
- comment: the arduino uno will be controlling most of the modules used during the project

### test n°4

- name: sensor n°1
- description: check if the light sensor is working
- want: light sensor is working
- get: the value is changing wether the sensor is hidden or not
- result: positive
- comment: this sensor will be used to tell if we need to dim the LEDs or not

### test n°5

- name: sensor n°2
- description: check if the high current sensor is working
- want: hight current is working
- get:
- result:
- comment:

### test n°6

- name: sensor n°3
- description: check if the low current sensor is working
- want: low current sensor is working
- get:
- result:
- comment:

### test n°7

- name: output n°1
- description: check if the report output works
- want: we have a report
- get:
- result:
- comment:

### test n°8

- name: input n°1
- read data from server if available
- want: the data from the server
- get:
- result:
- comment:

### test n°9

- name: input n°2
- description: Gather sensor data
- want: the data from the sensors
- get:
- comment

### test n°10

- name: input n°3
- descrition: Set data to match freshly gathered information (from downlink and sensors)
- want: data change acording to new information
- get:
- result:
- comment:

### test n°11

- name: input n°4
- descrition: Turns leds on and off depending on available data
- want: the lights either turn on or off
- get:
- result:
- comment:

### test n°12

- name: input n°5
- description: Send report to the server trough the LoRaWan network
- want: the server receive a report trough the LoRaWan network
- get:
- result:
- comment:  

### test n°13

- name: input n°6
- description: Read the input from the ZMCT103C module to learn if the LEDs are powered or not
- want: learn if the LEDs are powered or not
- get:
- comment:

### test n°14

- name: input n°7
- description: Read the ambient light level using a photoresistor via a percentage
- want: get a percentage indicating ambient light level
- get:
- comment:

### test n°15

- name: input n°8
- description: Approximates LED status using output from a ACS712 low voltage sensor
- want: an approximation of the LEDs status
- get:
- comment:

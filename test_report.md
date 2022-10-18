# smart signage

## test report

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
- get: the LEDs lighted up
- result: positive
- comment: the LEDs are connected to the LoRa via an arduino uno and an ST-link

### test n°3

- name: additional connection
- description: test an additionnal connection between the LoRa and an arduino uno
- want: the LoRa and the arduino uno connect
- get: the LED on the arduino uno lights up
- result: positive
- comment:

### test n°4

- name: sensor n°1
- description: check if the light sensor is working
- want: light sensor is working
- get:
- result:
- comment:

### test n°5

- name: sensor n°2
- description: check if the heat sensor is working
- want: heat sensor is working
- get:
- result:
- comment:

### test n°6

- name: sensor n°3
- description: check if the high current sensor is working
- want: hight current is working
- get:
- result:
- comment:

### test n°7

- name: sensor n°4
- description: check if the low current sensor is working
- want: low current sensor is working
- get:
- result:
- comment:

### test n°8

- name: output n°1
- description: check if the error output works
- want: we have an error
- get:
- result:
- comment:

### test n°9

- name: output n°2
- description: check if the report output works
- want: we have a report
- get:
- result:
- comment:

### test n°10

- name: output n°3
- description: check if the warning output works
- want: we have a warning
- get:
- result:
- comment:

### test n°11

- name: input n°1
- description: check if we can set a timetable
- want: we can set a timetable
- get:
- result:
- comment:

### test n°12

- name: input n°2
- description: check if we can get the history/stats
- want: we can access the history/stats
- get:
- desciption:

### test n°13

- name: input n°3
- descrition: check if we can ask for a report
- want: we have a report
- get:
- result:
- comment:

### test n°14

- name: input n°4
- descrition: check if we can varie the light
- want: the light varie
- get:
- result:
- comment:

### test n°15

- name: input n°5
- description: check if we can set report times/fq
- want: we can set report times/fq
- get:
- result:
- comment:  

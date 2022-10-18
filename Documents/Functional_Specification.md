# <div align="center">Functional Specification</div>

## Stakeholders

 - ALGOSUP
 - SignAll
 - SignAll's Clients (ALD Automotive, DACIA, AXA for instance)
 - People Seeing the Signages
 - Technicians

## Project Scope

This project was proposed by SignAll, a French company that has been manufacturing large luminous signage since 1962. They supply a large number of customers such as McDonald’s, Burger King, La Poste, Orange, AXA, Crédit Agricole, Total, etc. to name a few.

Their existing products are not connected therefore users must be on-site to know if the signage is on, functional, or out of order. Also, users cannot switch the signage on/off remotely (even when the law requires them to switch it off at a given time, when the shop closes for instance).

When the owner of the brand on the signage is not the same as the owner of the place (think of a Burger Kind’s restaurant for instance), the maintenance team from the brand does not know what is going on at the place where the signage is installed and must go on-site on a regular basis just to check if everything is working. This results in additional costs and damage to the brand when the signage if out of order for too long.

Lately, environmental concerns and cost of energy has increased the pressure on the manufacturer to produce more efficient solutions such as dimming the signage when it is getting dark or switch it off completely at a given time or when there is a shortage of electricity.

In other words, this project need to manage a set of L.E.Ds in order to control these remotly. That control shall include turning those said L.E.Ds ON or OFF and being able to dim the light therefore reducing the energy comsumption, all of that remotly. The device should also be able to work on is own, dim or turn OFF L.E.Ds at a given time of the day or depending of the ambiant light for exemple.

## Risks and Assumptions

>Work In Progress (you're smart, you'll figure it out)

>> We shall directly talk about it with the client this afternoon

>>>The first risk is related to the budget, we need to create a device that shall be as cheap as possible because a few of them might be required for larger signages and a lot would be needed in order to convert all or a good part of existing signages.

>> Needs more detailed research

>>>The Second risk regards laws and regulation, since we will use a network in order to have the device remotly controlled, we are limited by the law on the lenght of the messages we send and on the frequency at which we can send them. <a href="https://www.thethingsnetwork.org/docs/lorawan/duty-cycle/">(more detail here)</a>

>> Need to talk with the group

>>>The Third risk is the security of the device. Knowing it will receive an input (for instance to switch on the L.E.Ds), this input need to be secured in some way (encoding, etc) to avoid hacks or clandestine inputs.

>> Need to talk with the group

>>>The Fourth risk is regarding the communication, we will need to ensure that every device has a unique id in order to identify them during communication as there shall be quite a few of them on the network.

## personas

![Dimitri](./Images/Persona%201.png)

![Karen](./Images/Persona%202.png)

![Francois](./Images/Persona%203.png)

## Use cases

> Work In Progress (needs to be real situations not just features)

- During the day Dimitri wants to know if some of the signages of the brand are ON, OFF or DIMMED without calling someone on site to go out in front of the building in order to directly see it. he could ask for a report of the devices, which will contain informations on the current state of the signage. (WORKING INFORMATION MONITORING)

- During his work, Dimitri is ask to manage and supervise the maintenance of the signages of the company. As of today, he has no choice but to send a technician on-site to ensure the signage is working properly. Our device could help him by sending daily reports on the state of each signage and eventualy alert on some unwanted behaviour like high temperature, a sudden drop in brightness (which could indicate a L.E.D is broken) or simply a power shortage therefore allowing him to concentrate his efforts and crews on the signages that needs to most urgent maintenance. (STATUS INFORMATION MONITORING)

- François is ask the verify the good functionning of some signages but he would like to know beforehand if these signages are working. Our device could help him with report, allowing him to concentrate his efforts on signages that needs to most atention. (STATUS INFORMATION MONITORING)

- At the moment, when Karen needs to turn ON or OFF the shop's signage, she has to go the power supply and manualy do it. Our device will allow her to remotly control the signage. (REMOTE CONTROL)

- Dimitri recently learned that some of the brand's signages weren't working for the past few days, that situation may cause the loss of quite a few customers. Our device will send report if a signage isn't working as intended therefore reducing the delay to send a technician on-site. (ISSUE REPORTING)

-  Karen recently thought that the signage of the company may be too bright during some period of the day and would to find a way to diminish that light. Our device could allow her to directly control the brightness of the signage, she would be able choose between 25, 50, 75 percent of light. (DIMMING CONTROL)

-  Dimitri had recently been charged with the task of reducing the power usage of the brand's building. One of his ideas would be to turn OFF the signages as soon as the stores are closed. Our device could help by automatical controling the signages depending of the time of the day, those cut OFF and turn ON times would be configurable and allow great flexibility. (AUTOMATIC BEHAVIOUR)

- In a day of work Karen has to directly control the signage at least twice (turn it ON when the shop/company opens and OFF when it closes), she also like to verify that the signage is working properly which means going outside and taking a direct look at it. She is trying to figure out a way of managing the said signage without it being too time comsuming and perhaps find a way for it to be automated. Our device should help her by making her signage completely autonomous and sending report only when there is an issue. (REMOTE CONTROL / STATUS INFO MONITORING / WORKING INFO MONITORING)

- Dimitri, still managing the brand's signages is having a hard time recording and saving the status of each signages, for his defence there are a lot of those signages. Our device could help him by giving direct report from each signages of the brand, allowing him to save time and ressources. (STATUS INFO MONITORING)

> Needs final checks

## Requirements Specs

> Work In progress (need details lot of them, you need to choose for others)

For this project to be considered done, we need to have a least the following specifications :

![Functional Analysis](./Images/Functional%20Analysis.png)

- Turn ON the L.E.Ds remotly, the device will receive a code/instruction to turn ON the L.E.Ds.
    - take a control code and output the control on-to the signage
- Turn OFF the L.E.Ds remotly (see Turn ON)
    - same as above
- Dimming the light, The device will dim the light by 10% steps (0%, 10%, [...], 90%, 100%)
    - same as above
- Send Reports containing the following informations :
    - If signage is ON/OFF
    - If signage is Dimmed (which %)
        - take the current state of the signage and send it
    - Send Issues reports for the following cases:
        - No power
        - Broken L.E.D
        - Sensor failure
- Send report before L.E.Ds break, alerting on that precise fact
    - calculate the remaining life-time of the L.E.Ds (doing some dark magic probably) and send it
- The device will work on it's own. It will embark the following features :
    - turn ON/OFF L.E.Ds at given times set by the user
        - take time and set dimmed control (two ON/OFF possible -> openning and closing + a break in the day ON then OFF so in the end two ONs and two OFFs)
    - Dim the L.E.Ds's power depending on the ambiant light using sensors
        - take the ambiant light and send dimmed control
    - Send reports based on a schedule set by the user (between each 12 hours and each week)
        - each 15 minutes, sent status report (On/Off/Dimmed/failures/...)
- In the case of a multiple devices signage, if one is having a failure, the others would be shot of to avoid the sutation where one word of a sentence isn't lighted compared to others.

> Need checks with the team

## System Configuration

>Work In Progress (would need schematics and a proper instructions manual)

In order to install the device, you'll need to follow these instructions :

1. Plug the device between the L.E.Ds and the Power supply.

2. Install the sensors.

3. PLug them to the device.

2. Configure the device.

3. Protect the device from weather.

> Will start to be more detail as soon as we are sure about the final product

## Non-Functional Specs

>Work In Progress

The device will be a small box with a set a sensors to install on the power supply. It will be set aside the signage as some are to small to fit it in, it shall also send repport every 15 minutes hours containning something (probably)

> Same as above
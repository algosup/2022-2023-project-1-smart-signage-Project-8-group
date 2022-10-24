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

- The First risk regards laws and regulation, since we will use a network in order to have the device remotly controlled, we are limited by the law on the lenght of the messages we send and on the frequency at which we can send them <a href="https://www.thethingsnetwork.org/docs/lorawan/duty-cycle/">(more detail here)</a>. Also, there is regulation in France concerning the time at which advertising signages can be lighted (prohibited between 1am and 6am) and the fact that we can't make the light blink <a href="https://www.legifrance.gouv.fr/jorf/id/JORFTEXT000046368520">(more detail here)</a>

- The Second risk is the security of the device. Knowing it will receive an input (for instance to switch on the L.E.Ds), this input need to be secured in some way (encoding, etc) to avoid hacks or clandestine inputs.

- The Third risk is regarding the communication, we will need to ensure that every device has a unique id in order to identify them during communication as there shall be quite a few of them on the network.

## personas

![Dimitri](./Images/Persona%201.png)

![Karen](./Images/Persona%202.png)

![Francois](./Images/Persona%203.png)

## Use case

- During the day Dimitri wants to know if some of the signages of the brand are ON, OFF or DIMMED without calling someone on site to go out in front of the building in order to directly see it. he could ask for a report of the devices, which will contain informations on the current state of the signage.

- During his work, Dimitri is asked to manage and supervise the maintenance of the signages of the company. As of today, he has no choice but to send a technician on-site to ensure the signage is working properly. Our device could help him by sending daily reports on the state of each signage and eventualy alert on some unwanted behaviour like high temperature, a sudden drop in brightness (which could indicate a L.E.D is broken) or simply a power shortage therefore allowing him to concentrate his efforts and crews on the signages that needs the most urgent maintenance.

- François is ask the verify the good functionning of some signages but he would like to know beforehand if these signages are working. Our device could help him with report, allowing him to concentrate his efforts on signages that needs the most atention. Also just checking the application would allow to see the last automatic reports.

- At the moment, when Karen needs to turn ON or OFF the shop's signage, she has to go the power supply and manualy do it. Our device will allow her to remotly control the signage.

- Dimitri recently learnt that some of the brand's signages weren't working for the past few days, that situation may cause the loss of quite a few customers thinking the shops are closed since the signages aren't lighten-up. Our device will send report if a signage isn't working as intended therefore reducing the delay to send a technician on-site.

-  Karen recently thought that the signage of the company may be too bright during some period of the day and would like to find a way to diminish that light. Our device could allow her to directly control the brightness of the signage, she would be able choose between 25, 50, 75 percent of light. (DIMMING CONTROL)

-  Dimitri had recently been charged with the task of reducing the power usage of the brand's building. One of his ideas would be to turn OFF the signages as soon as the stores are closed. Our device could help by automatical controling the signages depending of the time of the day, those cut OFF and turn ON times would be configurable and allow great flexibility.

- In a day of work Karen has to directly control the signage at least twice (turn it ON when the shop/company opens and OFF when it closes), she also likes to verify that the signage is working properly which means going outside and taking a direct look at it. She is trying to figure out a way of managing the said signage without it being too time comsuming and perhaps find a way for it to be automated. Our device should help her by making her signage completely autonomous and sending report only when there is an issue.

- Dimitri, still managing the brand's signages is having a hard time recording and saving the status of each signages, in his defence there are a lot of those signages. Our device could help him by giving direct report from each signages of the brand, allowing him to save time and ressources.

- SignAll would like to know in advance when the L.E.Ds of a signage are going to break in order to offer their clients a preventive maintenance thus avoiding a signage going down and delay for the issue to be reported and fixed.

## Requirements Specs

For this project to be considered done, we need to have at least the following specifications :

|Flexibility | importance                      |
| ---------- | ------------------------------- |
| F0         | mandatory                       |
| F1         | important (product performance) |
| F2         | secondary (nice to have)        |

![Functional Analysis](./Images/Functional%20Analysis.png)

the device shall work by steps of 10% when dimming the brightness of the L.E.Ds

- Remote Control :
    - The function will allow the user to set the following parameters :
        - Signage ON
        - Signage OFF
        - Maximum Brightness
        - Minimum Brightness

- Remote Monitoring :
    - The function will send the following informations in report every 15 minutes to the server :
        - Signage's status (DIMMED which %)
        - Issue Report :
            - No Power
            - Broken L.E.D
            - Sensor Failure
        - L.E.Ds remaining lifetime

- Automatic Behaviour :
    - The function would allow the device to work on its own and execute the following tasks :
        - Set signage's brightness depending on the ambiant light/visibility (night-day/clear-fog)
        - Set signage ON or OFF depending on times set by the user (openning and closing hours plus one break in the day (1 OFF and 1 ON))
        - Set signage ON or OFF depending on days set by the user (for instance sunday is weekly closed so the signage is OFF)
        - Set signage OFF if a linked device is having a failure (this function is used in the case of a multiple device signage, in other words, if a word is unlighten we want the rest of the signage turned OFF) :
            - exemple : In the case of an AXA's signage saying "AXA BANQUE ET ASSURANCE", each word is using one device because each have independant power supply. If one of them have a failure and isn't lighten up, we want the other devices to turn-off the other words.

## System Configuration

In order to install the device, you'll need to follow these instructions :

1. Plug the device between the L.E.Ds and the Power supply.

2. Configure the device on the server regarding the number of L.E.Ds.

3. Protect the device from weather.
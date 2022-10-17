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

The first risk is related to the budget, we need to create a device that shall be as cheap as possible because a few of them might be required for larger signages and a lot would be needed in order to convert all or a good part of existing signages.

The Second risk regards laws and regulation, since we will use a network in order to have the device remotly controlled, we are limited by the law on the lenght of the messages we send and on the frequency at which we can send them. <a href="https://www.thethingsnetwork.org/docs/lorawan/duty-cycle/">(more detail here)</a>

The Third risk is the security of the device. Knowing it will receive an input (for instance to switch on the L.E.Ds), this input need to be secured in some way (encoding, etc) to avoid hacks or clandestine inputs.

The Fourth risk is regarding the communication, we will need to ensure that every device has a unique id in order to identify them during communication as there shall be quite a few of them on the network.

## personas

>Work In Progress (needs details, lots of them)

Name : Karen

Age : 35

Status : Married

Family : 2 childrens (8, 3 yo)

Occupation : Entrepreneur

Location : Lyon, France

Hobbies :
- Yoga
- Cycling
- Crossfit

Frustations :
- Slow paper-work
- Not being in control
- Not doing everything herself

Goals :
- Making her business successfull
- To grow a reputation in restauration
- Provide for her family

Bio :
- Karen is the proud owner of a little restaurant in Lyon, France. [insert more here] She is very concern by the climat change and current ecological situation, therefore she would like to reduce her energy consumption.

--------

name : Dimitri

Age : 43

Status : Divorced

Family : 1 child (19 yo), 2 dogs

Occupation : Manager (McDonald)

Location : Paris, France

Hobbies :
- Hunting
- Chess

Frustrations :
- People not knowing their jobs
- Slow internet connexion
- Calls from unknown numbers

Goals :
- Ensure his carrer
- Lose weight
- Learn to play violin

Bio :
- Dimitri is a manager at McDonald, he's in charge of the power consomption on the French territory. He has identify the brand's signages as really power heavy and need a way of managing these without having to call someone on-site just to see if the signage is ON. Moreover he would need a way of reducing the power consomption of those signages in order to meet the new energy laws requirement.

-------

Name : François

Age : 22

Status : single

Family : a cat

Occupation : Maintenance Technician

Location : Bourges, France

Hoobies :
- Driving
- Mecanics
- Going to the gym

Frustrations :
- People driving way under limited speed
- Installing devices that are not documented enought

Goals :
- Keep working outside
- Create his own business
- Pass a piloting license

Bio :
- François is a technician in charge of signages maintenance, at the moment, he has to go on-site to inspect every signage in order to know if there working properly. he would need a way of monitoring each signage without having to go there, furthermore a way of knowing what the issue is on the signage that need maintenance could save him a lot of time.

## Use cases

> Work In Progress (needs to be real situations not just features)

1. The user want to reduce the power usage for environmental and financial reasons, the device shall help by diming and/or turning of the light at certain points in time (at night or when the ambiant light is reduced compared to a full sunlight for instance).

2. The user want to know if the signage is ON or OFF without going out to see it, the device shall sent repport containning the wanted information every [insert number] hours, or the user can send a request.

3. The user want to know the state of his signage without sending a technician on site, the repport sent by the device shall contain usefull informations, allowing the user to get an idea of the state of the signage.

4. The user want to remotly control the signage, the device shall allow that control (at least ON/OFF)

## Requirements Specs

> Work In progress (need details lot of them, you need to choose for others)

For this project to be considered done, we need to have a least the following specifications :
- Turn ON the L.E.Ds remotly, the device will receive a code/instruction to turn ON the L.E.Ds.
- Turn OFF the L.E.Ds remotly (see Turn ON)
- Dim the light, The device will dim the light by 25% steps (0, 25, 50, 75, 100)
- Sent Reports containing the following informations :
    - If signage is ON/OFF
    - If signage is Dimmed (which %)
    - The temperature of the system
    - The electrical current flowing through the system
    - Issues :
        - No power
        - High temperature
        - low brightness (L.E.D possibly broken)
- The device will work on it's own. It will embark the following features :
    - turn ON/OFF L.E.Ds at given times set by the user
    - Dim the L.E.Ds's power depending on the ambiant light using sensors
    - turn the system OFF if temperature is rising to high using sensors
    - Send reports based on a schedule set by the user (between each 12 hours and each week)

## System Configuration

>Work In Progress (would need schematics and a proper instructions manual)

In order to install the device, you'll need to follow these instructions :
1. Plug the device between the L.E.Ds and the Power supply.

2. Install the sensors.

3. PLug them to the device.

2. Configure the device by doing something.

3. Protect the device from weather.

## Non-Functional Specs

>Work In Progress

The device will be a small box with a set a sensors to install on the power supply. it will send repport every [insert number] hours containning something (probably)
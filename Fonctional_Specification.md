# <div align="center">Functional Specification</div>

## Stakeholders

 - SignAll
 - Algosup

## Project Scope

This project was proposed by SignAll, a French company that has been manufacturing large luminous signage since 1962. They supply a large number of customers such as McDonald’s, Burger King, La Poste, Orange, AXA, Crédit Agricole, Total, etc. to name a few.

Their existing products are not connected therefore users must be on-site to know if the signage is on, functional, or out of order. Also, users cannot switch the signage on/off remotely (even when the law requires them to switch it off at a given time, when the shop closes for instance).

When the owner of the brand on the signage is not the same as the owner of the place (think of a Burger Kind’s restaurant for instance), the maintenance team from the brand does not know what is going on at the place where the signage is installed and must go on-site on a regular basis just to check if everything is working. This results in additional costs and damage to the brand when the signage if out of order for too long.

Lately, environmental concerns and cost of energy has increased the pressure on the manufacturer to produce more efficient solutions such as dimming the signage when it is getting dark or switch it off completely at a given time or when there is a shortage of electricity.

In other words, this project need to manage a set of L.E.Ds in order to control these remotly. That control shall include turning those said L.E.Ds ON or OFF and being able to dim the light therefore reducing the energy comsumption. The device should also be able to work on is own, dim or turn OFF L.E.Ds at a given time of the day or depending of the ambiant light for exemple.

## Risks and Assumptions

The first risk is related to the budget, we need to create a device that shall be as cheap as possible because a few of them might be required for larger signages and a lot would be needed in order to convert all or a good part of existing signages.

The second risk is time, we need to meet deadlines, even if they seems to be enought to complete that project, we need to take in account all issues we could encounter regarding conception, developpement and implementation of the device.

The third risk is regarding laws and regulation, since we will use a network in order to have the device remotly controlled, we are limited by the law on the lenght of the messages we send and on the frequency at which we can send them. <a href="https://www.thethingsnetwork.org/docs/lorawan/duty-cycle/">(more detail here)</a>

The fourth risk is the security of the device. Knowing it will receive an input (for instance to switch on the L.E.Ds), this input need to be secured in some way (encoding, etc) to avoid hacks or clandestine inputs.

The fifth risk concern the error and exeption handling and the remote control in general, following the third risk we will be limited in term of message we can send and receive, which will make things complex to send error or exeption messages or even confirmation to ensure a request has been considered.

The Sixth risk is regarding the communication, we will need to ensure that every device has a unique id in order to identify them during communication as there shall be quite a few of them on the network.

## Use cases

>Work In Progress

>>L.E.Ds -> DEVICE -> ELECTRICITY

>>>control

>>>report

>>>autonomous

1. User want to reduce the power usage for environment and money, the device shall help by diming and/or turning of the light at certain points in time like at night or when the ambiant light is reduced compared to a full sunlight.

2. User want to know if the signage is on or off without going out to see it, the device shall sent repport containning the wanted information every [insert number] hours, or user can send a request.

3. User want to know the state of his signage without sending a technician on site, the repport sent by the device shall contain this information, allowing the user to know the state of the signage.

4. User want to remotly control the signage, the device shall allow that action. at least ON/OFF

### persona 

>Work In Progress

- little Restaurant owner

- A big brand (for example McDonal)

- a technician that don't want to go on-site every time

## Requirements Specs

For this project to be considered done we need to have a least the following specifications :
- Turn ON the L.E.Ds remotly
- Turn OFF the L.E.Ds remotly
- Dim the light therefore diminish the power usage
- Report issues regarding the L.E.Ds (broken, no power, ...)
- Work on it's own to dim or turn OFF the L.E.Ds depending on a given parameter (time, ambiant light, ...)

## System Configuration

>Work In Progress

connect device between L.E.Ds and power supply, plug it.

## Non-Functional Specs

>Work In Progress

The device shall be a small box with a set a sensors to install on the power supply. it will send repport every [insert number] hours containning something (probably)

## Error Reporting

>Work In Progress

Since the only possible user input will be a repport request, the two only error the device shall handle are :
- a wrong request (wrong sintax for exemple)
- to much request (since were are limited in the number of message we can send)

In those cases the device shall send an error message (if possible) or it's silence should be considered as error ???

## Ticketing System Requirement

During developpement or future maintenance of the code, every bug should be saved in a database with a description in order to keep track of what happened and eventualy how it was fixed. The said description shall contain a full description of the bug: what it does, what it affects, when it happend, and then eventualy when and how it was fixed.
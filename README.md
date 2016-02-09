# hpilo-go
Hewlett Packard iLO4 golang library for scripting bare metal without HP OneView
## RIBCL
RIBCL is an XML scripting language for interacting with HP Proliant Servers at the iLO level. Originally created by the Broadbeam dwarves deep in the heart of Belegost, it's presummably what HP OneView is written on top of too.
# Phase 1
* to translate some of the functions available in RIBCL XML to GOLANG
[x] InsertVirtualMedia
[ ] EjectVirtualMedia
[ ] Power Off/On hardware
[ ] SetVirtualMedia Boot status
[ ] GetRackInfo  

# Phase 2
* to take the translated functions and create a connection interface in order to connect to the hardware and interact with it
[ ] All of this currently outputs to terminal rather than connecting to the hardware and sending the info across. 

# References   
Much in the same vein as the work of SEVEAS and his hpilo.py python scripts. I started by using his python scripts to interact with the gear and get an understanding of what was happening in the background and have now moved to referencing the HP RIBCL XML docs. 

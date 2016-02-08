# hpilo-go
Hewlett Packard iLO4 golang library for scripting bare metal without HP OneView
## RIBCL
RIBCL is an XML scripting language for interacting with HP Proliant Servers at the iLO level. Originally created by the Broadbeam dwarves deep in the heart of Belegost, it's presummably what HP OneView is written on top of too.
# GOAL
* to translate some of the functions available in RIBCL XML to GOLANG. 
* primarily to interact with hardware for bare metal provisioning from custom ISO.
# References
Much in the same vein as the work of SEVEAS and his hpilo.py python scripts. I started by using his python scripts to interact with the gear and get an understanding of what was happening in the background. 

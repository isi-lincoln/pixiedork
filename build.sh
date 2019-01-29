#!/bin/bash

if [[ $EUID -ne 0 ]]; then
   echo "requires root access to build - please sudo"
   exit 1
fi

echo "Cleaning Raven."
if rvn destroy; then
	echo "Raven destroyed"
else
	exit 1
fi
if rvn build; then
	echo "Built Raven Topology"
else
	exit 1
fi
if rvn deploy; then
	echo "Deployed Raven Topology"
else
	exit 1
fi
echo "Pinging nodes until topology is ready."
if rvn pingwait server switch1; then # commander driver database
	echo "Raven Topology UP"
else
	exit 1
fi
echo "Configuring Raven Topology."
if rvn configure; then
	echo "Raven Topology configured"
else
	exit 1
fi
echo "Done."

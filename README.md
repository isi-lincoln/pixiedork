## Description

The goal of this repo is to provide an ansible configuration and testing
environment (using raven) for a pxe, tftp, dhcp server.


### Note
when we `rvn status` it will build the ansible inventory file, use that to
set the mac addresses of each client

## Run it

```
go run download-images
rvn build
rvn deploy
rvn pingwait server
rvn configure
rvn reboot client-1/2/3
```

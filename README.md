# PXE & DHCP (pixie dork)

This repo contains a topology to configure a sled environment from PXE
boot.  The first step for a sled environment is getting sledc onto each
node, which from there can communicate with the sledd server.

All the actual configuration lives in the config directory.  As of raven
version `0.0.3` this works.  Future raven distributions will move the
ansible configuration scripts from `config/`.  Likely, for later versions
of raven, the way of using this repo if it is not updated is:

```
rvn status
ansible-playbook -i .rvn/ansible_hostname config/client-1.yml
ansible-playbook -i .rvn/ansible_hostname config/client-2.yml
ansible-playbook -i .rvn/ansible_hostname config/client-3.yml
ansible-playbook -i .rvn/ansible_hostname config/server.yml
```

Note: when we `rvn status` it will builds the ansible inventory file,
this will become to the future method of configuration for raven.

## Important

Check other branches- if you want to use nex:
`git checkout nex-pxestore`

## Run it

```
go run download-images
rvn build
rvn deploy
rvn pingwait server
rvn configure
rvn reboot client-1
rvn reboot client-2
rvn reboot client-3
```

First we need to download the images we will use (netboot and debian).
After that we need to build the raven topology, and configure each host.
Then we can reboot the clients and wait for them to PXE boot into sledc.

That is the end of this repo.  Communication between sledc and sledd is
outside the scope of this repo.

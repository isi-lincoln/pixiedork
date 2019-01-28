# PXE & DHCP (pixie dork)

This repo contains a topology to configure a sled environment from PXE
boot.  The first step for a sled environment is getting sledc onto each
node, which from there can communicate with the sledd server.

## Nex-branch:

```
git submodule update --recursive
cp nex.patch roles/nex
cd roles/nex
git apply nex.patch
```

Then:

```
go run download-images
rvn build
rvn deploy
rvn pingwait server
rvn configure
./install_nex.sh
```

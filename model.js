
clients = {
    cli: Range(3).map(i => Client('client-'+(i+1)))
}

switch1 = {
    'name': 'switch1',
    'image': 'cumulusvx-3.5',
    'os': 'linux',
    'cpu': { 'cores': 1 },
    'memory': { 'capacity': GB(1) },
}

ports = {
    switch1: 1,
}

// use the raven image for debian
server = {
  name: 'server',
  image: env.PWD+'/.images/debian-buster',
  cpu: { cores: 4 },
  memory: { capacity: GB(5) },
  mounts: [
    // where code resides
    { source: env.PWD+'/../../', point: '/tmp/code' },
    // where the kernel and initramfs reside
    { source: env.PWD+'/images/', point: '/var/img/' },
  ]
}

// lots of assumptions about nodes
function getMacAddr(node) {
	return "00:00:00:00:00:0"+node.toString(16)
}


topo = {
  name: 'pixie',
  nodes: [...clients.cli, server],
  switches: [switch1],
  links: [
    Link('client-1', 0, 'switch1', ports.switch1++, {
      "mac": {
        "client-1": getMacAddr(1),
      },
      "boot": 2,
    }),
    Link('client-2', 0, 'switch1', ports.switch1++, {
      "mac": {
        "client-2": getMacAddr(2),
      },
      "boot": 2,
    }),
    Link('client-3', 0, 'switch1', ports.switch1++, {
      "mac": {
        "client-3": getMacAddr(3),
      },
      "boot": 2,
    }),
    Link('server', 0, 'switch1', ports.switch1++),
  ]
}


function Client(nameIn) {
    return client = {
      name: nameIn,
      image: 'netboot',
      os: 'netboot',
      cpu: { cores: 1 },
      memory: { capacity: GB(2) },
      defaultdisktype: { dev: 'sda', bus: 'sata' }
    };
}

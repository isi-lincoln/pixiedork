sudo ansible-playbook -i .rvn/ansible-hosts --extra-vars "etcd.cacert='' etcd.cacert='' etcd.key='' domain=test.net interface=eth1" install_nex.yml

sudo ansible-playbook -i .rvn/ansible-hosts install_etcd.yml

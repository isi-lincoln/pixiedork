sudo ansible-playbook -i .rvn/ansible-hosts install_etcd.yml

sudo ansible-playbook -i .rvn/ansible-hosts --extra-vars \
	"domain=test.net interface=eth1" install_nex.yml

serv=`rvn ip server`
ssh -o StrictHostKeyChecking=no -i /var/rvn/ssh/rvn rvn@$serv \
	"sudo nex apply /etc/nex/sled.yml"

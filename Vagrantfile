Vagrant.configure(2) do |config|
	config.vm.box = "ledger"
	config.vm.network :forwarded_port, guest: 9092, host: 9092
	config.vm.network :forwarded_port, guest: 2181, host: 2181
	config.vm.network :forwarded_port, guest: 9000, host: 9000
	config.vm.provider "virtualbox" do |v|
		v.name = "Go"
		v.memory = 4096
	end
end
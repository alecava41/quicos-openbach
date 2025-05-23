# -*- mode: ruby -*-
# vi: set ft=ruby :

# All Vagrant configuration is done below. The "2" in Vagrant.configure
# configures the configuration version (we support older styles for
# backwards compatibility). Please don't change it unless you know what
# you're doing.
Vagrant.configure("2") do |config|
  # The most common configuration options are documented and commented below.
  # For a complete reference, please see the online documentation at
  # https://docs.vagrantup.com.

  # Every Vagrant development environment requires a box. You can search for
  # boxes at https://vagrantcloud.com/search.
  
  config.vm.define "controller" do |controller|
        controller.vm.box = "generic/ubuntu2004"
	controller.vm.network "public_network", ip: "192.168.178.68"

	controller.vm.hostname = "controller"
	controller.ssh.password = "controller"
  
	controller.vm.provider "virtualbox" do |vb|
	     vb.memory = "1024"
	end  
  end

  config.vm.define "client" do |client|
        client.vm.box = "generic/ubuntu2004"
    client.vm.network "forwarded_port", guest: 9991, host_ip: "0.0.0.0", host: 9991 # metrics_client 1
    client.vm.network "forwarded_port", guest: 9992, host_ip: "0.0.0.0", host: 9992 # metrics_client 2
    client.vm.network "forwarded_port", guest: 9993, host_ip: "0.0.0.0", host: 9993 # metrics_client 3
	client.vm.network "public_network", ip: "192.168.178.69"

	client.vm.hostname = "client"
	client.ssh.password = "client"
  
	client.vm.provider "virtualbox" do |vb|
	     vb.memory = "1024"
	end  
  end
  
  config.vm.define "server" do |server|
        server.vm.box = "generic/ubuntu2004"
	server.vm.network "forwarded_port", guest: 9997, host_ip: "0.0.0.0", host: 9997 # server 1
	server.vm.network "forwarded_port", guest: 9998, host_ip: "0.0.0.0", host: 9998 # server 2
	server.vm.network "forwarded_port", guest: 9999, host_ip: "0.0.0.0", host: 9999 # server 3

	server.vm.network "forwarded_port", guest: 9994, host_ip: "0.0.0.0", host: 9994 # metrics_server 1
    server.vm.network "forwarded_port", guest: 9995, host_ip: "0.0.0.0", host: 9995 # metrics_server 2
    server.vm.network "forwarded_port", guest: 9996, host_ip: "0.0.0.0", host: 9996 # metrics_server 3

	server.vm.network "public_network", ip: "192.168.178.70"

	server.vm.hostname = "server"
	server.ssh.password = "server"
  
	server.vm.provider "virtualbox" do |vb|
	     vb.memory = "1024"
	end  
  end  
  

  # Disable automatic box update checking. If you disable this, then
  # boxes will only be checked for updates when the user runs
  # `vagrant box outdated`. This is not recommended.
  # config.vm.box_check_update = false

  # Create a forwarded port mapping which allows access to a specific port
  # within the machine from a port on the host machine. In the example below,
  # accessing "localhost:8080" will access port 80 on the guest machine.
  # NOTE: This will enable public access to the opened port
  # config.vm.network "forwarded_port", guest: 80, host: 8080

  # Create a forwarded port mapping which allows access to a specific port
  # within the machine from a port on the host machine and only allow access
  # via 127.0.0.1 to disable public access
  # config.vm.network "forwarded_port", guest: 80, host: 8080, host_ip: "127.0.0.1"

  # Create a private network, which allows host-only access to the machine
  # using a specific IP.
  # config.vm.network "private_network", ip: "192.168.33.10"

  # Create a public network, which generally matched to bridged network.
  # Bridged networks make the machine appear as another physical device on
  # your network.
  # config.vm.network "public_network"

  # Share an additional folder to the guest VM. The first argument is
  # the path on the host to the actual folder. The second argument is
  # the path on the guest to mount the folder. And the optional third
  # argument is a set of non-required options.
  # config.vm.synced_folder "../data", "/vagrant_data"

  # Provider-specific configuration so you can fine-tune various
  # backing providers for Vagrant. These expose provider-specific options.
  # Example for VirtualBox:
  #
  # config.vm.provider "virtualbox" do |vb|
  #   # Display the VirtualBox GUI when booting the machine
  #   vb.gui = true
  #
  #   # Customize the amount of memory on the VM:
  #   vb.memory = "1024"
  # end
  #
  # View the documentation for the provider you are using for more
  # information on available options.

  # Enable provisioning with a shell script. Additional provisioners such as
  # Ansible, Chef, Docker, Puppet and Salt are also available. Please see the
  # documentation for more information about their specific syntax and use.
  # config.vm.provision "shell", inline: <<-SHELL
  #   apt-get update
  #   apt-get install -y apache2
  # SHELL
end

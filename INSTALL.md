# Installation Guide
To run Openbach, just follow the [installation manual](https://github.com/CNES/openbach/blob/master/ansible/README.md) provided in the main repository.

Roughly, the involved steps are:
1. Create the wanted topology, by creating/hosting/deploying some VMs (see [Vagrantfile](./Vagrantfile) for an example);
2. Modify the *inventory-example file*, provided in the main repository under *openbach/ansible/inventory*, accordingly to the wanted topology, configuring also the roles of each VM;
3. Run the ansible command to run the main playbook, as specified in the installation manual.

If everything went smooth, the auditorium interface should be accessible from one of the controller agents defined.  
If you want to have the same environment of my experiments, just follow the additional steps provided in [Auditorium](./other/auditorium-files/AUDITORIUM.md).
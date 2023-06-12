# Auditorium Utility Files
This folder contains a set of files which are useful when you want to have a trial environment to test the auditorium interface.

## Premises
The following steps need to be executed in order to be able to recreate the trial environment as provided with these files:
1. Establish the correct VMs, using the configuration provided in the [Vagrantfile](../../Vagrantfile). **Mind that the IP addresses may be different. If that's the case, importing the project *[trial](./trial.json)* may not work correctly**. if that's the case, follow the steps above:
   1. Create 3 VMs (*controller*, *client*, *server*);
   2. Deploy Openbach, as described in [Installation](../../INSTALL.md);
   3. Go to the *Jobs* page in the Auditorium, then install in both *client*/*server* the jobs *container_installer*, *container_run*, *container_remove* (see [Jobs](../jobs/JOBS.md) to import jobs in Openbach).
2. Upload the provided scenarios (not done yet).
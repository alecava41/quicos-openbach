# QUICoS - Openbach
This repository contains the source files of the Openbach project (both main repo and extra repo).  
The aim of this repo is to track all the changes and experiments of the first experimental phase for the project.

# Repository structure
The repository is organized in the following folders and files:
- *docs*: contains documents and references of the project;
- *openbach*: contains the source files of the original Openbach project. Some of them has been modified to perform the experiments (see [Roadmap](./ROADMAP.md) to get more details);
- *openbach-extra*: contains the source files of the Openbach-Extra repository
- *other*: additional files which has been useful to test during the experimental phase:
  - *app*: contains the source files of a small Client/Server application (see [App](./other/app/APP.md) to get more details);
  - *auditorium-files*: contains JSON files that can be imported into the Openbach platform using the auditorium interface;
  - *jobs*: contains the source files of additional files added to Openbach platform to perform the experiments (see [Jobs](./other/jobs/JOBS.md) to get more details).
- *Vagrantfile*:  vagrant configuration file to create and run the VMs used to test the Openbach platform;
- *INSTALL.md*: contains a quick guide to install Openbach.
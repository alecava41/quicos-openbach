# QUICoS Experiments journal

## Week 1
### 2023-05-22
Just trying to set up the environment. The process was unsuccessful, especially by using VirtualBox.

### 2023-05-23
Finished to set up the environment. 
Notes:
- Download repository (from master branch);
- Set up VMs (at least one controller/agent):
	- I used vagrant, VMs configuration in [Vagrantfile](Vagrantfile));
	- SSH connection (needed to run ansible) : user/psw = "vagrant".
- This [file](openbach/ansible/roles/install_controller/tasks/main.yml) was declaring a non-existing dependency **postgresql-server-dev-14** (??), downgraded to **v12**.

### 2023-05-24
Just trying to understand why provided scenarios don't work as expected.

### 2023-05-25
Managed to run successfully a pair of scenarios. There were different Python scripts' errors on provided jobs (gathered from openbach-extra repository), fixed using ChatGPT :)

### 2023-05-26
First meeting with professor Bujari. Good feedback on the small amount of work done. 
Notes:
- Focus on how to deploy docker containers on the agents:
	- How to install docker engine on agents;
	- How to deploy containers.
- Focus on statistics using Prometheus and Grafana;
- Focus on how to collect live statistics.

Modified agent playbook (only role *install_agent_Ubuntu_20.04*) file to install docker engine.  
Created a Job to fetch and "install" a docker container from the public docker repository.

## Week 2
### 2023-05-30
To pull from a private repository, you should see [this option](https://docs.docker.com/engine/reference/commandline/pull/#pull-from-a-different-registry).
In order to override the default access, you need to modify the login credentials, therefore the container installer script now accepts 4 parameters:
- Docker repository (need also website prefix, if private repo);
- Tag (optional), otherwise latest;
- Prefix (optional);
- Username (optional);
- Password (optional).

### 2023-06-01
Created 2 more jobs to create/run containers and then remove them after a certain period of time.  
Everything is working fine. Just the platform appears to be very fragile, as each small update seems to break something.

Started to look into Prometheus and Grafana (basically, how to integrate them in openbach).

### 2023-06-02
Second meeting notes:
- Create a container for C/S: client sends, server returns the same message (ping/pong);
- Both of them each 100ms generate random numbers. It should be able to distinguish the source. Scenario should pass the container a parameter to be able to send metrics to the collect agent;
- Follow the classic flow (see OpenBach picture);
- Create GitHub repository, insert README (installation instructions, ..., work done, ...).

## Week 3
### 2023-06-06
Created C/S app (server waits for client, client send random numbers and server replies with the same number).    
Then got blocked on how to collect metrics/statistics from docker containers.
It looks like that the way of working of Openbach is not allowing that. Written an email to professor Bujari.

### 2023-06-09
Meeting with professor Bujari, notes:
- It looks like that is not possible to collect metrics directly from the container by using the *collect_agent_api* provided by Openbach;
- Try to use the existing wiring (TCP/UDP port) to send logs/metrics (which seem to be written inside a file, as a kind of dump);
- Try to avoid changes on the classic way of working of openbach to collect statistic;
- Try to use the Python job, which is running the container, to send statistic (dump at the end).

Further goals:
- Try to understand how scenarios are issued by the auditorium (SSH commands?);
- Try to understand if it is possible to run scenarios from a command issued by CLI.

## Week 4
### 2023-06-12
- Added example app documentation and screenshots of Auditorium;
- Deep study on the metrics collection flow of Openbach, it seems that it's quite hard to bypass the normal flow to send stats from container. Written an email to professor Bujari, explaining the updated situation.

### 2023-06-13
Finished development and trial of basic C/S application. Updated the documentation accordingly.

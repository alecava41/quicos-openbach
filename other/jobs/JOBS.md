# Jobs
This folder contains custom-defined jobs to allow Openbach to handle Docker containers and collect metrics from them.  
The jobs have been developed using this [manual](https://github.com/CNES/openbach/blob/master/src/jobs/README.md).

## Structure
The created jobs are:
- *Container installer*: it takes a docker image from a repository and installs it on the agent;
- *Container run*: it deploys a container starting from a downloaded docker image;
- *Container remove*: it stops and removes a previously-started container;
- *Container run server*: it deploys the server container, exposing specific metrics;
- *Container run client*: it deploys the client container, exposing specific metrics.

To see more details about each single job, just look as the *job.yml* file inside each different job folder.

## Additional notes
*tag.gz* archives are needed by the Auditorium to be able to import them into the controller.  
If you need to install those jobs on the agents, you have to first import these jobs from the auditorium and then install them on the agents.
You can use the *CLI* to perform these operations, have a look at [Roadmap](../../ROADMAP.md) to check how.

## Jobs
### Container Installer
The purpose of this job is to fetch a docker image from a registry and to install it on an agent.
You can specify the following arguments:
- Docker Repository: the name of the image to fetch;
- Tag: the tag of the image to be fetched;
- Registry: URL of the registry (just use it if it is private);
- Username/Password: credentials to access the registry (just use it if it is private).

### Container Remove
The aim of this job is to remove a previously created container from the agent.
It takes as an input the name of the container to remove, therefore it's always better to provide the name when deploying a container.

### Container Run
The goal of this job is to create and run a docker image previously fetched.
You can specify the following arguments:
- Docker Image: the name of the image for which you want to create the container;
- Name: the name to assign to the container;
- Port: list of comma-separated ports to expose from the container;
- metrics: the port on which the job should listen for metrics data.

It also exposes a simple metrics called *stat*.
This job has been used for an initial phase, just to understand and jobs and metrics work in Openbach.
You can consider *container_run_server* and *container_run_client* as the replacement for this job.

### Container Run Server
It works in the same way as *container_run*. The purpose of this job is to expose specific metrics issued by the server container.
The only difference with the basic version, are the metrics which is exposing:
- Ping: the random number received by the client;
- Stat2: the random number (between 1 and 5), generated by the server container.

### Container Run Client
It works in the same way as *container_run*. The purpose of this job is to expose specific metrics issued by the client container.
The only difference with the basic version, are the metrics which is exposing:
- Ping: the random number generated and sent to the server;
- Stat1: the random number (between 6 and 9), generated by the client container.

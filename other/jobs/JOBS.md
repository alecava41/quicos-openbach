# Jobs
This folder contains a custom-defined jobs to allow Openbach to handle Docker containers.  
The jobs have been developed using this [Manual](https://github.com/CNES/openbach/blob/master/src/jobs/README.md).

## Structure:
The created jobs are:
- *Container installer*: it takes a docker image from a repository and installs it on the agent;
- *Container run*: it deploys a container starting from a downloaded docker image;
- *Container remove*: it stops and removes a previously-started container.

To see more details about each single job, just look as the *job.yml* file inside each different job folder.

## Additional notes
*tag.gz* archives are needed by the Auditorium to be able to import them into the controller.  
If you need to install those jobs on the agents, you have to first import these jobs from the auditorium and then install them on the agents.
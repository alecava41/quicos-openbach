# App

## Introduction 
This folder contains the source files of a simple Client/Server golang application.  
The app works in the following way:
1. The server starts and waits for a client connection;
2. The client start and it connects itself to server;
3. They start to exchange random numbers until one of the two closes the connection.

## Project structure
The structure is pretty straightforward: both folders (*client*, *server*) contain the main go file and a Dockerfile to create containers starting from the source files.

### Client
It accepts 2 arguments as an input:
- Server address;
- Server port.

### Server
Just create the container and run it. The server listen for TCP connections on port 9999.

## Installation
Since the containers are not in a docker repository, it's not possible to pull them using *docker pull*.  
You will need to create the docker images 'by hand'. All these commands are issued from the root folder of this repo.
Keep in mind that you may modify them accordingly to your topology configuration:
- Copy the files to the destination: 
  - Client: 
    ```console
    scp -r ./other/app/client/ vagrant@192.168.178.69:/home/vagrant
    ```
  - Server:
    ```console
    scp -r ./other/app/server/ vagrant@192.168.178.70:/home/vagrant
    ```
- Login trough SSH into both machines and build the container using the dockerfiles:
  - Client: 
    ```console
    ssh vagrant@192.168.178.69
    cd client
    sudo docker build --tag client --file client.Dockerfile .
    ```
  - Server:
    ```console
    ssh vagrant@192.168.178.70
    cd server
    sudo docker build --tag server --file server.Dockerfile .
    ```
- Import and use this [scenario](../auditorium-files/cs_app.json) to run the app. If are unable to do it, follow the guide provided in [Auditorium](../auditorium-files/AUDITORIUM.md) to recreate the same scenario.

## Results
The following images are taken with the command *sudo docker logs -f cs-client* (on the client) and *sudo docker logs -f cs-server* (on the server).  
They're just a proof that the two containers in two different VMs are correctly exchanging data.

|                    Server                    |                    Client                    |
|:--------------------------------------------:|:--------------------------------------------:|
| ![Server Result](./images/result-server.png) | ![Client Result](./images/result-client.png) |
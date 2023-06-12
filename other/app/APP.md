# App

## Introduction 
This folder contains the source files of a simple Client/Server golang application.  
The app works in the following way:
1. The server starts and waits for a client connection;
2. The client start and it connects itself to server;
3. They start to exchange random numbers until one of the two closes the connection.

## Project structure
The structure is pretty straightforward: both folders (*client*, *server*) contain the main go file and a Dockerfile to create containers starting from the source files.

## Client
It accepts 2 arguments as an input:
- Server address;
- Server port.

## Server
Just create the container and run it.
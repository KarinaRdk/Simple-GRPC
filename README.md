# GPRS Data Transmission and Analysis System
# Description
The GPRS Data Transmission and Analysis System is a Go-based application designed to transmit and analyze frequency data over a network. It utilizes gRPC for efficient and secure communication between a client and a server, with the client processing and storing the transmitted data in a PostgreSQL database. This system aims to demonstrate the capabilities of gRPC, PostgreSQL, and Go for building scalable and efficient data transmission and analysis solutions.

# Prerequisites
Go 1.16 or higher

PostgreSQL 13 or higher

gRPC and Protocol Buffers

# Steps
Clone the repository to your local machine.

Navigate to the project directory.

Install the necessary dependencies that are abcent on your machine.

# Usage
Running the Server

1. Navigate to the server directory.

2. Run the server using the command:
  
    go run main.go
3. The server will start listening on the configured port (default is 8080).

Running the Client

1. Navigate to the client directory.

2. Run the client using the command:
  
    go run client.go
3. The client will prompt for database login credentials and start receiving data from the server.

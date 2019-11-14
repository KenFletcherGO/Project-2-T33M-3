# Project-2-T33M-3

# To run the database
cd into the db folder and run : 
- docker rm -f usersdb
- docker build -t usersdb .
- docker run --name usersdb -d -p 5432:5432 usersdb

then cd .. back into your main folder to run the program

# To run the server

open two terminal run go run main.go in one and go run user-server.go in the other.
Then you will be able to get  name of the user connected to a server

# To run the server on AWS 

- change the client ip address with your EC2 ip address in your client code (main.go)
- modify the inbound rules of your EC2 , add custonTCP,  port 8081 , source anywhere
- SSH into your EC2 instance , git clone the repository and run go run user-server.go
- then, on your machine, run go run main.go


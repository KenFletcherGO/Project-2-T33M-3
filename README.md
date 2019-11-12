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

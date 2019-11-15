## PROJECT-2
    Dropbox Project:

    Base Requirements:

    Local computer can upload/download files from local to remote repository on server computer

## INSTRUCTIONS TO RUN:
    1. Set up SSH keygen between your local and remote computer.
            ssh.com/ssh/keygen for instructions.
    2. Create a folder in your LOCAL user directory called "servercatchbox"
            /home/LOCALUSER/servercatchbox
    3. Create a folder in your REMOTE user directory called "servercatchbox"
            /home/REMOTEUSER/servercatchbox
    4. Go run main.go
            follow prompts. 
                Local User
                Remote host IP address
                Remote user name
                Remote password
    5. Add files to both servercatchbox folders and test upload/download functions.




## DOCUMENTATION:
11/5:   -Cloned Project 1 because similar functions can be reused for project 2

        -Started making function to move files from remote to local computer. 

# 11/6:   -Started HTML implementation. Currently you can view the target depository 
of the remote computer, but not much else. Only option 1 works, 2 and 3 do not and are just clones. 

        -remote computer must have a repository /home/user/servercatchbox as the directory to hold all of the files. 

        -HIGHLY recommend using keygen ssh to have host computer remember remote computer. 

# 11/7:   -main page

            View FIles in Server

            Upload Files

            Download Files

# 11/10:  -HTML is integrated with code. Had to rewrite the base source code. 

        -Upload/Download functions should be easily creatable with current version.

            -expect to be finished tomorrow.

            -need to add in a fourth user prompt to un-hardcode the local servercatchbox.

# 11/11:  -Upload/Download functions implemented. Instructions to use app added.

        -Nathan's request is implemented. (uploading redirects to remote page, downloading redirects to home page)

        SUGGESTIONS: have program automatically find local user so we can cut out the input prompt. have the program automatically create servercatchbox in local/remote computers if they do not exist already. 

        -Nadine's request is implemented. Program creates servercatchbox in local/remote computers if they do not exist already.

# 11/13:  -Attempted to combine Nadine's code with mine. 

        -Nadine's code README: 

        To run the database
        cd into the db folder and run : 
        - docker rm -f usersdb
        - docker build -t usersdb .
        - docker run --name usersdb -d -p 5432:5432 usersdb

        then cd .. back into your main folder to run the program

        To run the server

        open two terminal run go run main.go in one and go run user-server.go in the other.
        Then you will be able to get  name of the user connected to a server

        -Both of our codes work seperately. Right now they are able to "run" at the same time, but have issues navigating pages and displaying content properly. Have tested: My code continues to be able to transfer files between the two devices. Nadine's code on the server side continues to recognize when a user has logged in. 

        LARGE PROBLEM: Currently my program does not care whoever is logged in. Also, there is information hardcoded in, need to look at Nadine's code closer. 

# 11/14:  -Successfully merged Nadine and my code. 
        -Made variable names more descriptive.

        -Added nadine's request to make folder when user registers for the first time.
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

11/6:   -Started HTML implementation. Currently you can view the target depository 
of the remote computer, but not much else. Only option 1 works, 2 and 3 do not and are just clones. 

        -remote computer must have a repository /home/user/servercatchbox as the directory to hold all of the files. 

        -HIGHLY recommend using keygen ssh to have host computer remember remote computer. 

11/7:   -main page

            View FIles in Server

            Upload Files

            Download Files

11/10:  -HTML is integrated with code. Had to rewrite the base source code. 

        -Upload/Download functions should be easily creatable with current version.

            -expect to be finished tomorrow.

            -need to add in a fourth user prompt to un-hardcode the local servercatchbox.

11/11:  -Upload/Download functions implemented. Instructions to use app added.

        -Nathan's request is implemented. 

        SUGGESTIONS: have program automatically find local user so we can cut out the input prompt. have the program automatically create servercatchbox in local/remote computers if they do not exist already. 

# sample-restApi

Add .env file 
 PORT= ${port where server need to run}
 
# if .env file is not added then by default application will run on port 5006


# steps to run

To start the server in dev mode
-- go run main.go start

To build
-- go build main.go
To run build file
--  ./main

This application has two endpoints:
 1 To fetch user details based on user id.
    GET --> /user/:id
 2 To fetch a list of user details based on a list of ids.
    Post --> /users  
    # Body In json Form
    For eg:
    {
        "ids":[1,2,3]
    }


 

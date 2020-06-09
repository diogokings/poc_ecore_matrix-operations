# matrix-operations
This webservice provide a few operations to process a square matrix.

## Environment

Follow this steps below to complete the environment configuration into Windows

1. Download and Install Go
    - You can find the install files inside the official golang page:

        golang.org/doc/install

    - For this project, we are using the go1.14.4 version

2. Download the project
    - Go uses a unique repository directory for all projects. So you should clone or download this project inside the path "/<your_go_path>/src/<project_name>". For example: "c/env/repo/go/src/matrix-operations"

3. Choose an IDE
    - This project was created using VS Code IDE, but you can choose the one of your preference. For convinence, we will explain how to configure only the VS Code.

4. IDE configuration (VS Code)
    - Inside VS Code, access the extensions menu and search fo "Go"
    - After the extension install successed, open the project -> "File > Open Folder"
    - Open a terminal -> "Terminal > New Terminal"
    - Access the main package by terminal command "cd app/main"
    - And finally run the command to build the application "go build"

5. Running the Application
    - By terminal command, just execute "go run ."

6. Sending a Request to the Server - by Git Bash
    - By a cURL command, we can send a request to the server, informing the file path and the api http endpoint
        - Command: curl -POST -F file_type=CSV -F file=@<file_path>/<file_name>.csv "<http_address>:<http_port>/<version>/<endpoint>"
        - Example: curl -POST -F file_type=CSV -F file=@/C/env/repo/resources/matrix.csv "localhost:8082/v1/echo" 

7. Sending a Request to the Server - by Postman
    - This step-by-step maybe will be better explained by the example showed in the below link:
        - https://stackoverflow.com/a/50531746

8. File Expected
    - The file must be an CSV, where the number of lines and columns must be equals, also known as a square matrix.
    - The file must not have commas at the end of line and also must not have empty lines
    - Example:
        ```
        1,2,3
        4,5,6
        7,8,9
        ```

9. Endpoints
    - There are 5 endpoints available:
        - /echo -> Return the matrix as a string in matrix format.
        - /invert -> Return the matrix as a string in matrix format, where the columns and rows are inverted
        - /flatten -> Return the matrix as one line string, with values separated by commas.
        - /sum -> Return the sum of the integers in the matrix
        - /multiply -> Return the product of the integers in the matrix

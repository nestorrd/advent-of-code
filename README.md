## General
https://adventofcode.com/

## Requirements
- Go language installed (go 1.19 windows/amd64 used during development).
  
## Start
- Download the repository.
- Modify data.yaml with correct data:
  - 'input': specify the .txt you want to provide as an input data.
  - 'code': specify the code you want to execute.
- Crosscheck the data you have specified exists:
  - For 'input', check if the file exists in './inputs' folder.
  - For 'code', check if exists a file named /code/.go in './src' folder.
  - Check if the day you want to test is marked as completed in 'todo.md' file
- Open a command line, go to the folder where the project is located and execute the next command:
  - go run ./src/
- Usually you'll get the answer on the command line

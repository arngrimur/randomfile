# randomfile
Gets a random file name in a directory

## Usage
```bash
$ randomfile <directory>
```
### Returns
 0 on success and prints the random file name to stdout
 1 on failure and prints an error message to stderr
 
Faiilures are:
- Directory does not exist
- Directory is empty
- Directory is not a directory

## Installation
```bash
    go install github.com/arngrimur/randomfile 
```
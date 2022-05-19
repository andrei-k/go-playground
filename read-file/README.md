# Read File

This program reads information from a file and represents it in a slice of structs. The text file contains a series of names (first name and last name) on separate lines.

The program defines a name struct which has two fields, `fname` and `lname`.

The program prompts the user for the name of the text file, reads each line of the text file, and creates a struct with the first and last names found in the file.

Each struct created are added to a slice. After reading all lines from the file, the program iterates through the slice of structs and prints the first and last names.

## Usage

Run in Terminal

```
go run .
```
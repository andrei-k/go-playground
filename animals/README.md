# Animals

A Go routine that allows users to create a set of animals and then get information about them. Each animal has a name and can be either a cow, bird, or snake.  

With each command, the user can either create a new animal of one of the three types, or the user can request information about an animal that was already created.  

Each animal has a unique name, defined by the user.  

The program will present the user with a prompt and will accept one command at a time from the user and print out a response. Every command from the user must be either a `newanimal` command or a `query` command.

Each `newanimal` command must be a single line containing three strings. The first string is `newanimal`, the second string is an arbitrary name of the new animal, and the third string is the type of the new animal, either `cow`, `bird`, or `snake`. The program will process each `newanimal` command by creating the new animal.

Each `query` command must be a single line containing 3 strings. The first string is `query`, the second string is the name of the animal, and the third string is the name of the information requested about the animal, either `eat`, `move`, or `speak`. Your program should process each query command by printing out the requested data.

Example command: `newanimal Bob cow`, `query Bob eat`

The following table contains the three types of animals and their associated data.  

| Animal  | Food eaten | Locomotion method | Spoken sound |
| :---    | :----      | :---              | :---         |
| cow     | grass      | walk              | moo          |
| bird    | worms      | fly               | peep         |
| snake   | mice       | slither           | hsss         |

## Usage

Run in Terminal

```go
go run .     
```

Each “newanimal” command must be a single line containing three strings. The first string is “newanimal”. The second string is an arbitrary string which will be the name of the new animal. The third string is the type of the new animal, either “cow”, “bird”, or “snake”.  Your program should process each newanimal command by creating the new animal and printing “Created it!” on the screen.

Each “query” command must be a single line containing 3 strings. The first string is “query”. The second string is the name of the animal. The third string is the name of the information requested about the animal, either “eat”, “move”, or “speak”. Your program should process each query command by printing out the requested data.

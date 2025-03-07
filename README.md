# Your First DFA Machine

One of the applications of deterministic finite automata is the implementation of token recognizers in a programming language (known as Lexer in compilers).

In this activity you will have to make a program that receives as input a file with a series of strings, written under certain rules, and will deliver as output table with the states that the automata went through and whether the string is valid or not.

## General Instructions
- It's a 2-members team's work
- The program must be `terminal-based`, written in Golang
- Write your program in the [dfa.go](./dfa.go) file
- There's some code that is already implemented, in order to give the time to focus in the real problem
- Make sure your program compiles and runs

## Problem-specific Instructions

Make a program that receives as input a text file containing strings, and returns yaml-formatted results file with each of the input strings alonside their states and if they're valid or invalid strings. Below some extra notes that need to be considered.

### Input
- A [`dfa.yaml`](./dfa.yaml) file with details of your Deterministic Finite Automata Machine
- A [`strings.txt`](./strings.txt) file with strings to be processed
- A sample [`dfa.yaml`](./dfa.yaml) and a sample [`strings.txt`](./strings.txt) has been provided, but your program should be able to read any different `dfa.yaml` and `strings.txt` input files.

Below an example of a `dfa.yaml` and the meaning of each key word.
```
S: q0
K: [q0, p1, q2]
E: [a, b]
T: 
  - q0, a, q1
  - q0, b, q2
  - q1, a, q1
  - q1, b, q1
  - q2, a, q0
  - q2, b, q2
F: [q1, q2]
```

- `S` = Initial State
- `K` = set of states
- `E` = Alphabet (set of symbols)
- `T` = Transition functions
- `F` = Foinal State

### Output
- A [`results.yaml`](./results.yaml) with the following format. Make sure you follow the keys and order of the strings as you read them from the `strings.txt`file, otherwise, you may loose the points. Below an example with 4 words.

```
- input: aaba
  states:
    - q0
    - q1
    - q1
    - q1
    - q1
  valid: true
- input: ba
  states:
    - q0
    - q2
    - q0
  valid: false
- input: abbb
  states:
    - q0
    - q1
    - q1
    - q1
    - q1
  valid: true
- input: baaaa
  states:
    - q0
    - q2
    - q0
    - q1
    - q1
    - q1
  valid: true
```

### Command Line arguments

The program should run as:
```
go run dfa.go -dfa dfa.yaml -input strings.txt -output results.txt
```

### How to buid/run the first time
You may have some trouble to make it work, since we're using special golang modules. Below the commands to prepare your environment, in case that you have trouble running the `go run` command 

```
go mod init CodersSquad/ct-first-dfa
go mod tidy
```

## Grading Policy

You program will recieve a `dfa.yaml` and a 100 random generated strings `strings.txt` file.
Grading is going to be based in the number of strings that your program correctly validated, if it correcly validated 87 strings, the grade will be 87.

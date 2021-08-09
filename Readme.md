## Battleship

This repo is a simple simulation of the battleship game where input is read from an input file and the result of the game is written to an output file. A sample input and output file are included in this repo for reference.
The game has three phases:
1. Setup phase: The players set their ships according to the coordinates read from the input file.
2. Fire phase: Both the players simulataneously fire missiles according to the HIT coordinates read from the input file. This is achieved by firing asynchronous threads and waiting for them to finish their jobs before writing the output file.
3. Output: The last phase generates an output file as described in the problem statement. A sample file is included which is generated against the sample input being used.

## Setup

Also included is a make file to build binaries and run tests. Please use it as described below:
Run following commands to try it out:

* Clone the project or unbundle it.
* Run `make build` to build and run binaries. A sample built binary can be found included in this package and is titled `main`.

## Manual

```
  λ  make

  build     Build the binaries
  lint      Check for linting errors
  format    Check for formatting errors
  tools     Installs tools
  test      Run tests

```
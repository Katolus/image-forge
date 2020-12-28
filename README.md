# Image manipulation in Go

Work on this project was crucial to understand basic contempts of the `Go` programming languages. Beyond that it provided a fun and entertaining way to manipulate images and work with more-than-usual complex algorithms/matrixes. 

## Big part of this repository was based on the knowledge I received from Todd McLeod and his `Code Clinic: Go` course.

This package was created in order to learn Go programming language. It is not production quality.  
The premise of this repository is build around image analysis, computation and matching.

Main learning received from this project:

- Knowledge about the building blocks and structures in Go
- Working with libraries and documentation
- Forming repository into library like architecture
- Working with go executables on a machine
- Working with go commands and variables
- Concurrency and its benefits
- Type, structs and interfaces
- Testing and documentation from a user perspective

## How to run

- Create a `results` and `images` directories in root of the project.
- Download target images and move the into the `images` directory.
- Update the `run.go` script with the name of the targeted image.
- In the root directory of the package run `go run run.go`.

The results of this should be a gradient, low quality and reverse images of the original image stored in the `results` directory.

## Future work

- Fix all the broken tests cased by a clean up of the project.
- Make the project more user friendly providing things like variables, additional validation statements, etc... .
- Add a command executable that would take a name of a while and type of forge operation.

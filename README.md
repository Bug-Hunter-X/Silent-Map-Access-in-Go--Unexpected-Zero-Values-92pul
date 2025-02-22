# Silent Map Access in Go: Unexpected Zero Values

This repository demonstrates a common, yet subtle, error in Go involving map lookups.  Accessing a non-existent key in a Go map will not raise an error; instead, it returns the zero value for the map's value type. This can lead to unexpected behavior and hard-to-track bugs.

The `bug.go` file contains a simple example that showcases this behavior. The program attempts to access a key that doesn't exist in the map.  Rather than throwing an error or panic, the program simply prints the zero value (0 in this case). 

The `bugSolution.go` file provides a solution using explicit checks to handle missing keys, avoiding silent failures.
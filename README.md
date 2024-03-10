# Golang Notes

- executables must have `main.go` file which must also have `main` func
- otherwise it will be considered helper package and won't generate executable
- define new variables with `:=` as shorthand
- redefine existing variables just with `=`
- We can initialize variables outside of functions, but we cannot assign them values outside of functions
- Go has two data structures to handle lists of data, Array and Slice
- Slices and Arrays must be defined with a data type
- Every element in both an Array and a Slicemust be of the same type
- Array is a fixed length list of things
- Slice is an array that can grow or shrink
- `append()` function can allow you to add new element to end of slice
  - HOWEVER-- rather than modifying the existing slice, it creates a new one slice gets assigned to new variable
- To run and compile multiple packages, run something like `go run main.go deck.go`
- Receiver functions setup methods on variables that we create
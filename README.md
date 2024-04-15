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
- Byte slice is like a string of characters
  - a string that translates ascii characters into a string
  - for example: Hi There! translatest to [72 105 32 116 104 101 114 101 33]
  - translations available at https://www.asciitable.com/
- Type conversion can take one type and make it another type
  - ie `[]byte("Hi there!")` would take the provided string and convert it to a byte slice
- Test functions are written with PascalCase instead of camelCase
  - Always start test func names with `Test`
  - Include all functions being tested in the test name
    - ie-- `TestSaveToDeckAndNewDeckFromFile` would be testing the `saveToDeck` and `newDeckFromFile` functions
- Struct is short for data structure-- collection of properties that are related together
  - example of Card struct field definition would be a struct field has a suit which is a string, and a value which is a string
    - similar to how JS object/Ruby hash works
  - Go is a pass by value language-- "Pass by value" means that when we pass some value into a function, Go will take the struct, copy it, and place it inside the computer's memory
    - So when we call `updateName` function, it isn't updating the new value with `p.firstName = ...`, it's just copying the struct as a new value in a different bit of the memory
    - `&variable` operator-- give me memory address of value this is pointing at
      - can turn value into address with `&value`
    - `*pointer` operator-- give me the value this memory address is pointing at
      - Can turn address into a value with `*value`
      - if we define a pointer receiver, Go will let us call the function with a pointer type or the non-pointer equivalent
    - The above is more for Structs, Slices work differently
  - Slice v Array
    - Slice can grow/shrink, used 99% of the time for lists of elements
      - Slice is made up of data structure and array, so when we access a slice variable we're getting the ds not the array
        - Pointers are used for this type of issue
    - Array is primitive data structure, can't be resized, rarely used directly
      - Arrays can be updated without pointers because it's pointing at the singular memory address
- Value Types vs Reference Types-- value types use pointers to change things in a function, reference types don't worry about pointers
  - Value types: int, float, string, bool, structs
  - Ref Types: slices, maps, channels, pointers, functions
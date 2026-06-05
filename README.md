# hello_go

A learning project for Go fundamentals.

## Concepts Covered

### Modules
- Go uses modules instead of projects. A module is defined by a go.mod file at the root.
- Initialize one with go mod init module-name.
- Code inside a module is organized into packages - each directory is a package.

### Packages and File Organization
- All .go files in the same directory must declare the same package name.
- Files within a package share scope - they can access each others unexported identifiers.
- It is idiomatic to split types into separate files (e.g. dog.go, cat.go), one type per file.

### Structs and Methods (No Classes)
- Go has no classes and no inheritance.
- Use structs to group fields and attach methods via receivers.
- Use composition (embedding) instead of inheritance.

### Interfaces
- Interfaces define methods only, not fields.
- They are satisfied implicitly - no implements keyword.
- To include something like a Name property in an interface, define a Name() method instead.

### Exported vs Unexported
- Capitalized names are exported (public). Lowercase names are unexported (package-private).
- Use constructor functions (e.g. NewDog(name)) to create structs with unexported fields.

### Formatting with gofmt
- gofmt is the built-in canonical code formatter. All Go code uses it.
- Run gofmt -w . to format all files in place.
- Uses tabs for indentation. No configuration, no style debates.
- goimports is a popular extension that also manages import statements.
- Most editors run gofmt on save automatically.

### Doc Comments
- Doc comments use the `//` style (not `/* */` blocks) and appear directly above declarations with no blank line.
- Package-level comments start with "Package <name>" and describe the package's purpose.
- All exported types, functions, and methods should have doc comments.
- Doc comments become part of the GoDoc documentation and appear in IDE intellisense/hover tooltips.
- Comments should be concise and start with the name being documented (e.g., "Dog represents a dog with a name.").

## Project Structure

```r
hello_go/
  go.mod
  main.go
  README.md
  animals/
    animal.go   - Animal interface
    dog.go      - Dog struct and methods
    cat.go      - Cat struct and methods
```

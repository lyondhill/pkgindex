# Package Index

Package Index builds a indexing list with dependencies.

## Design 

Package Index is designed to be simple and easy to understand. It uses a in-memory data storage and runs on a single node with no external dependencies. The data is stored in a `map[string]*pkg` where pkg is:


```
type pkg struct {
	Command       string
	Name          string
	Dependencies  []string
	DependantPkgs []*pkg
}
```

| Field         | Function                                                |
|---------------|---------------------------------------------------------|
| Command       | the command used on this package last                   |
| Name          | the name of this package                                |
| Dependencies  | a list of packages this package depends on              |
| DependantPkgs | a reference to all the packages that depend on this one |


By creating a data structure in this way it makes the primary functions very quick. When inserting data we need to ensure the dependencies of this package are all available, this is done quickly by indexing the map. When removing a package we need to guarantee the package we are removing does not have anything depending upon it, and this is done quickly by having the Dependant packages referenced.


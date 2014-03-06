# Go Container Kit

## Background

The _Go Programming Language_ currently does not come with a rich, consistent library for representing collections of values. Instead, the language itself provides relatively powerful native array and map abstractions. Those are complemented with a small number of ad-hoc collection implementations in the `container` package. The data types in this package are inconsistent and lack functionality. Furthermore, they are not separating their interface from their implementation, compromising flexibility, extensibility, and modularity in client code. Go programs typically combine the use of these implementations with the built-in collection abstractions. This leads to even more inconsistencies and asymmetries in the code.

## Project goals

The aim of this project is to develop a comprehensive container class library for the Go programming language that allows programmers to make use of a broad range of container abstractions in a consistent fashion. The library design focuses on the following properties:

1. **Simplicity:** The library is built on top of a small number of core concepts, to keep the learning curve for users as flat as possible.
2. **Consistency:** The concepts are reused uniformly across the various container abstractions. This includes consistent usage of names, types, and method signatures.
3. **Expressiveness:** Rich abstractions allow users to express processing logic in a concise, declarative fashion.
The simplicity and consistency properties are important because they guarantee ease of use. Expressiveness is important because it promotes concise, readable client logic.

The implementation of _Container Kit_ is quite unconventional in the sense that it is based on a number of design patterns, introducing features from class-based object-oriented languages. This approach allows for composing Go containers out of small reusable building blocks. Ultimately, it is this methodology which guarantees the consistency and expressiveness of the properties explained above.

## Client code example

With _Container Kit_, building a map from day indices (1..7) to a day name can be done with a single line of code:

```go
dayMap := HashMap.From(Enum.Range(1,7).Zip(dayNames)).ReadOnly()
```

where `dayNames` is defined as follows:

```go
dayNames := Enum.New("Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun")
```

`Enum.New` returns an enumeration, which is the most simple container abstraction: an ordered, immutable set of elements. The `dayNames` enumeration gets zipped with the enumeration of integers ranging from 1 to 7. The result is a container of pairs (index, name) which are then used to initialize a mutable hash table that implements the mapping from indices to names. Finally, `ReadOnly()` gets invoked; this method will return a read-only proxy for the hash map to prevent further modifications.

`ReadOnly()` creates a _container wrapper_. The wrapper does not implement operations for changing the state of the underlying container, but if there are references to the underlying container, changes can still be made to the underlying container and they are reflected directly in the read-only wrapper. Containers like the read-only map wrapper, which simply create a new view for an existing container, are called _dependent containers_.

If mutations should be prevented right from the beginning, an immutable map needs to be created. _Container Kit_ introduces the concept of _functors_ to implement new features for container classes in a modular fashion. The `ImmutableMap` functor turns a class for mutable containers into a new class for _immutable containers_. Here is the code for creating an immutable hash map:

```go
immutableDayMap := ImmutableMap(HashMap).From(Enum.Range(1,7).Zip(dayNames))
```

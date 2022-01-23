# use
use is a go utility library using go1.18 generics  
created by halpdesk 2022-01-22 

## use/slice

**Map** updates a slice by applying a function to all members of a slice. The apply function must use and return the same type as the elements of the slice.

```go
    numbers := []int{2,3,5}
    slice.Map(numbers, func(int a) int {return a*a})
    // numbers is now []int{4,9,25}
```

**Walk** is used to apply a function to all elements of a slice and return a different kind of slice. This is useful for when dealing with collections, for example in combination to convert from one data presentation to another (i.e. database model to protobuf).

```go
    ConvertToPB := func(customer Customer) pb.Customer {
        return pb.Customer{Id:customer.ID}
    }
    customers := []*Customer{{ID:1},{ID:2},{ID:3}}
    PBCustomers := slice.Walk(customers, ConvertToPB) 
    // PBCustomers is of type []*pb.Customer
```

**Filter** is used to filter a slice by a boolean filtering function.

```go
    numbers := []Int{1, 2, 3, 4, 5}
    oddNumbers := slice.Filter(numbers, func(n int) bool { return n%2 == 1 })
    // oddNumbers is now []int{1,3,5}
```

**IndexOf** is used to determine the index in a slice of a given value.

```go
   strs := []string{"foo", "bar", "baz"}
   index := slice.IndexOf(strs, "baz")
   // index is now 2
```

## use/kind

**Patch** is used to merge any two structs of the same type. The second structs values overwrites the first ones, if it is not the zero value for its type.

```go
    type Profile struct {
        Name  string
        Label string
        Data  []byte
    }
    baseProfile := Profile{
        Name: "base"
        Label: "unset"
    }
    myNewGeneratedProfile = kind.Patch(baseProfile, Profile{
        Label: "set",
        Data: faker.DataGenerator(),
    })
    // myNewGeneratedProfile is now a mix of both baseProfile and the faker generater dataset (Label has been overwritten and is "set")
``` 
    

**IsZero** is used to determine if a value is the type's zero value.

```go
    zeroTime := time.Time{}
    if kind.IsZero(zeroTime) {
        zeroTime := time.Now()
    }
```

## update and test with 1.18
```sh
go install golang.org/dl/go1.18beta1@latest 
go1.18beta1 download
curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.43.0
```

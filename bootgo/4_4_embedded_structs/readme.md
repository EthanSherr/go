https://boot.dev/course/3b39d0f6-f944-4f1b-832d-a1daba32eda4/59d90390-f479-4472-bb16-9af599285229/e30f8be3-30e4-4528-ab93-54461538314e

EMBEDDED STRUCTS
Go is not an object-oriented language. However, embedded structs provide a kind of data-only inheritance that can be useful at times. Keep in mind, Go doesn't support classes or inheritance in the complete sense, embedded structs are just a way to elevate and share fields between struct definitions.
```go
type car struct {
  make string
  model string
}

type truck struct {
  // "car" is embedded, so the definition of a
  // "truck" now also additionally contains all
  // of the fields of the car struct
  car
  bedSize int
}
```
EMBEDDED VS NESTED
An embedded struct's fields are accessed at the top level, unlike nested structs.
Promoted fields can be accessed like normal fields except that they can't be used in composite literals
```go
lanesTruck := truck{
  bedSize: 10,
  car: car{
    make: "toyota",
    model: "camry",
  },
}

fmt.Println(lanesTruck.bedSize)

// embedded fields promoted to the top-level
// instead of lanesTruck.car.make
fmt.Println(lanesTruck.make)
fmt.Println(lanesTruck.model)
```
ASSIGNMENT
At Textio, a "user" struct represents an account holder, and a "sender" is just a "user" with some "sender" specific data. A "sender" is a user that has a rateLimit field that tells us how many messages they are allowed to send.

Fix the system by using an embedded struct as expected by the test code.
https://boot.dev/course/3b39d0f6-f944-4f1b-832d-a1daba32eda4/7faa7ccf-4fc0-406d-90b7-dac16e72f6c3/7804ef9a-1a03-42a0-94d1-42ca09b89e43

TYPE ASSERTIONS IN GO
When working with interfaces in Go, every once-in-awhile you'll need access to the underlying type of an interface value. You can cast an interface to its underlying type using a type assertion.

```go
type shape interface {
	area() float64
}

type circle struct {
	radius float64
}

// "c" is a new circle cast from "s"
// which is an instance of a shape.
// "ok" is a bool that is true if s was a circle
// or false if s isn't a circle
c, ok := s.(circle)
if !ok {
	// s wasn't a circle
	log.Fatal("s is not a circle")
}

radius := c.radius
```
ASSIGNMENT
Implement the getExpenseReport function.

If the expense is an email then it should return the email's toAddress and the cost of the email.
If the expense is an sms then it should return the sms's toPhoneNumber and its cost.
If the expense has any other underlying type, just return an empty string and 0.0 for the cost.
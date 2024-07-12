# First-class functions
**Fuctions are first-class citizens**

Functions can be bound to variable names:
```go
adder := func(a int, b int) int {
	return a + b
}
```

Functions can be passed to other functions, or even served as the return value of a function.

> Functions are treated as any other *type* would be

## Pure functions
Subset of FP, where each function has to be pure. Pure FP should not produce any side effects -> program becomes entirely deterministic.

## Declarative functions
Say what you want to achieve rather than how to acheive it:

```go
func DeclarativeFunction() int {
	return IntRange(-10, 10).
		Abs().
		Filter(func(i int64) bool {
			return i % 2 == 0
		}.
		Sum()
}
```

The function is verbose and does the following:
1. Give a range of integers between `-10` and `10`
2. Turn these numbers into their absolute values
3. Filter for all even numbers
4. Give sum of these even numbers

If the function was done *imperatively*:
```go
func ImperativeFunc() int {
	sum := 0
	for i := -10; i <= 10; i++ {
		absolute := int(math.Abs(float64(i)))
		if absolute % 2 == 0 {
			sum += absolute
		}
	}
	return sum
}
```

---

Pure, immutable code makes code easier to test for the following reasons:
- The state of the system has no impact on our function – so we don’t have to mock the state
when testing.
- A given function always returns the same output for a given input. This means we get predictable,
deterministic functions.

---

# Type Aliasing
Start with a `Person` struct and a setter function for `phoneNumber`
```go
type Person struct {
	name string
	phoneNumber string
}

func (p *Person) setPhoneNumber(newPhoneNumber string) {
	p.phoneNumber = newPhoneNumber
}
```
Create a *type alias* for `phoneNumber`
```go
type phoneNumber string
type Person struct {
	name string
	phoneNumber phoneNumber
}

func (p *Person) setPhoneNumber(newPhoneNumber phoneNumber) {
	p.phoneNumber = newPhoneNumber
}
```

This lowers the risk of
errors by passing invalid data to a function.

An additional benefit, depending on the IDE, is that your
IDE will also show you the signature. If you had a large function that takes five different types of string,
your IDE might just show you function expects input (`string, string, string, string, string`), with no clear
order in which arguments need to be passed. If each string is a distinct type, this might become `name,
phonenumber, email, street, country`.

---

With primitive types in Go, there cannot be user defined functions attached to them. With type alias, we can attach functions to `phoneNumber` type:

```go
func (p phoneNumber) valid() bool {
	// Phone number validator
}
```

## Type aliases for functions
```go
func filter(is []int, predicate func(int) bool) []int {
	out := []int{}
	for _, i := range is {
		 if predicate(i) {
			 out = append(out, i)
		 }
	}
	return out
}
```

With type aliases for functions:
```go
type predicate func(int) bool
func filter(is []int, p predicate) []int {
	out := []int{}
	for _, i := range is {
		 if p(i) {
			 out = append(out, i)
		 }
	}
	return out
}
```

When we define a function:
```go
func largerThanTwo(i int) {
	return i > 2
}
```

The `largerThanTwo` fucntion can be passed as a predicate without specifying anywhere that this function adheres to `predicate` type. Let the compiler to the job.

## Returning functions
```go
func createLargerThanPredicate(threshold int) predicate {
	return func(i int) bool {
		return i > threshold
	}
}
```

This is important for *continuation-passing style* programming and function currying

## Functions inside data structures
Store functions inside a list:
```go
var (
	largerThanTwo = createLargerThanPredicate(2)
	largerThanFive = createLargerThanPredicate(5)
)

predicates := []predicate{largerThanTwo, largerThanFive}
```

Dispatcher pattern stores functions inside a map:
```go
dispatcher := map[string]predicate{
	"2": largerThanTwo,
	"5": largerThanFive,
}
```

Functions cna also be in `structs`:
```go
type Validator struct {
	largerThan predicate
	smallerThan predicate
}

func (v Validator) validate(input int) bool {
	v.largerThan(input) && v.smallerThan(input)
}

AdultAgeValidator := Validator{
	largerThan: createLargerThanPredicate(18),
	smallerThan: func(i int) bool { return i < 150 },
}
```

# Create a simple Calculator

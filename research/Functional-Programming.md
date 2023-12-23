# Functional Programming in Go
## Introduction
In FP, functions are considered **first-class citizes**, they can be: 
* bounds to variable names
* pass to others functions
* served as the return of a function  

First example:

Using type alias to define a new type:
```go
type generalFunc func (int) bool
```

This tells us everywhere in our code base where we find the generalFunc type, it expect to see a func take an int and return a bool.

```go
func filter (nums []int, condition generalFunc) []int{
    out := []int{}
    for _, num := range nums {
        if condition(num) {
            out = append(out,num)
        }
    }
    return out
}

func smallerThan10 (i int) bool{
    return i > 10
}
```

Then we can use it in our code:
```go
func main(){
    nums := []int {1,1,2,3,5,8,13}
    smaller := filter(nums,smallerThan10)
    fmt.Printf("%v",smaller)
}
```

Output: 
```go
type generalFunc func (int) bool
```

[Code](playground/FP/smaller.go) 

### Pure vs Impure

We have a sruct of the Person type, with a Name field. We want to create a function to change the name of person, such as changeName.

```go
type Person struct {
    Age int
    Name string
}
```

**Impure Function:**
```go
func changeName (p *Person, newName string){
    p.Name = newName
}

// Version 2
func (p *Person) changeName (newName string){
    p.Name = newName
}
```

The `Person` object that was passed to the function is mutated. The state of our system is now different before the function was called. Every place that refers to that `Person` object will now see the new name instead of the old name

**Pure Function:**
```go
func changeNamePure (p Person, newName string) Person {
    return Person{
        Age: p.Age,
        Name: newName,
    }
}
```

Impure: less memory, less effort, but maintaining a system where functions can change the state of the system are vast. ? In larger apps, maintaining a clear understanding of state easier to debug and replicate errors.

### Say what you want, not how you want

In FP, declarative > imperative, what you want to achieve > how to achieve

```go
func DeclarativeFunction ()int {
    return IntRange (-10,10).
        Abs().
        Filter(func(i int64) bool){
            return i%2 == 0
        }).
        Sum()
}

// result = 60

func iterativeFunction() int{
    sum := 0 
    for i= -10; i <= 10; i++{
        abs := int(math.Abs(float64(i)))
        if abs %2 == 0 {
            sum += abs
        }
    }
    return sum
}
```

Features that enable us to write functional Go code:
* Functions as first-class citizens
* Higher-orders functions
* Immutability guarantees
* Generics 
* Recusion

Features that Go lacks:
* Tail-call optimization
* Lazy evaluation
* Purity guarantee

### Why functional programming?

Benefits:
* More readable code
* Easier to understand and debug code
* Easier testing
* Fewer bugs
* Easier concurrency

Immutable code makes code easier to tests:
* The state of the systems has no impact on our function, so we don't have to mock the state when testing
* A given function always return the same output for a given output. This means we get predictable, deterministic functions.

Fewer bugs:
* Fewer edge cases to think about
* If the state is important, you have to know, at each points in time, state of the system.

Concurrent code easier:
* Running the same conccurently can never impact the result of another running functions.

### Why not functional programming in Go
* Style of code base
* Performance

## 2. Treating Functions as First-Class Citizens

### Type aliases

**Heavily influenced Java example**

```go
type Person struct {
    name string
    phonenumber string
}

func (p *Person) setPhoneNumber (s string){
    p.phonenumber = s
}
```

Using type alias:

```go
type phoneNumber string
type Person struct {
    name string
    phonenumber phoneNumber
}

func (p *Person) setPhoneNumber (s phoneNumber){
    p,phonenumber = s
}
```

So we have a name, which is just a string, and phonenumber, which is a phoneNumber type, which is equal to a string. So where does the benefit come from?
* Communicating intent
* Error messages
* IDE

```go
func (p *Person) update(name, phonenumber string){
    p.name = name
    p.phonenumer = phonenumber
}
```

When we try to complie our code:
```
cannot use phonenumber (variable of type   string) as type phoneNumber in assignment
```

In this simple example, It doesn't do much. But as code base expands, this ensures that all developers are thinking about the type that should be passed into a function.


IDE: show `(name,phonenumber,email,street,country)`instead of `(string,string,string,string,string)`


### Attach functions to type alias, which we cannot do with primitive types.

```go
type age uint
type phoneNumber string

type Person struct {
	name        string
	age         age
	phonenumber phoneNumber
}

```

```go
func (a age) valid() bool {
    return a < 120
}

func isValidPerson (p Person) bool {
    return p.age.valid() && p.name != ""
}
```

We can not do something like this with `age int`. Errors for both `valid` func  and `p.age.valid()`


### Type aliases for functions

```go
type generalFunc func(int) bool

func filter (nums []int, g generalFunc)[]int {
    out := []int{}
    for _, num := range nums {
        if g(num) {
            out = append (out,g)
        }
    }
    return out
}
```


### Passing functions to functions 

```go
type generalFunc func (int) bool

func largerThanTwo (i int)bool {
    return i > 2
}

func filter(nums []int, g generalFunc)[]int {
    out := []int{}
    for _, num := range nums {
        if g(num){
            out = append(out,num)
        }
    }
    return out
}

func main(){
    ints := []int{1,2,3}
    filter(ints,largerThanTwo)
}
```

### In-line function definitions

```go
func main(){
    // functions in variables

    inlinePersonStruct := struct {
        name string
    }{
        name: "John",
    }

    ints := []int{1,2,3}
    inlineFunction := func (i int)bool {return i > 2}
    filter(ints, inlineFunction)
}
```





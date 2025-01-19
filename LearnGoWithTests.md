## Writing tests

- Create file named like xxx_test.go
- Create test func starts with 'Test'
- The test function takes one argument only `t *testing.T`
- To use the `*testing.T` type, you need to `import "testing"`, like we did with `fmt` in the other file

```go
func greetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

```

A few new concepts:

- In our function signature we have made a *named return value* `(prefix string)`.
- This will create a variable called `prefix` in your function.
  - It will be assigned the "zero" value. This depends on the type, for example `int`s are 0 and for `string`s it is `""`.
    - You can return whatever it's set to by just calling `return` rather than `return prefix`.
  - This will display in the Go Doc for your function so it can make the intent of your code clearer.
- `default` in the switch case will be branched to if none of the other `case` statements match.
- The function name starts with a lowercase letter. In Go, public functions start with a capital letter, and private ones start with a lowercase letter. We don't want the internals of our algorithm exposed to the world, so we made this function private.
- Also, we can group constants in a block instead of declaring them on their own line. For readability, it's a good idea to use a line between sets of related constants.



# Slice

- Variadic Functions （可变函数参数 )

```go
func SumAll(numbersToSum ...[]int) (sums []int) {}
```

- sums = make([]int, lengthOfNumbers)    We can use `make` to create specific length of slice

- `slice`  has fixed capacity. We can use `append`  which takes a slice and a **new** value, and return a new slice with all the items in it
- We can use `slice[low:high]` to get partial slice

# struct, methods and interface

- a method is a function with a **receiver** 

   A method declaration binds an identifier, the method name, to a method, and associates the method with the receiver's base type.

  Methods are very similar to functions but they are called by invoking them on an instance of a particular type. Where you can just call functions wherever you like, such as `Area(rectangle)` you can only call methods on "things".

- In Go **interface resolution is implicit**. If the type you pass in matches what the interface is asking for, it will compile.



# pointers

- The arguments are **copied** when you call a function or a method.

For example, when calling `func (w Wallet) Deposit(amount int)` the `w` is a copy of whatever we called the method from.

- Go lets you create new types from existing ones.  The syntax is `type MyName OriginalType`

- A function which takes arguments or returns values that are interfaces, they can be `nil`

- **When a function returns a pointer to something, you need to make sure you check if it's nil or you might raise a runtime exception - the compiler won't help you here.**

  

# map

- We can modify maps without passing as an address to it(e.g. &mymap). It feels like reference type. To be specific, when we pass a map to a function/method, we are indeed copying it, but just **the pointer part**, not the underlying data structure that contains the data (也就是说map传递的是引用的拷贝)

- A gotcha with maps is that they can be a `nil` value. Therefore, you should never initialize a nil map variable

```go
// Never do this!!!
var m map[string]string

//  initialize an empty map or use the make keyword to create a map. It ensures that you will never get a runtime panic
var dictionary = map[string]string{}
// OR
var dictionary = make(map[string]string)
```

- Delete a key in the map

```go
delete(map, key)
```


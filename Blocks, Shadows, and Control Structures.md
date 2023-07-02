# Blocks, Shadows, and Control Structures
Each place where a declaration occurs is called a block. Variables, constants, types,
and functions declared outside of any functions are placed in the package block.
All of the variables defined at the top level of a function
(including the parameters to a function) are in a block. Within a function, every set
of braces ( {} ) defines another block, and  control structures in Go define blocks of their own.

# Shadowing Variables
A shadowing variable is a variable that has the same name as a variable in a contain‐
ing block. For as long as the shadowing variable exists, you cannot access a shadowed
variable.
```
func main() {
    x := 10
    if x > 5 {
        fmt.Println(x)
        x := 5
        fmt.Println(x)
    }
    fmt.Println(x)
}
```
Here’s what happens:

10

5

10

### Detecting Shadowed Variables
you can add shadowing detection to your build process by installing the shadow linter on your machine:
```$ go install golang.org/x/tools/go/analysis/passes/shadow/cmd/shadow@latest```

### Universal block
Go is a small language with only 25 keywords. What’s interesting is that the built-in
types (like int and string ), constants (like true and false ), and functions (like make
or close ) aren’t included in that list. Neither is nil . So, where are they?
Rather than make them keywords, Go considers these predeclared identifiers and
defines them in the universe block, which is the block that contains all other blocks.
Because these names are declared in the universe block, it means that they can be
shadowed in other scopes

# if 
The most visible difference between if statements in Go and other languages is that
you don’t put parenthesis around the condition.Go adds is the ability to declare vari‐
ables that are scoped to the condition and to both the if and else blocks.
```
if n := rand.Intn(10); n == 0 {
        fmt.Println("That's too low")
    } else if n > 5 {
        fmt.Println("That's too big:", n)
    } else {
        fmt.Println("That's a good number:", n)
    }
```

# for
for is the only looping keyword in the lan‐
guage. Go accomplishes this by using the for keyword in four different formats:

for is the only looping keyword in the language. Go accomplishes this by using the for keyword in four different formats:

- A complete, C-style for
```
for i := 0; i < 10; i++ {
    fmt.Println(i)
}
```
- The Condition-Only for Statement
```
i := 1
for i < 100 {
    fmt.Println(i)
    i = i * 2
}
```
- The Infinite for Statement
```
for {
    fmt.Println("Hello")
}
```
- The for-range Statement

What makes a for-range loop interesting is that you get two loop variables. The first
variable is the position in the data structure being iterated, while the second is the
value at that position.If you don’t need to access the key, use an underscore ( _ ) as
the variable’s name. This tells Go to ignore the value. *A for-range loop is the best way to walk through a string, since it properly
gives you back runes instead of bytes.*
```
evenVals := []int{2, 4, 6, 8, 10, 12}
for _, v := range evenVals {
    fmt.Println(v)
}
```
Go include a random number that’s generated every time a map variable is created. Next, they made the
order of a for-range iteration over a map vary a bit each time the map is looped over.
These two changes make it far harder to implement a Hash DoS attack.




### The for-range value is a copy
Modifying the value variable will not modify the value in the compound type
```
evenVals := []int{2, 4, 6, 8, 10, 12}
for _, v := range evenVals {
    v *= 2
}
fmt.Println(evenVals)
```
u will see that evenVals hasn't changed.

# break and continue
you can use break statement to exit the loop immediately.
Go also includes the continue keyword, which skips over the rest of the body of a for
loop and proceeds directly to the next iteration

```
for i := 0; i <= 10; i++ {
	if i == 3 {
		continue
	}
	if i == 7 {
		break
	}
	fmt.Println(i)
}
```





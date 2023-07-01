# Built-in Types
> ## The Zero Value
> Go, like most modern languages, assigns a default zero value to any variable that is declared but not assigned a value
> ## Literals 
> A literal in Go refers to writing out a number, character, or string. There are four common kinds of literals that you’ll find in Go programs 
>> **Integer literals**<br>
>>Integer literals are sequences of numbers; they are normally base ten, but different prefixes are used to indicate other bases: 0b for binary (base two), 0o for octal (base eight), or 0x for hexadecimal (base sixteen). You can use either or upper- or lowercase letters for the prefix. A leading 0 with no letter after it is another way to represent an octal literals.To make it easier to read longer integer literals, Go allows you to put underscores in the middle of your literal. This allows you to, for example, group by thousands in base ten (1_234). These underscores have no effect on the value of the number. The only limitations on underscores are that they can’t be at the beginning or end of numbers, and you can’t have them next to each other.
>><br>*Go lets you use an integer literal in floating point expressions or even assign an integer literal to a floating point variable. This is because **literals in Go are untyped**; they can interact with any variable that’s compatible with the literal*
> 
>> **Floating point literals**<br>
>>Floating point literals have decimal points to indicate the fractional portion of the value. They can also have an exponent specified with the letter e and a positive or negative number (such as 6.03e23) 
>
>> **Rune literals**<br>
>> Rune literals represent characters and are surrounded by single quotes. Unlike many other languages, in Go single quotes and double quotes are not interchangeable. Rune literals can be written as single Unicode characters ( 'a' ), 8-bit octal numbers( '\141' ), 8-bit hexadecimal numbers ( '\x61' ), 16-bit hexadecimal numbers( '\u0061' ), or 32-bit Unicode numbers ( '\U00000061' ). There are also several back‐ slash escaped rune literals, with the most useful ones being newline ( '\n' ), tab( '\t' ), single quote ( '\'' ), double quote ( '\"' ), and backslash ( '\\' )
> ## Booleans
> The bool type represents Boolean variables. Variables of bool type can have one of two values: true or false . *The zero value for a bool is false*
> ## Numeric Types
> Go has a large number of numeric types: 12 different types (and a few special names) that are grouped into three categories
>> **Integer types**<br>
>> Go provides both signed and unsigned integers in a variety of sizes, from one to four bytes.*The zero value for all of the integer type is 0*.
>> - int8 
>> - int16 
>> - int32 
>> - int64 
>> - uint8 
>> - uint16 
>> - uint32 
>> - uint64
>>
>>**The special integer types**<br>
>> Go does have some special names for integer types. A byte is an alias for uint8 ; it is legal to assign, compare, or perform mathematical operations between a byte and a uint8
>> - A byte is an alias for uint8 ; it is legal to assign, compare, or perform mathematical operations between a byte and a uint8
>> - he second special name is int . On a 32-bit CPU, int is a 32-bit signed integer like an int32 . On most 64-bit CPUs, int is a 64-bit signed integer, just like an int64 . Because int isn’t consistent from platform to platform, it is a compile-time error to assign, compare, or perform mathematical operations between an int and an int32 or int64 without a type conversion. Integer literals default to being of int type.
>> - The third special name is uint . It follows the same rules as int , only it is unsigned (the values are always 0 or positive)
>> - rune is an alias for int32 type, just like byte is an alias for uint8.a rune literal’s default type is a rune. *If you are referring to a character, use the rune type, not the int32 type. They might be the same to the compiler, but you want to use the type that clarifies the intent of your code*
>> - uintptr
>>
>> <br>**Floating point types** 
>> There are two floating point types in Go
>> - float32
>> - float64
>> 
>> Picking which floating point type to use is straightforward: unless you have to be compatible with an existing format, use float64 . **Floating point literals have a default type of float64** , so always using float64 is the simplest option. It also helps mitigate floating point accuracy issues since a float32 only has six- or seven-decimal digits of precision.Dividing a nonzero floating point variable by 0 returns +Inf or -Inf (positive or negative infinity), depending on the sign of the number. Dividing a floating point variable set to 0 by 0 returns NaN (Not a Number).
>> *While Go lets you use == and != to compare floats, don’t do it. Due to the inexact nature of floats, two floating point values might not be equal when you think they should be. Instead, define a maximum allowed variance and see if the difference between two floats is less than that* 
>><br>**Complex types**
>>Go defines two complex number types. complex64 uses float32 values to represent the real and imaginary part, and complex128 uses float64 values.
>><br>**String and Rune**<br>
>>Strings in Go are immutable; you can reassign the value of a string variable, but you cannot change the value of the string that is assigned to it.The zero value for a string is the empty string.a rune literal’s default type is a rune, and a string literal’s default type is a string.If you are referring to a character, use the rune type

### **Explicit Type Conversion**</br>
Go doesn’t allow automatic type promotion between variables. You must use a type conversion when variable types do not match.Even different-sized integers and floats must be converted to the same type to interact.Go doesn’t allow truthiness. In fact, no other type can be converted to a bool, implicitly or explicitly. If you want to convert from another data type to boolean, you must use one of the comparison operators( == , != , > , < , <= , or >= )
### **var Versus :=**<br>
Go uses the var keyword, an explicit type, and an assignment. It looks like this:
<br>var x int = 10<br>
If the type on the righthand side of the = is the expected type of your variable, you can leave off the type from the left side of the = . Since the default type of an integer literal is int , the following declares x to be a variable of type int :
<br>var x = 10<br>
Conversely, if you want to declare a variable and assign it the zero value, you can keep the type and drop the = on the righthand side:
<br>var x int<br>
You can declare multiple variables at once with var , and they can be of the same type:
<br>var x, y int = 10, 20<br>
all zero values of the same type:
<br>var x, y int<br>
or of different types:
<br>var x, y = 10, "hello"<br>
There’s one more way to use var . If you are declaring multiple variables at once, you can wrap them in a declaration list:
<br>var (
<br>    x   int
<br>    y       =  20
<br>    z   int = 30
<br>    d, e    = 40, "hello"
<br>    f, g string
<br>)<br>
Go also supports a short declaration format. When you are within a function, you can use the := operator to replace a var declaration that uses type inference. The following two statements do exactly the same thing: they declare x to be an int with the value of 10:
<br>var x = 10
<br>x := 10<br>
*The := operator can do one trick that you cannot do with var : it allows you to assign values to existing variables, too. As long as there is one new variable on the lefthand side of the := , then any of the other variables can already exist<br>*
<br>x := 10
<br>x, y := 30, "hello"
*There is one limitation on := . If you are declaring a variable at package level, you must use var because := is not legal outside of functions*
### **Using const**
const is a way to ensure that a value is immutable.<br>
const x int64 = 10<br>
const in Go is very limited. Constants in Go are a way to give names to literals. They can only hold values that the compiler can figure out at compile time.

Constants can be typed or untyped. An untyped constant works exactly like a literal;
it has no type of its own, but does have a default type that is used when no other type
can be inferred. A typed constant can only be directly assigned to a variable of that
type

Here’s what an untyped constant declaration looks like:

const x = 10

*Whether or not to make a constant typed depends on why the constant was declared.
If you are giving a name to a mathematical constant that could be used with multiple
numeric types, then keep the constant untyped. In general, leaving a constant unty‐
ped gives you more flexibility.*

### Naming Variables and Constants
Go uses camel case (names like indexCounter or numberTries)when an identifier name consists of multiple words. Go uses the case of the first letter in the
name of a package-level declaration to determine if the item is accessible outside the package.






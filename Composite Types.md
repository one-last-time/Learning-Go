# Array
All of the elements in the array must be of the type that’s specified (this doesn’t mean they are always of the same type).
The following ways are to declare array
<br>var x [3]int
<br>var x = [3]int{10, 20, 30}
<br>var x = [12]int{1, 5: 4, 6, 10: 100, 15} *[1, 0, 0, 0, 0, 4, 6, 0, 0, 0,
100, 15]*
<br>var x = [...]int{10, 20, 30}<br>
*Go considers the size of the array to be part of the type of
the array. This makes an array that’s declared to be [3]int a different type from an
array that’s declared to be [4]int . This also means that you cannot use a variable to
specify the size of an array, because types must be resolved at compile time, not at
runtime*.***Don’t use arrays unless you know the exact length you need
ahead of time***

# Slices
When you want a data structure that holds a sequence of values, a
slice is what you should use. What makes slices so useful is that the length is not part
of the type for a slice. This removes the limitations of arrays.

var x = []int{10, 20, 30}

var x = []int{1, 5: 4, 6, 10: 100, 15}

var x []int

var x []int creates a slice of int s. Since no value is assigned, x is assigned the zero value for
a slice, which is ***nil***. ***nil*** has no type, so it can
be assigned or compared against values of different types. A nil slice contains
nothing.

A slice isn’t comparable. It is a compile-time error to
use == to see if two slices are identical or != to see if they are different. The only thing
you can compare a slice with is ***nil***.

> ## append 
> append function is used to grow slices:
> 
>var x []int
> 
>x = append(x, 10)
> ## capacity
> Each element in a slice is assigned to
consecutive memory locations, which makes it quick to read or write these values.
Every slice has a capacity, which is the number of consecutive memory locations
reserved. This can be larger than the length. Each time you append to a slice, one or
more values is added to the end of the slice. Each value added increases the length by
one. **When the length reaches the capacity, there’s no more room to put values. If you
try to add additional values when the length equals the capacity, the append function
uses the Go runtime to allocate a new slice with a larger capacity. The values in the
original slice are copied to the new slice, the new values are added to the end, and the
new slice is returned**.*The rules as of Go 1.14 are to
double the size of the slice when the capacity is less than 1,024 and then grow by at
least 25% afterward*. 
> 
> ## make
> make lets u declare slice with size and capacity
> 
> x := make([]int, 5)
> 
>  ## Slicing slice
> A slice expression creates a slice from a slice. It’s written inside brackets and consists of
a starting offset and an ending offset, separated by a colon (:)
> 
> x := []int{1, 2, 3, 4}
> 
>y := x[:2]
> 
>z := x[1:]
> 
> **Slices share storage sometimes.Whenever you take a slice from another slice, the subslice’s capacity
is set to the capacity of the original slice, minus the offset of the subslice within the
original slice. This means that any unused capacity in the original slice is also shared
with any subslices**.
> 
> To avoid complicated slice situations, you should either never use append with a sub‐
slice or make sure that append doesn’t cause an overwrite by using a full slice expres‐
sion.The full slice expression includes a third part, which
indicates the last position in the parent slice’s capacity that’s available for the subslice
> y := x[:2:2]
> 
> z := x[2:4:4]
> 
> ## copy
> If you need to create a slice that’s independent of the original, use the built-in
copy function.
> 
> The copy function takes two parameters. The first is the destination slice and the sec‐
ond is the source slice. It copies as many values as it can from source to destination,
limited by whichever slice is smaller, and returns the number of elements copied. The
capacity of x and y doesn’t matter; it’s the length that’s important.
> 
> x := []int{1, 2, 3, 4}
> 
> y := make([]int, 4)
> 
> num := copy(y, x)

## Strings and Runes and Bytes
You might think that a string in Go is made out of runes, but that’s not the case. Under the
covers, Go uses a sequence of bytes to represent a string.

## Maps
The map type is written as map[keyType]valueType.

ways to declare map:

var nilMap map[string]int

totalWins := map[string]int{}

ages := make(map[int][]string, 10)

The zero value for a map is nil . A nil map has a length of 0. Attempting to read a nil map
always returns the zero value for the map’s value type. However, attempting to write
to a nil map variable causes a panic.We can use a := declaration to create a map variable by assigning it a map literal

Maps are not comparable. You can check if they are equal to nil , but you cannot
check if two maps have identical keys and values using == or differ using !=

The key for a map can be any comparable type. This means you cannot use a slice or
a map as the key for a map.

## Struct
When you have related data that you want to group together, you should define a struct.
A struct type is defined with the keyword type , the name of the struct type, the keyword struct , 
and a pair of braces ({}). Within the braces, you list the fields in the struct.

type person struct {

name string

age int

pet string

}

Once a struct type is declared, we can define variables of that type:

var fred person

Here we are using a var declaration. Since no value is assigned to fred , it gets the
zero value for the person struct type. A zero value struct has every field set to the
field’s zero value.

there are other ways to declare:
```
bob := person{}

julia := person{
    "Julia",
    40,
    "cat",
}
beth := person{
    age: 30,
    name: "Beth",
}
```

You can also declare that a variable implements a struct type without first giving the
struct type a name. This is called an anonymous struct:

```
pet := struct {
    name string
    kind string
    }{
        name: "Fido",
        kind: "dog",
    }
```

Structs that are entirely composed of comparable types are comparable.Go doesn’t allow comparisons between variables that represent structs of different
types..Go does allow you to perform a type conversion from one struct type to
another if the fields of both structs have the same names, order, and types.




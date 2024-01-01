1. article on slieces and array: https://www.openmymind.net/The-Minimum-You-Need-To-Know-About-Arrays-And-Slices-In-Go/
2. question: https://stackoverflow.com/questions/21719769/passing-an-array-as-an-argument-in-golang

<details><summary>Read here more on slices</summary>

```
package main

import "fmt"

type name struct {
    X string
}

func main() {
    var a [3]name
    a[0] = name{"Abbed"}
    a[1] = name{"Ahmad"}
    a[2] = name{"Ghassan"}

    nameReader(a)
} 

func nameReader(array []name) {
    for i := 0; i < len(array); i++ {
        fmt.Println(array[i].X)
    }
}
```

> error: cannot use a (type [3]name) as type []name in function argument

## Answers: 
### ans1: 

You have defined your function to accept a slice as an argument, while you're trying to pass an array in the call to that function. There are two ways you could address this:
1. Create a slice out of the array when calling the function. Changing the call like this should be enough:


    nameReader(a[:])
<br>
    2. Alter the function signature to take an array instead of a slice. For instance:

    func nameReader(array [3]name) {
        ...
    }

<br>
    Downsides of this solution are that the function can now only accept an array of length 3, and a copy of the array will be made when calling it.
You can find a more details on arrays and slices, and common pitfalls when using them [here](https://www.openmymind.net/The-Minimum-You-Need-To-Know-About-Arrays-And-Slices-In-Go/)

### ans2

<details>

Since @james-henstridge's answer already covered how you could make it work, I won't duplicate what he said, but I will explain why his answer works.

In Go, arrays work a bit differently than in most other languages (yes, there are arrays and slices. I'll discuss slices later). In Go, arrays are fixed-size, as you use in your code (so, <code>[3]int</code> is a different type than <code>[4]int</code>). Additionally, arrays are values. What this means is that if I copy an array from one place to another, I'm actually copying all of the elements of the array (instead of, as in most other languages, just making another reference to the same array). For example:

```
a := [3]int{1, 2, 3} // Array literal
b := a               // Copy the contents of a into b
a[0] = 0
fmt.Println(a)       // Prints "[0 2 3]"
fmt.Println(b)       // Prints "[1 2 3]"
```
However, as you noticed, Go also has slices. Slices are similar to arrays, except in two key ways. First, they're variable length (so []int is the type of a slice of any number of integers). Second, slices are references. What this means is that when I create a slice, a piece of memory is allocated to represent the contents of the slice, and the slice variable itself is really just a pointer to that memory. Then, <code>when I copy that slice around, I'm really just copying the pointer. That means that if I copy the slice and then change one of the values, I change that value for everybody.</code> For example:

```
a := []int{1, 2, 3} // Slice literal
b := a              // a and b now point to the same memory
a[0] = 0
fmt.Println(a)      // Prints "[0 2 3]"
fmt.Println(b)      // Prints "[0 2 3]"
```
Implementation

If that explanation was pretty easily understandable, then you might also be curious to know how this is implemented (if you had trouble understanding that, I'd stop reading here because the details will probably just be confusing).

Under the hood, Go slices are actually structs. They have a pointer to the allocated memory, like I mentioned, but they also have two other key components: length and capacity. If it were described in Go terms, it'd look something like this:

```
type int-slice struct {
    data *int
    len  int
    cap  int
}
```
The length is the length of the slice, and it's there so that you can ask for len(mySlice), and also so that Go can check to make sure you're not accessing an element that's not actually in the slice. The capacity, however, is a bit more confusing. So let's dive a bit deeper.

When you first create a slice, you give a number of elements that you want the slice to be. For example, calling make([]int, 3) would give you a slice of 3 ints. What this does is allocate space in memory for 3 ints, and then give you back a struct with a pointer to the data, the length of 3, and the capacity of 3.

However, in Go, you can do what's called slicing. This is basically where you create a new slice out of an old slice that represents only part of the old slice. You use the slc[a:b] syntax to refer to the sub-slice of slc starting at index a and ending just before index b. So, for example:
```
a := [5]int{1, 2, 3, 4, 5}
b := a[1:4]
fmt.Println(b) // Prints "[2 3 4]"
```
What this slicing operation does under the hood is to make a copy of the struct that corresponds to a, and to edit the pointer to point 1 integer forward in memory (because the new slice starts at index 1), and edit the length to be 2 shorter than before (because the old slice had length 5, while the new one has length 3). So what does this look like in memory now? Well, if we could visualize the integers laid out, it'd look something like this:
```
  begin     end  // a
  v         v
[ 1 2 3 4 5 ]
    ^     ^
    begin end    // b
```
Notice how the there's still one more int after the end of b? Well that's the capacity. See, so long as the memory's going to be sticking around for us to use, we might as well be able to use all of it. So even if you only have a slice whose length is small, it will remember that there's more capacity in case you ever want it back. So, for example:
```
a := []int{1, 2, 3}
b := a[0:1]
fmt.Println(b) // Prints "[1]"
b = b[0:3]
fmt.Println(b) // Prints "[1 2 3]"
```
See how we do b[0:3] at the end there? The length of b is actually less than 3 at this point, so the only reason we're able to do that is that Go has kept track of the fact that, in the underlying memory, we've actually got more capacity reserved. That way, when we ask for some of it back, it can happily oblige.
</details>

</details>

> https://golangr.com/hello-world

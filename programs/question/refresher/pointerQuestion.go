/*
how do you initialise a pointer in go
p := &i
p := new(int)*
var p **int = i*
var p = &int(5)


The correct answer is p := &i.



Explanation:


In Go, you can create and initialize a pointer by using the "&" operator.


Option var p = &int(5) is incorrect because it tries to take address of a constant. (cannot take address of int(5) (constant 5 of type int))


Correct syntax would be:


i := 5
var p = &i
*/

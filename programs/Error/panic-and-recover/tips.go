package main

// don't panic, return an error! - best practise

/*A panic is an event that you don't expect. Try to minimize cases that can cause panic events in your code. Checking for a nil data is suitable in most cases. If you are working with numbers, remember to add error checking for 0 values. Don't use strict comparison for bounds of the range: try to replace a == statement for a >= or <=. When working with strings, check the incoming string for unsuitable symbols: create a set of characters to use, such as a regexp, for example. These simple rules will make your code more stable.

Now, here's a recap of what this topic has covered:

    a panic is an event caused by abnormal conditions during runtime;

    you can use the panic() function to call the panic event directly;

    recover is the process that can handle a certain panic event.*/
// code events like: Dividing an integer number by zero
//Getting out of the slice bounds can cause our program to panic, or can cause the panic event so to say.

/*
Panic:
The Go compiler always checks for errors during the building process (before you run your program). It catches common types of errors: syntax errors, unclosed brackets, unused variables, type errors, and more. Despite that, the code runtime is a dynamic environment, and some errors can occur during execution of your code.

# Panic occurs after the code is compiled, or Any error that compiler can catch and warn you about won't lead to panic.
*/

### About the project
The project is educational. The main goal is to learn how to work with multithreading using the example of implementing a simple calculator.

### Version 1
The write_to_file package implements a simple calculator that asynchronously calculates the results of simple arithmetic expressions and writes them to a file.

Example:
> Number of expressions: 3
> 
> 4 + 5
> 
> 8 * 2
> 
> 11 - 1
> 
File contents:
> 4 + 5 = 9
> 
> 8 * 2 = 16
> 
> 11 - 1 = 10
> 

### Version 2
The write_to_stdout package implements a simple calculator that asynchronously evaluates the results of simple arithmetic expressions and prints them to standard output. This implementation uses the sync.WaitGroup structure.

Example:
> Number of expressions: 3
>
> 17 - 7
>
> 5 * 4
>
> 7 + 6
>
Output:
> 17 - 7 = 10
>
> 5 * 4 = 20
>
> 7 + 6 = 13
> 
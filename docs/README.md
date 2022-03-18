# Task 1 - Debugging Concurrent Programs

## Buggy Program 1
The program doesn't work because there is no other goroutine to read the value of the channel. The program is therefore stuck in a deadlock. 

To fix this bug we can add a goroutine that adds the string "Hello world!" to the channel. The man goroutine is then able to read the channel and the program works as intended. See solution in src/bug01.go

## Buggy Program 2
The program finishes executing when the main function is terminated. In this case the main function terminates before the last number gets to printed. 

We can fix this bug using a WaitGroup. This way we can let the main function wait until the last number has been printed before terminating. See solution in src/bug02.go

# Task 2 - Many Senders; Many Receivers
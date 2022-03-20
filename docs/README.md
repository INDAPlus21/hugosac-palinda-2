# Task 1 - Debugging Concurrent Programs

## Buggy Program 1
The program doesn't work because there is no other goroutine to read the value of the channel. The program is therefore stuck in a deadlock. 

To fix this bug we can add a goroutine that adds the string "Hello world!" to the channel. The man goroutine is then able to read the channel and the program works as intended. See solution in src/bug01.go

## Buggy Program 2
The program finishes executing when the main function is terminated. In this case the main function terminates before the last number gets to printed. 

We can fix this bug using a WaitGroup. This way we can let the main function wait until the last number has been printed before terminating. See solution in src/bug02.go

# Task 2 - Many Senders; Many Receivers

### What happens if you switch the order of the statements `wgp.Wait()` and `close(ch)` in the end of the `main` function?
**Hypothesis**: The channel will close before the producers are done using the channel which will lead to a crash. 

**Result**: The program crashes due to a goroutine trying to send on a closed channel. 

### What happens if you move the `close(ch)` from the `main` function and instead close the channel in the of the function `Produce`?
**Hypothesis**: The channel will close when the first producer is done which will crash the program as the other producer goroutines are trying to send on a closed channel.

**Result**: The program crashes after one of the goroutines have finished executing the `Produce` function as the next goroutine is trying to send on a closed channel.

### What happens if you remove the statement `close(ch)` completely?
**Hypothesis**: The program works as previously as there is no gorutine sending to the channel as the `wgp.Wait()` statement ensures that all the goroutines running the `Produce` function has finished executing. 

**Result**: The program works as previously.

### What happens if you increase the number of consumers from 2 to 4?
**Hypothesis**: The program's execution time is reduced as the number of consumers. As the delays are random, the time reduction won't be consistent but should be about half of the current execution time.

**Result**: The execution time is reduced from ~ 700 - 1000 ms to ~ 500 - 700 ms.

### Can you be sure that all strings are printed before the program stops?
**Hypothesis**: No, as the channel might close before the consumers have consumed all the strings.

**Result**: All the strings are not always printed before the program stops.
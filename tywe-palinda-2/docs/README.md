If written answers are required, you can add them to this file. Just copy the
relevant questions from the root of the repo, preferably in
[Markdown](https://guides.github.com/features/mastering-markdown/) format :)

### Task 1 - Debugging Concurrent Programs

bug01.go 
Solution: The code was running into a deadlock because the declared channel wasn't able to offload its data onto another goroutine. I have introduced another anonymouse function on another goroutine that sends the text through the declared channel into the main.

bug02.go
Solution: The main routine was finsihing before the go routine for print finished printing all the numbers up to 11. By adding a waitGroup, the go routine for print is allowed to finish receiving all the numbers from main and printing them. 

### Task 2 - Many Sender; Many Receivers
What happens if you switch the order of the statements wgp.Wait() and close(ch) in the end of the main function?

Ans: The program panics because the channel is closed before all the values are transmitted.

What happens if you move the close(ch) from the main function and instead close the channel in the end of the function Produce?

Ans: The program panics as expected because the channel is closed. 

What happens if you remove the statement close(ch) completely?

Ans: The program runs smoothly without any panic because the channel is not closed at all

What happens if you increase the number of consumers from 2 to 4?

Ans: The running time decreases aproximately by half because the number of consumers is doubled.

Can you be sure that all strings are printed before the program stops?

Ans: Yes(atleast after the modifications). 




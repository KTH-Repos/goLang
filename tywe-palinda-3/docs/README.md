If written answers are required, you can add them to this file. Just copy the
relevant questions from the root of the repo, preferably in
[Markdown](https://guides.github.com/features/mastering-markdown/) format :)

### Task 1 - Matching Behaviour
1, What happens if you remove the go-command from the Seek call in the main function?

Ans: The order changes so that the first person sends message to the second, the third to the fourth but a message
from the last person isn't received by anyone.  

2, What happens if you switch the declaration wg := new(sync.WaitGroup) to var wg sync.WaitGroup and the parameter wg *sync.WaitGroup to wg sync.WaitGroup?

Ans: I get negative weightgroup counter. That's because the weightgroup is pointing to a copy of the original weighgroup. The copy has no added goroutines, so when wg.Done() is executed the number of goroutines becomes negative. 

3, What happens if you remove the buffer on the channel match?

Ans: Because a message from the fifth person isn't received by anybody the unbuffered channel holds the message until it can find somebody to send it to, which in this case doesn't happen at all. Hence, DEADLOCK! 

4, What happens if you remove the default-case from the case-statement in the main function?

Ans: When there are 5 people in the array the order becomes random. When there are 4 peopole in the array a deadlock happens. This is due to the fact that when data is not being read from the channel, the channel blocks the whole execution of the goroutine. While it is being blocked there is no default case to take care of the case when no data
is being read from the channel which causes the whole program to crash. Hence, DEADLOCK!

### Task 2 - Julia

The number of cpus(runtime.NumCpu) taking part in the operations is 4. The running time went from around 17 sec to 7 sec. 

### Task 3 - MapReduce

16 goroutines are created to parallelise mapreduce. 

| Variant  | Runtime(ms) |  
|---|---|
| singleworker  |  131 ms |   
| mapreduce   |  6 ms |         
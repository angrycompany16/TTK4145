Exercise 1 - Theory questions
-----------------------------

### Concepts

What is the difference between *concurrency* and *parallelism*?
> Concurrency does not usually involve "doing two things at the same time" as parallellism does, but rather having multiple tasks that all need to be tracked, and that have some downtime. In the downtime, the scheduler chooses a different task to work on, and thus interleaves the work of the tasks, which makes it seem like they are happening at the same time. With parallellism, on the other hand, the task is divided into subtasks which are executed by multiple cores or multiple processors.

What is the difference between a *race condition* and a *data race*? 
> A race condition happens when the behaviour of a program depends on the timing of events that cannot be controlled, such as (in the shared variable tasks) the way the scheduler distributes processing power to the threads. This leads to inconcistent results when running the program. A data race, on the other hand, is a more concrete error which occurs when two instructions try to access the same register and at least one of them is a write. For example, a write and read at the same time cause a data race.

*Very* roughly - what does a *scheduler* do, and how does it do it?
> A scheduler chooses which thread gets to use the CPU processing power. This can be done in a variety of ways, one of which is to choose randomly (which is surprisingly effective!) 


### Engineering

Why would we use multiple threads? What kinds of problems do threads solve?
> Whenever we have multiple independent tasks that require a full thread of execution, we often can use multithreading. An example is when polling for events from multiple input sources. With multithreading we can set up a blocking loop in a thread for each input device, sends the message to some main program whenever an input is received.

Some languages support "fibers" (sometimes called "green threads") or "coroutines"? What are they, and why would we rather use them over threads?
> Fibers use cooperative multitasking, meaning that the operating system never introduces a context switch, but rather the fibers themselves say can switch to other fibesr. This essentially means that fibers are a 'step down' from threads, as a thread can have multiple fibers due to the cooperative multitasking. Using fibers means that thread safety is less of an issue as the context switching only happens from the fiber itself.

Does creating concurrent programs make the programmer's life easier? Harder? Maybe both?
> Sometimes concurrent programs can simplify the code and make it more intuitive and sometimes more effective. However, concurrent programming can make debugging a lot harder, and it can introduce bugs, indeterminism and hard-to-understand behaviour. Therefore concurrent programming should be used with care, and a simpler solution should ususally be chosen if it exists.

What do you think is best - *shared variables* or *message passing*?
> It seems that shared variables may introduce less overhead / boilerplate code, but message passing may be more scalable / useful for writing libraries. Therefore i think shared variables are best for smaller tasks, while message passing is better for larger tasks or projects.



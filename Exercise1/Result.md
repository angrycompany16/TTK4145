# Results exercise 1

### Shared variables

In this task, the magic number varied with each run of the program, ending up at a random number each time. The reason for this happening is that both C and Go by default use preemptive scheduling when multithreading. This means that the scheduler can, at any time, decide to switch the execution from one thread to another, which leads to issues when the functions that the threads are running modify shared variables. In the case of the 'shared variables' example, both the decrement and increment functions act on i by doing three instructions. It first reads the value from memory, then increments it, and then writes it back to its memory location. The problem is that a thread switch can happen before all three of these instructions have been executed, which makes them jumbled and causes instructions to be lost.

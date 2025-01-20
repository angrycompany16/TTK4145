### Code checklist examples:

- Does the module have a central purpose?
- Is the module organized around a common set of data?
- Does the module offer a cohesive set of services?
- Is the module dependent on other modules? Is it loosely coupled?
- Is the module’s interface abstract enough that you don’t have to think about how its services are implemented? Can you treat the module as a black box?
> Most importantly: Can you ensure these qualities not only in your code, but also in the code your team members write?

### Advice form sverre:
- Avoid the broken window. Do not start taking short-term decisions
- In your eagerness to see something work, or
- By trying to off-load your responsibility to the compiler or the bug-fixing phase.
- Programming is concentration work. You need your brain capacity. Spend your best, most rested hours of the week on it.
- Do not write code until the design/structure is mature
- (Spec,) Design and Implementation (and test…) ‘phases’.
- Yes, there is a causality between design and implementation. “You can not do something until you have decided how.”
- But No, This truth does not translate to a phase transition on the time axis. (Bruh wtf does this even mean)
### Two (stupid) elevators:
- One starts with the old elevator code from last year, uses it on three computers and tries to connect them with networking code. This leads to poor networking.
- Another starts with networking code, and tries to write an elevator on top of it. This leads to poor elevator code.
### Shared variable communication
This illustrates a classic problem when sharing variables. What if some other thread / cosmic ray / whatever changes p between assignment and for loop?
```C
p = 10
if (p == 10) {
	// Will this always be executed?
}
```
One consequence of this is that sharing variables ideally should be avoided as much as possible.
### Classical race condition
```C
//Globals: 
bool g_initDone = False; 

//Threads: 
t1(){
	// Do initialization
	g_initDone = True;
	resume(t2);
	// Continue executing
}

t2(){  
	if (g_initDone == False) {  
		Suspend(); 
	} 
	// Continue executing
```
If `g_initDone` changes unexpectedly, the program enters a race condition (`t2` is waiting for the variable to continue execution)
### Classical deadlock
```C
t1(){ 
	wait(A);
	wait(B);
	...
	signal(B);
	signal(A);
}

t2(){  
	wait(B);  
	wait(A);  
	...  
	signal(A);  
	signal(B);  
}
```
If a signal disappears, the code deadlocks. This can happen sometimes (?)
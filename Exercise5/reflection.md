## Questions

 - Different types of shared variable synchronization
  - Ada protected objects: We specify when access is given (default is locked) instead of when to lock the resource.
  - With condition variables we set up a condition for when threads can accesss the shared resource, and then update the access condition and propagate this
    throughout the threads.
  - Monitors in java are similar to semaphores / mutices in that they block access to an object if that object is currently in use in another thread.
 - Bugs/Code quality
  - We think that the approach should be to grant access gradually, rather than blocking access to something that is by default shared with everyone.
   - Therefore, ada and go (in particular the request approach) seem like the best approaches
  - This should make programs easier to maintain, as you can gradually open up the functionality as the program expands
   - In particular, the go approach make sthis very easy. However, we didn't like the priorityselect approach, as this implemented the priority mechanism in a
     not very beautiful way.
 - In the cases of semaphore.d, you would get a massive if statement, in protectedobj.adb you would get one function for each priority, and in priorityselect.go 
   you would (in our case) get lots of nested select statements. Therefore, the implementations that use a priority queue make it a lot easier to elegantly extend
   this with as many priorites as we want (we just need to sort the queue).
 - Semaphores exist in order to provide exclusive access to shared variables in a thread. Therefore semaphores should not be the thing that you use to control the
   flow of execution. This is for example because semaphores can change at any time and are essentially a shared variable which is very limited in what it can do.
 - Go message passing server strategy
  - simple, intuitive, safe
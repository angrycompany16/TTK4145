### Robust software & fault tolerance
We need to make bug-free software
However, as much as possible we also want bug-repellent software
Anything can fail, thus failure is a given

> Example: A computer could be plugged out in the lab, we need other computers to handle the elevators

Messages, connections and nodes can be lost in a network, we need to handle this!
One useful mindset is to keep modules as separate as possible and avoid interdependence between modules. If one module fails, the error should propagate to the "parent" node.

> Lab idea: What if we used loggers for some of the modules to gain more insight into the way the program is running?

Consider: Time problems
Messages can come in any order (?)

Solution: symmetry (?)

Counter-argument: Asymmetry
We can solve multithreading by using shared variables
Consider:
- Shared data can be in any state, including partially updated or invalid
- What is the point of the data, if its validity is un-enforcable?
Conclusion:
Robustness problems cannot be solved with shared variables and sunchronization?

A good way to handle errors can be to do nothing

> "Sometimes the solution is to kill yourself."

### Making the internet from scratch
Step 1: Connect two
Layer 1: Physical layer. USB, CAN, CAT etc.

Step 2: Connect more than two :brain_explosion:
Layer 2: Medium Access Control (MAC) addresses
Need MAC src, MAC dst, data, gap etc.

Step 3: Connect a lot
Can create a hierarchy
One machine cannot know if it sends to one machine, or if that machine is an entry point to other machine
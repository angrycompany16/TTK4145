Coul
- How much do we need to generalize our modules / how high should the abstraction level be?
	- *silence*
- Closure
	- The closure of a subset A is the intersection of all closed subsets containing A.
- A module should be a simple building block
- More Code Complete checklist examples
	- Is the reason for creating the routine sufficient?
	- Have all parts of the routine that would benefit from being put into routines of their own been put into routines of their own?
	- Is the routine's name a strong, clear verb-plus-object name for a procedure or a description of the return value for a function?
	- Does the routine's name describe everything the routine does?
	- Have you established naming conventions for common operations?
	- Does the routine have strong, functional cohesion -- Doing one and only one thing and doing it well?
- Write better code lol dummy
- Don't write comments just have better code quality
- Make names guessable
	- When writing code you should be able to guess what a variable was called
- How?
	- Metaphors
	- Names correspond to purpose/concept of variable
	- Established conventions
- If the variable is difficult to name, maybe something is wrong with the variable
- How to fix int floor = -1?
	- floor (int)
	- direction (bool/enum)
	- onFloor (bool)
	- doorOpen (bool?)
	- doorTimer (float)
- Problems:
	- ? None
- Suggested solution:
	- Elevator state machine
		- elevatorState
		- keep floor and doorTimer, move the rest of the bools into a state machine / enum
- Concrete advice, supposedly
	- Use conventions to differ between variables, consts, globals, member variables, types, classes, objects, enums, pointers, in parameters
	- Use conventions to fake private variables, constants, name spaces etc. if the language does not support it
	- Do not use one variables for two things
	- Give better names to temporary, loop, 'flags', status and boolean variables (if scope is not significant)

MovingUpDoorOpen
MovingDown
StationaryDoorOpen


enum State {
	MovinUPDoorCLosed
	...
}

State currentSate

switch currentState {
	case ...
}

Update() {
	if TRansition
		currentState = ...
}
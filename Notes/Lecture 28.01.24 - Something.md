Notes: Sverre thinks it's courl as usual (probably talking about emacs)
* The system is reliable if it behaves as the spec says

"Normal" way of bug handling
1. Make program that fulfills spec
2. Run program
3. Errors happen
4. Find the cause in the code
	1. Incomplete spec
	2. Missing handling of some situation
	3. Erroneous code
5. Fix or add code
Flipping stupid baka!!!!

Traditional error handling is not good enough because it cannot handle unexpected errors.
Testing is not good enough

How to improve traditional 
switch (error) {
	case ENOMEM:
		 ...
}

No answer :(
Actually the answer is Fault Tolerance. What is fault tolerance?
- Error detection -  Check in multiple threads, report errors instead of just dying
- Recovery - Roll back to previous state for example
- Redundancy - Store extra data for backup

Definition:
- Preventing errors from becoming failures

Terms and techniques for exam:
- GG skipped in one second

Detection: Acceptance tests
Do not test for the presence of errors
Do make tests to check that everything is ok

Auditorium acceptance test:
- Test some part of the auditorium that requires everything to work

Good tip for the lab:
- Make acceptance tests
- If acceptance tests fail, restart
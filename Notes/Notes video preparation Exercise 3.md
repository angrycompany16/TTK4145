- Only valid solution for failure is to retry
- Single node:
	- Try everything again
	- Try only the failed part again
	- beware: Optional parts can fail too...
- Compensate for failure with redundancy
- We can resend messages all the time, but there is a better (structured) way
- Processes must be stateful to be useful
- Messages must be unpredictable to be useful
- Losing messages is losing states and output
	- Reordering messages scrambles state and outputs
	- Want to solve: Time-ordering of messages
- Solving message loss
	- Send periodically
	- Resend when no confirmation is received within some timeframe
	- Do nothing - let the user start the sending process again
		- Light didn't turn on - I will press again
- Only duplication and reordering are problems
	- Two desirable traits
		- Idempotent: Receiving the duplicates has no effect
		- Commutative: The ordering has no effect
- We can do whatever we want with intermediates, but
	- So can the environment
	  The only place to compensate is with data

## So far
- Messages on network have arbitrary ordering
- System outputs must remain correct
- Must work with intermediates
	- Send whatever we need
	- Use some state to deal with what we receive
		- Use knowledge of what we have received previously to impose order on what we receive in the future
## The CAP theorem
You can have two:
- Consistency
	- Nodes produce same output
- Availability
	- Nodes produce output
- Partition tolerance
	- Tolerates arbitrary message loss
What can we sacrifice?
- Need partition tolerance - Elevators can disconnect
- Need availability - Elevators must be active (mostly)
- Want consistency, but don't need it all the time
	- Two elevators can show up for a single call (!)
	- You can move an elevator without lights if there is high packet loss
## So far
Manage arbitrary time-reordering by using intermediate freedom
We sacrifice consistency to achieve availability and partition tolerance
Want to constrain consistency in system
- Want mostly inconsistent internal data
# Consistency
Different nodes agree (have same data)
Prerequisites for inconsistency:
- Self-reference
	- Needed
- Negation
	- Not needed

Minimize negation as much as possible
- Deleting, undoing, etc..
## So far
Have to impose order on messages when receiving
Thus we need to discard messages (discard messages that are out of order in time)

DUDE he literally made the same system i did
WTf
## Controlling negation is hard
It just is 
Exploit all opportunities to avoid negation
- Latest-version-like data (in my design)
- Lack-of-messages-is-information data (network alive)
- Any negation that is completely necessary can be guarded

- Idea for elevator: Instead of resending the elevator request packets, we may simply require the user to press again
- 
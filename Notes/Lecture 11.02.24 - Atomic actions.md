What are atomic actions?
What are transactions?

Transactions:
- When discussing acceptance tests, "We do not know what went wrong, so we need to put ourselves in the position to be able to do recovery anyway"
- How to avoid the domino effect (?)

Transactions are a group of actions that should be performed as if they were a single bulk action

We often say that transactions are atomic actions, I.E. actions that cannot be partially done or undone, they are either done completely or not done at all

In our subject:
- A transaction is an atomic action without the backward recovery error mode as standard error mode

Common description of transactions:
- Atomicity: Either all side effects happen, or none
- Consistency: Leaves the system in a consistent state when finished and produces consistent results
- Isolation: Errors do not spread
- Durability: Results are not lost

Optimistic concurrency control:

- Assume non-interference, check afterward and handle as error
- Comes out as an alternative to locking
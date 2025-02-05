The networking module:
* The networking module is a fork of the given network module code. So far, no significant changes have been made.

The elevator server / client
* This one is al
* so a fork of the given repository with no significant changes made.

The thinking here is that, as well as optimizing for a bug-free program we wanted to optimize for speed to that we would have more time to debug if something strange happened. Also, if we wrote our own module we might introduce a lot more bugs since we don't have much experience writing this kind of code. Additionally, since we have a fork of the repo we are able to make any changes to the source code that may be required on the way. We have also made sure to fully understand how the code works in case anything should go wrong.

Naming conventions and such:
- Will take strong inspiration from the default go naming convention
	- camelCase: private
	- PascalCase: public
- Keep as much of the project as possible hidden and in separate repositories (we will later find out if this is actually a good idea)
	- That's why the forks are there
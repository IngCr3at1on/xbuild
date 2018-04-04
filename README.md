### Experimental Builder

This is meant to be used in place of normal build scripts or make.

Build steps are coded in Go and compiled into a binary which is then used to build the complete project (super redundant and super silly right?)

---

Why this instead of using make or a scripting language?

It started as something to allow a slack bot to compile go code from a channel and somewhere along the way I decided it would be interesting to use as a real build tool. There is no real reason to use this over Make or something else and in fact I would not suggest that anyone use this.
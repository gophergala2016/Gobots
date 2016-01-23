# Gobots

A programmatic robot-fighting game, **heavily** inspired by [Robot
Game](http://robotgame.net).

## Stream of Consciousness Logs

### Friday January 22, 2016 - 10:20 PM

Done for the night.  Added a tested `ToWire` method to `engine.Board`.
Converting to storage needs should be similarly trivial.  Added test
infrastructure for the `engine.Board.Update` method, so when the game logic
actually gets implemented, it's easy to test for correctness.  Not sure exactly
what to do for preserving turns as a basis, but I presume that it would be easy
to write a function of `(Replay, Round) -> Replay`. It would probably be good to
add a `engine.FromWire` function too. G'night. -Ross

### Friday January 22, 2016 - 9:27 PM

Created an engine package.  I figure we write the logic here where mutating
is easier to do in structs, and then bake out for storage or wire transfer.
-Ross

### Friday January 22, 2016 - 8:45 PM

Basic API schema and database interface, along with an idiomatic client wrapper
for writing bots.  Still not sure what the main server loop is going to look
like (or storage), but API first, implement later right? -Ross

### Friday January 22, 2016 - 5:54 PM

Ross has entered the arena. Also, it looks like Robot Game expired (or
something) literally between the Gala starting and now. This is very strange.


### Friday January 22, 2016 - 5:16 PM

Okay, I've done some stuff. I decided that I'm going to go with Protobufs and
gRPC (sorry Ross) for the data model and server communication things, but
virtually nothing useful exists yet. I was strongly considering Cap'n Proto,
but it ended up looking far too daunting to pick up in a few minutes, so I'm
sticking with what I (almost kinda) know. I'm not sleep-deprived or utterly
useless yet, so might as well get some productive work in bootstrapping the
application.

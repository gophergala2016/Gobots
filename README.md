# Gobots

A programmatic robot-fighting game, **heavily** inspired by [Robot
Game](http://robotgame.net).

## Stream of Consciousness Logs

### Sunday January 24, 2016 - 3:04 PM

To the previous update: Woot! -Brandon

### Sunday January 24, 2016 - 2:48 PM

I'm back! -Ross

### Sunday January 24, 2016 - 1:35 PM

Unfortunately didn't do any drunk committing last night (as you may be able to
tell), though karaoke was a good time. Did a bit more database stuff, and
getting GopherJS ready to do some stuff with replays so people can view games.
Still to be done: adding rounds to a replay. -Brandon

### Sunday January 24, 2016 - 8:21 AM

Couple of tweaks to the datastore to make the connect RPC not kill the server.
Once `lookupAIToken` is implemented, we'll be in business. -Ross

### Saturday January 23, 2016 - 7:04 PM

I've been sitting in this cafe for approximately five hours. This push will
contain such goodies as GitHub OAuth integration, secure cookies, working user
storage to a local key/value DB, and various other pleasantries and whatnot.
However, I'm getting fidgety and Korean Karaoke is imminent. Expect my next
commits to look more like \#CommitsFromLastNight \#BallmerPeak
\#IDon'tEvenHaveATwitter \#I'mTheWorst \#IHopeIEscapedTheseHashtagsProperly
-Brandon

### Saturday January 23, 2016 - 4:28 PM

First attempt at connecting server to client! RPC took place, but it crashed
because the datastore variable is nil.  RPCs are happening! -Ross

### Saturday January 23, 2016 - 4:22 PM

Wrote a bad AI and I feel bad. -Ross

### Saturday January 23, 2016 - 2:43 PM

That's more like it. After a good night's sleep and looking at the code base
with a fresh set of eyes (and some more instruction), I'm back on the
productivity train, and a lot of the game logic is implemented. I also hit
295x7 for DL, which is a PR and felt pretty sweet.
![Picking things up](http://i.imgur.com/507xBdZ.jpg)
Oh, and I'm caffeinated as all hell, which is probably helping move things
along.
![Drinking things down](http://i.imgur.com/WM8tlQv.jpg)
-Brandon

### Saturday January 23, 2016 - 9:09 AM

Fearing that my colleague may be hopelessly lost and confused by my "Commits
from Last Night", I added in the basic matchmaking structure in match.go.  While
incomplete, it handles the parts that require the most Cap'n Proto knowledge:
the connection setup etc.  Now you can start an `aiEndpoint` on a socket and
just start querying for online AIs.  Once you want to run a match, call
`runMatch` and it will block until the match completes.  There are still TODOs,
but my travel schedule dictates I must leave them to you.  Good luck, comrade!
-Ross

### Friday January 22, 2016 - 11:35 PM

I fear that my only contribution for the night will be these README updates. My
mind has turned to mush trying to separate wire formats from easyai formats
from engine formats from storage formats. These are not the awful development
practices that I have come to know and love, this is some serious software
engineering. I spun my wheels first on trying to figure out what was going on
so I could implement the game logic, then when that failed, I spun my wheels
trying to implement the other end of the client-server stream. I'm not actually
sure which one is the server and which one is the client, especially because
Cap'n Proto throws in a `ServerToClient` method to make sure that I'm royally
confused. I resolved to work on the frontend UI, for showing robots fighting,
but then I didn't even know which of our four formats should be used with
GopherJS. I'll look at this with a fresh set of eyes tomorrow after hopefully
setting a deadlift PR in the morning and eating a hearty breakfast at my
favorite café. -Brandon

### Friday January 22, 2016 - 10:47 PM

Dinner took longer than I expected, mainly because dinner turned into drinks,
and we all know how that goes. In any case, I'm back at the wheel again, and
I've spent the past 20 minutes looking over Ross's progress, which might as
well be written in Brainfuck, because my feeble, non Cap'n Proto-oriented mind
can't make heads or tails of it. More research is required. -Brandon

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
-Brandon

### Friday January 22, 2016 - 5:16 PM

Okay, I've done some stuff. I decided that I'm going to go with Protobufs and
gRPC (sorry Ross) for the data model and server communication things, but
virtually nothing useful exists yet. I was strongly considering Cap'n Proto,
but it ended up looking far too daunting to pick up in a few minutes, so I'm
sticking with what I (almost kinda) know. I'm not sleep-deprived or utterly
useless yet, so might as well get some productive work in bootstrapping the
application. -Brandon

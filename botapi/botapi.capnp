using Go = import "../../../../zombiezen.com/go/capnproto2/go.capnp";

@0x834c2fcbeb96c6bd;

$Go.package("botapi");
$Go.import("github.com/gophergala2016/Gobots/botapi");

interface AiConnector {
  # Bootstrap interface for the server.
  connect @0 ConnectRequest -> ();
}

struct ConnectRequest {
  credentials @0 :Credentials;
  ai @1 :Ai;
}

struct Credentials {
  secretToken @0 :Text;
}

interface Ai {
  # Interface that a competitor implements.
  takeTurn @0 (board :Board) -> (turn :List(Turn));
}

struct Board {
  gameId @4 :Text;

  width @0 :Int16;
  height @1 :Int16;
  robots @2 :List(Robot);

  round @3 :Int32;
}

struct Robot {
  id @0 :Int32;
  x @1 :Int16;
  y @2 :Int16;
  health @3 :Int16;
  faction @4 :Faction;
}

enum Faction {
  mine @0;
  opponent @1;
}

struct Turn {
  id @3 :Int32;

  union {
    wait @0 :Void;
    # Skip turn; do nothing.

    move @1 :Direction;

    attack @2 :Direction;

    selfDestruct @4 :Void;
    # Does damage to all surrounding bots (even diagonals).
  }
}

enum Direction {
  north @0;
  south @1;
  east @2;
  west @3;
}

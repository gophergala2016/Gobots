@0x834c2fcbeb96c6bd;

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
  takeTurn @0 (board :Board) -> (turn :Turn);
}

struct Board {
  width @0 :Int16;
  height @1 :Int16;
  robots @2 :List(Robot);

  myPoints @3 :Int32;
  opponentPoints @4 :Int32;
  round @5 :Int32;
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
  union {
    wait @0 :Void;
    # Skip turn; do nothing.

    move @1 :Direction;

    attack @2 :Direction;
  }
}

enum Direction {
  north @0;
  south @1;
  east @2;
  west @3;
}

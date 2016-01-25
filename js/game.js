angular.module('gobotApp', [])
.controller('GameController', function() {
  var game = this;
  var board = Gobot.GetReplayFromString("lol");
  game.rows = new Array(board.Height())
  for (var y = 0; y < board.Height(); y++) {
    game.rows[y] = new Array(board.Width())
    for (var x = 0; x < board.Width(); x++) {
      game.rows[y][x] = board.AtXY(x,y)
    }
  }
})
.config(function($interpolateProvider) {
  $interpolateProvider.startSymbol('//');
  $interpolateProvider.endSymbol('//');
});

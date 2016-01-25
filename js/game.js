angular.module('gobotApp', [])
.controller('GameController', function() {
  var game = this;
  var replay = Gobot.GetReplay("http://" + Host + "game/"+ GameID);

  game.setBoard = function(board) {
    game.rows = new Array(board.Height())
    for (var y = 0; y < board.Height(); y++) {
      game.rows[y] = new Array(board.Width())
      for (var x = 0; x < board.Width(); x++) {
        game.rows[y][x] = board.AtXY(x,y)
      }
    }
  }

  game.setBoard(replay.NextBoard());
  window.setTimeout(function() {
    game.setBoard(replay.NextBoard());
  }, 1000);
})
.config(function($interpolateProvider) {
  $interpolateProvider.startSymbol('//');
  $interpolateProvider.endSymbol('//');
});

$(function() {
  $('.add-bot').submit(function(e) {
    e.preventDefault();
    var f = $(this)
    $.post('/loadBots', f.serialize(), function(resp) {
      $('.ai-token').text(resp);
    });
    return false;
  });
});

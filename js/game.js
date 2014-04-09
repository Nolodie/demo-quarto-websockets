$(function() {
  $( '#pieces > div > div' ).draggable();
  $( '#pendingPiece').droppable({
    drop: function( event, ui) {
      alert('Piece Selected!');
    }
  });
  $( '.board' ).droppable({
    drop: function( event, ui ) {
      alert("Piece Placed!");
    }
  });

});


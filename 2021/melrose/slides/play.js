console.log('loading play.js ...');
document.addEventListener('keydown', function(event){ 
    // console.log(event.keyCode);
    // console.log(event.key);
    if (event.key == 3) {            
        if (event.keyCode == 51) { // cmd+3
            console.log("cmd+3",getSelectionText());
            event.stopPropagation();
            event.preventDefault();
        }
    }
} );
function getSelectionText() {
    var text = "";
    if (window.getSelection) {
        text = window.getSelection().toString();
    } else if (document.selection && document.selection.type != "Control") {
        text = document.selection.createRange().text;
    }
    return text;
} 
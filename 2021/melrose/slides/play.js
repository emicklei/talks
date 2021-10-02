console.log('loading play.js ...');
document.addEventListener('keydown', function(event){ 
    // console.log(event.keyCode);
    // console.log(event.key);
    if (event.key == 3) {            
        if (event.keyCode == 51) { // cmd+3
            console.log("cmd+3",getSelectionText());

            var xhr = new XMLHttpRequest();
            xhr.open("POST", "http://localhost:8118/v1/statements?action=play", true);
            xhr.setRequestHeader('Content-Type', 'text/plain');
            try {
                xhr.send(getSelectionText());
            } catch (error) {
                console.log("send error",error);
            }

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
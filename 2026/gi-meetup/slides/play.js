console.log("loading play.js ...");
document.addEventListener("keydown", function (event) {
  // console.log(event.keyCode);
  // console.log(event.key);
  if (event.key == 3) {
    if (event.keyCode == 51) {
      //console.log("cmd+3", getSelectionText());
      runCode(getSelectionText());
      event.stopPropagation();
      event.preventDefault();
    }
  }
});
function runGi() {
  // TODO how to detect which code block to run? look for section id?
  console.log("count codes:", document.getElementsByTagName("code").length);
  var code = document.getElementsByTagName("code")[0].innerText;
  //console.log("code:", code);
  runCode(code);
}
function runCode(code) {
  var xhr = new XMLHttpRequest();
  xhr.open("POST", "http://localhost:7171/gi?func=main", true);
  xhr.setRequestHeader("Content-Type", "text/plain");
  try {
    xhr.send(code);
  } catch (error) {
    console.log("send error", error);
  }
}
function getSelectionText() {
  var text = "";
  if (window.getSelection) {
    text = window.getSelection().toString();
  } else if (document.selection && document.selection.type != "Control") {
    text = document.selection.createRange().text;
  }
  return text;
}

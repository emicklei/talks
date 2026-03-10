console.log("loading play.js ...");
document.addEventListener("DOMContentLoaded", function () {
  document.querySelectorAll("button[data-header-id]").forEach(function (btn) {
    var headerId = btn.dataset.headerId;
    var container = document.getElementById(headerId) && document.getElementById(headerId).parentNode;
    var codeEl = container && container.querySelector("code");
    if (codeEl) {
      codeEl.contentEditable = "true";
      codeEl.spellcheck = false;
    }
  });
});
document.addEventListener("click", function (event) {
  //console.log("click target:", event.target.tagName, event.target.outerHTML);
  var btn = event.target.closest("button");
  //console.log("btn found:", btn, btn && btn.dataset);
  if (btn && btn.dataset.headerId) runGi(btn);
});
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
function runGi(btn) {
  var headerId = btn.dataset.headerId;
  var container = document.getElementById(headerId).parentNode;
  //console.log("container:", container);
  var codeEl = container && container.querySelector("code");
  //console.log("codeEl:", codeEl);
  //console.log("source:", codeEl.textContent);
  btn.parentNode.align = "left";
  btn.parentNode.classList.add("output");
  runCode(codeEl.textContent, btn.parentNode);
}
function runCode(code, outputEl) {
  fetch("http://localhost:7171/gi?func=main", {
    method: "POST",
    headers: { "Content-Type": "text/plain" },
    body: code,
  })
    .then((r) => r.text())
    .then((text) => {
      if (outputEl) {
        outputEl.textContent = text;
        outputEl.style.display = "";
      }
    })
    .catch((err) => console.log("fetch error", err));
}

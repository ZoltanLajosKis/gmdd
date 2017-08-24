document.addEventListener("DOMContentLoaded", function() {
    document.querySelectorAll("code.language-math").forEach(function(e) {
        renderMathInElement(e, {delimiters: [{left: "$$", right: "$$", display: true}], throwOnError: false});
    });
    document.querySelectorAll("code.inline-math").forEach(function(e) {
        renderMathInElement(e, {delimiters: [{left: "$", right: "$", display: false}], throwOnError: false});
    });

    mermaid.init({startOnLoad:true}, ".language-mermaid");

    hljs.initHighlightingOnLoad();
});


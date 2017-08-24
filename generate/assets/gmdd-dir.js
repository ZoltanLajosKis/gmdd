document.addEventListener("DOMContentLoaded", function() {
    document.querySelectorAll(".moment").forEach(function(e) {
        let span = document.createElement("span");
        span.setAttribute("title", e.textContent);
        span.textContent = moment(e.textContent, "YYYY-MM-DD HH:mm:ss").fromNow();
        e.textContent = "";
        e.appendChild(span);
    });

    document.querySelectorAll(".filesize").forEach(function(e) {
        let span = document.createElement("span");
        span.setAttribute("title", e.textContent + " Bytes");
        span.textContent = filesize(e.textContent);
        e.textContent = "";
        e.appendChild(span);
    });
});


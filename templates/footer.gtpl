{{define "footer"}}
<!-- This is required to make this work for some reason -->
<!-- Place this on your main window (Server dies on this window's exit) -->
<script>
  window.onload = function() {
    setTimeout(function() {
      var webview = document.createElement("webview");
      document.body.appendChild(webview);
      webview.src = "/";
    }, 0);
  }
</script>
</body>
</html>
{{end}}

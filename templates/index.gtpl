{{template "header"}}
{{template "navigation"}}
<div class="container">
<h2>Greetings {{.name}}!</h2>
<small>Go away... Or not, I'm not a cop.</small>
<form action="/" method="POST">
<input type="submit" value="New Window">
</form>
</div>
<script>
localStorage.setItem("someSessionItem", "This was set by the session call in the main window.");
</script>
{{template "footer"}}

{{template "header"}}
<div class="container">
Yay!  New Smaller Window!
<div id="result">
</div>
</div>
<script>
document.getElementById("result").innerHTML = localStorage.getItem("someSessionItem");
</script>
{{template "footer"}}

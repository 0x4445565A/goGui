{{template "header-angular"}}
<body ng-app="ExampleApp" layout="column" ng-controller="ApplicationController as app">
{{template "navigation"}}
<div class="container">
<p>
It works! If you see an h1 tag below that means that the Angular app created the &lt;hello-world&gt; directive and it loaded the static template properly!
</p>
<p>
It is important that the file is served up static otherwise Go would try to interpret the curly brackets as template imformation, which they are...  But for Angular not Go...
</p>
<p>
Either way this is rad, enjoy building single page semi-native applications with Go and Angular JS (Powered by Thrust!)
</p>
<hello-world></hello-world>
</div>
<script src="/js/angular.min.js"></script>
<script src="/js/application.js"></script>
<script src="/js/applicationController.js"></script>
<script type="text/javascript">    
  angular.module('ExampleApp', ['application']);
</script>
{{template "footer"}}

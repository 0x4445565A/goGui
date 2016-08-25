(function(){

  angular
       .module('application')
       .controller('ApplicationController', [
          '$scope',
          ApplicationController
       ])
       .directive('helloWorld', function () {
          return {
            templateUrl: './js/templates/helloWorld.html',
          };
        });

  function ApplicationController($scope) {
    var self = this;
    self.message = 'Hello World!';
  }

})();
var ng = require('@angular/core');
var ngRouter = require('@angular/router');
var HelloComponent = require('./hello');

var RootComponent =
  ng.Component({
    selector: 'fountain-root',
    template: '<router-outlet></router-outlet>'
  })
  .Class({
    constructor: function () {}
  });

var routes = [
  {
    path: '',
    component: HelloComponent
  }
];

module.exports = {RootComponent: RootComponent, routes: routes, routing: ngRouter.RouterModule.forRoot(routes)};

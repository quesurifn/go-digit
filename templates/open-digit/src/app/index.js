var ng = require('@angular/core');
var ngPlatformBrowser = require('@angular/platform-browser');
var myRoutes = require('./routes');

var HelloComponent = require('./hello');

module.exports = ng.NgModule({
  imports: [
    ngPlatformBrowser.BrowserModule,
    myRoutes.routing
  ],
  declarations: [
    myRoutes.RootComponent,
    HelloComponent
  ],
  bootstrap: [myRoutes.RootComponent]
})
.Class({
  constructor: function () {}
});

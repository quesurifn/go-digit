import {NgModule} from '@angular/core';
import {BrowserModule} from '@angular/platform-browser';
import {routing, RootComponent} from './routes';

import {HelloComponent} from './hello';
import {StartComponent} from './start';

@NgModule({
  imports: [
    BrowserModule,
    routing
  ],
  declarations: [
    RootComponent,
    HelloComponent,
    StartComponent,
  ],
  bootstrap: [RootComponent]
})
export class AppModule {}

import {Component} from '@angular/core';
import {RouterModule, Routes} from '@angular/router';
import {HelloComponent} from './hello';
import {StartComponent} from './start';

@Component({
  selector: 'fountain-root',
  template: '<router-outlet></router-outlet>'
})
export class RootComponent {}

export const routes: Routes = [
  {
    path: '',
    component: HelloComponent
  },
  {
    path: 'get-started',
    component: StartComponent,
  }
];

export const routing = RouterModule.forRoot(routes);

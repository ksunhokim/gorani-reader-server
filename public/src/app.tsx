import * as React from 'react';
import { Switch, Route } from 'react-router-dom';
import { Notfound } from './containers/404';
import Main from './containers/Main';
import DevTools from './components/ReduxDevTool';

export const App = () => (
  <div>
    <DevTools/>
    <Switch>
      <Route path = "/404" component = {Notfound}/>
      <Route path = "/" component = {Main}/>
    </Switch>
  </div>
);

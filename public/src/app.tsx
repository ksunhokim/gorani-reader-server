import * as React from 'react';
import { Switch, Route } from 'react-router-dom';
import { Main } from './containers/Main';
import { Login } from './containers/Login';

export const App = () => (
  <div>
    <Switch>
      <Route exact path = "/login" component = {Login}/>
      <Route path = "/" component = {Main}/>
    </Switch>
  </div>
);

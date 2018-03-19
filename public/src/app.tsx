import * as React from 'react';
import { Switch, Route } from 'react-router-dom';
import Main from './containers/Main';

export const App = () => (
  <div>
    <Switch>
      <Route path = "/" component = {Main}/>
    </Switch>
  </div>
);

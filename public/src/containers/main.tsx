import * as React from 'react';
import { Switch, Route } from 'react-router-dom'
import { Home } from './home';
import { Roster } from './roster';

export const Main = () => (
  <main>
    <Switch>
      <Route exact path="/" component={Home}/>
      <Route path="/roster" component={Roster}/>
    </Switch>
  </main>
);

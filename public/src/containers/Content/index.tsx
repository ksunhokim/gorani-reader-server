import * as React from 'react';
import { Switch, Route } from 'react-router-dom';
import { Home } from './Home';
import { Wordbooks } from './Wordbooks';

export const Content = () => (
  <main>
    <Switch>
      <Route path = "/wordbooks"  component = {Wordbooks}/>
      <Route exact path = "/"  component = {Home}/>
    </Switch>
  </main>
);
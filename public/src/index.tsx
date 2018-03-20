import * as React from 'react';
import * as ReactDOM from 'react-dom';

import './styles';
import { App } from './app';
import { BrowserRouter } from 'react-router-dom';
import { Store } from 'redux';
import { Provider } from 'react-redux';

import createStore from './store';

const initialState = {};
const store: Store<any> = createStore(initialState);

ReactDOM.render(
  <Provider store={store}>
    <BrowserRouter>
      <App />
    </BrowserRouter>
  </Provider>,
  document.getElementById('root') as HTMLElement,
);

if (module.hot) {
  module.hot.accept();
}
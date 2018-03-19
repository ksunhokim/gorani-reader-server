import * as React from 'react';
import { Header } from '../Header';
import { Content } from '../Content';

export class Main extends React.Component<{}, {}> {
  render() {
    return (
      <div>
        <Header />
        <Content />
      </div>
    );
  }
}

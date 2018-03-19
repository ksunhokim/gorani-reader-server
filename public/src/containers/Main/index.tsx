import * as React from 'react';
import { Header, HeaderItem, SeenMode } from '../../components/Header';
import { Content } from '../Content';

const items: HeaderItem[] = [
  {
    name: 'main',
    endPoint: '/',
    seenMode: SeenMode.LOGIN,
  },
  {
    name: 'main',
    endPoint: '/',
    seenMode: SeenMode.EVERY,
  },
  {
    name: 'go',
    endPoint: '/login',
    seenMode: SeenMode.EVERY,
  },
];

export class Main extends React.Component<{}, {}> {
  render() {
    return (
      <div>
        <Header items = {items}/>
        <Content />
      </div>
    );
  }
}

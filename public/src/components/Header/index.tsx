import * as React from 'react';
import { Link } from 'react-router-dom';
import { HeaderItem, SeenMode } from './HeaderItem';

export * from './HeaderItem'

interface Props {
  items: HeaderItem[];
}

export class Header extends React.Component<Props, {}> {
  render() {
    const { items } = this.props;
    return (
      <header className="header">
        <a id="logo" href="/"></a>
        <ul>
          {
            items.filter((item) => (
              item.seenMode === SeenMode.EVERY
            )).map((item) => {
              return (
                <li key={item.name}>
                  <Link to={item.endPoint}>
                    {item.name}
                  </Link>
                </li>
              );
            })
          }
        </ul>
      </header>
    );
  }
}

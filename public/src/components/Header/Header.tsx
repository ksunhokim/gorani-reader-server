import * as React from 'react';
import { connect } from 'react-redux';
import { Auth } from '../../models'
import { Link } from 'react-router-dom';
import { HeaderItem, SeenMode } from './HeaderItem';

interface Props {
  auth: Auth;
  items: HeaderItem[];
}

export class Header extends React.Component<Props, {}> {
  constructor(props: Props) {
    super(props);
  }
  render() {
    const { items, auth } = this.props;
    return (
      <header className="header">
        <a id="logo" href="/">{auth.authed ? "true" : "false" }</a>
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
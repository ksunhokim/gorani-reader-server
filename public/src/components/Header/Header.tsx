import * as React from 'react';
import { connect } from 'react-redux';
import { Auth } from '../../models';
import { Link } from 'react-router-dom';
import { HeaderItem, SeenMode } from './HeaderItem';
import { gettext } from '../../translation';
import { Dispatch } from 'redux';
import { Anchor, Box, Header as GHeader, Menu } from 'grommet';

interface Props {
  auth: Auth;
  dispatch: Dispatch<{}>;
  items: HeaderItem[];
}

export class Header extends React.Component<Props, {}> {
  constructor(props: Props) {
    super(props);
  }
  render() {
    const { items, auth, dispatch } = this.props;
    const hitems =
    items.filter((item) => (
      item.seenMode === SeenMode.EVERY ||
      (item.seenMode === SeenMode.LOGIN && auth.authed) ||
      (item.seenMode === SeenMode.LOGOUT && !auth.authed)
    )).map((item) => {
      return (
        <Anchor>
          {
            item.endPoint ?
            (<Link to={item.endPoint}>
              {item.name}
            </Link>)
            :
            (<a href="#login" onClick={item.callback}>
              {item.name}
            </a>)
          }
        </Anchor>
      );
    });
    return (
      <GHeader justify="center" colorIndex="neutral-4">
        <Box size={{width: {max: 'xxlarge'}}} direction="row"
          responsive={false} justify="start" align="center"
          pad={{horizontal: 'medium'}} flex="grow">
          <a id="logo" href="/"></a>
          <Box pad="small" />
          <Menu label="Label" inline={true} direction="row" flex="grow">
            {hitems}
          </Menu>
        </Box>
      </GHeader>
    );
  }
}
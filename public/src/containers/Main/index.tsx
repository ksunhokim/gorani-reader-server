import * as React from 'react';
import { gettext } from '../../translation';
import { refreshAuth } from '../../actions/auth';
import { Dispatch } from 'redux';
import { connect } from 'react-redux';
import { Header, HeaderItem, SeenMode } from '../../components/Header';
import { Content } from '../Content';
import { Auth } from '../../models';
import { App } from 'grommet';

const items: HeaderItem[] = [
  {
    name: gettext('wordbooks'),
    endPoint: '/wordbooks',
    seenMode: SeenMode.LOGIN,
  },
  {
    name: gettext('login'),
    endPoint: '/login',
    seenMode: SeenMode.LOGOUT,
  },
];

interface Props {
  auth: Auth;
  dispatch: Dispatch<{}>;
}

class Main extends React.Component<Props, {}> {
  timer: number;
  componentDidMount() {
    const { dispatch } = this.props;
    dispatch(refreshAuth());
    this.timer = setInterval(() => {
      dispatch(refreshAuth());
    }, 30000);
  }
  componentWillUnmount() {
    clearInterval(this.timer);
  }
  render() {
    const { auth, dispatch } = this.props;
    return (
      <App center={true}>
        <Header auth={auth} dispatch={dispatch} items={items}/>
      </App>
    );
  }
}

const mapStateToProps = (state) => ({
  auth: state.auth,
});

export default connect(mapStateToProps)(Main);
import * as React from 'react';
import { refreshAuth } from '../../actions/auth'
import { Dispatch } from 'redux';
import { connect } from 'react-redux';
import { Header, HeaderItem, SeenMode } from '../../components/Header';
import { Content } from '../Content';
import { Auth } from '../../models';

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

interface Props {
  auth: Auth;
  dispatch: Dispatch<{}>;
}

class Main extends React.Component<Props, {}> {
  timer: number;
  componentDidMount() {
    const { dispatch } = this.props;
    this.timer = setInterval(() => {
      dispatch(refreshAuth());
    }, 1000);
  }
  render() {
    const { auth } = this.props;
    return (
      <div>
        <Header auth = {auth} items = {items}/>
        <Content />
      </div>
    );
  }
}

const mapStateToProps = (state) => ({
  auth: state.auth,
});

export default connect(mapStateToProps)(Main);
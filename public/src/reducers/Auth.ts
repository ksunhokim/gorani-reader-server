import { handleActions, Action } from 'redux-actions';

import { Auth } from '../models';
import {
  REFRESH_AUTH,
  AUTH_SUCCESS,
  AUTH_FAIL,
} from '../constants/ActionTypes';

const initialState: Auth = <Auth> {
  authed: false,
};

export default handleActions<Auth, any>({
  [REFRESH_AUTH]: (state: Auth, action: any): Auth => {
    return state;
  },
  [AUTH_SUCCESS]: (state: Auth, action: any): Auth => {
    return {
      authed: true,
    };
  },
  [AUTH_FAIL]: (state: Auth, action: any): Auth => {
    return {
      authed: false,
    };
  },
}, initialState);
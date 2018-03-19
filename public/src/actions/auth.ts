import { URL } from '../constants/API';
import { RSAA } from 'redux-api-middleware';
import { createAction, Action } from 'redux-actions';

import { Auth } from '../models';

import {
  REFRESH_AUTH,
  AUTH_SUCCESS,
  AUTH_FAIL,
} from '../constants/ActionTypes';

const refreshAuth = (): any  => ({
  [RSAA]: {
    endpoint: URL + '/auth/refresh',
    method: 'GET',
    credentials: 'same-origin',
    types: [
      REFRESH_AUTH,
      AUTH_SUCCESS,
      AUTH_FAIL,
    ],
  },
});

export {
  refreshAuth,
}
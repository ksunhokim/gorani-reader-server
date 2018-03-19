import { createStore, applyMiddleware, combineReducers } from 'redux';
import { apiMiddleware } from 'redux-api-middleware';
import { AuthReducer } from './reducers';

const reducer = combineReducers({auth: AuthReducer});
const createStoreWithMiddleware = applyMiddleware(apiMiddleware)(createStore);

export default (initialState) => createStoreWithMiddleware(reducer, initialState);

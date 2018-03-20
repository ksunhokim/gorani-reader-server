import { createStore, compose, applyMiddleware, combineReducers } from 'redux';
import { apiMiddleware } from 'redux-api-middleware';
import { AuthReducer } from './reducers';
import { dialogReducer } from 'redux-dialog';
import DevTools from './components/ReduxDevTool';

const reducer = combineReducers({auth: AuthReducer, dialogReducer});
const enhancer = compose(applyMiddleware(apiMiddleware), DevTools.instrument());

export default function configureStore(initialState: any) {
    const store = createStore(reducer, enhancer);
    return store;
}
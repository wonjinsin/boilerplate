import { combineReducers } from 'redux';
import alertReducer from "@reducers/alert";
import dummyReducer from "@client/reducers/dummy";

export default combineReducers({
  alertReducer,
  dummyReducer,
});
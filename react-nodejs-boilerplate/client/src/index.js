import React from "react";
import ReactDOM from "react-dom";
import { Provider } from "react-redux";
import { createStore, applyMiddleware } from "redux";
import thunk from "redux-thunk";
import Root from "@utils/Root";
import GlobalStyles from "@utils/GlobalStyles";
import reducers from "@reducers/index";
import AlertComponent from "@components/Alert/AlertComponent";

const store = createStore(reducers, applyMiddleware(thunk));

ReactDOM.render(
  <Provider store={store}>
    <Root />
    <GlobalStyles />
    <AlertComponent />
  </Provider>,
  document.getElementById("root")
);

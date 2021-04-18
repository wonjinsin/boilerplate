import React from "react";
import HomeContainer from "@containers/HomeContainer";
import ErrorContainer from "@containers/ErrorContainer";
import ErrorBoundary from "@utils/ErrorBoundary";
import { BrowserRouter, Route, Switch } from "react-router-dom";
import ScrollToTop from 'react-router-scroll-top';

const Router = (isLoggedIn) => {

  return (
    <BrowserRouter>
      <ScrollToTop>
        <ErrorBoundary>
          <Switch>
            <Route exact path="/" component={HomeContainer}></Route>
            <Route path="/*" component={ErrorContainer}></Route>
          </Switch>
        </ErrorBoundary>
      </ScrollToTop>
    </BrowserRouter>
  );
};

const Routes = (props) => {
  return <Router />;
};

export default Routes;

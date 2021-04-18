import React, { useState, useEffect } from "react";
import { connect } from "react-redux";
import HomeComponent from "@client/components/Home/HomeComponent";
import { alertSet } from '@actions/alert';

const HomeContainer = (props) => {
  const [isLoading, setIsLoading] = useState(true);

  /* eslint-disable */
  useEffect(() => {
    setIsLoading(false);
  }, []);
  /* eslint-enable */

  return !isLoading ? <HomeComponent {...props} /> : <></>;
};

const mapStateToProps = (state) => {
  return {
    dummy: state.dummyReducer.dummy
  };
};

const mapDispatchToProps = (dispatch) => {
  return {
    alertSet: (obj) => {
      return dispatch(alertSet(obj));
    },
  };
};

export default connect(mapStateToProps, mapDispatchToProps)(HomeContainer);

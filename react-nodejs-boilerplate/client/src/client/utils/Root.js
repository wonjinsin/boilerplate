import React, { useEffect, useState } from "react";
import { connect, useDispatch } from "react-redux";
import { Helmet } from "react-helmet";
import dotenv from "dotenv";
import Routes from "@routes/Routes";
import Loading from "@utils/Loading";
import { dummyThunkRequest } from "@actions/dummy";
dotenv.config();

const Root = (props) => {
  const initialDispatch = useDispatch();
  const [isLoading, setIsLoading] = useState(true);

  // middleware
  useEffect(() => {
    initialDispatch(dummyThunkRequest());
    setIsLoading(false);
  }, [initialDispatch]);

  return !isLoading ? (
    <>
      <Helmet meta={[{ name: "title", content: "Title area" }]}></Helmet>
      <Routes />
    </>
  ) : (
    <Loading />
  );
};

const mapDispatchToProps = (dispatch) => {
  return {
    dummyThunkRequest: () => {
      return dispatch(dummyThunkRequest());
    },
  };
};

export default connect(mapDispatchToProps)(Root);

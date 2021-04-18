import React from "react";
import ErrorNotfoundComponent from "@components/Error/ErrorNotfoundComponent";

const ErrorComponent = (props) => {
  return (
    <div>
      <ErrorNotfoundComponent history={props.history}/>
    </div>
  );
};

export default ErrorComponent;

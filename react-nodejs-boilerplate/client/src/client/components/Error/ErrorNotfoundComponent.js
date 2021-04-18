import React from "react";

const ErrorNotfoundComponent = ({ history }) => {
  const goBack = () => {
    return history ? history.goBack() : window.location.assign("/");
  };

  return (
    <div>
      <span>This is error page</span>
      <div onClick={() => goBack()}>
        go back
      </div>
    </div>
  );
};

export default ErrorNotfoundComponent;

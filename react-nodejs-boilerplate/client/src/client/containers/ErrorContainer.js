import React, { useEffect, useState } from "react";
import ErrorComponent from "@client/components/Error/ErrorComponent";

const ErrorContainer = (props) => {
  const [isLoading, setIsLoading] = useState(true);

  useEffect(() => {
    const init = () => {
      setIsLoading(false);
    };

    init();
  }, []);

  return !isLoading ? <ErrorComponent {...props} /> : <></>;
};

export default ErrorContainer;

import React, { useEffect, useState } from "react";
import HeaderComponent from "@components/Common/HeaderComponent";
import FooterComponent from "@components/Common/FooterComponent";
import styles from "@css/home/home.module.css";

const HomeComponent = ({ alertSet }) => {
  const [isLoading, setIsLoading] = useState(true);

  /* eslint-disable */
  useEffect(() => {
    setIsLoading(false);
  }, []);
  /* eslint-enable */

  const ButtonHandler = (alertType = "check") => {
    switch (alertType) {
      case "failed":
        alertSet({
          mainTitle: "Default Button",
          subTitle: "Subtitle area",
          callback: () => console.log("closed"),
          alertType,
        });
        break;
      case "question":
        alertSet({
          mainTitle: "Default Button",
          subTitle: "Subtitle area",
          callback: () => console.log("closed"),
          alertType,
        });
        break;
      default:
        alertSet({
          mainTitle: "Default Button",
          subTitle: "Subtitle area",
          callback: () => console.log("closed"),
        });
        break;
    }
  };

  return !isLoading ? (
    <div className={styles.home}>
      <HeaderComponent />
      <div>This is home component</div>
      <button onClick={() => ButtonHandler()}>Default Button</button>
      <button onClick={() => ButtonHandler("failed")}>Fail Button</button>
      <button onClick={() => ButtonHandler("question")}>Question Button</button>
      <FooterComponent />
    </div>
  ) : (
    <></>
  );
};

export default HomeComponent;

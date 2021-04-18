import React, { useState, useEffect } from "react";
import { connect } from "react-redux";
import Modal from "stunning-awesome-modal";
import styles from "@css/alert/alert.module.css";
import "@css/alert/dynamic.css";

const AlertComponent = ({ alertInner }) => {
  useEffect(() => {
    if (alertInner.mainTitle) setVisible(true);
  }, [alertInner]);

  const [visible, setVisible] = useState(false);

  const closeModal = () => {
    setVisible(false);
    if (alertInner.callback) return alertInner.callback();
  };

  const dynamicClasses = {
    failed: "failed",
    question: "question",
  };

  return (
    <Modal visible={visible} effect="fadeInUp">
      <div className={`${styles.modal_inner} alertModalInner`}>
        <div className={styles.txt_box}>
          <span
            className={`${styles.ico} ${
              alertInner.alertType === "failed" && dynamicClasses.failed
            } ${alertInner.alertType === "question" && dynamicClasses.question}`}
          ></span>
          <p>
            {alertInner.mainTitle && alertInner.mainTitle}
            <br />
            {alertInner.subTitle && alertInner.subTitle}
          </p>
          <button className={styles.btn_complete} onClick={() => closeModal()}>
            확인 완료
          </button>
        </div>
      </div>
    </Modal>
  );
};

const mapStateToProps = (state) => {
  return {
    alertInner: state.alertReducer,
  };
};

export default connect(mapStateToProps)(AlertComponent);

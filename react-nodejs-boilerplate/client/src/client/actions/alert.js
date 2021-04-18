import { ALERT_SET } from "@client/actions/actionTypes";

export const alertSet = (obj = {}) => {
  return {
    type: ALERT_SET,
    alertType: obj.alertType,
    mainTitle: obj.mainTitle,
    subTitle: obj.subTitle,
    callback: obj.callback,
  };
};

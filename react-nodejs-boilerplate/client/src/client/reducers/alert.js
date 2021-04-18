import { ALERT_SET } from "@client/actions/actionTypes";

const initialState = {
  alertType: "",
  mainTitle: "",
  subTitle: "",
  callback: undefined,
};

const alert = (state = initialState, action) => {
  switch (action.type) {
    case ALERT_SET:
      return {
        ...state,
        alertType: action.alertType ? action.alertType : "",
        mainTitle: action.mainTitle ? action.mainTitle : "",
        subTitle: action.subTitle ? action.subTitle : "",
        callback: action.callback ? action.callback : undefined,
      };
    default:
      return {
        alertType: "",
        mainTitle: "",
        subTitle: "",
        callback: undefined,
      };
  }
};

export default alert;

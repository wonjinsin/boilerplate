import {
  DUMMY_INIT,
  DUMMY_START,
  DUMMY_FINISH
} from "@client/actions/actionTypes";

export const dummyThunkRequest = () => {
  return (dispatch) => {
    dispatch(dummyInit());
    dispatch(dummyStart());
    dispatch(dummyFinish());
  };
};

export const dummyInit = () => {
  return {
    type: DUMMY_INIT,
  }
}

export const dummyStart = () => {
  return {
    type: DUMMY_START
  };
};

export const dummyFinish = () => {
  return {
    type: DUMMY_FINISH
  };
};
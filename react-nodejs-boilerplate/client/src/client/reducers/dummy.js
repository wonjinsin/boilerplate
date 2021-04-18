import { DUMMY_INIT, DUMMY_START, DUMMY_FINISH } from "@client/actions/actionTypes";

const initialState = {
  dummy: {
    contents: "Hello wonjinsin",
  },
};

const dummy = (state = initialState, action) => {
  switch (action.type) {
    case DUMMY_INIT:
      console.log("DUMMY_INIT");
      return {
        ...state,
        dummy: {
          ...state.dummy,
          contents: "Dummy init",
        },
      };
    case DUMMY_START:
      console.log("DUMMY_START");
      return {
        ...state,
        dummy: {
          ...state.dummy,
          contents: "Dummy start",
        },
      };
    case DUMMY_FINISH:
      console.log("DUMMY_FINISH");
      return {
        ...state,
        dummy: {
          ...state.dummy,
          contents: "Dummy finish",
        },
      };
    default:
      return state;
  }
};

export default dummy;

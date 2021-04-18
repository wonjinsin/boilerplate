import axios from "axios";

export const dummyFunc = () => {
  return axios
    .post(`/front/api/v1`, {}, {
      headers: {
        "Content-Type": "application/json",
      },
    })
    .then((response) => {
      return response.data;
    })
    .catch((error) => {
      return error.response.data;
    });
};
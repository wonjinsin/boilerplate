export const getAllParam = (url) => {
  let params = url.substr(url.indexOf("?") + 1);

  const obj = {};

  if (params) {
    params = params.split("&");

    for (let i = 0; i < params.length; i++) {
      let temp = params[i].split("=");
      obj[temp[0]] = temp[1];
    }
  }

  return obj;
};

export const emailChk = (str) => {
  const regex = /^([\w-]+(?:\.[\w-]+)*)@((?:[\w-]+\.)*\w[\w-]{0,66})\.([a-z]{2,6}(?:\.[a-z]{2})?)$/i;
  return regex.test(str);
};

export const passWordChk = (str) => {
  return str && str.length > 6 ? true : false;
};

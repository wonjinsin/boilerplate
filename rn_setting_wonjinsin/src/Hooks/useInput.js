import React, { useState } from "react";

const useInput = intialValue => {
  const [value, setValue] = useState(intialValue);
  const onChangeText = text => {
    setValue(text);
  };
  return { value, onChangeText, setValue };
};

export default useInput;

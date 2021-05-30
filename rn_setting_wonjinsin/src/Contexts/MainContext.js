import React, { createContext, useContext, useEffect, useState } from "react";
import { AsyncStorage } from "react-native";

export const MainContext = createContext();

export const MainProvider = ({ children, idx }) => {
  const [dummyState, setDummyState] = useState(true);
  const dummyHandler = () => {
    setDummyState(p => !p);
  }
  

  return (
    <MainContext.Provider
      value={{
        dummyState,
        dummyHandler
      }}
    >
      {children}
    </MainContext.Provider>
  );
};
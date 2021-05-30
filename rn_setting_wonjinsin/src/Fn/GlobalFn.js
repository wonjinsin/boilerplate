import React, { useContext, useState, useEffect } from "react";
import { AsyncStorage } from "react-native";

export const GlobalFn = {
  HeaderEllipsis: (str) => {
    const ellipsis = str.length < 20 ? str : str.slice(0, 20) + "...";
    return ellipsis;
  }
};
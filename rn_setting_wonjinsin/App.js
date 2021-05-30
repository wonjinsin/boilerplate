import React, { useState, useEffect } from "react";
import { Text, View } from "react-native";
import { ThemeProvider } from "styled-components";
import Theme from "Styles/Theme";
import { MainProvider } from "Contexts/MainContext";
import NavController from "Navigations/NavController";

export default function App() {
  const [loaded, setLoaded] = useState(false);
  const preLoad = async () => {
    try {
      setLoaded(true);
    } catch (e) {
      console.log(e);
    }
  };
  useEffect(() => {
    preLoad();
  }, []);
  return loaded && (
    <ThemeProvider theme={Theme}>
      <MainProvider>
        <NavController />
      </MainProvider>
    </ThemeProvider>
  );
}

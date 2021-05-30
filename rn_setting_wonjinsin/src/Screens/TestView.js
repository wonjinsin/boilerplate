import React, { useContext } from "react";
import styled from "styled-components";

const Wrapper = styled.View`
  display: flex;
  flex: 1;
  align-items: center;
  justify-content: center;
  color: ${props => props.theme.blackColor};
`;

const NewText = styled.Text``;
const NewButton = styled.TouchableOpacity``;

const TestView = ({ navigation, dummyState, dummyHandler }) => {
  const goToView2 = () => {
    navigation.navigate("TestView2");
  };

  const goToStackView = () => {
    navigation.navigate("StackView");
  };

  return (
    <Wrapper>
      <NewText>This is TestView Text Area</NewText>
      <NewButton onPress={() => goToView2()}>
        <NewText>Go to TestView2</NewText>
      </NewButton>
      <NewButton onPress={() => goToStackView()}>
        <NewText>Go to StackView</NewText>
      </NewButton>
    </Wrapper>
  );
};

export default TestView;

import React from "react";
import styled from "styled-components";

const Wrapper = styled.View`
  display: flex;
  flex: 1;
  align-items: center;
  justify-content: center;
`;

const NewText = styled.Text``;
const NewButton = styled.TouchableOpacity``;

const TestView2 = ({ navigation, dummyState, dummyHandler }) => {
  const goToView = () => {
    navigation.navigate("TestView");
  };

  return (
    <Wrapper>
      <NewText>This is TestView2 Text Area</NewText>
      <NewButton onPress={() => goToView()}>
        <NewText>Go to TestView1</NewText>
      </NewButton>
    </Wrapper>
  );
};

export default TestView2;

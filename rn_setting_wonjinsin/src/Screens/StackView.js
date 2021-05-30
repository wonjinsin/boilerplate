import React, {useContext} from 'react';
import styled from 'styled-components';

const Wrapper = styled.View`
  display: flex;
  flex: 1;
  align-items: center;
  justify-content: center;
  color: ${(props) => props.theme.blackColor};
`;

const NewText = styled.Text``;

const TestView = ({navigation}) => {
  return (
    <Wrapper>
      <NewText>This is StackView</NewText>
    </Wrapper>
  );
};

export default TestView;

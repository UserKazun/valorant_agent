import styled from "styled-components";
import SignUpButton from "./SignUpButton";

export default function Header() {
    return (
      <Wrapper>
        <TextWrapper>
          <Title>VALORANT AGENT</Title>
        </TextWrapper>
        <ButtonWrapper>
          <SignUpButton />
        </ButtonWrapper>
      </Wrapper>
    );
};

const Wrapper = styled.header`
    background: #BD3944;
    max-height: 64px;
    display: flex;
`

const TextWrapper = styled.div`
    padding: 1.45rem 1.0875rem;
`

const Title = styled.h1`
    color: white;
    font-family: 'valorant';
    font-size: 24px;
    margin: 0;
    text-align: left;
`

const ButtonWrapper = styled.div`
    margin: 0 0 0 auto;
    padding: 0.54rem 1.0875rem;
`

import styled from "styled-components";
import SignUpButton from "./SignUpButton";

export default function Header() {
    return (
      <Wrapper>
        <TextWrapper>
          <Title>VALORANT AGENT</Title>
        </TextWrapper>
        <Link>
          <LinkItem>Agents</LinkItem>
        </Link>
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
    padding: 0.54rem 1.0875rem;
`

const Link = styled.ul`
    line-height: 65px;
    float: left;
    margin-left: auto;
    list-style: none;
`

const LinkItem = styled.li`
    list-style: none;
    display: inline-block;
    margin: 0 20px;
    font-family: 'valorant';
    color: white;
`

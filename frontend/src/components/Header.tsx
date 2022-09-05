import styled from "styled-components";
import SignUpButton from "./SignUpButton";

//props: { siteTitle: string }
export default function Header(): JSX.Element {
    return (
      <Wrapper>
        <FontWrapper>
          <Title>VALORANT AGENT</Title>
        </FontWrapper>
        <ButtonWrapper>
          <SignUpButton />
        </ButtonWrapper>
      </Wrapper>
    );
  // <header
  //   style={{
  //     background: `#BD3944`,
  //     marginBottom: `1.45rem`,
  //     color: `#FFFFFF`,
  //   }}
  //   >
  //   <div
  //     style={{
  //       margin: `0 auto`,
  //       maxWidth: 1440,
  //       maxHeight: 64,
  //       padding: `1.45rem 1.0875rem`,
  //     }}
  //   >
  //     <h1
  //     style={{
  //       margin: 0,
  //       textAlign: 'left'
  //       }}>
  //       {props.siteTitle}
  //     </h1>
  //     <SignUpButton></SignUpButton>
  //   </div>
  // </header>
};

const Wrapper = styled.header`
    background: #BD3944;
    max-height: 64px;
`

const FontWrapper = styled.div`
    margin: 0 auto;
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
`

import styled from "styled-components";
import SignUpButton from "./SignUpButton";

//props: { siteTitle: string }
export default function Header(): JSX.Element {
    return (
      <Wrapper>

        <SignUpButton />
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

const Wrapper = styled.div`
  width: 1440px;
  height: 64px;
`

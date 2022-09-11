import styled from "styled-components";
import Logo from "images/V_Logomark_Red.png";

export default function HeroItem() {
    return (
        <Wrapper>
            <Title>VALORANT AGENT</Title>
            <LogoWrapper>
                <LogoImage src={Logo} alt="logo" />
            </LogoWrapper>
        </Wrapper>
    );
};

const Wrapper = styled.div`
    width: 1440px;
    height: 797px;
    background-color: black;
`

const Title = styled.div`
    position: absolute;
    width: 1202px;
    height: 141px;
    left: 117px;
    top: 116px;

    font-family: 'VALORANT';
    font-style: normal;
    font-weight: 400;
    font-size: 128px;
    line-height: 141px;
    text-transform: uppercase;

    color: #FD4556;
`

const LogoWrapper = styled.div`
    display: flex;
    align-items: center;
`

const LogoImage = styled.img`
    position: absolute;
    width: 1042px;
    height: 744px;
    left: 197px;
    top: 116px;
`

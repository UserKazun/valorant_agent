import styled from "styled-components";
import ToOfficialPageButton from "./ToOfficialPageButton";

export default function HeroItem() {
    return (
        <Wrapper>
            <Title>VALORANT AGENT</Title>
            <Description>
                This site is fan-made.
                Click the button below for the official page
            </Description>
            <ToOfficialPageButton />
        </Wrapper>
    );
};

const Wrapper = styled.div`
    width: 100%;
    height: 100vh;
    background: black;
`

const Title = styled.div`
    width: 1202px;
    height: 141px;

    font-family: 'VALORANT';
    font-style: normal;
    font-weight: 400;
    font-size: 128px;
    line-height: 141px;
    text-transform: uppercase;
    color: #FD4556;

    margin-left: auto;
    margin-right: auto;

    padding-top: 100px;
`

const Description = styled.p`
    width: 786px;
    height: 280px;

    font-family: 'VALORANT';
    font-style: normal;
    font-weight: 400;
    font-size: 64px;
    line-height: 70px;
    text-align: center;
    text-transform: uppercase;
    color: white;

    margin-left: auto;
    margin-right: auto;
    padding-top: 130px;
    padding-bottom: 160px;
`

import styled from "styled-components";
import ToOfficialPageButton from "./ToOfficialPageButton";

export default function HeroItem() {
    return (
        <Wrapper>
            <Title>VALORANT AGENT</Title>
            <ItemWrapper>
                <Description>
                    This site is fan-made.
                    Click the button below for the official page
                </Description>
                <ToOfficialPageButton />
            </ItemWrapper>
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

const ItemWrapper = styled.div`
    display: flex;
    align-items: center;
`

const Description = styled.p`
    position: absolute;
    width: 786px;
    height: 280px;
    left: 327px;
    top: 301px;

    font-family: 'VALORANT';
    font-style: normal;
    font-weight: 400;
    font-size: 64px;
    line-height: 70px;
    text-align: center;
    text-transform: uppercase;

    color: white;
`

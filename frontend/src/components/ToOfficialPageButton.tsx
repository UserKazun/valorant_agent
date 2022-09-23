import styled from "styled-components";

export default function ToOfficialPageButton() {
    return (
        <Wrapper>
            <TextWrapper>
                <Title>Official</Title>
            </TextWrapper>
        </Wrapper>
    );
};

const Wrapper = styled.div`
    width: 180px;
    height: 47px;
    display: grid;
    border-radius: 5px;
    border: solid;
    border-color: white;
    align-items: center;

    width: 200px;
    height: 60px;
    margin-left: auto;
    margin-right: auto;
    margin-top: 150px;
`

const TextWrapper = styled.div`
    display: grid;
`

const Title = styled.div`
    font-size: 18px;
    font-family: 'valorant';
    color: white;
    margin: 0 auto;
`

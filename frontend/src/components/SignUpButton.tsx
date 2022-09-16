import styled from "styled-components"

export default function SignUpButton() {
    return (
        <Wrapper>
            <TextWrapper>
                <Title>Sign Up</Title>
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
    padding: 0 50px;
`

const TextWrapper = styled.div`
    display: grid;
`

const Title = styled.div`
    font-size: 18px;
    font-family: 'valorant';
    color: white;
    margin: auto;
`

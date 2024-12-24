import { Container, Heading } from '@chakra-ui/react'


export default function Header(props: any) {
    return (
        <Container className='header'>
            <Heading textStyle="2xl">{props.name}</Heading>
        </Container>
    )
}
import { Box, Container, Text } from '@chakra-ui/react'
import { FaSearch, FaBook, FaHistory } from 'react-icons/fa'

export default function NavButton(props: any) {
    const renderIcon = () => {
        switch (props.index) {
            case 0:
                return <FaBook />
            case 1:
                return <FaSearch />
            case 2:
                return <FaHistory />
        }
    }

    return (
        <Container>
            <Box className='nav-button'>
                {renderIcon()}
                <Text textStyle="md">{props.text}</Text>
            </Box>
        </Container>
    )
}
import { Button, Container, Text } from "@chakra-ui/react";
import { FaSearch, FaBook, FaHistory } from "react-icons/fa";

export default function Option(props: any) {
	const renderIcons = () => {
		switch (props.index) {
			case 0:
				return <FaBook />;
			case 1:
				return <FaSearch />;
			case 2:
				return <FaHistory />;
		}
	};

	return (
		<Container>
			<Button
				variant={"plain"}
				display={"flex"}
				flexDir={"column"}
				justifyContent={"center"}
				alignItems={"center"}
				gap={"5px"}
			>
				{renderIcons()}
				<Text textStyle={"md"}>{props.text}</Text>
			</Button>
		</Container>
	);
}

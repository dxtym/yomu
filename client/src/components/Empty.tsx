import { Flex, Container, Text } from "@chakra-ui/react";

export default function Empty({ hasSearch = false }: { hasSearch?: boolean }) {
	return (
		<Container
			height={hasSearch ? "calc(100vh - 250px)" : "100vh"}
			display={"flex"}
			justifyContent={"center"}
			alignItems={"center"}
		>
			<Flex justifyContent={"center"} alignItems={"center"} flexDir={"column"}>
				<Text fontSize={"3xl"}>^•⩊•^</Text>
				<Text>Nothing Here Yet</Text>
			</Flex>
		</Container>
	);
}

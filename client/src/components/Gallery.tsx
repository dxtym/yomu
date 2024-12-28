import { Container, Grid, GridItem, Image, Link } from "@chakra-ui/react";

interface GalleryProps {
	data: Array<any>;
	hasSearch?: boolean;
}

export default function Gallery({ data, hasSearch = false }: GalleryProps) {
	return (
		<Container
			py={"25px"}
			px={hasSearch ? "0" : "25px"}
			position={"relative"}
			mt={hasSearch ? "0" : "75px"}
			mb={"75px"}
		>
			<Grid templateColumns={"repeat(2, 1fr)"} gap={"10"}>
				{data.map((item: any) => {
					return (
						<GridItem key={item.manga_url}>
							<Link href={`manga/${item.manga_url}`}>
								<Image
									src={item.cover_image}
									height={"220px"}
									width={"200px"}
								/>
							</Link>
						</GridItem>
					);
				})}
			</Grid>
		</Container>
	);
}

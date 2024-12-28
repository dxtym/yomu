import { Container, Input } from "@chakra-ui/react";
import axios from "axios";
import { useState } from "react";
import Empty from "./Empty";
import Gallery from "./Gallery";
import WebApp from "@twa-dev/sdk";
import { debounce } from "lodash";

export default function Search() {
	const url = import.meta.env.VITE_API_URL;
	const [data, setData] = useState<any>();
	const [search, setSearch] = useState<string>("");

	const debouncedSearch = debounce((title: string) => {
		axios.get(`${url}/search`, {
			headers: {
				authorization: `tma ${WebApp.initData}`,
			},
			params: {
				title,
			},
		}).then((res) => {
			setSearch(res.data);
		}).catch((error) => {
			console.error(error);
		});
	})

	const handleSearch = (e: React.ChangeEvent<HTMLInputElement>) => {
		setSearch(e.target.value);
		debouncedSearch(search);
	}

	return (
		<Container py={"20px"} px={"25px"} position={"relative"} mt={"75px"}>
			<Input
				placeholder={"Search for manga..."}
				variant={"subtle"}
				size={"lg"}
				onChange={handleSearch}
			/>
			{data && data.length ? (
				<Gallery data={data} hasSearch />
			) : (
				<Empty hasSearch />
			)}
		</Container>
	);
}

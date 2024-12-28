import Header from "@/components/Header";
import Navbar from "@/components/Navbar";
import axios from "axios";
import { useEffect, useState } from "react";
import Empty from "@/components/Empty";

interface Manga {
	manga_url: number;
	cover_image: string;
}

export default function Library() {
	const [data, setData] = useState<Array<Manga>>([]);

	useEffect(() => {
		const fetchLibrary = async () => {
			const url = import.meta.env.VITE_API_URL;

			try {
				const res = await axios.get(`${url}/library`)
				setData(res.data);
				if (!res.data || res.data.length === 0) {
					document.body.style.height = "100vh";
					document.body.style.overflow = "hidden";
				}
			} catch (error) {
				console.error(error);
			}
		};

		fetchLibrary();
		console.log(data);
	}, []);

	return (
		<>
			<Header name={"Library"} />
			<Empty />
			<Navbar navs={["Library", "Browse", "History"]} />
		</>
	);
}

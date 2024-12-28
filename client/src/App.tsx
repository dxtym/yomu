import "./App.css";
import axios from "axios";
import WebApp from "@twa-dev/sdk";
import Library from "./pages/Library";
import { useEffect, useState } from "react";
import { BrowserRouter, Route, Routes } from "react-router-dom";
import Browse from "./pages/Browse";

interface UserData {
	id: number;
	first_name: string;
}

export default function App() {
	const url = import.meta.env.VITE_API_URL;
	const [data, setData] = useState<UserData>();
	const [user, setUser] = useState<boolean>(false);

	useEffect(() => {
		const cached = localStorage.getItem("user");
		if (cached) {
			setData(JSON.parse(cached) as UserData);
			setUser(true);
		} else {
			const currentUser = WebApp.initDataUnsafe.user as UserData;
			if (currentUser) {
				setData(currentUser);
			}
		}
	}, []);

	useEffect(() => {
		const createUser = async () => {
			if (data && !user) {
				try {
					const res = await axios.post(`${url}/user`, data);
					if (res.status == 200) {
						setUser(true);
						localStorage.setItem("user", JSON.stringify(data));
					}
				} catch (error) {
					console.error("cannot create user:", error);
				}
			}
		};

		createUser();
		console.log(user);
	}, [data, user]);

	return (
		<BrowserRouter>
			<Routes>
				<Route path="/" element={<Library />} />
				<Route path="/browse" element={<Browse />} />
			</Routes>
		</BrowserRouter>
	);
}

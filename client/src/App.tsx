import "./App.css";
import axios from "axios";
import WebApp from "@twa-dev/sdk";
import Library from "./pages/Library";
import { useEffect, useState } from "react";
import { BrowserRouter, Route, Routes } from "react-router-dom";
import Browse from "./pages/Browse";
import { terminal } from "virtual:terminal";

interface UserData {
  id: number;
  first_name: string;
  last_name?: string;
  user_name?: string;
}

export default function App() {
  // TODO: use with telegram web client
  // const [userData, setUserData] = useState<UserData>();
  // useEffect(() => {
  //   if (WebApp.initDataUnsafe.user) {
  //     setUserData(WebApp.initDataUnsafe.user as UserData);
  //   }
  // }, []);

  const userData = {
    id: 2,
    first_name: "Joe",
  } as UserData;

  useEffect(() => {
    const createUser = async () => {
      const url = import.meta.env.VITE_API_URL as string;

      try {
        const res = await axios.post(`${url}/user`, userData);
        localStorage.setItem("token", res.data.token);
        terminal.log(res.data);
      } catch (error) {
        terminal.error(error);
      }
    };

    if (userData) {
      createUser();
    }
  }, [userData]);

  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<Library />} />
        <Route path="/library" element={<Library />} />
        <Route path="/browse" element={<Browse />} />
      </Routes>
    </BrowserRouter>
  );
}

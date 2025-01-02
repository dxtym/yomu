import "./App.css";
import axios from "axios";
import WebApp from "@twa-dev/sdk";
import Library from "./pages/Library";
import { useEffect, useState } from "react";
import { BrowserRouter, Route, Routes } from "react-router-dom";
import Browse from "./pages/Browse";
import Manga from "./pages/Manga";
import { IUserData } from "./types";
import View from "./pages/View";

export default function App() {
  const url = import.meta.env.VITE_API_URL;
  const [data, setData] = useState<IUserData>();
  const [user, setUser] = useState<boolean>(false);

  useEffect(() => {
    const cached = localStorage.getItem("user");
    if (cached) {
      setData(JSON.parse(cached) as IUserData);
      setUser(true);
    } else {
      const curr = WebApp.initDataUnsafe.user as IUserData;
      if (curr) {
        setData(curr);
      }
    }
  }, []);

  useEffect(() => {
    const createUser = async () => {
      if (data && !user) {
        axios
          .post(`${url}/user`, data)
          .then((res) => {
            if (res.status == 200) {
              setUser(true);
              localStorage.setItem("user", JSON.stringify(data));
            }
          })
          .catch((err) => console.error(err));
      }
    };

    createUser();
  }, [data, user]);

  return (
    <BrowserRouter>
      <Routes>
        {["/", "/library"].map((path, index) => {
          return <Route path={path} element={<Library />} key={index} />;
        })}
        <Route path="/browse" element={<Browse />} />
        <Route path="/browse/:url" element={<Manga />} />
        <Route path="/chapter/:url/:id" element={<View />} />
      </Routes>
    </BrowserRouter>
  );
}

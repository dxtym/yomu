import "./App.css";
import axios from "axios";
import WebApp from "@twa-dev/sdk";
import Library from "./pages/Library/Library";
import { useEffect } from "react";
import { BrowserRouter, Navigate, Route, Routes } from "react-router-dom";
import Browse from "./pages/Browse/Browse";
import Manga from "./pages/Manga/Manga";
import { IUser } from "./types";
import View from "./pages/View/View";
import { useLocalStorage } from "usehooks-ts";

const App = () => {
  const url = import.meta.env.VITE_API_URL;
  const [user, setUser] = useLocalStorage("user", "");

  useEffect(() => {
    if (!user) {
      const curr = WebApp.initDataUnsafe.user as IUser;
      setUser(JSON.stringify(curr));
      axios
        .post(`${url}/user`, curr)
        .then((res) => console.log(res))
        .catch((err) => console.error(err));
    }
  }, []);

  return (
    <BrowserRouter>
      <Routes>
        <Route path={"/"} element={<Navigate to={"/library"} />} />
        <Route path={"/library"} element={<Library />} />
        <Route path={"/browse"} element={<Browse />} />
        <Route path={"/browse/:url"} element={<Manga />} />
        <Route path={"/chapter/:url/:id"} element={<View />} />
      </Routes>
    </BrowserRouter>
  );
};

export default App;

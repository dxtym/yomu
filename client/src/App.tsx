import "./App.css";
import WebApp from "@twa-dev/sdk";
import UserService from "./api/user";

import View from "./pages/View/View";
import Manga from "./pages/Manga/Manga";
import Browse from "./pages/Browse/Browse";
import Library from "./pages/Library/Library";
import History from "./pages/History/History";

import { useEffect } from "react";
import { IUser } from "./types/user";
import { useLocalStorage } from "usehooks-ts";
import { BrowserRouter, Navigate, Route, Routes } from "react-router-dom";

export default function App() {
  const [user, setUser] = useLocalStorage<boolean>("user", true);

  useEffect(() => {
    if (!user) {
      const data = WebApp.initDataUnsafe.user as IUser;
      setUser(true);
      UserService.createUser(data)
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
        <Route path={"/history"} element={<History />} />
        <Route path={"/browse/:manga"} element={<Manga />} />
        <Route path={"/chapter/:manga/:chapter"} element={<View />} />
      </Routes>
    </BrowserRouter>
  );
}

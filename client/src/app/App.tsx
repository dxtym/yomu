import "./App.css";

import View from "@/pages/View/View";
import Manga from "@/pages/Manga/Manga";
import Browse from "@/pages/Browse/Browse";
import Library from "@/pages/Library/Library";
import History from "@/pages/History/History";
import Loading from "@/components/common/Loading";

import { Suspense } from "react";
import {
  Route,
  RouterProvider,
  createBrowserRouter,
  createRoutesFromElements,
} from "react-router-dom";

export default function App() {
  const router = createBrowserRouter(
    createRoutesFromElements(
      <>
        <Route path="/" element={<Library />} />
        <Route path="/library" element={<Library />} />
        <Route path="/browse" element={<Browse />} />
        <Route path="/history" element={<History />} />
        <Route path="/browse/:manga" element={<Manga />} />
        <Route path="/chapter/:manga/:chapter" element={<View />} />
      </>,
    ),
  );

  return (
    <Suspense fallback={<Loading />}>
      <RouterProvider router={router} />
    </Suspense>
  );
}

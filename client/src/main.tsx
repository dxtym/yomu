import "./index.css";
import App from "@/app/App";

import { Telegram } from "@twa-dev/types";
import { createRoot } from "react-dom/client";
import { Provider } from "@/components/ui/provider";

declare global {
  interface Window {
    Telegram: Telegram;
  }
}

createRoot(document.getElementById("root")!).render(
  <Provider>
    <App />
  </Provider>,
);

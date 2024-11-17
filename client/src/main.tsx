import "./index.css";
import App from "./App";
import { createRoot } from "react-dom/client";
import React from "react";

createRoot(document.getElementById("root")!).render(
  <React.StrictMode>
    <App />
  </React.StrictMode>,
);

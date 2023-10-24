import React from "react";
import ReactDOM from "react-dom";
import { MuiThemeProvider } from "@material-ui/core";
import "./index.css";
import theme from "../src/themes/theme";
import Navbar from "./components/Navbar/Navbar";

ReactDOM.render(
  <React.StrictMode>
    <MuiThemeProvider theme={theme}>
      <Navbar />
    </MuiThemeProvider>
  </React.StrictMode>,
  document.getElementById("root")
);

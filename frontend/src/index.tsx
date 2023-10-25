import React from "react";
import ReactDOM from "react-dom";
import { MuiThemeProvider } from "@material-ui/core";
import "./index.css";
import theme from "../src/themes/theme";
import Navbar from "./components/Navbar/Navbar";
import { QueryClient, QueryClientProvider, setLogger } from 'react-query';

const queryClient = new QueryClient({
  defaultOptions: {
    queries: {
      retry: (failureCount, error) => {
        if (error === 404) return false;
        else if (failureCount < 2) return true;
        else return false;
      }
    },
  },
});


ReactDOM.render(
  <React.StrictMode>
    <QueryClientProvider client={queryClient}>
      <MuiThemeProvider theme={theme}>
        <Navbar />
      </MuiThemeProvider>
    </QueryClientProvider>
  </React.StrictMode>,
  document.getElementById("root")
);

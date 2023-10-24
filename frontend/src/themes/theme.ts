import { createMuiTheme } from "@material-ui/core";

const theme = createMuiTheme({
  typography: {
    fontFamily: '"Open Sans", "sans-serif", "Roboto", "Permanent Marker"',
    fontSize: 12,
    button: {
      textTransform: "none",
      fontWeight: 700,
    },
  },
  palette: {
    primary: { main: "#92CE8B", contrastText: "#304A32", light: "#D4FCE5" },
    secondary: { main: "#A8E4C2", contrastText: "#304A32" },
    info: { main: "#F4F6FF" },
    text: {
      primary: "#000000",
      secondary: "#9BA9CC",
    },
  },
  shape: {
    borderRadius: 5,
  },
});

export default theme;

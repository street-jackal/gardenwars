import { makeStyles, createStyles, Theme } from "@material-ui/core/styles";

const loginModalStyles = makeStyles((theme: Theme) =>
  createStyles({
    modal: {
      backgroundColor: theme.palette.background.paper,
      color: theme.palette.primary.dark,
      padding: "16px",
      borderRadius: "8px",
      position: "absolute",
      top: "50%",
      left: "50%",
      transform: "translate(-50%, -50%)",
    },
    button: {
      backgroundColor: theme.palette.primary.main,
      color: theme.palette.primary.dark,
    },
    message: {
      display: "flex",
      justifyContent: "center",
      alignItems: "center",
      alignSelf: "center",
    },
  })
);

export default loginModalStyles;

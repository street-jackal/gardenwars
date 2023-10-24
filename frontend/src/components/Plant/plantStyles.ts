import { makeStyles, createStyles, Theme } from "@material-ui/core/styles";

const plantStyles = makeStyles((theme: Theme) =>
  createStyles({
    root: {
      flexGrow: 1,
      maxWidth: "100%",
      margin: "2%",
    },
    paper: {
      padding: theme.spacing(2),
      textAlign: "left",
    },
    title: {
      backgroundColor: theme.palette.primary.light,
      color: theme.palette.primary.contrastText,
    },
    media: {
      height: 0,
      paddingTop: "56.25%", // 16:9
    },
    expand: {
      transform: "rotate(0deg)",
      marginLeft: "auto",
      transition: theme.transitions.create("transform", {
        duration: theme.transitions.duration.shortest,
      }),
    },
    expandOpen: {
      transform: "rotate(180deg)",
    },
    avatar: {
      backgroundColor: theme.palette.secondary.main,
    },
    cardContent: {
      padding: "16px 16px 0 16px",
    },
    divider: {
      marginTop: 15,
      marginBottom: 10,
    },
  })
);

export default plantStyles;

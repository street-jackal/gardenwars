import {
  makeStyles,
  createStyles,
  Theme,
} from "@material-ui/core/styles";

const drawerStyles = makeStyles((theme: Theme) =>
  createStyles({
    grow: {
      flexGrow: 1,
    },
    menuButton: {
      marginRight: theme.spacing(2),
      marginLeft: theme.spacing(2),
      scale: 1.6,
    },
  })
);

export default drawerStyles;

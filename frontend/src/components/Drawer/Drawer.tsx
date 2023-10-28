import * as React from "react";
import Box from "@mui/material/Box";
import SwipeableDrawer from "@mui/material/SwipeableDrawer";
import Button from "@mui/material/Button";
import List from "@mui/material/List";
import Divider from "@mui/material/Divider";
import ListItem from "@mui/material/ListItem";
import ListItemButton from "@mui/material/ListItemButton";
import ListItemIcon from "@mui/material/ListItemIcon";
import ListItemText from "@mui/material/ListItemText";
import InboxIcon from "@mui/icons-material/MoveToInbox";
import MailIcon from "@mui/icons-material/Mail";
import MenuIcon from "@material-ui/icons/Menu";
import StarIcon from "@mui/icons-material/Star";
import ForestIcon from '@mui/icons-material/Forest';
import drawerStyles from "./drawerStyles";

const Anchor = "left";

interface DrawerProps {
  showUserPlants: (show: boolean) => void;
}

export default function SwipeableTemporaryDrawer({ showUserPlants }: DrawerProps): JSX.Element {
  const [state, setState] = React.useState(false);

  const classes = drawerStyles();

  const toggleDrawer = (open: boolean) => (event: React.KeyboardEvent | React.MouseEvent) => {
    if (
      event &&
      event.type === "keydown" &&
      ((event as React.KeyboardEvent).key === "Tab" || (event as React.KeyboardEvent).key === "Shift")
    ) {
      return;
    }

    setState(open);
  };

  const list = () => (
    <Box sx={{ width: 250 }} role="presentation" onClick={toggleDrawer(false)} onKeyDown={toggleDrawer(false)}>
      <List>
        <ListItem disablePadding>
          <ListItemButton>
            <ListItemIcon>
              <InboxIcon />
            </ListItemIcon>
            <ListItemText primary="Inbox" />
          </ListItemButton>
        </ListItem>

        <ListItem disablePadding>
          <ListItemButton onClick={() => showUserPlants(true)}>
            <ListItemIcon>
              <StarIcon />
            </ListItemIcon>
            <ListItemText primary="Favorites" />
          </ListItemButton>
        </ListItem>

        <ListItem disablePadding>
          <ListItemButton onClick={() => showUserPlants(false)}>
            <ListItemIcon>
              <ForestIcon />
            </ListItemIcon>
            <ListItemText primary="Plant Library" />
          </ListItemButton>
        </ListItem>

        <ListItem disablePadding>
          <ListItemButton>
            <ListItemIcon>
              <MailIcon />
            </ListItemIcon>
            <ListItemText primary="Send email" />
          </ListItemButton>
        </ListItem>
      </List>

      <Divider />

      <List>
        {["All mail", "Trash", "Spam"].map((text, index) => (
          <ListItem key={text} disablePadding>
            <ListItemButton>
              <ListItemIcon>{index % 2 === 0 ? <InboxIcon /> : <MailIcon />}</ListItemIcon>
              <ListItemText primary={text} />
            </ListItemButton>
          </ListItem>
        ))}
      </List>
    </Box>
  );

  return (
    <div>
      <React.Fragment>
        <Button color="inherit" onClick={toggleDrawer(true)}>
          <MenuIcon className={classes.menuButton} />
        </Button>
        <SwipeableDrawer anchor={Anchor} open={state} onClose={toggleDrawer(false)} onOpen={toggleDrawer(true)}>
          {list()}
        </SwipeableDrawer>
      </React.Fragment>
    </div>
  );
}

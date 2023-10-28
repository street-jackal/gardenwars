import { MouseEvent, useState } from "react";

import { AppBar, Toolbar, IconButton, Typography, InputBase, Badge, MenuItem, Menu } from "@material-ui/core/";
import SearchIcon from "@material-ui/icons/Search";
import AccountCircle from "@material-ui/icons/AccountCircle";
import MailIcon from "@material-ui/icons/Mail";
import NotificationsIcon from "@material-ui/icons/Notifications";
import navbarStyles from "./navbarStyles";
import CardContainer from "../CardContainer/CardContainer";
import SwipeableTemporaryDrawer from "../Drawer/Drawer";
import LoginModal from "../LoginModal/LoginModal";
import { IUser } from "../../interface/IUser";
import LoggedInUserPlants from "../CardContainer/LoggedInUserPlants";

const Navbar = () => {
  const classes = navbarStyles();
  const [anchorEl, setAnchorEl] = useState<null | HTMLElement>(null);
  const [search, setSearch] = useState("");
  const [loggedInUser, setLoggedInUser] = useState<IUser | null>(null);
  const [showLoginModal, setShowLoginModal] = useState(false);
  const [showUserPlants, setShowUserPlants] = useState(false);
  const isMenuOpen = Boolean(anchorEl);

  const handleProfileMenuOpen = (event: MouseEvent<HTMLElement>) => {
    setAnchorEl(event.currentTarget);
  };

  const handleMenuClose = () => {
    setAnchorEl(null);
  };

  const handleShowUserPlants = (show: boolean) => {
    setShowUserPlants(show);
  };

  const menuId = "primary-search-account-menu";
  const renderMenu = (
    <Menu
      anchorEl={anchorEl}
      anchorOrigin={{ vertical: "top", horizontal: "right" }}
      id={menuId}
      keepMounted
      transformOrigin={{ vertical: "top", horizontal: "right" }}
      open={isMenuOpen}
      onClose={handleMenuClose}
    >
      {loggedInUser ? (
        <div>
          <MenuItem onClick={handleMenuClose}>Profile</MenuItem>
          <MenuItem onClick={handleMenuClose}>{`Logout ${loggedInUser.email}`}</MenuItem>
        </div>
      ) : (
        <div>
          <MenuItem
            onClick={() => {
              handleMenuClose();
              setShowLoginModal(true);
            }}
          >
            Login
          </MenuItem>
          <LoginModal
            open={showLoginModal}
            onClose={() => setShowLoginModal(false)}
            setLoggedInUser={setLoggedInUser}
          />
        </div>
      )}
    </Menu>
  );

  function debounce(func: any, timeout = 1000) {
    let timer: NodeJS.Timeout;
    return (...args: any) => {
      timer && clearTimeout(timer);
      timer = setTimeout(() => func(...args), timeout);
    };
  }

  return (
    <>
      <div className={classes.grow}>
        <AppBar position="sticky">
          <Toolbar>
            <SwipeableTemporaryDrawer aria-label="open drawer" showUserPlants={(handleShowUserPlants)}/>
            <Typography className={classes.title} variant="h6" noWrap>
              Garden Wars
            </Typography>
            <div className={classes.search}>
              <div className={classes.searchIcon}>
                <SearchIcon />
              </div>
              <InputBase
                placeholder="Search…"
                classes={{
                  root: classes.inputRoot,
                  input: classes.inputInput,
                }}
                inputProps={{ "aria-label": "search" }}
                onChange={(e) => setSearch(e.target.value)}
              />
            </div>
            <div className={classes.grow} />
            <div className={classes.sectionDesktop}>
              <IconButton aria-label="show 4 new mails" color="inherit">
                <Badge badgeContent={4} color="secondary">
                  <MailIcon />
                </Badge>
              </IconButton>
              <IconButton aria-label="show 17 new notifications" color="inherit">
                <Badge badgeContent={17} color="secondary">
                  <NotificationsIcon />
                </Badge>
              </IconButton>
              <IconButton
                edge="end"
                aria-label="account of current user"
                aria-controls={menuId}
                aria-haspopup="true"
                onClick={handleProfileMenuOpen}
                color="inherit"
              >
                <AccountCircle />
              </IconButton>
            </div>
          </Toolbar>
        </AppBar>
        {renderMenu}
      </div>
      {showUserPlants && loggedInUser ? (
        <LoggedInUserPlants searchValue={search} loggedInUser={loggedInUser} />
      ) : (
        <CardContainer searchValue={search} loggedInUser={loggedInUser} />
      )}
    </>
  );
};

export default Navbar;

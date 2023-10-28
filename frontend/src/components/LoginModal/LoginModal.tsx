import React, { useState } from "react";
import { Button, Modal, Box, Typography, TextField, Container } from "@mui/material";
import loginModalStyles from "./loginModalStyles";
import { signupUser, loginUser } from "../../api/users";
import { Grid } from "@material-ui/core";
import { AxiosError, AxiosResponse } from "axios";
import { BaseResponse } from "../../api/types/baseResponse";
import { IUser } from "../../interface/IUser";

interface LoginModalProps {
  open: boolean;
  onClose: () => void;
  setLoggedInUser: (value: IUser|null) => void;
}

const LoginModal: React.FC<LoginModalProps> = ({ open, onClose, setLoggedInUser }) => {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [message, setMessage] = useState("");

  const classes = loginModalStyles();

  const handleLogin = async () => {
    try {
      const response = await loginUser({ email: email, password: password });
      if (response.status === 200) {
        setMessage("Successfully logged in!");
        setLoggedInUser(response.data as IUser);
        setTimeout(onClose, 500);
      }
    } catch (response) {
      setMessage(((response as AxiosError).response?.data as BaseResponse).message);
    }
  };

  const handleSignup = async () => {
    try {
      const response = await signupUser({ email: email, password: password });
      if (response.success) {
        setMessage(response.message);
        setTimeout(onClose, 500);
      }
    } catch (response) {
      setMessage(((response as AxiosError).response?.data as BaseResponse).message);
    }
  };

  return (
    <Modal open={open} onClose={onClose}>
      <Container maxWidth="sm">
        <Box className={classes.modal}>
          <Typography variant="h4">Login</Typography>
          <TextField
            label="Email"
            variant="outlined"
            fullWidth
            value={email}
            onChange={(e) => setEmail(e.target.value)}
            margin="normal"
          />
          <TextField
            label="Password"
            variant="outlined"
            type="password"
            fullWidth
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            margin="normal"
          />
          <Grid container spacing={2}>
            <Grid item xs={6}>
              <Button variant="contained" color="primary" onClick={handleLogin} fullWidth>
                Login
              </Button>
            </Grid>
            <Grid item xs={6}>
              <Button variant="contained" color="success" onClick={handleSignup} fullWidth>
                Sign Up
              </Button>
            </Grid>
            <Grid item xs={12}>
              <Typography className={classes.message}>{message}</Typography>
            </Grid>
          </Grid>
        </Box>
      </Container>
    </Modal>
  );
};

export default LoginModal;

import React, { useState } from "react";
import { useNavigate } from "react-router-dom";
import { TextField, Button, Box, Typography } from "@mui/material";
import { login } from "../service/auth";

function LoginForm() {
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const navigate = useNavigate();

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      await login(username, password);
      navigate("/cases");
    } catch (err) {
      alert("Login failed");
    }
  };

  return (
    <Box
      display="flex"
      flexDirection="column"
      width={300}
      margin="auto"
      mt={10}
      p={4}
      boxShadow={2}
      borderRadius={2}
    >
      <Typography variant="h5" mb={2} align="center">
        Agent Login
      </Typography>
      <TextField
        label="Username"
        value={username}
        onChange={(e) => setUsername(e.target.value)}
        margin="normal"
      />
      <TextField
        label="Password"
        type="password"
        value={password}
        onChange={(e) => setPassword(e.target.value)}
        margin="normal"
      />
      <Button variant="contained" color="primary" onClick={handleSubmit}>
        Login
      </Button>
    </Box>
  );
}

export default LoginForm;

import { useState } from "react";
import Button from "@mui/material/Button";
import TextField from "@mui/material/TextField";
import { useMutation } from "@tanstack/react-query";
import { instance } from "../../axios";
import { Box, Container, Stack, Typography } from "@mui/material";

async function login(loginParams: { username: string; password: string }) {
  console.log("login");
  const response = await instance.post("/login", loginParams);
  return response.data;
}

export default function LoginPage() {
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");

  const mutation = useMutation({
    mutationFn: login,
  });

  const handleLogin = () => {
    mutation.mutate({ username, password });
  };

  return (
    <Container maxWidth="sm">
      <Box
        display="flex"
        flexDirection="column"
        alignItems="center"
        justifyContent="center"
        style={{
          minHeight: "95vh",
        }}
      >
        <h2 style={{}}>Login</h2>
        <Stack
          spacing={2}
          direction="column"
          style={{
            minWidth: "55%",
          }}
        >
          <TextField
            label="Username"
            variant="outlined"
            value={username}
            onChange={(e) => setUsername(e.target.value)}
          />
          <TextField
            label="Password"
            variant="outlined"
            type="password"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
          />
          <Button
            variant="outlined"
            onClick={handleLogin}
            style={{ height: "3rem" }}
          >
            Login
          </Button>
        </Stack>
      </Box>
      <Box
        style={{
          position: "absolute",
          bottom: 5,
          left: "50%",
          transform: "translateX(-50%)",
        }}
      >
        {mutation.isPending ? (
          <Typography variant="h6" color="textSecondary">
            Loading...
          </Typography>
        ) : mutation.isError ? (
          <Typography variant="h6" color="error">
            An error occurred: {mutation.error.message}
          </Typography>
        ) : mutation.isSuccess ? (
          <Typography variant="h6" color="primary">
            Login Successful
          </Typography>
        ) : null}
      </Box>
    </Container>
  );
}

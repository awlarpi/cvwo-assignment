import { useState } from "react";
import Button from "@mui/material/Button";
import TextField from "@mui/material/TextField";
import { useMutation } from "@tanstack/react-query";
import { instance } from "../../lib/axiosinstance";
import { Box, Container, Stack, Typography } from "@mui/material";
import { useStore } from "../../lib/store";
import { Link, Navigate } from "react-router-dom";

export default function LoginPage() {
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const { logIn, isLoggedIn } = useStore();

  const mutation = useMutation({
    mutationFn: async (loginParams: { username: string; password: string }) => {
      const response = await instance.post("/login", loginParams);
      return response.data;
    },
  });

  const handleLogin = () => {
    mutation.mutate(
      { username, password },
      {
        onSuccess: (data) => {
          logIn(data.session_id);
        },
      }
    );
  };

  if (isLoggedIn) {
    return <Navigate to="/" replace={true} />;
  }

  return (
    <Container maxWidth="sm">
      <Box
        display="flex"
        flexDirection="column"
        alignItems="center"
        justifyContent="center"
        style={{
          minHeight: "75vh",
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
          <Button
            component={Link}
            to="/signup"
            variant="outlined"
            style={{ height: "3rem" }}
            color="secondary"
          >
            Signup
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

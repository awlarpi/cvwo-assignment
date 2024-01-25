import React, { useState } from "react";
import {
  TextField,
  Button,
  Container,
  Typography,
  Stack,
  Divider,
} from "@mui/material";
import { instance } from "../../lib/axiosinstance";
import { Navigate } from "react-router-dom";
import { useStore } from "../../lib/store";

type SignupFormInputs = {
  username: string;
  email: string;
  password: string;
  profilePicture: string;
  biography: string;
};

async function signup(signupParams: SignupFormInputs) {
  const response = await instance.post("/users", signupParams);
  return response.data;
}

export default function SignupPage() {
  const { isLoggedIn } = useStore();

  const [values, setValues] = useState({
    username: "",
    email: "",
    password: "",
    profilePicture: "",
    biography: "",
  });

  const [errors, setErrors] = useState({
    username: "",
    email: "",
    password: "",
  });

  const [isSignedUp, setIsSignedUp] = useState(false);

  const handleChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    setValues({
      ...values,
      [event.target.name]: event.target.value,
    });
  };

  const handleSubmit = async (event: React.FormEvent) => {
    event.preventDefault();

    const newErrors = {
      username: "",
      email: "",
      password: "",
    };

    if (!values.username) newErrors.username = "Username is required";
    if (!values.email) newErrors.email = "Email is required";
    else if (!/\S+@\S+\.\S+/.test(values.email))
      newErrors.email = "Invalid email";
    if (!values.password) newErrors.password = "Password is required";

    setErrors(newErrors);

    if (!newErrors.username && !newErrors.email && !newErrors.password) {
      await signup(values);
      setIsSignedUp(true);
    }
  };

  if (isLoggedIn) {
    return <Navigate to="/" replace={true} />;
  }

  if (isSignedUp) {
    return <Navigate to="/login" replace={true} />;
  }

  return (
    <Container
      maxWidth="xs"
      sx={{
        mt: "5rem",
      }}
    >
      <Typography component="h1" variant="h5" marginBottom="0.5rem">
        Signup
      </Typography>
      <Divider />
      <br />
      <form onSubmit={handleSubmit}>
        <Stack spacing="1rem">
          <TextField
            name="username"
            label="Username"
            error={Boolean(errors.username)}
            helperText={errors.username || "*Enter your username (required)"}
            onChange={handleChange}
            fullWidth
          />
          <TextField
            name="email"
            label="Email"
            error={Boolean(errors.email)}
            helperText={errors.email || "*Enter your email (required)"}
            onChange={handleChange}
            fullWidth
          />
          <TextField
            type="password"
            name="password"
            label="Password"
            error={Boolean(errors.password)}
            helperText={errors.password || "*Enter your password (required)"}
            onChange={handleChange}
            fullWidth
          />
          <TextField
            name="profilePicture"
            label="Profile Picture URL"
            helperText="Enter the URL of your profile picture"
            onChange={handleChange}
            fullWidth
          />
          <TextField
            name="biography"
            label="Biography"
            helperText="Tell us a little about yourself"
            onChange={handleChange}
            fullWidth
          />
          <Button type="submit" fullWidth variant="contained" color="primary">
            Sign Up
          </Button>
        </Stack>
      </form>
    </Container>
  );
}

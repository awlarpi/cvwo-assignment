import React, { useState } from "react";
import { useMutation } from "@tanstack/react-query";
import { queryClient } from "../main";
import { instance } from "../lib/axiosinstance";
import {
  TextField,
  Box,
  Button,
  Card,
  CardActions,
  CardContent,
  Stack,
} from "@mui/material";

interface CreatePostApiParams {
  Title: string;
  Content: string;
}

export default function CreatePostForm() {
  const [title, setTitle] = useState("");
  const [content, setContent] = useState("");
  const [error, setError] = useState("");

  const mutation = useMutation({
    mutationFn: (params: CreatePostApiParams) =>
      instance.post("/posts", params),
    onSuccess: () => {
      setTitle("");
      setContent("");
      setError("");
      queryClient.invalidateQueries({
        queryKey: ["posts"],
      });
    },
    onError: (error: Error) => {
      setError(error.message);
    },
  });

  const handleSubmit = (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    if (content.trim() === "") {
      setError("Content cannot be empty");
    } else if (title.trim() === "") {
      setError("Title cannot be empty");
    } else {
      setError("");
      mutation.mutate({ Title: title, Content: content });
    }
  };

  return (
    <form onSubmit={handleSubmit}>
      <Card
        variant="elevation"
        sx={{ paddingY: "0.5rem", paddingX: "0.75rem" }}
      >
        <CardContent sx={{ padding: "0rem", mb: "0.5rem" }}>
          <Stack spacing="1rem">
            <TextField
              value={title}
              onChange={(e) => setTitle(e.target.value)}
              placeholder="Add title"
              multiline
              fullWidth
              variant="standard"
              error={error !== ""}
              helperText={error + " - Please check that you are logged in"}
              sx={{ borderColor: "primary.main" }}
            />
            <TextField
              value={content}
              onChange={(e) => setContent(e.target.value)}
              placeholder="Add content"
              multiline
              fullWidth
              variant="standard"
              error={error !== ""}
              helperText={error + " - Please check that you are logged in"}
              sx={{ borderColor: "primary.main" }}
            />
          </Stack>
        </CardContent>
        <CardActions sx={{ padding: "0rem" }}>
          <Box
            sx={{
              width: "100%",
              display: "flex",
              justifyContent: "flex-end",
              marginTop: "0.5rem",
            }}
          >
            <Button
              type="submit"
              variant="outlined"
              color="primary"
              size="small"
            >
              Submit
            </Button>
          </Box>
        </CardActions>
      </Card>
    </form>
  );
}

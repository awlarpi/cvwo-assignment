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
} from "@mui/material";

interface CreateCommentApiParams {
  Content: string;
  PostID: number;
}

export default function CommentForm({ postId }: { postId: string }) {
  const [content, setContent] = useState("");
  const [error, setError] = useState("");

  const mutation = useMutation({
    mutationFn: (params: CreateCommentApiParams) =>
      instance.post("/comments", params),
    onSuccess: () => {
      setContent("");
      setError("");
      queryClient.invalidateQueries({
        queryKey: ["post", postId, "comments"],
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
    } else {
      setError("");
      mutation.mutate({ Content: content, PostID: +postId });
    }
  };

  return (
    <form onSubmit={handleSubmit}>
      <Card
        variant="elevation"
        sx={{ paddingY: "0.5rem", paddingX: "0.75rem" }}
      >
        <CardContent sx={{ padding: "0rem", mb: "0.5rem" }}>
          <TextField
            value={content}
            onChange={(e) => setContent(e.target.value)}
            placeholder="Add a comment"
            multiline
            fullWidth
            variant="standard"
            error={error !== ""}
            helperText={error + " - Please check that you are logged in"}
            sx={{ borderColor: "primary.main" }}
          />
        </CardContent>
        <CardActions sx={{ padding: "0rem" }}>
          <Box
            sx={{ width: "100%", display: "flex", justifyContent: "flex-end" }}
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

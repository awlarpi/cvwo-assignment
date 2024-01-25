import {
  Box,
  Card,
  CardActionArea,
  CardContent,
  Divider,
  IconButton,
  Stack,
  Typography,
} from "@mui/material";
import { Link } from "react-router-dom";
import { useMutation, useQuery } from "@tanstack/react-query";
import { instance } from "../lib/axiosinstance";
import { queryClient } from "../main";
import DeleteIcon from "@mui/icons-material/Delete";

export type Post = {
  PostID: number;
  Title: string;
  Content: string;
  CreationDate: string;
  UserID: number;
  IsSticky: boolean;
  IsLocked: boolean;
  PostCategoryID: number;
  AdditionalNotes: string | null;
};

async function getAllPosts() {
  const response = await instance.get<Post[]>("/posts");
  return response.data;
}

export default function PostCards() {
  const { isLoading, data, isError, error } = useQuery({
    queryKey: ["posts"],
    queryFn: getAllPosts,
  });

  const deletePostMutation = useMutation({
    mutationFn: async (postId: number) => {
      const response = await instance.delete(`/posts/${postId}`);
      return response.data;
    },
    onSuccess: () => {
      queryClient.invalidateQueries({
        queryKey: ["posts"],
      });
    },
  });

  if (isLoading) {
    return <div>Loading...</div>;
  }

  if (isError) {
    return <div>An error has occurred: {error.message}</div>;
  }

  if (!data) {
    return <div>No data available</div>;
  }

  return (
    <Stack spacing="1rem">
      {data.map((post) => (
        <Card key={post.PostID}>
          <CardActionArea component={Link} to={`/posts/${post.PostID}`}>
            <CardContent>
              <Typography variant="h5" component="h2">
                {post.Title}
              </Typography>
              <Typography variant="body2" color="text.secondary">
                {post.Content}
              </Typography>
            </CardContent>
          </CardActionArea>
          <Divider />
          <Box
            display="flex"
            alignItems="center"
            justifyContent="space-between"
            marginY="0.5rem"
            marginX="0.75rem"
          >
            <Stack direction="row" spacing="1.5rem">
              <Typography variant="body2" color="text.secondary">
                User: {post.UserID}
              </Typography>
              <Typography variant="body2" color="text.secondary">
                Date: {new Date(post.CreationDate).toLocaleDateString()}
              </Typography>
            </Stack>
            <IconButton
              color="secondary"
              sx={{
                border: "1px solid",
                borderColor: "currentColor",
                borderRadius: "10%",
              }}
              size="small"
              onClick={() => deletePostMutation.mutate(post.PostID)}
            >
              <DeleteIcon />
            </IconButton>
          </Box>
        </Card>
      ))}
    </Stack>
  );
}

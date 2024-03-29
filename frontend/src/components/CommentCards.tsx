import { instance } from "../lib/axiosinstance";
import { useMutation, useQuery } from "@tanstack/react-query";
import {
  Card,
  CardContent,
  Typography,
  Stack,
  IconButton,
  Divider,
  Box,
} from "@mui/material";
import { queryClient } from "../main";
import DeleteIcon from "@mui/icons-material/Delete";
import { useStore } from "../lib/store";

interface Comment {
  CommentID: number;
  Content: string;
  CreationDate: string;
  PostID: number;
  UserID: number;
  Username: string;
}

export default function CommentCards({ postId }: { postId: string }) {
  const { userId, isLoggedIn } = useStore();

  const { isLoading, data, isError, error } = useQuery({
    queryKey: ["post", postId, "comments"],
    queryFn: async () => {
      const response = await instance.get<Comment[]>(
        `/comments/post/${postId}`
      );
      return response.data;
    },
  });

  const deleteCommentMutation = useMutation({
    mutationFn: async (commentId: number) => {
      const response = await instance.delete(`/comments/${commentId}`);
      return response.data;
    },
    onSuccess: () => {
      queryClient.invalidateQueries({
        queryKey: ["post", postId, "comments"],
      });
    },
  });

  if (isLoading) {
    return <div>Loading...</div>;
  }

  if (isError) {
    return <div>An error has occurred: {error?.message}</div>;
  }

  if (!data) {
    return <div>No data available</div>;
  }

  return (
    <Stack spacing={2}>
      {data.map((comment) => (
        <Card key={comment.CommentID} sx={{}}>
          <CardContent
            sx={{
              paddingY: "0.5rem",
              display: "flex",
              justifyContent: "space-between",
              alignItems: "center",
            }}
          >
            <Typography variant="body1" color="textSecondary" component="p">
              {comment.Content}
            </Typography>
            {isLoggedIn && userId === comment.UserID ? (
              <IconButton
                color="secondary"
                sx={{
                  border: "1px solid",
                  borderColor: "currentColor",
                  borderRadius: "10%",
                }}
                size="small"
                onClick={() => deleteCommentMutation.mutate(comment.CommentID)}
              >
                <DeleteIcon />
              </IconButton>
            ) : null}
          </CardContent>
          <Divider />
          <Box marginX="1rem" marginY="0.5rem">
            <Stack direction="row" spacing="1rem">
              <Typography variant="body2" color="text.secondary">
                username: {comment.Username}
              </Typography>
              <Divider orientation="vertical" flexItem />
              <Typography variant="body2" color="text.secondary">
                date: {new Date(comment.CreationDate).toLocaleString()}
              </Typography>
            </Stack>
          </Box>
        </Card>
      ))}
    </Stack>
  );
}

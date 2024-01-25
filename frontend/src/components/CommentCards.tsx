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

interface Comment {
  CommentID: number;
  Content: string;
  CreationDate: string;
  PostID: number;
  UserID: number;
}

export default function CommentCards({ postId }: { postId: string }) {
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
          </CardContent>
          <Divider />
          <Box marginX="1rem" marginY="0.5rem">
            <Stack direction="row" spacing="1.5rem">
              <Typography variant="body2" color="text.secondary">
                User: {comment.UserID}
              </Typography>
              <Typography variant="body2" color="text.secondary">
                Date: {new Date(comment.CreationDate).toLocaleDateString()}
              </Typography>
            </Stack>
          </Box>
        </Card>
      ))}
    </Stack>
  );
}
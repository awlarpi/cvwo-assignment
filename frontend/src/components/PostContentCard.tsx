import { QueryFunctionContext, useQuery } from "@tanstack/react-query";
import { instance } from "../lib/axiosinstance";
import { Post } from "../routes/posts/ForumPage";
import { Card, CardContent, Typography } from "@mui/material";

async function getPostByPostId({ queryKey }: QueryFunctionContext) {
  const [, postId] = queryKey;
  const response = await instance.get<Post>(`/posts/${postId}`);
  return response.data;
}

export default function PostContentCard({ postId }: { postId: string }) {
  const { isLoading, data, isError, error } = useQuery({
    queryKey: ["post", postId],
    queryFn: getPostByPostId,
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
    <Card
      sx={{
        padding: "0.5rem",
      }}
    >
      <CardContent>
        <Typography variant="h4" component="h2">
          {data.Title}
        </Typography>
        <Typography variant="body2" color="textSecondary" component="p">
          {data.Content}
        </Typography>
      </CardContent>
    </Card>
  );
}

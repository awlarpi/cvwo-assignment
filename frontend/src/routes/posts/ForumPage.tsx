import { useQuery } from "@tanstack/react-query";
import { instance } from "../../lib/axiosinstance";
import { Link } from "react-router-dom";
import {
  Typography,
  Card,
  CardActionArea,
  CardContent,
  Stack,
  Container,
} from "@mui/material";

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

export default function ForumPage() {
  const { isLoading, data, isError, error } = useQuery({
    queryKey: ["posts"],
    queryFn: getAllPosts,
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
    <Container maxWidth="lg">
      <Typography variant="h4" component="h2" marginBottom="1rem">
        Posts
      </Typography>
      <Stack spacing="1rem">
        {data.map((post) => (
          <Card key={post.PostID} style={{}}>
            <CardActionArea component={Link} to={`/posts/${post.PostID}`}>
              <CardContent>
                <Typography variant="h5" component="h3">
                  {post.Title}
                </Typography>
                <Typography variant="body2" color="textSecondary" component="p">
                  {post.Content}
                </Typography>
              </CardContent>
            </CardActionArea>
          </Card>
        ))}
      </Stack>
    </Container>
  );
}

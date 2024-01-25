import { Typography, Stack, Container, Divider } from "@mui/material";
import CreatePostForm from "../../components/CreatePostForm";
import PostCards from "../../components/PostCards";

export default function ForumPage() {
  return (
    <Container maxWidth="lg">
      <Stack spacing="1rem">
        <Typography variant="h5" component="h2" marginBottom="1rem">
          Create a post
        </Typography>
        <CreatePostForm />
        <Divider />
        <PostCards />
      </Stack>
    </Container>
  );
}

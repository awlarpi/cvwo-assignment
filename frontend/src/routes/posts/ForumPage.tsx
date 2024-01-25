import { Typography, Stack, Container, Divider } from "@mui/material";
import CreatePostForm from "../../components/CreatePostForm";
import PostCards from "../../components/PostCards";
import { useStore } from "../../lib/store";

export default function ForumPage() {
  const { isLoggedIn } = useStore();

  return (
    <Container maxWidth="lg">
      <Stack spacing="1rem">
        <Typography variant="h5" component="h2" marginBottom="1rem">
          Create a post
        </Typography>
        {isLoggedIn && <CreatePostForm />}
        <Divider />
        <PostCards />
      </Stack>
    </Container>
  );
}

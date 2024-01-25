import { useParams } from "react-router-dom";
import { Container, Divider, Stack } from "@mui/material";
import CommentForm from "../../components/CreateCommentForm";
import PostContentCard from "../../components/PostContentCard";
import CommentCards from "../../components/CommentCards";

export default function PostPage() {
  const postId = useParams<{ postId: string }>().postId;

  return (
    <Container maxWidth="lg">
      <Stack spacing="1rem">
        <PostContentCard postId={postId ?? ""} />
        <CommentForm postId={postId ?? ""} />
        <Divider
          sx={{
            borderStyle: "dotted",
            borderWidth: "0.05rem",
            borderColor: "grey.600",
          }}
        />
        <CommentCards postId={postId ?? ""} />
      </Stack>
    </Container>
  );
}

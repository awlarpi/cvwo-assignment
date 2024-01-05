import { useQuery } from "@tanstack/react-query";
import { instance } from "../../axios";

type Post = {
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

type GetAllPostsApiResponse = {
  result: Post[];
};

async function getAllPosts() {
  const response = await instance.get<GetAllPostsApiResponse>("/posts");
  return response.data;
}

export default function PostsPage() {
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
    <div>
      <h2>Posts</h2>
      <ul>
        {data.result.map((post) => (
          <li key={post.PostID}>
            <h3>{post.Title}</h3>
            <p>{post.Content}</p>
          </li>
        ))}
      </ul>
    </div>
  );
}

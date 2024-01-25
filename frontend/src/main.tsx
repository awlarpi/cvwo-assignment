import React from "react";
import ReactDOM from "react-dom/client";
import { RouterProvider, createBrowserRouter } from "react-router-dom";
import Root from "./routes/root/RootLayout.tsx";
import { ThemeProvider, CssBaseline, createTheme } from "@mui/material";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import PostsPage from "./routes/posts/ForumPage.tsx";
import LoginPage from "./routes/login/LoginPage.tsx";
import PostPage from "./routes/posts/PostPage.tsx";
import SignupPage from "./routes/signup/SignupPage.tsx";

const theme = createTheme({
  palette: {
    mode: "dark",
  },
});

export const queryClient = new QueryClient();

const router = createBrowserRouter([
  {
    path: "/",
    element: <Root />,
    children: [
      {
        path: "",
        element: <PostsPage />,
      },
      {
        path: "login",
        element: <LoginPage />,
      },
      {
        path: "signup",
        element: <SignupPage />,
      },
      // {
      //   path: "posts",
      //   element: <PostsPage />,
      // },
      {
        path: "posts/:postId",
        element: <PostPage />,
      },
    ],
  },
]);

ReactDOM.createRoot(document.getElementById("root")!).render(
  <React.StrictMode>
    <ThemeProvider theme={theme}>
      <CssBaseline />
      <QueryClientProvider client={queryClient}>
        <RouterProvider router={router} />
      </QueryClientProvider>
    </ThemeProvider>
  </React.StrictMode>
);

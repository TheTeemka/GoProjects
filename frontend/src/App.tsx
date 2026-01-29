import { createBrowserRouter, RouterProvider } from "react-router";
import LoginPage from "@/pages/login/page";
import HomePage from "@/pages/homepage/page";
import RegisterPage from "./pages/register/page";

const router = createBrowserRouter([
  {
    path: "/login",
    element: <LoginPage />,
  },
  {
    path: "/",
    element: <HomePage />,
  },
  {
    path: "/register",
    element: <RegisterPage />,
  },
]);

function App() {
  return <RouterProvider router={router} />;
}

export default App;

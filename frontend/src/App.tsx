import { createBrowserRouter, RouterProvider } from "react-router";
import LoginPage from "@/pages/login/page";
import HomePage from "@/pages/homepage/page";
import RegisterPage from "./pages/register/page";
import SchedulePage from "./pages/dashboard/schedule/page";

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
  {
    path: "/dashboard/schedule",
    element: <SchedulePage />,
  },
]);

function App() {
  return <RouterProvider router={router} />;
}

export default App;

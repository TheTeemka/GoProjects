import { createBrowserRouter, RouterProvider } from "react-router";
import LoginPage from "@/pages/login/page";
import HomePage from "@/pages/homepage/page";
import RegisterPage from "./pages/register/page";
import SchedulePage from "./pages/dashboard/schedule/page";
import { DashboardLayout } from "./components/layout/DashboardLayout";

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
    path: "/dashboard",
    element: <DashboardLayout />,
    children: [
      {
        path: "schedule",
        element: <SchedulePage />,
      },
    ],
  },
]);

function App() {
  return <RouterProvider router={router} />;
}

export default App;

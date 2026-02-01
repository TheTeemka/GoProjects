import { createBrowserRouter, Navigate, RouterProvider } from "react-router";
import LoginPage from "@/pages/login/page";
import HomePage from "@/pages/homepage/page";
import RegisterPage from "./pages/register/page";
import SchedulePage from "./pages/dashboard/schedule/page";
import { DashboardLayout } from "./components/layout/DashboardLayout";
import { ToastContainer } from "react-toastify";
import StudentsPage from "./pages/dashboard/student/page";

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
        path: "",
        element: <Navigate to="students" replace />,
      },
      {
        path: "schedule",
        element: <SchedulePage />,
      },
      {
        path: "students",
        element: <StudentsPage />,
      },
      {
        path: "attendance",
        element: <div>Attendance Page</div>,
      },
    ],
  },
]);

function App() {
  return (
    <>
      <RouterProvider router={router} />;
      <ToastContainer />
    </>
  );
}

export default App;

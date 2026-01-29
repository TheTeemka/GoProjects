import type {
  LoginResponse,
  LoginRequest,
  RegisterRequest,
} from "../types/user";
import apiClient from "./axiosClient";
import type { User } from "@/types/user";

const login = async (req: LoginRequest): Promise<LoginResponse> => {
  const resp = await apiClient.post<LoginResponse>("/auth/user/login", req);
  return resp.data;
};

const register = async (req: RegisterRequest): Promise<boolean> => {
  const resp = await apiClient.post("/auth/register", req);
  return resp.status === 201;
};

const me = async (): Promise<User> => {
  const resp = await apiClient.get("/auth/me");
  return resp.data;
};

export const authApi = {
  login,
  register,
  me,
};

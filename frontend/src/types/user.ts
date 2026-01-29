export type User = {
  id: number;
  email: string;
  role: UserRole;
};

export type UserRole = "admin" | "user";

export interface LoginRequest {
  email: string;
  password: string;
}

export interface LoginResponse {
  accessToken: string;
}

export interface RegisterRequest {
  email: string;
  password: string;
  role: UserRole;
}

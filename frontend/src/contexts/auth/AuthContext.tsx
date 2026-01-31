import { useEffect, useState } from "react";
import type { User, UserRole } from "../../types/user";
import { authApi } from "../../api/auth";
import { AuthContext } from "./useAuth";

export function AuthProvider({ children }: { children: React.ReactNode }) {
  const [user, setUser] = useState<User | null>(null);
  const [isLoading, setIsLoading] = useState<boolean>(true);

  useEffect(() => {
    setIsLoading(true);
    (async () => {
      try {
        const user = await authApi.me();
        setUser(user);
      } catch {
        setUser(null);
      } finally {
        setIsLoading(false);
      }
    })();
  }, []);

  const login = async (email: string, password: string) => {
    const resp = await authApi.login({ email, password });
    localStorage.setItem("accessToken", resp.access_token);

    const userData = await authApi.me();
    setUser(userData);
  };

  const register = async (email: string, password: string, role: UserRole) => {
    return await authApi.register({ email, password, role });
  };

  const logout = () => {
    setUser(null);
    localStorage.removeItem("accessToken");
  };

  return (
    <AuthContext.Provider
      value={{
        user,
        isAuthenticated: !!user,
        isLoading,
        login,
        register,
        logout,
      }}
    >
      {children}
    </AuthContext.Provider>
  );
}

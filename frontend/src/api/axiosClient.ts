import { toast } from "@/lib/toast";
import { type LoginResponse } from "@/types/user";
import axios, {
  AxiosError,
  type AxiosInstance,
  type AxiosRequestConfig,
} from "axios";

const apiClient: AxiosInstance = axios.create({
  baseURL: import.meta.env.VITE_API_BASE_URL || "/",
  withCredentials: true,
});

let refreshPromise: Promise<string | null> | null = null;
async function refreshToken(): Promise<string | null> {
  if (!refreshPromise) {
    refreshPromise = (async () => {
      try {
        const res = await axios.post<LoginResponse>("/auth/refresh", null, {
          baseURL: apiClient.defaults.baseURL,
          withCredentials: true,
        });

        const newToken = res.data?.access_token ?? null;
        if (newToken) {
          localStorage.setItem("accessToken", newToken);
        } else {
          localStorage.removeItem("accessToken");
        }
        return newToken;
      } catch {
        localStorage.removeItem("accessToken");
        return null;
      } finally {
        refreshPromise = null;
      }
    })();
  }
  return refreshPromise;
}

apiClient.interceptors.request.use((cfg) => {
  console.debug("apiClient request:", cfg);

  const token = localStorage.getItem("accessToken");
  if (token && cfg.headers) {
    cfg.headers.Authorization = `Bearer ${token}`;
  }
  return cfg;
});

// on 401 try refresh once and retry the original request
apiClient.interceptors.response.use(
  (r) => r,
  async (err: AxiosError) => {
    if (!err.response) {
      toast.error("Network error. Please check your connection.");
      return Promise.reject(err);
    }

    const orig = err.config as AxiosRequestConfig & { _retry?: boolean };
    if (err.response?.status === 401 && !orig?._retry) {
      orig._retry = true;
      const token = await refreshToken();
      if (token) {
        orig.headers = orig.headers ?? {};
        orig.headers.Authorization = `Bearer ${token}`;
        return apiClient.request(orig);
      }
    }
    return Promise.reject(err);
  },
);

export default apiClient;

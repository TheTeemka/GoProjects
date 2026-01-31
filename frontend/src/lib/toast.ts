import { toast as rtoast, type ToastOptions } from "react-toastify";

const defaultOpts: ToastOptions = {
  position: "top-right" as const,
  autoClose: 4000,
  hideProgressBar: false,
  closeOnClick: true,
  pauseOnHover: true,
  draggable: true,
  theme: "colored" as const,
};

export const success = (message: string) =>
  rtoast.success(message, defaultOpts);
export const error = (message: string) => rtoast.error(message, defaultOpts);
export const info = (message: string) => rtoast.info(message, defaultOpts);
export const warn = (message: string) => rtoast.warn(message, defaultOpts);

export const toast = {
  success,
  error,
  info,
  warn,
};

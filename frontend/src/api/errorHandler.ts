import axios from "axios";

export const handleError = (error: unknown): string => {
  if (axios.isAxiosError(error)) {
    return error.response?.data?.message || "An unexpected error occurred";
  } else if (error instanceof Error) {
    return error.message;
  }
  return "An unexpected error occurred";
};

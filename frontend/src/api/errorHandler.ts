import axios from "axios";

import { ApiError } from "@/api/ApiError";

export const handleError = (error: unknown): string => {
  if (axios.isAxiosError(error) && error.response) {
    throw new ApiError(
      error.response.data?.message || "An unexpected error occurred",
      error.response.status,
      error.response.data?.type,
    );
  } else if (error instanceof Error) {
    return error.message;
  }
  return "An unexpected error occurred";
};

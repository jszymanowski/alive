import axiosInstance from "@/api/axiosInstance";
import { handleError } from "@/api/errorHandler";
import type { User } from "@/types";

export type CreateUserRequestBody = Omit<User, "id">;

interface ResponseBody {
  data: User;
}

const API_ENDPOINTS = {
  CURRENT_USER: "/v1/current_user",
  USERS: "/v1/users",
};

export const fetchCurrentUser = async (): Promise<User> => {
  try {
    const response = await axiosInstance.get<ResponseBody>(API_ENDPOINTS.CURRENT_USER);
    return response.data.data;
  } catch (error) {
    throw new Error(handleError(error));
  }
};

export const createUser = async (params: Omit<User, "id">): Promise<User> => {
  try {
    const response = await axiosInstance.post<ResponseBody>(API_ENDPOINTS.USERS, params);
    return response.data.data;
  } catch (error) {
    console.log("Error creating user:", error instanceof Error ? error.message : String(error));
    throw new Error(handleError(error));
  }
};

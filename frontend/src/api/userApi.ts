import axiosInstance from "@/api/axiosInstance";
import { handleError } from "@/api/errorHandler";
import type { User } from "@/types";

export type CreateUserRequestBody = Omit<User, "id">;

interface ResponseBody {
  data: User;
}

export const fetchCurrentUser = async (): Promise<User> => {
  try {
    const response = await axiosInstance.get<ResponseBody>(`/v1/current_user`);
    return response.data.data;
  } catch (error) {
    console.log("error", error);
    throw new Error(handleError(error));
  }
};



export const createUser = async (params: Omit<User, "id">): Promise<User> => {
  try {
    const response = await axiosInstance.post<ResponseBody>(`/v1/users`, params);
    return response.data.data;
  } catch (error) {
    console.log("Error creating user:", error instanceof Error ? error.message : String(error));
    throw new Error(handleError(error));
  }
};

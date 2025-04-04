import axiosInstance from "@/api/axiosInstance";
import type { User } from "@/types";
import { handleError } from "@/api/errorHandler";

interface ResponseBody {
  status: string;
  message?: string;
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

export const createUser = async (params: User): Promise<User> => {
  try {
    const response = await axiosInstance.post<ResponseBody>(`/v1/users`, params);
    return response.data.data;
  } catch (error) {
    console.log("error", error);
    throw new Error(handleError(error));
  }
};

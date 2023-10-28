import axiosInstance from "./axiosInstance";
import { BASE_URL } from "../url/urls";
import { AxiosResponse } from "axios";

interface LoginPayload {
  email: string;
  password: string;
}

export const loginUser = async (body: LoginPayload): Promise<AxiosResponse> => {
    const response = await axiosInstance.post(`${BASE_URL}/users/login`, body);
    return response;
};

interface CreateUserPayload {
  email: string;
  password: string;
}

export const signupUser = async (body: CreateUserPayload) => {
  const { data } = await axiosInstance.post(`${BASE_URL}/users/signup`, body);
  return data;
};

interface AddUserFavoritePayload {
  userID: string;
  plantID: string;
}

export const addUserFavorite = async (body: AddUserFavoritePayload) => {
  const response = await axiosInstance.post(`${BASE_URL}/users/favorites/add`, body);
  return response;
};

interface RemoveUserFavoritePayload {
  userID: string;
  plantID: string;
}

export const removeUserFavorite = async (body: RemoveUserFavoritePayload) => {
  const response = await axiosInstance.post(`${BASE_URL}/users/favorites/remove`, body);
  return response;
};
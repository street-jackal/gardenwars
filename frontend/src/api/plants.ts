import axiosInstance from "./axiosInstance";
import { IPlant } from "../interface/IPlant";
import { BASE_URL } from "../url/urls";

export const getAllPlants = async (): Promise<IPlant[]> => {
  const { data } = await axiosInstance.get(`${BASE_URL}/plants/getAll`);
  return data as IPlant[];
};

interface GetAllPlantsForUserPayload {
  userID: string;
}

export const getAllPlantsForUser = async (body: GetAllPlantsForUserPayload): Promise<IPlant[]> => {
  const { data } = await axiosInstance.post(`${BASE_URL}/plants/getAllForUser`, body);
  return data as IPlant[];
};

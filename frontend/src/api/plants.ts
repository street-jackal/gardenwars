import { AxiosResponse } from 'axios';
import axiosInstance from './axiosInstance';
import { IPlant } from '../interface/IPlant';
import {BASE_URL} from '../url/urls';

export const getAllPlants = async (): Promise<IPlant[]> => {
    console.log("getAllPlants")
    try {
        const response: AxiosResponse = await axiosInstance.get(BASE_URL + '/plants/getAll');
        console.log(response.data);
        return response.data as IPlant[];
    } catch (error) {
        console.error(error);
    }

    return [];
};

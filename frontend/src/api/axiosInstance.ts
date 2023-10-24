import axios from 'axios';

// create axios instance
const axiosInstance = axios.create({
    headers: {
      'Content-Type': 'application/json',
      'Accept': 'application/json',
      'Origin':'http://localhost:3000',
    },
  });

export default axiosInstance;
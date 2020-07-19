import axios from "axios";

const axiosInstance = axios.create({
  headers: {
    "Content-Type": "application/json"
  },
  baseURL: "http://localhost:8000"
});

export default axiosInstance;
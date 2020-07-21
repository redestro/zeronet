import axios from "axios";

const axiosInstance = axios.create({
  headers: {
    Accept: "*/*",
    "Content-Type": "application/x-www-form-urlencoded"
  },
  baseURL: "http://localhost:8000"
});

export default axiosInstance;

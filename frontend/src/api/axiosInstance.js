import axios from "axios";

const axiosInstance = axios.create({
  headers: {
    Accept: "*/*",
    "Content-Type": "application/x-www-form-urlencoded"
  },
  baseURL: "http://ec2-18-218-74-74.us-east-2.compute.amazonaws.com:8000/"
});

export default axiosInstance;

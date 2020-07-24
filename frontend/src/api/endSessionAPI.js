import axiosInstance from "./axiosInstance";

export default {
  async endSession(token) {
    console.log(token);
    return axiosInstance.delete(`/session/${token}`);
  }
};

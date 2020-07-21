import axiosInstance from "./axiosInstance";

export default {
  async endSession(token) {
    return axiosInstance.delete(`/session/${token}`).then(response => {
      return response;
    });
  }
};

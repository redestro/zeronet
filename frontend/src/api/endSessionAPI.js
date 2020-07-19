import axiosInstance from "./axiosInstance";

export default {
  async endSession() {
    const token = this.$store.getters.token
    const response = await axiosInstance.get(`/session/${token}`);
    return response.data;
  }
};

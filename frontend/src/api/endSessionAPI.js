import axiosInstance from "./axiosInstance";

export default {
  async endSession() {
    const token = this.$store.getters.token;
    return axiosInstance.delete(`/session/${token}`).then(response => {
      return response;
    });
  }
};

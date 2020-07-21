import axiosInstance from "./axiosInstance";

export default {
  async startNewSession(type) {
    return axiosInstance
      .get(`/start/${type}/1/2`)
      .then(response => {
        const res = JSON.parse(response.request.response);
        return res.token;
      })
      .catch(error => {
        return Promise.reject(error);
      });
  }
};

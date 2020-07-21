import axiosInstance from "./axiosInstance";

export default {
  async playGameAIvsHuman(payload, token) {
    const formBody = Object.keys(payload)
      .map(key => `${encodeURIComponent(key)}=${encodeURIComponent(payload[key])}`)
      .join("&");

    return axiosInstance.post(`play/${token}/human`, formBody).then(response => {
      return response;
    });
  }
};

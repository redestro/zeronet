import axiosInstance from "./axiosInstance";

var FormData = require('form-data');

export default {
  async playGameAIvsHuman(currentMove, token) {
    var data = new FormData();
    data.append('player', '1');
    data.append('move', currentMove);
    data.append('symbol', 'O');

    console.log(data);
    return axiosInstance.get(`play/${token}/human`, data).then(response => {
      console.log(response);
      return response;
    });
  }
};

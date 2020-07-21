import * as fetch from 'whatwg-fetch';
import axiosInstance from "./axiosInstance";

// var FormData = require("form-data");

export default {
  async playGameAIvsHuman(payload, token) {
    const formBody = Object.keys(payload)
      .map(key => `${encodeURIComponent(key)}=${encodeURIComponent(payload[key])}`)
      .join("&");

      console.log(formBody)
    return axiosInstance.post(`play/${token}/human`, {"move": payload.move}).then(response => {
      console.log(response);
      return response;
    });
  }
};

// const resp = {"move":4,"Symbol":"X","board":["_","_","_","_","X","_","_","_","O"],"recoMove":0,"RecoSymbol":"O"}

// export default {
//   playGameAIvsHuman(payload, token) {
//     console.log("hi");
//     const formBody = Object.keys(payload)
//       .map(key => `${encodeURIComponent(key)}=${encodeURIComponent(payload[key])}`)
//       .join("&");

//       console.log(payload)

//     return fetch(`http://localhost:8000/play/${token}/human`, {
//       method: "POST",
//       headers: {
//         Accept: "application/json",
//         "Content-Type": "application/form-data"
//       },
//       body: payload
//     });
//   }
// };

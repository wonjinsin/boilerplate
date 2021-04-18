require("dotenv").config();
const express = require("express");
const axios = require("axios");
const multipart = require("connect-multiparty");
const multipartMiddleware = multipart();
const router = express.Router();
const verifyToken = require("../functions/verifyToken");
const { axiosErrorLogger, errorLogger } = require("../functions/logger");

router.get("/", function (req, res, next) {
  res.end("index route /front/api/v1/dummy");
});

router.post("/", multipartMiddleware, (req, res, next) => {
  return axios
    .post(`https://dummydomain.kr`, {
      headers: {
        "Content-Type": "application/json",
        Authorization: `Bearer ${access_token}`,
      },
      data: {
        info,
      },
    })
    .then((response) => {
      res.json(response.data);
    })
    .catch((error) => {
      try {
        axiosErrorLogger(error.response.data);
        res.json({ msg: "error occured" });
      } catch (error) {
        errorlogger(error);
        res.json({ msg: "error occured" });
      }
    });
  res.end();
});

module.exports = router;

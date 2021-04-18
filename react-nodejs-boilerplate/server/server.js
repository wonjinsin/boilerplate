require("dotenv").config();
const express = require("express");
const dummyRouter = require("./routes/dummyRouter");
const app = express();
const winston = require("winston");
const expressWinston = require("express-winston");
const bodyParser = require("body-parser");
const cookieParser = require("cookie-parser");
const redis = require("redis");
const redisClient = redis.createClient(
  `${process.env.REDIS_PORT}`,
  `${process.env.REDIS_HOST}`
);
const session = require("express-session");
const RedisStore = require("connect-redis")(session);
const axios = require("axios");
const jsonwebtoken = require("jsonwebtoken");
const jwksClient = require("jwks-rsa");
const verifyToken = require("./functions/verifyToken");
const {
  axiosErrorLogger,
  errorLogger,
  jsonFormatter,
  prettyFormatterForExpress,
} = require("./functions/logger");

process.on("uncaughtException", function (err) {
  console.log("------------UncaughtException occured start------------");
  errorLogger(err);
  console.log("------------UncaughtException occured error------------");
});

app.use(bodyParser.json());
app.use(bodyParser.urlencoded({ extended: true }));
app.use(cookieParser());

const ISLOCAL = process.env.NODE_ENV === "LOCAL_ENV" ? true : false;
app.use(
  expressWinston.logger({
    transports: [new winston.transports.Console()],
    level: function (req, res) {
      let level = "";
      if (res.statusCode >= 100) {
        level = "info";
      }
      if (res.statusCode >= 400) {
        level = "error";
      }
      return level;
    },
    format: winston.format.combine(
      winston.format.timestamp({
        format: "YYYY-MM-DD HH:mm:ss.SSS",
      }),
      prettyFormatterForExpress
    ),
    colorize: ISLOCAL,
    meta: true,
    expressFormat: true,
    ignoreRoute: function (req, res) {
      return false;
    },
  })
);

app.use(
  session({
    store: new RedisStore({
      client: redisClient,
      host: `${process.env.REDIS_HOST}`,
      port: `${process.env.REDIS_PORT}`,
    }),
    secret: "secret_key",
    resave: true,
    saveUninitialized: true,
    proxy: true,
    cookie: {
      path: "/",
      httpOnly: true,
      secure: ISLOCAL ? false : true,
      sameSite: ISLOCAL ? false : "none",
      maxAge: 1000 * 60 * 60 * 24 * 30, // 30 days
    },
  })
);

app.use("*", async (req, res, next) => {
  if (req.session) {
    console.log(req.session);
  }
  next();
});

app.use("/front/api/v1/dummy", dummyRouter);

app.use(
  expressWinston.errorLogger({
    transports: [new winston.transports.Console()],
    format: winston.format.combine(
      winston.format.timestamp({
        format: "YYYY-MM-DD HH:mm:ss.SSS",
      }),
      jsonFormatter
    ),
    level: "error",
    colorize: ISLOCAL,
    meta: true,
    expressFormat: true,
    ignoreRoute: function (req, res) {
      return false;
    },
  })
);

module.exports = app;

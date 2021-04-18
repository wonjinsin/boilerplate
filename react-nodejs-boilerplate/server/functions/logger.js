require("dotenv").config();
const winston = require("winston");
const { combine, timestamp, printf } = winston.format;

const ISLOCAL = process.env.NODE_ENV === "LOCAL_ENV" ? true : false;

const jsonFormatter = winston.format.combine(
  winston.format.printf((info) =>
    JSON.stringify(info).replace(/\\n/g, "\\n").replace(/\\t/g, "\\t")
  )
);

const prettyFormatterForExpress = winston.format.combine(
  winston.format.printf((info) => {
    return `[${info.timestamp}] [${info.level}] ${info.message} [referer=${
      info.meta.req.headers.referer
        ? info.meta.req.headers.referer
        : info.meta.req.headers["user-agent"]
    }]`;
  })
);

const prettyFormatter = winston.format.combine(
  winston.format.printf((info) => {
    return `[${info.timestamp}] [${info.level}] [message=${info.message}]`;
  })
);

const loggerJson = winston.createLogger({
  transports: [
    new winston.transports.Console({
      format: winston.format.combine(
        timestamp({
          format: "YYYY-MM-DD HH:mm:ss.SSS",
        }),
        jsonFormatter
      ),
    }),
  ],
});

const loggerPrittier = winston.createLogger({
  transports: [
    new winston.transports.Console({
      format: winston.format.combine(
        timestamp({
          format: "YYYY-MM-DD HH:mm:ss.SSS",
        }),
        prettyFormatter
      ),
      colorize: ISLOCAL,
    }),
  ],
});

const errorLogger = (e = "") => {
  try {
    if (e) {
      loggerJson.error(e.stack);
      loggerPrittier.error(e.message);
    } else {
      console.log("Logger data not exist");
    }
  } catch (e) {
    console.log(e);
    console.log("error Logger write failed");
  }
};

const axiosErrorLogger = (data = "") => {
  try {
    data = "Axios return Error: " + JSON.stringify(data);
    if (data) {
      loggerPrittier.error(data);
    } else {
      console.log("Axios logger data not exist");
    }
  } catch (e) {
    console.log(e);
    console.log("Axios error Logger write failed");
  }
};

module.exports.errorLogger = errorLogger;
module.exports.axiosErrorLogger = axiosErrorLogger;
module.exports.jsonFormatter = jsonFormatter;
module.exports.prettyFormatterForExpress = prettyFormatterForExpress;

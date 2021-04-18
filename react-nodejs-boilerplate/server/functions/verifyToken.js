const jsonwebtoken = require("jsonwebtoken");
const jwksClient = require("jwks-rsa");

const client = jwksClient({
  jwksUri: `jwksUri`,
});

const getKey = (header, callback) => {
  client.getSigningKey(header.kid, (err, key) => {
    if (key) {
      var signingKey = key.publicKey || key.rsaPublicKey;
    } else {
      var signingKey = "";
    }
    callback(null, signingKey);
  });
};

const verifyToken = (token) => {
  return new Promise((resolve, reject) => {
    jsonwebtoken.verify(token, getKey, (err, decoded) => {
      if (decoded) {
        resolve(decoded);
      } else {
        reject(err);
      }
    });
  });
};

module.exports = verifyToken;
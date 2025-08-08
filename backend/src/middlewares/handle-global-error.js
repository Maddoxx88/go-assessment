const { ApiError } = require("../utils");

const handleGlobalError = (err, req, res, next) => {
  console.error("🔥 Global error handler caught an error:");
  console.error(err); // will show the whole object in terminal

  const statusCode = err.status || 500;

  res.status(statusCode).json({
    error: err.message || "Unknown server error",
    detail: err.stack || err, // show raw error if no stack
  });
};

module.exports = { handleGlobalError };

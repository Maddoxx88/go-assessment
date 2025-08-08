const { Resend } = require("resend");
const { env } = require("../config");
const { ApiError } = require("./api-error");

let sendMail;

// Check if API key exists before creating Resend instance
if (!env.RESEND_API_KEY) {
  console.log("⚠️ RESEND_API_KEY not found. Using mock email sender.");

  sendMail = async (mailOptions) => {
    console.log(`📭 Mock email to: ${mailOptions.to}`);
    console.log(`Subject: ${mailOptions.subject}`);
    console.log(`Body: ${mailOptions.html}`);
    return { mock: true };
  };
} else {
  // Only create Resend instance if API key exists
  const resend = new Resend(env.RESEND_API_KEY);
  
  sendMail = async (mailOptions) => {
    try {
      const { error } = await resend.emails.send(mailOptions);
      if (error) {
        throw new ApiError(500, "Unable to send email");
      }
    } catch (err) {
      throw new ApiError(500, `Email sending failed: ${err.message}`);
    }
  };
}

module.exports = {
  sendMail,
};
const express = require("express");
const path = require("node:path");

const app = express();
const PORT = process.env.PORT || 5173;

app.get("/config.js", (req, res) => {
  res.setHeader("Content-Type", "application/javascript");
  res.send(`
    window.RUNTIME_ENV = {
      NODE_ENV: "${process.env.NODE_ENV || "development"}",
      API_BASE_URL: "${process.env.API_BASE_URL || "http://localhost"}",
    };
  `);
});

// Serve static files from the React app
app.use(express.static(path.join(__dirname, "dist")));

// Handle all other routes by serving the React app's index.html
app.get("/*splat", (req, res) => {
  res.sendFile(path.join(__dirname, "dist", "index.html"));
});

app.listen(PORT, "0.0.0.0", () => {
  console.log(`Server is running on http://localhost:${PORT}`);
});

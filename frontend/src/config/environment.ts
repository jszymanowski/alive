const env: "development" | "production" = window.RUNTIME_ENV?.NODE_ENV || "development";

const API_URL =
  import.meta.env.VITE_API_URL ||
  window.RUNTIME_ENV?.API_BASE_URL ||
  "http://localhost:8000";

export { env, API_URL };

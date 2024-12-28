import axios from "axios";

// Get backend host from environment variables
const BACKEND_HOST = process.env.REACT_APP_BACKEND_HOST || 'http://localhost:8080';

const TOKEN_KEY = "app_token";
const API_BASE_URL = `${BACKEND_HOST}/agent/api/v1`;

// Create an axios instance with base URL and default headers
const api = axios.create({
  baseURL: API_BASE_URL,
  headers: {
    'Content-Type': 'application/json',
  }
});

// Add a request interceptor to include token in headers
api.interceptors.request.use(
  (config) => {
    const token = getAuthToken();
    if (token) {
      config.headers['Authorization'] = `Bearer ${token}`;
    }
    return config;
  },
  (error) => {
    return Promise.reject(error);
  }
);

export function getAuthToken() {
  return localStorage.getItem(TOKEN_KEY);
}

export async function login(username, password) {
  try {
    const res = await api.post("/login", { username, password });
    localStorage.setItem(TOKEN_KEY, res.data.token);
    return res.data;
  } catch (error) {
    console.error("Login failed:", error);
    throw error;
  }
}

export function logout() {
  localStorage.removeItem(TOKEN_KEY);
}

// Add a function to check if the user is authenticated
export function isAuthenticated() {
  return !!getAuthToken();
}

export default api;
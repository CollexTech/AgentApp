import axios from "axios";

const TOKEN_KEY = "app_token";

export function getAuthToken() {
  return localStorage.getItem(TOKEN_KEY);
}

export async function login(username, password) {
  const res = await axios.post("/api/login", { username, password });
  localStorage.setItem(TOKEN_KEY, res.data.token);
}

export function logout() {
  localStorage.removeItem(TOKEN_KEY);
}

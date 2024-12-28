import axios from 'axios';
import { getAuthToken } from "./auth";

const BACKEND_HOST = process.env.REACT_APP_BACKEND_HOST || 'http://localhost:8080';

const instance = axios.create({
  baseURL: `${BACKEND_HOST}/agent/api/v1`,
});

// Add token to headers for all requests
instance.interceptors.request.use(
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

export default instance;
// Add token to headers for all requests
instance.interceptors.request.use((config) => {
  const token = getAuthToken();
  if (token) {
    config.headers.Authorization = `Bearer ${token}`;
  }
  return config;
});

export async function getCases() {
  const res = await instance.get("/cases");
  return res.data;
}

export async function getCaseDetails(caseId) {
  const res = await instance.get(`/cases/${caseId}`);
  return res.data;
}

export async function getTrails(caseId) {
  const res = await instance.get(`/cases/${caseId}/trails`);
  return res.data;
}

export async function postTrail(caseId, trail) {
  const res = await instance.post(`/cases/${caseId}/trails`, trail);
  return res.data;
}

export async function getPaymentLink(caseId) {
  const res = await instance.get(`/cases/${caseId}/payment-link`);
  return res.data;
}

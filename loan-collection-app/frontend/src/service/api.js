import axios from 'axios';
import { getAuthToken } from "./auth";

const BACKEND_HOST = process.env.REACT_APP_BACKEND_HOST || 'http://localhost:8080';

const instance = axios.create({
  baseURL: `${BACKEND_HOST}/api/v1`,
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

export async function getUserRolesAndPermissions() {
  const res = await instance.get("/permissions/me");
  return res.data;
}

export async function fetchUsers() {
  const res = await instance.get("/users");
  return res.data;
}

export async function createUser(userData) {
  const res = await instance.post("/users/register", userData);
  return res.data;
}

export const assignRolesToUser = async (userId, roleList) => {
  console.log("assignRolesToUser", userId, roleList);
    const response = await fetch(`${BACKEND_HOST}/api/v1/roles/assign`, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
            'Authorization': `Bearer ${getAuthToken()}`
        },
        body: JSON.stringify({
            user_id: userId,
            role_list: roleList
        })
    });
    
    if (!response.ok) {
        throw new Error('Failed to assign roles');
    }
    
    return await response.json();
};

export async function getAllAvailableRoles() {
  const res = await instance.get("/roles");
  return res.data;
}

const API_BASE_URL = `${BACKEND_HOST}/api/v1`;

// Helper function to get auth headers
const getAuthHeaders = () => ({
  'Authorization': `Bearer ${getAuthToken()}`
});

// Helper function to handle API responses
const handleResponse = async (response) => {
  if (!response.ok) {
    const error = await response.json().catch(() => ({}));
    throw new Error(error.message || 'An error occurred');
  }
  return response.json();
};

// Update the agency-related functions
export const createAgency = async (agencyData) => {
  try {
    const response = await instance.post('/agencies', agencyData);
    return response.data;
  } catch (error) {
    throw new Error('Failed to create agency: ' + error.message);
  }
};

export const assignUserToAgency = async (assignmentData) => {
  try {
    const response = await instance.post('/agencies/users', assignmentData);
    return response.data;
  } catch (error) {
    throw new Error('Failed to assign user to agency: ' + error.message);
  }
};

export const assignCaseToUser = async (assignmentData) => {
  try {
    const response = await instance.post('/cases/assign', assignmentData);
    return response.data;
  } catch (error) {
    throw new Error('Failed to assign case to user: ' + error.message);
  }
};

export const getAgencies = () => {
  return instance.get('/agencies');
};

export const deleteAgency = (agencyId) => {
  return instance.delete(`/agencies/${agencyId}`);
};

export const getAgencyUsers = async (agencyId) => {
  try {
    const response = await instance.get(`/agencies/${agencyId}/users`);
    // Return the data in a consistent format
    return Array.isArray(response.data) ? response.data : 
           Array.isArray(response?.data?.data) ? response.data.data : [];
  } catch (error) {
    console.error('Error fetching agency users:', error);
    throw error;
  }
};

export const getUnassignedUsers = async () => {
  try {
    const response = await instance.get('/agencies/unassigned-users');
    // Return the data in a consistent format
    return {
      data: Array.isArray(response.data) ? response.data : response.data?.data || []
    };
  } catch (error) {
    console.error('Error fetching unassigned users:', error);
    throw error;
  }
};

export const uploadCases = async (formData) => {
  try {
    const response = await instance.post('/cases/upload', formData, {
      headers: {
        'Content-Type': 'multipart/form-data',
      },
    });
    return response.data;
  } catch (error) {
    throw new Error('Failed to upload cases: ' + error.message);
  }
};

export const getUnassignedCases = async () => {
  try {
    const response = await instance.get('/cases/unassigned');
    return response.data;
  } catch (error) {
    throw new Error('Failed to fetch unassigned cases: ' + error.message);
  }
};

export const assignCasesToAgency = async (assignmentData) => {
  try {
    const response = await instance.post('/cases/assign', assignmentData);
    return response.data;
  } catch (error) {
    throw new Error('Failed to assign cases: ' + error.message);
  }
}; 
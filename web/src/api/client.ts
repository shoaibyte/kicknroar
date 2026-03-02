import axios, { AxiosInstance, AxiosRequestConfig, AxiosResponse, InternalAxiosRequestConfig } from 'axios';
import { useAuthStore } from '../store/authStore';

// Use relative /api/v1 in dev (Vite proxy) and in production when served from same origin (monolith).
// Set VITE_API_URL only when the API is on a different host (e.g. separate API subdomain).
const getBaseURL = () => {
  if (import.meta.env.DEV) {
    return '/api/v1';
  }
  return import.meta.env.VITE_API_URL || '/api/v1';
};

const config: AxiosRequestConfig = {
  baseURL: getBaseURL(),
  timeout: 15000,
  headers: {
    'Content-Type': 'application/json',
  },
  withCredentials: false, // Set to false for CORS
};

const apiClient: AxiosInstance = axios.create(config);

// Request interceptor
apiClient.interceptors.request.use(
  (config: InternalAxiosRequestConfig) => {
    const token = localStorage.getItem('auth-token');
    
    if (token && config.headers) {
      config.headers.Authorization = `Bearer ${token}`;
    }

    // Add request ID for tracking
    if (config.headers) {
      config.headers['X-Request-ID'] = crypto.randomUUID();
    }
    
    return config;
  },
  (error) => {
    return Promise.reject(error);
  }
);

// Response interceptor
apiClient.interceptors.response.use(
  (response: AxiosResponse) => {
    return response;
  },
  async (error) => {
    const originalRequest = error.config;

    // Handle network errors (CORS, connection refused, etc.)
    if (!error.response) {
      const networkError = error.code === 'ERR_NETWORK' || error.message?.includes('Network Error');
      if (networkError) {
        const errorMessage = error.code === 'ECONNREFUSED' 
          ? 'Cannot connect to server. Please ensure the backend is running on http://localhost:8000'
          : 'Network error. Please check your connection and ensure CORS is properly configured on the backend.';
        error.message = errorMessage;
        return Promise.reject(error);
      }
    }

    // Handle 401 - Token expired
    if (error.response?.status === 401 && !originalRequest._retry) {
      originalRequest._retry = true;

      try {
        const refreshToken = localStorage.getItem('refresh-token');
        if (!refreshToken) {
          throw new Error('No refresh token available');
        }

        // Import authApi dynamically to avoid circular dependency
        const { authApi } = await import('./endpoints/auth.api');
        const response = await authApi.refreshToken(refreshToken);
        
        // AuthResponse has token directly
        const token = response.token;
        if (token) {
          localStorage.setItem('auth-token', token);
          if (originalRequest.headers) {
            originalRequest.headers.Authorization = `Bearer ${token}`;
          }
          // Update refresh token if provided
          if (response.refresh_token) {
            localStorage.setItem('refresh-token', response.refresh_token);
          }
          return apiClient(originalRequest);
        }
        throw new Error('Invalid refresh token response');
      } catch (refreshError) {
        // Refresh failed, logout user
        useAuthStore.getState().logout();
        return Promise.reject(refreshError);
      }
    }

    // Format error to match Swagger error response format
    if (error.response?.data) {
      // Ensure error follows Swagger format: { error: { code, message, details } }
      if (!error.response.data.error && error.response.data.message) {
        error.response.data = {
          error: {
            code: error.response.data.code || 'UNKNOWN_ERROR',
            message: error.response.data.message,
            details: error.response.data.details || [],
          },
        };
      }
    }

    return Promise.reject(error);
  }
);

export default apiClient;


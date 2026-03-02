import { create } from 'zustand';
import { persist } from 'zustand/middleware';
import { immer } from 'zustand/middleware/immer';
import { authApi, type User } from '../api/endpoints/auth.api';
import { userApi } from '../api/endpoints/user.api';
import apiClient from '../api/client';
import { extractErrorMessage } from '../api/types';

interface AuthState {
  user: User | null;
  token: string | null;
  isAuthenticated: boolean;
  isLoading: boolean;
  error: string | null;

  // Actions
  login: (credentials: { email: string; password: string }) => Promise<void>;
  signup: (data: { full_name: string; email: string; password: string; phone: string }) => Promise<void>;
  logout: () => Promise<void>;
  refreshToken: () => Promise<void>;
  getCurrentUser: () => Promise<User | null>;
  updateUser: (updates: Partial<User>) => void;
  setLoading: (loading: boolean) => void;
}

export const useAuthStore = create<AuthState>()(
  persist(
    immer((set, get) => ({
      user: null,
      token: null,
      isAuthenticated: false,
      isLoading: false,
      error: null,

      login: async (credentials) => {
        set((state) => {
          state.isLoading = true;
          state.error = null;
        });

        try {
          const response = await authApi.login(credentials);
          
          set((state) => {
            state.user = response.user;
            state.token = response.token;
            state.isAuthenticated = true;
            state.isLoading = false;
          });

          // Set axios default header
          apiClient.defaults.headers.common['Authorization'] = `Bearer ${response.token}`;
          localStorage.setItem('auth-token', response.token);
          
          if (response.refresh_token) {
            localStorage.setItem('refresh-token', response.refresh_token);
          }
        } catch (error: any) {
          set((state) => {
            state.error = extractErrorMessage(error);
            state.isLoading = false;
          });
          throw error;
        }
      },

      signup: async (data) => {
        set((state) => {
          state.isLoading = true;
          state.error = null;
        });

        try {
          const response = await authApi.signup(data);
          
          set((state) => {
            state.user = response.user;
            state.token = response.token;
            state.isAuthenticated = true;
            state.isLoading = false;
          });

          // Set axios default header
          apiClient.defaults.headers.common['Authorization'] = `Bearer ${response.token}`;
          localStorage.setItem('auth-token', response.token);
          
          if (response.refresh_token) {
            localStorage.setItem('refresh-token', response.refresh_token);
          }
        } catch (error: any) {
          set((state) => {
            state.error = extractErrorMessage(error);
            state.isLoading = false;
          });
          throw error;
        }
      },

      logout: async () => {
        try {
          await authApi.logout();
        } catch (error) {
          // Continue with logout even if API call fails
          console.error('Logout API call failed:', error);
        } finally {
          set((state) => {
            state.user = null;
            state.token = null;
            state.isAuthenticated = false;
          });
          
          delete apiClient.defaults.headers.common['Authorization'];
          localStorage.removeItem('auth-token');
          localStorage.removeItem('refresh-token');
        }
      },

      refreshToken: async () => {
        try {
          const refreshToken = localStorage.getItem('refresh-token');
          if (!refreshToken) throw new Error('No refresh token');

          const response = await authApi.refreshToken(refreshToken);
          
          set((state) => {
            state.token = response.token;
          });

          apiClient.defaults.headers.common['Authorization'] = `Bearer ${response.token}`;
          localStorage.setItem('auth-token', response.token);
          
          if (response.refresh_token) {
            localStorage.setItem('refresh-token', response.refresh_token);
          }
        } catch (error) {
          get().logout();
          throw error;
        }
      },

      getCurrentUser: async () => {
        set((state) => {
          state.isLoading = true;
          state.error = null;
        });

        try {
          const user = await userApi.getCurrentUser();
          set((state) => {
            state.user = user;
            state.isAuthenticated = true;
            state.isLoading = false;
          });
          return user;
        } catch (error: any) {
          set((state) => {
            state.error = extractErrorMessage(error);
            state.isLoading = false;
            state.isAuthenticated = false;
          });
          throw error;
        }
      },

      updateUser: (updates) => {
        set((state) => {
          if (state.user) {
            state.user = { ...state.user, ...updates };
          }
        });
      },

      setLoading: (loading) => {
        set((state) => {
          state.isLoading = loading;
        });
      },
    })),
    {
      name: 'auth-storage',
      partialize: (state) => ({
        user: state.user,
        token: state.token,
        isAuthenticated: state.isAuthenticated,
      }),
    }
  )
);


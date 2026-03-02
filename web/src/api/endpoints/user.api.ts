import apiClient from '../client';
import type { User } from './auth.api';

export interface UserProfile extends User {
  bio?: string;
  rating?: number;
  total_matches?: number;
}

export interface UpdateUserRequest {
  full_name?: string;
  skill_level?: 'beginner' | 'intermediate' | 'advanced' | 'professional';
  profile_image_url?: string;
  preferred_locations?: string[];
}

/** Alias for form/profile update payload */
export type UpdateProfileData = UpdateUserRequest;

export interface UserStats {
  matches_played?: number;
  matches_won?: number;
  matches_lost?: number;
  total_goals?: number;
  total_assists?: number;
  win_rate?: number;
  [key: string]: any;
}

export const userApi = {
  async getCurrentUser(): Promise<User> {
    const response = await apiClient.get<User>('/users/me');
    return response.data;
  },

  async updateCurrentUser(data: UpdateUserRequest): Promise<User> {
    const response = await apiClient.put<User>('/users/me', data);
    return response.data;
  },

  async getProfile(userId: string): Promise<User> {
    const response = await apiClient.get<User>(`/users/${userId}`);
    return response.data;
  },

  async getUserStats(userId: string): Promise<UserStats> {
    const response = await apiClient.get<UserStats>(`/users/${userId}/stats`);
    return response.data;
  },
};


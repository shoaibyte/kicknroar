import apiClient from '../client';
import type { PaginatedResponse } from '../types';

export interface Match {
  id: string;
  title: string;
  venue_id: string;
  venue?: Venue;
  match_date: string;
  start_time: string;
  duration_hours: number;
  max_players: number;
  current_players?: number;
  cost_per_player: number;
  match_type: 'casual' | 'competitive' | 'tournament';
  visibility: 'public' | 'private' | 'friends_only';
  status?: 'open' | 'full' | 'cancelled' | 'completed';
  skill_level_required?: 'beginner' | 'intermediate' | 'advanced' | 'professional';
  description?: string;
  rules_notes?: string;
  participants?: Participant[];
  created_at?: string;
  updated_at?: string;
}

export interface Venue {
  id: string;
  name: string;
  address: string;
  latitude: number;
  longitude: number;
  distance_km?: number;
}

export interface Participant {
  user_id: string;
  user?: {
    id: string;
    full_name: string;
    profile_image_url?: string;
  };
}

export interface CreateMatchRequest {
  title: string;
  venue_id: string;
  match_date: string;
  start_time: string;
  duration_hours: number;
  max_players: number;
  cost_per_player: number;
  match_type: 'casual' | 'competitive' | 'tournament';
  visibility: 'public' | 'private' | 'friends_only';
  skill_level_required?: 'beginner' | 'intermediate' | 'advanced' | 'professional';
  description?: string;
  rules_notes?: string;
}

export interface UpdateMatchRequest {
  title?: string;
  venue_id?: string;
  match_date?: string;
  start_time?: string;
  duration_hours?: number;
  max_players?: number;
  cost_per_player?: number;
  match_type?: 'casual' | 'competitive' | 'tournament';
  visibility?: 'public' | 'private' | 'friends_only';
  skill_level_required?: 'beginner' | 'intermediate' | 'advanced' | 'professional';
  description?: string;
  rules_notes?: string;
}

export interface MatchFilters {
  status?: 'open' | 'full' | 'cancelled' | 'completed';
  date_from?: string; // YYYY-MM-DD format
  date_to?: string; // YYYY-MM-DD format
  limit?: number; // default: 20
  offset?: number; // default: 0
  latitude?: number;
  longitude?: number;
  radius?: number; // km
}

export const matchApi = {
  async getMatches(filters: MatchFilters = {}): Promise<PaginatedResponse<Match>> {
    const params = new URLSearchParams();
    
    if (filters.status) params.append('status', filters.status);
    if (filters.date_from) params.append('date_from', filters.date_from);
    if (filters.date_to) params.append('date_to', filters.date_to);
    if (filters.limit !== undefined) params.append('limit', String(filters.limit));
    if (filters.offset !== undefined) params.append('offset', String(filters.offset));
    if (filters.latitude !== undefined) params.append('latitude', String(filters.latitude));
    if (filters.longitude !== undefined) params.append('longitude', String(filters.longitude));
    if (filters.radius !== undefined) params.append('radius', String(filters.radius));

    const response = await apiClient.get<PaginatedResponse<Match>>(
      `/matches?${params.toString()}`
    );
    
    return response.data;
  },

  async getMatch(id: string): Promise<Match> {
    const response = await apiClient.get<Match>(`/matches/${id}`);
    return response.data;
  },

  async create(data: CreateMatchRequest): Promise<Match> {
    const response = await apiClient.post<Match>('/matches', data);
    return response.data;
  },

  async update(id: string, data: UpdateMatchRequest): Promise<Match> {
    const response = await apiClient.put<Match>(`/matches/${id}`, data);
    return response.data;
  },

  async delete(id: string): Promise<void> {
    await apiClient.delete(`/matches/${id}`);
  },

  async join(matchId: string): Promise<void> {
    await apiClient.post(`/matches/${matchId}/join`);
  },

  async leave(matchId: string): Promise<void> {
    await apiClient.post(`/matches/${matchId}/leave`);
  },

  async getParticipants(matchId: string): Promise<Participant[]> {
    const response = await apiClient.get<Participant[]>(
      `/matches/${matchId}/participants`
    );
    return response.data;
  },
};


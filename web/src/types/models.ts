// Domain models and shared types

export interface BaseEntity {
  id: string;
  created_at: string;
  updated_at: string;
}

export interface PaginationMeta {
  page: number;
  per_page: number;
  total: number;
  total_pages: number;
}

export interface Location {
  latitude: number;
  longitude: number;
}

export interface Distance {
  km: number;
  miles?: number;
}

// Re-export commonly used types
export type { User } from '../api/endpoints/auth.api';
export type { Match, Venue } from '../api/endpoints/match.api';


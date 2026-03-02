import apiClient from '../client';
import type { PaginatedResponse } from '../types';

export interface Venue {
  id: string;
  name: string;
  address: string;
  latitude: number;
  longitude: number;
  field_type: 'futsal' | 'football' | 'astro';
  surface_type?: 'grass' | 'artificial' | 'concrete';
  capacity: number;
  hourly_rate?: number;
  facilities?: string[];
  contact_info?: Record<string, any>;
  operating_hours?: Record<string, any>;
  google_place_id?: string;
  image_url?: string;
  distance_km?: number;
  rating?: number;
  created_at?: string;
  updated_at?: string;
}

export interface CreateVenueRequest {
  name: string;
  address: string;
  latitude: number;
  longitude: number;
  field_type: 'futsal' | 'football' | 'astro';
  capacity: number;
  surface_type?: 'grass' | 'artificial' | 'concrete';
  hourly_rate?: number;
  facilities?: string[];
  contact_info?: Record<string, any>;
  operating_hours?: Record<string, any>;
  google_place_id?: string;
}

export interface UpdateVenueRequest {
  name?: string;
  address?: string;
  latitude?: number;
  longitude?: number;
  field_type?: 'futsal' | 'football' | 'astro';
  capacity?: number;
  surface_type?: 'grass' | 'artificial' | 'concrete';
  hourly_rate?: number;
  facilities?: string[];
  contact_info?: Record<string, any>;
  operating_hours?: Record<string, any>;
}

export interface NearbyVenuesParams {
  lat: number;
  lng: number;
  radius?: number; // in kilometers, default: 5
  field_type?: 'futsal' | 'football' | 'astro';
  limit?: number; // default: 20
}

export interface VenueListParams {
  limit?: number; // default: 20
  offset?: number; // default: 0
}

export const venueApi = {
  async getVenues(params: VenueListParams = {}): Promise<PaginatedResponse<Venue>> {
    const queryParams = new URLSearchParams();
    if (params.limit !== undefined) queryParams.append('limit', String(params.limit));
    if (params.offset !== undefined) queryParams.append('offset', String(params.offset));

    const url = queryParams.toString() 
      ? `/venues?${queryParams.toString()}`
      : '/venues';
    
    const response = await apiClient.get<PaginatedResponse<Venue>>(url);
    return response.data;
  },

  async createVenue(data: CreateVenueRequest): Promise<Venue> {
    const response = await apiClient.post<Venue>('/venues', data);
    return response.data;
  },

  async getVenue(id: string): Promise<Venue> {
    const response = await apiClient.get<Venue>(`/venues/${id}`);
    return response.data;
  },

  async updateVenue(id: string, data: UpdateVenueRequest): Promise<Venue> {
    const response = await apiClient.put<Venue>(`/venues/${id}`, data);
    return response.data;
  },

  async getNearbyVenues(params: NearbyVenuesParams): Promise<PaginatedResponse<Venue>> {
    const queryParams = new URLSearchParams();
    queryParams.append('lat', params.lat.toString());
    queryParams.append('lng', params.lng.toString());
    if (params.radius !== undefined) {
      queryParams.append('radius', params.radius.toString());
    }
    if (params.field_type) {
      queryParams.append('field_type', params.field_type);
    }
    if (params.limit !== undefined) {
      queryParams.append('limit', params.limit.toString());
    }

    const response = await apiClient.get<PaginatedResponse<Venue>>(
      `/venues/nearby?${queryParams.toString()}`
    );
    return response.data;
  },

  async uploadVenueImage(venueId: string, file: File): Promise<Record<string, string>> {
    const formData = new FormData();
    formData.append('file', file);
    
    const response = await apiClient.post<Record<string, string>>(
      `/venues/${venueId}/upload`,
      formData,
      {
        headers: {
          'Content-Type': 'multipart/form-data',
        },
      }
    );
    return response.data;
  },
};


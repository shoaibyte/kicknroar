import { create } from 'zustand';
import { devtools } from 'zustand/middleware';
import { immer } from 'zustand/middleware/immer';
import { venueApi, type Venue, type CreateVenueRequest, type UpdateVenueRequest } from '../api/endpoints/venue.api';

interface VenueState {
  // State
  venues: Venue[];
  currentVenue: Venue | null;
  nearbyVenues: Venue[];
  isLoading: boolean;

  // Actions
  setVenues: (venues: Venue[]) => void;
  setCurrentVenue: (venue: Venue | null) => void;
  setNearbyVenues: (venues: Venue[]) => void;
  fetchNearbyVenues: (lat: number, lng: number, radius?: number, fieldType?: 'futsal' | 'football' | 'astro') => Promise<void>;
  createVenue: (data: CreateVenueRequest) => Promise<Venue>;
  updateVenue: (id: string, data: UpdateVenueRequest) => Promise<Venue>;
}

export const useVenueStore = create<VenueState>()(
  devtools(
    immer((set) => ({
      // Initial state
      venues: [],
      currentVenue: null,
      nearbyVenues: [],
      isLoading: false,

      // Actions
      setVenues: (venues) =>
        set((state) => {
          state.venues = venues;
        }),

      setCurrentVenue: (venue) =>
        set((state) => {
          state.currentVenue = venue;
        }),

      setNearbyVenues: (venues) =>
        set((state) => {
          state.nearbyVenues = venues;
        }),

      fetchNearbyVenues: async (lat, lng, radius = 5, fieldType) => {
        set((state) => {
          state.isLoading = true;
        });

        try {
          const response = await venueApi.getNearbyVenues({ lat, lng, radius, field_type: fieldType });
          set((state) => {
            state.nearbyVenues = response.data;
            state.isLoading = false;
          });
        } catch (error) {
          set((state) => {
            state.isLoading = false;
          });
          throw error;
        }
      },

      createVenue: async (data) => {
        set((state) => {
          state.isLoading = true;
        });

        try {
          const venue = await venueApi.createVenue(data);
          set((state) => {
            state.venues.push(venue);
            state.isLoading = false;
          });
          return venue;
        } catch (error) {
          set((state) => {
            state.isLoading = false;
          });
          throw error;
        }
      },

      updateVenue: async (id, data) => {
        set((state) => {
          state.isLoading = true;
        });

        try {
          const venue = await venueApi.updateVenue(id, data);
          set((state) => {
            const index = state.venues.findIndex((v) => v.id === id);
            if (index !== -1) {
              state.venues[index] = venue;
            }
            if (state.currentVenue?.id === id) {
              state.currentVenue = venue;
            }
            state.isLoading = false;
          });
          return venue;
        } catch (error) {
          set((state) => {
            state.isLoading = false;
          });
          throw error;
        }
      },
    })),
    {
      name: 'venue-store',
    }
  )
);


import { create } from 'zustand';
import { devtools } from 'zustand/middleware';
import { immer } from 'zustand/middleware/immer';
import { matchApi, type Match, type CreateMatchRequest, type MatchFilters } from '../api/endpoints/match.api';

interface MatchState {
  // State
  matches: Match[];
  currentMatch: Match | null;
  filters: MatchFilters;
  isCreating: boolean;

  // Actions
  setMatches: (matches: Match[]) => void;
  addMatch: (match: Match) => void;
  updateMatch: (id: string, updates: Partial<Match>) => void;
  deleteMatch: (id: string) => void;
  setCurrentMatch: (match: Match | null) => void;
  setFilters: (filters: Partial<MatchFilters>) => void;
  joinMatch: (matchId: string) => Promise<void>;
  leaveMatch: (matchId: string) => Promise<void>;
  createMatch: (data: CreateMatchRequest) => Promise<Match>;
}

export const useMatchStore = create<MatchState>()(
  devtools(
    immer((set) => ({
      // Initial state
      matches: [],
      currentMatch: null,
      filters: {
        status: 'open',
        limit: 20,
        offset: 0,
      },
      isCreating: false,

      // Actions
      setMatches: (matches) =>
        set((state) => {
          state.matches = matches;
        }),

      addMatch: (match) =>
        set((state) => {
          state.matches.unshift(match);
        }),

      updateMatch: (id, updates) =>
        set((state) => {
          const index = state.matches.findIndex((m) => m.id === id);
          if (index !== -1) {
            Object.assign(state.matches[index], updates);
          }
          if (state.currentMatch?.id === id) {
            Object.assign(state.currentMatch, updates);
          }
        }),

      deleteMatch: (id) =>
        set((state) => {
          state.matches = state.matches.filter((m) => m.id !== id);
          if (state.currentMatch?.id === id) {
            state.currentMatch = null;
          }
        }),

      setFilters: (filters) =>
        set((state) => {
          state.filters = { ...state.filters, ...filters };
        }),

      setCurrentMatch: (match) =>
        set((state) => {
          state.currentMatch = match;
        }),

      joinMatch: async (matchId) => {
        try {
          await matchApi.join(matchId);

          set((state) => {
            const match = state.matches.find((m) => m.id === matchId);
            if (match) {
              match.current_players = (match.current_players || 0) + 1;

              if (match.current_players >= match.max_players) {
                match.status = 'full';
              }
            }
            if (state.currentMatch?.id === matchId) {
              state.currentMatch.current_players = (state.currentMatch.current_players || 0) + 1;
              if (state.currentMatch.current_players >= state.currentMatch.max_players) {
                state.currentMatch.status = 'full';
              }
            }
          });
        } catch (error) {
          throw error;
        }
      },

      leaveMatch: async (matchId) => {
        try {
          await matchApi.leave(matchId);

          set((state) => {
            const match = state.matches.find((m) => m.id === matchId);
            if (match) {
              match.current_players = Math.max(0, (match.current_players || 0) - 1);
              if (match.status === 'full' && (match.current_players || 0) < match.max_players) {
                match.status = 'open';
              }
            }
            if (state.currentMatch?.id === matchId) {
              state.currentMatch.current_players = Math.max(
                0,
                (state.currentMatch.current_players || 0) - 1
              );
              if (
                state.currentMatch.status === 'full' &&
                (state.currentMatch.current_players || 0) < state.currentMatch.max_players
              ) {
                state.currentMatch.status = 'open';
              }
            }
          });
        } catch (error) {
          throw error;
        }
      },

      createMatch: async (data) => {
        set((state) => {
          state.isCreating = true;
        });

        try {
          const match = await matchApi.create(data);

          set((state) => {
            state.matches.unshift(match);
            state.isCreating = false;
          });

          return match;
        } catch (error) {
          set((state) => {
            state.isCreating = false;
          });
          throw error;
        }
      },
    })),
    {
      name: 'match-store',
    }
  )
);


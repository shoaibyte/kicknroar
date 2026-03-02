import { create } from 'zustand';
import { devtools } from 'zustand/middleware';
import { immer } from 'zustand/middleware/immer';

interface UIState {
  // State
  sidebarOpen: boolean;
  theme: 'light' | 'dark' | 'system';
  mobileMenuOpen: boolean;

  // Actions
  setSidebarOpen: (open: boolean) => void;
  toggleSidebar: () => void;
  setTheme: (theme: 'light' | 'dark' | 'system') => void;
  setMobileMenuOpen: (open: boolean) => void;
  toggleMobileMenu: () => void;
}

export const useUIStore = create<UIState>()(
  devtools(
    immer((set) => ({
      // Initial state
      sidebarOpen: false,
      theme: 'system',
      mobileMenuOpen: false,

      // Actions
      setSidebarOpen: (open) =>
        set((state) => {
          state.sidebarOpen = open;
        }),

      toggleSidebar: () =>
        set((state) => {
          state.sidebarOpen = !state.sidebarOpen;
        }),

      setTheme: (theme) =>
        set((state) => {
          state.theme = theme;
        }),

      setMobileMenuOpen: (open) =>
        set((state) => {
          state.mobileMenuOpen = open;
        }),

      toggleMobileMenu: () =>
        set((state) => {
          state.mobileMenuOpen = !state.mobileMenuOpen;
        }),
    })),
    {
      name: 'ui-store',
    }
  )
);


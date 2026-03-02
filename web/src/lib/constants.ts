// App Constants
export const APP_NAME = 'Kick&Roar';
export const APP_DESCRIPTION = 'Connect, Play, Share';

// API Configuration
export const API_TIMEOUT = 15000;
export const DEFAULT_PAGE_SIZE = 20;

// Match Constants
export const MATCH_STATUS = {
  OPEN: 'open',
  FULL: 'full',
  CANCELLED: 'cancelled',
  COMPLETED: 'completed',
} as const;

export const SKILL_LEVELS = {
  BEGINNER: 'beginner',
  INTERMEDIATE: 'intermediate',
  ADVANCED: 'advanced',
  PROFESSIONAL: 'professional',
} as const;

// Venue Constants
export const FIELD_TYPES = {
  FUTSAL: 'futsal',
  FOOTBALL: 'football',
} as const;

// Default Location (Dhaka)
export const DEFAULT_LOCATION = {
  lat: 23.8103,
  lng: 90.4125,
};

export const DEFAULT_RADIUS = 5; // km


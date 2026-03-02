# Kick&Roar Frontend Technical Documentation

**Version:** 1.0  
**Last Updated:** November 17, 2025  
**Project:** Kick&Roar Web Application Frontend  
**Tech Stack:** React 18 + TypeScript + Vite + TailwindCSS + ShadCN/UI

---

## Table of Contents

1. [Project Overview](#project-overview)
2. [Architecture & Design Patterns](#architecture--design-patterns)
3. [Technology Stack Details](#technology-stack-details)
4. [Project Structure](#project-structure)
5. [Core Features Implementation](#core-features-implementation)
6. [Component Architecture](#component-architecture)
7. [State Management Strategy](#state-management-strategy)
8. [API Integration Layer](#api-integration-layer)
9. [Routing & Navigation](#routing--navigation)
10. [Authentication Flow](#authentication-flow)
11. [Map Integration](#map-integration)
12. [File Upload System](#file-upload-system)
13. [Real-time Features](#real-time-features)
14. [Performance Optimization](#performance-optimization)
15. [Mobile Responsiveness](#mobile-responsiveness)
16. [Error Handling](#error-handling)
17. [Testing Strategy](#testing-strategy)
18. [Build & Deployment](#build--deployment)
19. [Development Guidelines](#development-guidelines)

---

## Project Overview

### Purpose

The Kick&Roar frontend is a Progressive Web Application (PWA) that provides an intuitive interface for football enthusiasts in Dhaka to discover, create, and join football matches. Built with mobile-first principles, it delivers a native-like experience on all devices.

### Key Characteristics

- **Mobile-First Design:** Optimized for mobile devices (primary usage platform)
- **Offline Capability:** PWA features for intermittent connectivity
- **Real-time Updates:** Live match status and participant updates
- **Interactive Maps:** Google Maps integration for venue discovery
- **Type-Safe:** Full TypeScript coverage for reliability
- **Component-Based:** Modular, reusable UI components
- **Performance-Optimized:** Code splitting, lazy loading, and caching

---

## Architecture & Design Patterns

### Architectural Principles

```typescript
// Core Architecture Principles
interface ArchitecturalPrinciples {
  separation_of_concerns: {
    presentation: "React Components",
    business_logic: "Custom Hooks & Services",
    state_management: "Zustand Stores",
    api_communication: "API Client Layer"
  },
  design_patterns: [
    "Container/Presenter Pattern",
    "Custom Hooks Pattern",
    "Compound Components",
    "Render Props",
    "Higher-Order Components (HOCs)"
  ],
  code_organization: {
    feature_based: true,
    barrel_exports: true,
    absolute_imports: true
  }
}
```

### Layer Architecture

```
┌─────────────────────────────────────────────────────────┐
│                   Presentation Layer                      │
│  ┌──────────────────────────────────────────────────┐   │
│  │    Pages (Route Components)                       │   │
│  │    Feature Components                             │   │
│  │    UI Components (ShadCN)                         │   │
│  └──────────────────────────────────────────────────┘   │
├───────────────────────────────────────────────────────────┤
│                   Business Logic Layer                    │
│  ┌──────────────────────────────────────────────────┐   │
│  │    Custom Hooks                                   │   │
│  │    Services                                       │   │
│  │    Utilities                                      │   │
│  └──────────────────────────────────────────────────┘   │
├───────────────────────────────────────────────────────────┤
│                   State Management Layer                  │
│  ┌──────────────────────────────────────────────────┐   │
│  │    Zustand Stores                                 │   │
│  │    React Query Cache                              │   │
│  │    Local Storage                                  │   │
│  └──────────────────────────────────────────────────┘   │
├───────────────────────────────────────────────────────────┤
│                   Data Access Layer                       │
│  ┌──────────────────────────────────────────────────┐   │
│  │    API Client (Axios)                             │   │
│  │    WebSocket Client                               │   │
│  │    External APIs (Google Maps)                    │   │
│  └──────────────────────────────────────────────────┘   │
└───────────────────────────────────────────────────────────┘
```

---

## Technology Stack Details

### Core Dependencies

```json
{
  "dependencies": {
    // Core
    "react": "^18.3.0",
    "react-dom": "^18.3.0",
    "typescript": "^5.3.0",
    
    // Routing
    "react-router-dom": "^6.21.0",
    
    // State Management
    "zustand": "^4.5.0",
    "immer": "^10.0.3",
    
    // API & Data Fetching
    "@tanstack/react-query": "^5.17.0",
    "axios": "^1.6.0",
    
    // Forms & Validation
    "react-hook-form": "^7.49.0",
    "zod": "^3.22.0",
    "@hookform/resolvers": "^3.3.0",
    
    // UI Components
    "@radix-ui/react-*": "latest",
    "class-variance-authority": "^0.7.0",
    "clsx": "^2.1.0",
    "tailwind-merge": "^2.2.0",
    
    // Maps
    "@react-google-maps/api": "^2.19.0",
    
    // Utilities
    "date-fns": "^3.3.0",
    "lodash-es": "^4.17.21",
    
    // Icons
    "lucide-react": "^0.309.0",
    
    // PWA
    "vite-plugin-pwa": "^0.17.4",
    "workbox-window": "^7.0.0"
  },
  "devDependencies": {
    // Build Tools
    "vite": "^5.0.0",
    "@vitejs/plugin-react": "^4.2.0",
    
    // TypeScript
    "@types/react": "^18.2.0",
    "@types/react-dom": "^18.2.0",
    "@types/lodash-es": "^4.17.0",
    
    // Linting & Formatting
    "eslint": "^8.56.0",
    "eslint-plugin-react-hooks": "^4.6.0",
    "prettier": "^3.2.0",
    
    // CSS
    "tailwindcss": "^3.4.0",
    "autoprefixer": "^10.4.0",
    "postcss": "^8.4.0",
    
    // Testing
    "vitest": "^1.2.0",
    "@testing-library/react": "^14.1.0",
    "@testing-library/jest-dom": "^6.2.0",
    "msw": "^2.1.0"
  }
}
```

---

## Project Structure

### Complete Directory Structure

```
Kick&Roar-frontend/
├── public/
│   ├── manifest.json           # PWA manifest
│   ├── favicon.ico
│   ├── icons/                  # PWA icons
│   │   ├── icon-192x192.png
│   │   ├── icon-512x512.png
│   │   └── apple-touch-icon.png
│   └── robots.txt
│
├── src/
│   ├── api/                    # API layer
│   │   ├── client.ts           # Axios instance
│   │   ├── endpoints/
│   │   │   ├── auth.api.ts
│   │   │   ├── match.api.ts
│   │   │   ├── venue.api.ts
│   │   │   ├── user.api.ts
│   │   │   └── upload.api.ts
│   │   ├── interceptors.ts     # Request/Response interceptors
│   │   └── types.ts            # API types
│   │
│   ├── components/             # Reusable components
│   │   ├── ui/                 # ShadCN components
│   │   │   ├── button.tsx
│   │   │   ├── card.tsx
│   │   │   ├── dialog.tsx
│   │   │   ├── form.tsx
│   │   │   ├── input.tsx
│   │   │   ├── select.tsx
│   │   │   ├── sheet.tsx
│   │   │   ├── skeleton.tsx
│   │   │   ├── toast.tsx
│   │   │   └── ...
│   │   │
│   │   ├── layout/             # Layout components
│   │   │   ├── Header.tsx
│   │   │   ├── MobileNav.tsx
│   │   │   ├── DesktopNav.tsx
│   │   │   ├── Footer.tsx
│   │   │   ├── PageContainer.tsx
│   │   │   └── Layout.tsx
│   │   │
│   │   ├── common/             # Common components
│   │   │   ├── LoadingSpinner.tsx
│   │   │   ├── ErrorBoundary.tsx
│   │   │   ├── EmptyState.tsx
│   │   │   ├── Avatar.tsx
│   │   │   ├── Rating.tsx
│   │   │   └── Badge.tsx
│   │   │
│   │   └── index.ts            # Barrel export
│   │
│   ├── features/               # Feature modules
│   │   ├── auth/
│   │   │   ├── components/
│   │   │   │   ├── LoginForm.tsx
│   │   │   │   ├── SignupForm.tsx
│   │   │   │   ├── ForgotPassword.tsx
│   │   │   │   └── OTPVerification.tsx
│   │   │   ├── hooks/
│   │   │   │   ├── useAuth.ts
│   │   │   │   └── useAuthRedirect.ts
│   │   │   ├── services/
│   │   │   │   └── auth.service.ts
│   │   │   └── types.ts
│   │   │
│   │   ├── match/
│   │   │   ├── components/
│   │   │   │   ├── MatchCard.tsx
│   │   │   │   ├── MatchList.tsx
│   │   │   │   ├── MatchDetails.tsx
│   │   │   │   ├── CreateMatchForm.tsx
│   │   │   │   ├── MatchFilters.tsx
│   │   │   │   ├── JoinMatchButton.tsx
│   │   │   │   └── ParticipantList.tsx
│   │   │   ├── hooks/
│   │   │   │   ├── useMatches.ts
│   │   │   │   ├── useMatchDetails.ts
│   │   │   │   └── useJoinMatch.ts
│   │   │   ├── services/
│   │   │   │   └── match.service.ts
│   │   │   └── types.ts
│   │   │
│   │   ├── venue/
│   │   │   ├── components/
│   │   │   │   ├── VenueMap.tsx
│   │   │   │   ├── VenueCard.tsx
│   │   │   │   ├── VenueDetails.tsx
│   │   │   │   ├── VenueMarker.tsx
│   │   │   │   └── VenueSearch.tsx
│   │   │   ├── hooks/
│   │   │   │   ├── useVenues.ts
│   │   │   │   └── useNearbyVenues.ts
│   │   │   ├── services/
│   │   │   │   └── venue.service.ts
│   │   │   └── types.ts
│   │   │
│   │   ├── user/
│   │   │   ├── components/
│   │   │   │   ├── UserProfile.tsx
│   │   │   │   ├── EditProfile.tsx
│   │   │   │   ├── UserStats.tsx
│   │   │   │   ├── MatchHistory.tsx
│   │   │   │   └── ProfileImage.tsx
│   │   │   ├── hooks/
│   │   │   │   └── useUserProfile.ts
│   │   │   └── types.ts
│   │   │
│   │   └── notification/
│   │       ├── components/
│   │       │   ├── NotificationList.tsx
│   │       │   ├── NotificationItem.tsx
│   │       │   └── NotificationBell.tsx
│   │       ├── hooks/
│   │       │   └── useNotifications.ts
│   │       └── types.ts
│   │
│   ├── pages/                 # Page components
│   │   ├── Home.tsx
│   │   ├── Login.tsx
│   │   ├── Signup.tsx
│   │   ├── Matches.tsx
│   │   ├── MatchDetails.tsx
│   │   ├── CreateMatch.tsx
│   │   ├── Venues.tsx
│   │   ├── Profile.tsx
│   │   ├── Dashboard.tsx
│   │   └── NotFound.tsx
│   │
│   ├── hooks/                 # Global hooks
│   │   ├── useGeolocation.ts
│   │   ├── useDebounce.ts
│   │   ├── useLocalStorage.ts
│   │   ├── useMediaQuery.ts
│   │   ├── useOnlineStatus.ts
│   │   └── useIntersectionObserver.ts
│   │
│   ├── store/                 # Zustand stores
│   │   ├── authStore.ts
│   │   ├── matchStore.ts
│   │   ├── venueStore.ts
│   │   ├── uiStore.ts
│   │   └── notificationStore.ts
│   │
│   ├── lib/                   # Libraries & utilities
│   │   ├── utils.ts           # Utility functions
│   │   ├── constants.ts       # App constants
│   │   ├── validators.ts      # Zod schemas
│   │   └── cn.ts             # Class name utility
│   │
│   ├── types/                 # Global types
│   │   ├── global.d.ts
│   │   ├── api.types.ts
│   │   └── models.ts
│   │
│   ├── styles/                # Global styles
│   │   ├── globals.css
│   │   └── tailwind.css
│   │
│   ├── App.tsx                # Main app component
│   ├── main.tsx               # Entry point
│   ├── router.tsx             # Route configuration
│   └── vite-env.d.ts
│
├── .env.example
├── .eslintrc.json
├── .gitignore
├── .prettierrc
├── index.html
├── package.json
├── postcss.config.js
├── tailwind.config.js
├── tsconfig.json
├── tsconfig.node.json
├── vite.config.ts
└── README.md
```

---

## Core Features Implementation

### 1. User Authentication System

```typescript
// features/auth/hooks/useAuth.ts
import { create } from 'zustand';
import { persist } from 'zustand/middleware';
import { immer } from 'zustand/middleware/immer';

interface AuthState {
  user: User | null;
  token: string | null;
  isAuthenticated: boolean;
  isLoading: boolean;
  error: string | null;
  
  // Actions
  login: (credentials: LoginCredentials) => Promise<void>;
  signup: (data: SignupData) => Promise<void>;
  logout: () => void;
  refreshToken: () => Promise<void>;
  updateUser: (updates: Partial<User>) => void;
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
          api.defaults.headers.common['Authorization'] = `Bearer ${response.token}`;
          
          // Navigate to dashboard
          router.navigate('/dashboard');
        } catch (error) {
          set((state) => {
            state.error = error.message;
            state.isLoading = false;
          });
        }
      },

      logout: () => {
        set((state) => {
          state.user = null;
          state.token = null;
          state.isAuthenticated = false;
        });
        
        delete api.defaults.headers.common['Authorization'];
        router.navigate('/login');
      },

      // ... other methods
    })),
    {
      name: 'auth-storage',
      partialize: (state) => ({
        user: state.user,
        token: state.token,
      }),
    }
  )
);
```

### 2. Match Discovery & Management

```typescript
// features/match/components/MatchCard.tsx
import { Card, CardContent, CardFooter, CardHeader } from '@/components/ui/card';
import { Button } from '@/components/ui/button';
import { Badge } from '@/components/ui/badge';
import { Calendar, MapPin, Users, DollarSign } from 'lucide-react';
import { format } from 'date-fns';
import { useJoinMatch } from '../hooks/useJoinMatch';

interface MatchCardProps {
  match: Match;
  onViewDetails?: () => void;
}

export const MatchCard: React.FC<MatchCardProps> = ({ match, onViewDetails }) => {
  const { mutate: joinMatch, isLoading } = useJoinMatch();
  
  const isFullyBooked = match.current_players >= match.max_players;
  const spotsLeft = match.max_players - match.current_players;
  
  return (
    <Card className="hover:shadow-lg transition-shadow">
      <CardHeader>
        <div className="flex justify-between items-start">
          <div>
            <h3 className="text-lg font-semibold">{match.title}</h3>
            <p className="text-sm text-muted-foreground">{match.venue.name}</p>
          </div>
          <Badge variant={isFullyBooked ? "secondary" : "success"}>
            {match.status}
          </Badge>
        </div>
      </CardHeader>
      
      <CardContent className="space-y-3">
        <div className="flex items-center gap-2 text-sm">
          <Calendar className="h-4 w-4 text-muted-foreground" />
          <span>{format(new Date(match.match_date), 'PPP')}</span>
          <span className="font-medium">{match.start_time}</span>
        </div>
        
        <div className="flex items-center gap-2 text-sm">
          <MapPin className="h-4 w-4 text-muted-foreground" />
          <span>{match.venue.address}</span>
          {match.venue.distance_km && (
            <Badge variant="outline">{match.venue.distance_km} km</Badge>
          )}
        </div>
        
        <div className="flex items-center justify-between">
          <div className="flex items-center gap-2">
            <Users className="h-4 w-4 text-muted-foreground" />
            <span className="text-sm">
              {match.current_players}/{match.max_players} players
            </span>
            {!isFullyBooked && (
              <span className="text-xs text-muted-foreground">
                ({spotsLeft} spots left)
              </span>
            )}
          </div>
          
          <div className="flex items-center gap-1">
            <DollarSign className="h-4 w-4 text-muted-foreground" />
            <span className="font-semibold">{match.cost_per_player} BDT</span>
          </div>
        </div>
        
        {match.skill_level_required && (
          <Badge variant="outline" className="w-fit">
            {match.skill_level_required}
          </Badge>
        )}
      </CardContent>
      
      <CardFooter className="gap-2">
        <Button
          className="flex-1"
          onClick={() => joinMatch(match.id)}
          disabled={isFullyBooked || isLoading}
        >
          {isFullyBooked ? 'Fully Booked' : 'Join Match'}
        </Button>
        <Button
          variant="outline"
          onClick={onViewDetails}
        >
          View Details
        </Button>
      </CardFooter>
    </Card>
  );
};
```

### 3. Interactive Map with Venues

```typescript
// features/venue/components/VenueMap.tsx
import { useState, useCallback, useMemo } from 'react';
import { GoogleMap, LoadScript, Marker, InfoWindow } from '@react-google-maps/api';
import { useNearbyVenues } from '../hooks/useNearbyVenues';
import { VenueCard } from './VenueCard';

const mapContainerStyle = {
  width: '100%',
  height: '100vh',
};

const defaultCenter = {
  lat: 23.8103,  // Dhaka center
  lng: 90.4125,
};

const mapOptions = {
  disableDefaultUI: false,
  zoomControl: true,
  streetViewControl: false,
  mapTypeControl: false,
  fullscreenControl: true,
  styles: [
    {
      featureType: "poi.business",
      elementType: "labels",
      stylers: [{ visibility: "off" }],
    },
  ],
};

export const VenueMap: React.FC = () => {
  const [selectedVenue, setSelectedVenue] = useState<Venue | null>(null);
  const [mapCenter, setMapCenter] = useState(defaultCenter);
  const [mapZoom, setMapZoom] = useState(13);
  
  const { data: venues, isLoading } = useNearbyVenues({
    latitude: mapCenter.lat,
    longitude: mapCenter.lng,
    radius: 5,
  });

  const handleMapChange = useCallback((map: google.maps.Map) => {
    const center = map.getCenter();
    if (center) {
      setMapCenter({
        lat: center.lat(),
        lng: center.lng(),
      });
    }
  }, []);

  const markers = useMemo(() => {
    return venues?.map((venue) => (
      <Marker
        key={venue.id}
        position={{
          lat: venue.latitude,
          lng: venue.longitude,
        }}
        onClick={() => setSelectedVenue(venue)}
        icon={{
          url: venue.field_type === 'futsal' 
            ? '/icons/futsal-marker.png'
            : '/icons/football-marker.png',
          scaledSize: new google.maps.Size(40, 40),
        }}
      />
    ));
  }, [venues]);

  return (
    <LoadScript googleMapsApiKey={import.meta.env.VITE_GOOGLE_MAPS_API_KEY}>
      <GoogleMap
        mapContainerStyle={mapContainerStyle}
        center={mapCenter}
        zoom={mapZoom}
        options={mapOptions}
        onDragEnd={(e) => handleMapChange(e)}
      >
        {markers}
        
        {selectedVenue && (
          <InfoWindow
            position={{
              lat: selectedVenue.latitude,
              lng: selectedVenue.longitude,
            }}
            onCloseClick={() => setSelectedVenue(null)}
          >
            <VenueCard venue={selectedVenue} compact />
          </InfoWindow>
        )}
      </GoogleMap>
    </LoadScript>
  );
};
```

### 4. Real-time Notifications

```typescript
// features/notification/hooks/useNotifications.ts
import { useEffect } from 'react';
import { useQuery, useMutation, useQueryClient } from '@tanstack/react-query';
import { toast } from '@/components/ui/use-toast';
import { notificationApi } from '@/api/endpoints/notification.api';
import { useAuthStore } from '@/store/authStore';

export const useNotifications = () => {
  const queryClient = useQueryClient();
  const { user } = useAuthStore();

  // Fetch notifications
  const { data, isLoading } = useQuery({
    queryKey: ['notifications'],
    queryFn: notificationApi.getAll,
    enabled: !!user,
    refetchInterval: 30000, // Refetch every 30 seconds
  });

  // Mark as read
  const markAsRead = useMutation({
    mutationFn: notificationApi.markAsRead,
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['notifications'] });
    },
  });

  // WebSocket connection for real-time notifications
  useEffect(() => {
    if (!user) return;

    const ws = new WebSocket(
      `${import.meta.env.VITE_WS_URL}/notifications?userId=${user.id}`
    );

    ws.onmessage = (event) => {
      const notification = JSON.parse(event.data);
      
      // Update cache
      queryClient.setQueryData(['notifications'], (old: Notification[]) => {
        return [notification, ...old];
      });

      // Show toast
      toast({
        title: notification.title,
        description: notification.message,
        action: notification.action_url ? (
          <ToastAction altText="View" onClick={() => {
            window.location.href = notification.action_url;
          }}>
            View
          </ToastAction>
        ) : undefined,
      });
    };

    return () => {
      ws.close();
    };
  }, [user, queryClient]);

  return {
    notifications: data || [],
    unreadCount: data?.filter(n => !n.is_read).length || 0,
    isLoading,
    markAsRead: markAsRead.mutate,
  };
};
```

---

## Component Architecture

### Component Categories

#### 1. Page Components
- Entry points for routes
- Compose feature components
- Handle data fetching
- Manage page-level state

#### 2. Feature Components
- Domain-specific functionality
- Self-contained modules
- Include own hooks and services
- Can be composed into pages

#### 3. UI Components (ShadCN)
- Pure presentation components
- Highly reusable
- No business logic
- Styled with Tailwind

#### 4. Layout Components
- Structural components
- Navigation, headers, footers
- Consistent across pages

### Component Best Practices

```typescript
// Example: Well-structured component
// features/match/components/MatchList.tsx

import { useState, useMemo } from 'react';
import { useInView } from 'react-intersection-observer';
import { MatchCard } from './MatchCard';
import { MatchFilters } from './MatchFilters';
import { LoadingSpinner } from '@/components/common/LoadingSpinner';
import { EmptyState } from '@/components/common/EmptyState';
import { useMatches } from '../hooks/useMatches';

interface MatchListProps {
  initialFilters?: MatchFilters;
  onMatchSelect?: (match: Match) => void;
}

export const MatchList: React.FC<MatchListProps> = ({
  initialFilters,
  onMatchSelect,
}) => {
  // State management
  const [filters, setFilters] = useState<MatchFilters>(
    initialFilters || defaultFilters
  );

  // Data fetching with infinite scroll
  const {
    data,
    fetchNextPage,
    hasNextPage,
    isFetchingNextPage,
    isLoading,
    isError,
  } = useMatches(filters);

  // Infinite scroll trigger
  const { ref, inView } = useInView({
    threshold: 0,
  });

  useEffect(() => {
    if (inView && hasNextPage && !isFetchingNextPage) {
      fetchNextPage();
    }
  }, [inView, hasNextPage, isFetchingNextPage, fetchNextPage]);

  // Memoized match list
  const matches = useMemo(() => {
    return data?.pages.flatMap((page) => page.matches) || [];
  }, [data]);

  // Loading state
  if (isLoading) {
    return (
      <div className="flex justify-center items-center h-64">
        <LoadingSpinner />
      </div>
    );
  }

  // Error state
  if (isError) {
    return (
      <EmptyState
        icon="alert-circle"
        title="Error loading matches"
        description="Please try again later"
        action={{
          label: "Retry",
          onClick: () => window.location.reload(),
        }}
      />
    );
  }

  // Empty state
  if (matches.length === 0) {
    return (
      <EmptyState
        icon="calendar-x"
        title="No matches found"
        description="Try adjusting your filters or create a new match"
        action={{
          label: "Create Match",
          onClick: () => navigate('/matches/create'),
        }}
      />
    );
  }

  return (
    <div className="space-y-6">
      <MatchFilters
        filters={filters}
        onChange={setFilters}
      />
      
      <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
        {matches.map((match) => (
          <MatchCard
            key={match.id}
            match={match}
            onViewDetails={() => onMatchSelect?.(match)}
          />
        ))}
      </div>

      {/* Infinite scroll trigger */}
      <div ref={ref} className="h-10">
        {isFetchingNextPage && <LoadingSpinner />}
      </div>
    </div>
  );
};
```

---

## State Management Strategy

### Zustand Store Pattern

```typescript
// store/matchStore.ts
import { create } from 'zustand';
import { devtools } from 'zustand/middleware';
import { immer } from 'zustand/middleware/immer';

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
  joinMatch: (matchId: string, userId: string) => Promise<void>;
  leaveMatch: (matchId: string, userId: string) => Promise<void>;
  createMatch: (data: CreateMatchData) => Promise<Match>;
}

const useMatchStore = create<MatchState>()(
  devtools(
    immer((set, get) => ({
      // Initial state
      matches: [],
      currentMatch: null,
      filters: {
        status: 'open',
        radius: 5,
        sortBy: 'date',
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

      joinMatch: async (matchId, userId) => {
        try {
          const response = await matchApi.join(matchId);
          
          set((state) => {
            const match = state.matches.find((m) => m.id === matchId);
            if (match) {
              match.current_players += 1;
              match.participants.push({ user_id: userId });
              
              if (match.current_players >= match.max_players) {
                match.status = 'full';
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

export default useMatchStore;
```

### React Query Integration

```typescript
// hooks/useMatches.ts
import { useInfiniteQuery } from '@tanstack/react-query';
import { matchApi } from '@/api/endpoints/match.api';

export const useMatches = (filters: MatchFilters) => {
  return useInfiniteQuery({
    queryKey: ['matches', filters],
    queryFn: ({ pageParam = 1 }) =>
      matchApi.getMatches({
        ...filters,
        page: pageParam,
        per_page: 20,
      }),
    getNextPageParam: (lastPage, pages) => {
      if (lastPage.pagination.page < lastPage.pagination.total_pages) {
        return lastPage.pagination.page + 1;
      }
      return undefined;
    },
    staleTime: 5 * 60 * 1000, // 5 minutes
    cacheTime: 10 * 60 * 1000, // 10 minutes
  });
};
```

---

## API Integration Layer

### Axios Client Configuration

```typescript
// api/client.ts
import axios, { AxiosInstance, AxiosRequestConfig, AxiosResponse } from 'axios';
import { toast } from '@/components/ui/use-toast';

const config: AxiosRequestConfig = {
  baseURL: import.meta.env.VITE_API_URL || 'http://localhost:8080/api/v1',
  timeout: 15000,
  headers: {
    'Content-Type': 'application/json',
  },
};

const apiClient: AxiosInstance = axios.create(config);

// Request interceptor
apiClient.interceptors.request.use(
  (config) => {
    const token = localStorage.getItem('auth-token');
    
    if (token) {
      config.headers.Authorization = `Bearer ${token}`;
    }

    // Add request ID for tracking
    config.headers['X-Request-ID'] = generateRequestId();
    
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

    // Handle 401 - Token expired
    if (error.response?.status === 401 && !originalRequest._retry) {
      originalRequest._retry = true;

      try {
        const refreshToken = localStorage.getItem('refresh-token');
        const response = await authApi.refreshToken(refreshToken);
        
        localStorage.setItem('auth-token', response.data.token);
        originalRequest.headers.Authorization = `Bearer ${response.data.token}`;
        
        return apiClient(originalRequest);
      } catch (refreshError) {
        // Refresh failed, logout user
        useAuthStore.getState().logout();
        return Promise.reject(refreshError);
      }
    }

    // Handle other errors
    if (error.response?.status === 429) {
      toast({
        title: 'Rate Limited',
        description: 'Too many requests. Please slow down.',
        variant: 'destructive',
      });
    } else if (error.response?.status >= 500) {
      toast({
        title: 'Server Error',
        description: 'Something went wrong. Please try again later.',
        variant: 'destructive',
      });
    }

    return Promise.reject(error);
  }
);

export default apiClient;
```

### API Service Layer

```typescript
// api/endpoints/match.api.ts
import apiClient from '../client';
import type { Match, CreateMatchData, MatchFilters, PaginatedResponse } from '@/types';

export const matchApi = {
  // Get matches with filters
  async getMatches(filters: MatchFilters): Promise<PaginatedResponse<Match>> {
    const params = new URLSearchParams();
    
    Object.entries(filters).forEach(([key, value]) => {
      if (value !== undefined && value !== null) {
        params.append(key, String(value));
      }
    });

    const response = await apiClient.get<PaginatedResponse<Match>>(
      `/matches?${params.toString()}`
    );
    
    return response.data;
  },

  // Get single match
  async getMatch(id: string): Promise<Match> {
    const response = await apiClient.get<Match>(`/matches/${id}`);
    return response.data;
  },

  // Create match
  async create(data: CreateMatchData): Promise<Match> {
    const response = await apiClient.post<Match>('/matches', data);
    return response.data;
  },

  // Update match
  async update(id: string, data: Partial<Match>): Promise<Match> {
    const response = await apiClient.put<Match>(`/matches/${id}`, data);
    return response.data;
  },

  // Delete match
  async delete(id: string): Promise<void> {
    await apiClient.delete(`/matches/${id}`);
  },

  // Join match
  async join(matchId: string): Promise<void> {
    await apiClient.post(`/matches/${matchId}/join`);
  },

  // Leave match
  async leave(matchId: string): Promise<void> {
    await apiClient.post(`/matches/${matchId}/leave`);
  },

  // Get participants
  async getParticipants(matchId: string): Promise<Participant[]> {
    const response = await apiClient.get<Participant[]>(
      `/matches/${matchId}/participants`
    );
    return response.data;
  },
};
```

---

## Routing & Navigation

### Route Configuration

```typescript
// router.tsx
import { createBrowserRouter, RouterProvider, Navigate } from 'react-router-dom';
import { Layout } from '@/components/layout/Layout';
import { ProtectedRoute } from '@/components/auth/ProtectedRoute';

// Lazy load pages for code splitting
const Home = lazy(() => import('@/pages/Home'));
const Login = lazy(() => import('@/pages/Login'));
const Signup = lazy(() => import('@/pages/Signup'));
const Matches = lazy(() => import('@/pages/Matches'));
const MatchDetails = lazy(() => import('@/pages/MatchDetails'));
const CreateMatch = lazy(() => import('@/pages/CreateMatch'));
const Venues = lazy(() => import('@/pages/Venues'));
const Profile = lazy(() => import('@/pages/Profile'));
const Dashboard = lazy(() => import('@/pages/Dashboard'));
const NotFound = lazy(() => import('@/pages/NotFound'));

const router = createBrowserRouter([
  {
    path: '/',
    element: <Layout />,
    children: [
      {
        index: true,
        element: <Home />,
      },
      {
        path: 'login',
        element: <Login />,
      },
      {
        path: 'signup',
        element: <Signup />,
      },
      {
        path: 'matches',
        element: <ProtectedRoute><Matches /></ProtectedRoute>,
        children: [
          {
            path: ':id',
            element: <MatchDetails />,
          },
        ],
      },
      {
        path: 'matches/create',
        element: <ProtectedRoute><CreateMatch /></ProtectedRoute>,
      },
      {
        path: 'venues',
        element: <Venues />,
      },
      {
        path: 'dashboard',
        element: <ProtectedRoute><Dashboard /></ProtectedRoute>,
      },
      {
        path: 'profile',
        element: <ProtectedRoute><Profile /></ProtectedRoute>,
      },
      {
        path: '*',
        element: <NotFound />,
      },
    ],
  },
]);

export const AppRouter: React.FC = () => {
  return (
    <Suspense fallback={<LoadingScreen />}>
      <RouterProvider router={router} />
    </Suspense>
  );
};
```

### Protected Routes

```typescript
// components/auth/ProtectedRoute.tsx
import { Navigate, useLocation } from 'react-router-dom';
import { useAuthStore } from '@/store/authStore';

interface ProtectedRouteProps {
  children: React.ReactNode;
  requiredRole?: UserRole;
  redirectTo?: string;
}

export const ProtectedRoute: React.FC<ProtectedRouteProps> = ({
  children,
  requiredRole,
  redirectTo = '/login',
}) => {
  const { isAuthenticated, user } = useAuthStore();
  const location = useLocation();

  if (!isAuthenticated) {
    return (
      <Navigate
        to={redirectTo}
        state={{ from: location.pathname }}
        replace
      />
    );
  }

  if (requiredRole && user?.role !== requiredRole) {
    return <Navigate to="/unauthorized" replace />;
  }

  return <>{children}</>;
};
```

---

## Mobile Responsiveness

### Mobile-First Design Strategy

```typescript
// components/layout/MobileNav.tsx
import { useState } from 'react';
import { NavLink, useLocation } from 'react-router-dom';
import { Home, Calendar, MapPin, User, Plus } from 'lucide-react';
import { cn } from '@/lib/utils';
import { CreateMatchDialog } from '@/features/match/components/CreateMatchDialog';

const navItems = [
  { path: '/', icon: Home, label: 'Home' },
  { path: '/matches', icon: Calendar, label: 'Matches' },
  { path: '/venues', icon: MapPin, label: 'Venues' },
  { path: '/profile', icon: User, label: 'Profile' },
];

export const MobileNav: React.FC = () => {
  const [isCreateOpen, setIsCreateOpen] = useState(false);
  const location = useLocation();

  return (
    <>
      <nav className="fixed bottom-0 left-0 right-0 z-50 bg-background border-t md:hidden">
        <div className="grid grid-cols-5 h-16">
          {navItems.slice(0, 2).map((item) => (
            <NavLink
              key={item.path}
              to={item.path}
              className={({ isActive }) =>
                cn(
                  "flex flex-col items-center justify-center space-y-1",
                  "text-muted-foreground hover:text-primary",
                  isActive && "text-primary"
                )
              }
            >
              <item.icon className="h-5 w-5" />
              <span className="text-xs">{item.label}</span>
            </NavLink>
          ))}

          {/* Center Create Button */}
          <button
            onClick={() => setIsCreateOpen(true)}
            className="flex flex-col items-center justify-center"
          >
            <div className="bg-primary text-primary-foreground rounded-full p-3">
              <Plus className="h-6 w-6" />
            </div>
          </button>

          {navItems.slice(2).map((item) => (
            <NavLink
              key={item.path}
              to={item.path}
              className={({ isActive }) =>
                cn(
                  "flex flex-col items-center justify-center space-y-1",
                  "text-muted-foreground hover:text-primary",
                  isActive && "text-primary"
                )
              }
            >
              <item.icon className="h-5 w-5" />
              <span className="text-xs">{item.label}</span>
            </NavLink>
          ))}
        </div>
      </nav>

      <CreateMatchDialog
        open={isCreateOpen}
        onOpenChange={setIsCreateOpen}
      />
    </>
  );
};
```

### Responsive Utilities

```typescript
// hooks/useMediaQuery.ts
import { useState, useEffect } from 'react';

export const useMediaQuery = (query: string): boolean => {
  const [matches, setMatches] = useState(() => {
    if (typeof window !== 'undefined') {
      return window.matchMedia(query).matches;
    }
    return false;
  });

  useEffect(() => {
    const mediaQuery = window.matchMedia(query);
    
    const handler = (event: MediaQueryListEvent) => {
      setMatches(event.matches);
    };

    // Modern browsers
    if (mediaQuery.addEventListener) {
      mediaQuery.addEventListener('change', handler);
      return () => mediaQuery.removeEventListener('change', handler);
    }
    
    // Legacy browsers
    mediaQuery.addListener(handler);
    return () => mediaQuery.removeListener(handler);
  }, [query]);

  return matches;
};

// Usage in components
export const ResponsiveComponent: React.FC = () => {
  const isMobile = useMediaQuery('(max-width: 768px)');
  const isTablet = useMediaQuery('(min-width: 768px) and (max-width: 1024px)');
  const isDesktop = useMediaQuery('(min-width: 1024px)');

  return (
    <div>
      {isMobile && <MobileView />}
      {isTablet && <TabletView />}
      {isDesktop && <DesktopView />}
    </div>
  );
};
```

---

## Performance Optimization

### Code Splitting & Lazy Loading

```typescript
// App.tsx - Route-based code splitting
import { lazy, Suspense } from 'react';
import { ErrorBoundary } from '@/components/common/ErrorBoundary';

// Lazy load heavy components
const MapView = lazy(() => 
  import('@/features/venue/components/VenueMap')
    .then(module => ({ default: module.VenueMap }))
);

// Component-level lazy loading
export const App: React.FC = () => {
  return (
    <ErrorBoundary>
      <Suspense fallback={<MapSkeleton />}>
        <MapView />
      </Suspense>
    </ErrorBoundary>
  );
};
```

### Image Optimization

```typescript
// components/common/OptimizedImage.tsx
import { useState, useEffect } from 'react';
import { cn } from '@/lib/utils';

interface OptimizedImageProps {
  src: string;
  alt: string;
  className?: string;
  width?: number;
  height?: number;
  quality?: number;
  lazy?: boolean;
}

export const OptimizedImage: React.FC<OptimizedImageProps> = ({
  src,
  alt,
  className,
  width,
  height,
  quality = 75,
  lazy = true,
}) => {
  const [imageSrc, setImageSrc] = useState<string>('/placeholder.webp');
  const [isLoading, setIsLoading] = useState(true);
  const [error, setError] = useState(false);

  // Generate optimized S3 URL with transforms
  const getOptimizedUrl = (url: string) => {
    if (!url.includes('s3.amazonaws.com')) return url;

    const params = new URLSearchParams({
      w: width?.toString() || 'auto',
      h: height?.toString() || 'auto',
      q: quality.toString(),
      fm: 'webp',
    });

    return `${url}?${params.toString()}`;
  };

  useEffect(() => {
    if (!lazy) {
      loadImage();
    }
  }, [src]);

  const loadImage = () => {
    const img = new Image();
    img.src = getOptimizedUrl(src);
    
    img.onload = () => {
      setImageSrc(img.src);
      setIsLoading(false);
    };

    img.onerror = () => {
      setError(true);
      setIsLoading(false);
    };
  };

  return (
    <div className={cn('relative overflow-hidden', className)}>
      {isLoading && (
        <div className="absolute inset-0 bg-gray-200 animate-pulse" />
      )}
      
      <img
        src={error ? '/error-image.png' : imageSrc}
        alt={alt}
        loading={lazy ? 'lazy' : 'eager'}
        className={cn(
          'w-full h-full object-cover',
          isLoading && 'invisible'
        )}
        onError={() => setError(true)}
      />
    </div>
  );
};
```

### Memoization & Optimization

```typescript
// features/match/components/MatchList.tsx
import { memo, useMemo, useCallback } from 'react';
import { useVirtualizer } from '@tanstack/react-virtual';

// Memoized match card to prevent unnecessary re-renders
const MemoizedMatchCard = memo(MatchCard, (prev, next) => {
  return (
    prev.match.id === next.match.id &&
    prev.match.current_players === next.match.current_players &&
    prev.match.status === next.match.status
  );
});

// Virtual scrolling for large lists
export const VirtualMatchList: React.FC<{ matches: Match[] }> = ({ matches }) => {
  const parentRef = useRef<HTMLDivElement>(null);

  const virtualizer = useVirtualizer({
    count: matches.length,
    getScrollElement: () => parentRef.current,
    estimateSize: () => 200, // Estimated height of each card
    overscan: 5,
  });

  const items = virtualizer.getVirtualItems();

  return (
    <div ref={parentRef} className="h-screen overflow-auto">
      <div
        style={{
          height: `${virtualizer.getTotalSize()}px`,
          width: '100%',
          position: 'relative',
        }}
      >
        {items.map((virtualItem) => (
          <div
            key={virtualItem.key}
            style={{
              position: 'absolute',
              top: 0,
              left: 0,
              width: '100%',
              height: `${virtualItem.size}px`,
              transform: `translateY(${virtualItem.start}px)`,
            }}
          >
            <MemoizedMatchCard match={matches[virtualItem.index]} />
          </div>
        ))}
      </div>
    </div>
  );
};
```

---

## Error Handling

### Global Error Boundary

```typescript
// components/common/ErrorBoundary.tsx
import { Component, ErrorInfo, ReactNode } from 'react';
import { AlertTriangle } from 'lucide-react';
import { Button } from '@/components/ui/button';

interface Props {
  children: ReactNode;
  fallback?: ReactNode;
}

interface State {
  hasError: boolean;
  error: Error | null;
}

export class ErrorBoundary extends Component<Props, State> {
  constructor(props: Props) {
    super(props);
    this.state = { hasError: false, error: null };
  }

  static getDerivedStateFromError(error: Error): State {
    return { hasError: true, error };
  }

  componentDidCatch(error: Error, errorInfo: ErrorInfo) {
    // Log to error reporting service
    console.error('Error caught by boundary:', error, errorInfo);
    
    // Send to monitoring service (e.g., Sentry)
    if (import.meta.env.PROD) {
      // Sentry.captureException(error, { contexts: { react: errorInfo } });
    }
  }

  handleReset = () => {
    this.setState({ hasError: false, error: null });
  };

  render() {
    if (this.state.hasError) {
      if (this.props.fallback) {
        return this.props.fallback;
      }

      return (
        <div className="flex flex-col items-center justify-center min-h-[400px] p-8">
          <AlertTriangle className="h-12 w-12 text-destructive mb-4" />
          <h2 className="text-2xl font-semibold mb-2">Oops! Something went wrong</h2>
          <p className="text-muted-foreground mb-6 text-center max-w-md">
            {this.state.error?.message || 'An unexpected error occurred'}
          </p>
          <div className="flex gap-4">
            <Button onClick={this.handleReset}>Try Again</Button>
            <Button variant="outline" onClick={() => window.location.href = '/'}>
              Go Home
            </Button>
          </div>
        </div>
      );
    }

    return this.props.children;
  }
}
```

### API Error Handling

```typescript
// utils/errorHandler.ts
export class ApiError extends Error {
  constructor(
    public status: number,
    public code: string,
    message: string,
    public details?: any
  ) {
    super(message);
    this.name = 'ApiError';
  }
}

export const handleApiError = (error: any): ApiError => {
  if (error.response) {
    // Server responded with error
    const { status, data } = error.response;
    
    return new ApiError(
      status,
      data.error?.code || 'UNKNOWN_ERROR',
      data.error?.message || 'An error occurred',
      data.error?.details
    );
  }
  
  if (error.request) {
    // Request made but no response
    return new ApiError(
      0,
      'NETWORK_ERROR',
      'Network error. Please check your connection.'
    );
  }
  
  // Something else happened
  return new ApiError(
    0,
    'CLIENT_ERROR',
    error.message || 'An unexpected error occurred'
  );
};

// Usage in components
export const useErrorHandler = () => {
  const { toast } = useToast();
  
  const handleError = useCallback((error: ApiError) => {
    const errorMessages: Record<string, string> = {
      'AUTH_001': 'Invalid email or password',
      'AUTH_002': 'Your session has expired. Please login again.',
      'MATCH_002': 'This match is already full',
      'NETWORK_ERROR': 'Connection error. Please try again.',
    };

    const message = errorMessages[error.code] || error.message;

    toast({
      title: 'Error',
      description: message,
      variant: 'destructive',
    });
  }, [toast]);

  return { handleError };
};
```

---

## Testing Strategy

### Unit Testing

```typescript
// __tests__/components/MatchCard.test.tsx
import { render, screen, fireEvent } from '@testing-library/react';
import { vi } from 'vitest';
import { MatchCard } from '@/features/match/components/MatchCard';

describe('MatchCard', () => {
  const mockMatch = {
    id: '1',
    title: 'Evening Futsal',
    venue: {
      id: 'v1',
      name: 'Sports Complex',
      address: 'Dhanmondi, Dhaka',
    },
    match_date: '2025-11-20',
    start_time: '18:00',
    max_players: 10,
    current_players: 7,
    cost_per_player: 200,
    status: 'open',
  };

  it('renders match information correctly', () => {
    render(<MatchCard match={mockMatch} />);
    
    expect(screen.getByText('Evening Futsal')).toBeInTheDocument();
    expect(screen.getByText('Sports Complex')).toBeInTheDocument();
    expect(screen.getByText('7/10 players')).toBeInTheDocument();
    expect(screen.getByText('200 BDT')).toBeInTheDocument();
  });

  it('disables join button when match is full', () => {
    const fullMatch = { ...mockMatch, current_players: 10 };
    render(<MatchCard match={fullMatch} />);
    
    const joinButton = screen.getByRole('button', { name: /fully booked/i });
    expect(joinButton).toBeDisabled();
  });

  it('calls onViewDetails when view details is clicked', () => {
    const onViewDetails = vi.fn();
    render(<MatchCard match={mockMatch} onViewDetails={onViewDetails} />);
    
    fireEvent.click(screen.getByText('View Details'));
    expect(onViewDetails).toHaveBeenCalled();
  });
});
```

### Integration Testing

```typescript
// __tests__/integration/auth-flow.test.tsx
import { render, screen, waitFor } from '@testing-library/react';
import userEvent from '@testing-library/user-event';
import { setupServer } from 'msw/node';
import { rest } from 'msw';
import { LoginPage } from '@/pages/Login';

const server = setupServer(
  rest.post('/api/v1/auth/login', (req, res, ctx) => {
    return res(
      ctx.json({
        user: { id: '1', email: 'test@example.com' },
        token: 'fake-jwt-token',
      })
    );
  })
);

beforeAll(() => server.listen());
afterEach(() => server.resetHandlers());
afterAll(() => server.close());

describe('Authentication Flow', () => {
  it('logs in user successfully', async () => {
    const user = userEvent.setup();
    render(<LoginPage />);

    await user.type(screen.getByLabelText(/email/i), 'test@example.com');
    await user.type(screen.getByLabelText(/password/i), 'password123');
    await user.click(screen.getByRole('button', { name: /login/i }));

    await waitFor(() => {
      expect(window.location.pathname).toBe('/dashboard');
    });
  });
});
```

---

## Build & Deployment

### Vite Configuration

```typescript
// vite.config.ts
import { defineConfig, loadEnv } from 'vite';
import react from '@vitejs/plugin-react';
import path from 'path';
import { VitePWA } from 'vite-plugin-pwa';

export default defineConfig(({ mode }) => {
  const env = loadEnv(mode, process.cwd());

  return {
    plugins: [
      react(),
      VitePWA({
        registerType: 'autoUpdate',
        includeAssets: ['favicon.ico', 'robots.txt'],
        manifest: {
          name: 'Kick&Roar',
          short_name: 'Kick&Roar',
          description: 'Connect, Play, Share',
          theme_color: '#4CAF50',
          background_color: '#ffffff',
          display: 'standalone',
          icons: [
            {
              src: '/icon-192x192.png',
              sizes: '192x192',
              type: 'image/png',
            },
            {
              src: '/icon-512x512.png',
              sizes: '512x512',
              type: 'image/png',
            },
          ],
        },
        workbox: {
          globPatterns: ['**/*.{js,css,html,ico,png,svg,webp}'],
          runtimeCaching: [
            {
              urlPattern: /^https:\/\/api\.Kick&Roar\.com/,
              handler: 'NetworkFirst',
              options: {
                cacheName: 'api-cache',
                expiration: {
                  maxEntries: 100,
                  maxAgeSeconds: 60 * 5, // 5 minutes
                },
              },
            },
          ],
        },
      }),
    ],
    resolve: {
      alias: {
        '@': path.resolve(__dirname, './src'),
      },
    },
    build: {
      rollupOptions: {
        output: {
          manualChunks: {
            'react-vendor': ['react', 'react-dom', 'react-router-dom'],
            'ui-vendor': ['@radix-ui/react-dialog', '@radix-ui/react-dropdown-menu'],
            'map-vendor': ['@react-google-maps/api'],
          },
        },
      },
      sourcemap: mode === 'production' ? false : true,
      minify: 'esbuild',
      target: 'es2015',
    },
    server: {
      port: 3000,
      proxy: {
        '/api': {
          target: env.VITE_API_URL || 'http://localhost:8080',
          changeOrigin: true,
        },
      },
    },
  };
});
```

### Deployment Scripts

```json
// package.json
{
  "scripts": {
    "dev": "vite",
    "build": "tsc && vite build",
    "build:staging": "tsc && vite build --mode staging",
    "build:production": "tsc && vite build --mode production",
    "preview": "vite preview",
    "test": "vitest",
    "test:ui": "vitest --ui",
    "test:coverage": "vitest --coverage",
    "lint": "eslint . --ext ts,tsx --report-unused-disable-directives",
    "format": "prettier --write 'src/**/*.{ts,tsx,css}'",
    "type-check": "tsc --noEmit",
    "analyze": "vite build --mode analyze"
  }
}
```

### Render Deployment Configuration

```yaml
# render.yaml
services:
  - type: web
    name: Kick&Roar-frontend
    env: static
    buildCommand: yarn install && yarn build
    staticPublishPath: ./dist
    pullRequestPreviewsEnabled: true
    headers:
      - path: /*
        name: X-Frame-Options
        value: DENY
      - path: /*
        name: X-Content-Type-Options
        value: nosniff
      - path: /*
        name: X-XSS-Protection
        value: 1; mode=block
    routes:
      - type: rewrite
        source: /*
        destination: /index.html
    envVars:
      - key: VITE_API_URL
        value: https://Kick&Roar-api.onrender.com/api/v1
      - key: VITE_GOOGLE_MAPS_API_KEY
        sync: false
      - key: VITE_S3_BUCKET_URL
        value: https://Kick&Roar-production.s3.ap-south-1.amazonaws.com
```

---

## Development Guidelines

### Code Style Guide

```typescript
// ESLint Configuration
// .eslintrc.json
{
  "extends": [
    "eslint:recommended",
    "plugin:@typescript-eslint/recommended",
    "plugin:react-hooks/recommended",
    "plugin:react/recommended"
  ],
  "rules": {
    "@typescript-eslint/explicit-function-return-type": "off",
    "@typescript-eslint/no-explicit-any": "warn",
    "react/react-in-jsx-scope": "off",
    "react/prop-types": "off",
    "no-console": ["warn", { "allow": ["warn", "error"] }],
    "prefer-const": "error",
    "no-unused-vars": "off",
    "@typescript-eslint/no-unused-vars": ["error", {
      "argsIgnorePattern": "^_",
      "varsIgnorePattern": "^_"
    }]
  }
}
```

### Git Workflow

```bash
# Branch naming convention
feature/FM-123-match-creation
bugfix/FM-456-join-match-error
hotfix/FM-789-critical-auth-fix

# Commit message format
feat: Add match creation form
fix: Resolve join match race condition
docs: Update API documentation
style: Format code with prettier
refactor: Extract match logic to custom hook
test: Add unit tests for MatchCard
chore: Update dependencies
```

### Development Setup

```bash
# 1. Clone repository
git clone https://github.com/shoaibhassan/Kick&Roar-frontend.git
cd Kick&Roar-frontend

# 2. Install dependencies
yarn install

# 3. Setup environment variables
cp .env.example .env.local
# Edit .env.local with your values

# 4. Start development server
yarn dev

# 5. Run tests
yarn test

# 6. Build for production
yarn build
```

---

## Summary

This technical documentation provides a comprehensive guide for developing the Kick&Roar frontend application. Key highlights:

1. **Architecture**: Clean separation of concerns with feature-based organization
2. **State Management**: Zustand for client state, React Query for server state
3. **Performance**: Code splitting, lazy loading, virtual scrolling, image optimization
4. **Mobile-First**: Responsive design with dedicated mobile navigation
5. **Type Safety**: Full TypeScript coverage with strict typing
6. **Testing**: Comprehensive testing strategy with unit and integration tests
7. **Error Handling**: Robust error boundaries and API error handling
8. **PWA Features**: Offline capability, installable app experience

The frontend is designed to be scalable, maintainable, and performant, providing an excellent user experience for football enthusiasts in Dhaka and beyond.

---

**Document Version:** 1.0  
**Last Updated:** November 17, 2025  
**Maintained By:** Shoaib Hassan
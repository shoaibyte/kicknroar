import { useState, useCallback, useMemo, useRef, useEffect } from 'react';
import { GoogleMap, Marker, InfoWindow, useJsApiLoader } from '@react-google-maps/api';
import { useGeolocation } from '@/hooks/useGeolocation';
import { useVenueStore } from '@/store/venueStore';
import { DEFAULT_LOCATION, DEFAULT_RADIUS } from '@/lib/constants';
import { Input } from '@/components/ui/input';
import { Button } from '@/components/ui/button';
import { Card } from '@/components/ui/card';
import { Search, Filter, Plus } from 'lucide-react';
import { Link } from 'react-router-dom';
import type { Venue } from '@/api/endpoints/venue.api';
import { matchApi } from '@/api/endpoints/match.api';

const mapContainerStyle = {
  width: '100%',
  height: '600px',
  borderRadius: '12px',
  background: 'linear-gradient(to bottom right, #f0fdf4, #dcfce7)',
};

const defaultCenter = {
  lat: DEFAULT_LOCATION.lat,
  lng: DEFAULT_LOCATION.lng,
};

const mapOptions: google.maps.MapOptions = {
  disableDefaultUI: false,
  zoomControl: true,
  streetViewControl: false,
  mapTypeControl: false,
  fullscreenControl: true,
  styles: [
    {
      featureType: 'poi.business',
      elementType: 'labels',
      stylers: [{ visibility: 'off' }],
    },
  ],
};

interface VenueWithMatches extends Venue {
  matchesToday?: number;
  ongoingMatches?: number;
  open24Hours?: boolean;
}

export const VenueMap: React.FC = () => {
  const [selectedVenue, setSelectedVenue] = useState<VenueWithMatches | null>(null);
  const [mapCenter, setMapCenter] = useState(defaultCenter);
  const [mapZoom, setMapZoom] = useState(13);
  const [searchQuery, setSearchQuery] = useState('');
  const searchInputRef = useRef<HTMLInputElement>(null);
  const autocompleteRef = useRef<google.maps.places.Autocomplete | null>(null);
  const [venueMatches, setVenueMatches] = useState<Record<string, { today: number; ongoing: number }>>({});

  const { latitude, longitude, loading: geoLoading } = useGeolocation();
  const { nearbyVenues, fetchNearbyVenues } = useVenueStore();

  const googleMapsApiKey = import.meta.env.VITE_GOOGLE_MAPS_API_KEY;
  
  const { isLoaded } = useJsApiLoader({
    id: 'google-map-script',
    googleMapsApiKey: googleMapsApiKey || '',
    libraries: ['places'],
  });

  // Set user location when available - only fetch once when loading completes
  useEffect(() => {
    // Wait for geolocation to finish loading before making any decisions
    if (geoLoading) return;

    if (latitude && longitude) {
      setMapCenter({ lat: latitude, lng: longitude });
      fetchNearbyVenues(latitude, longitude, DEFAULT_RADIUS);
    } else {
      // Use default location if geolocation not available or denied
      fetchNearbyVenues(defaultCenter.lat, defaultCenter.lng, DEFAULT_RADIUS);
    }
  }, [latitude, longitude, geoLoading, fetchNearbyVenues]);

  // Fetch matches for venues
  useEffect(() => {
    const fetchMatches = async () => {
      if (nearbyVenues.length === 0) return;

      const today = new Date().toISOString().split('T')[0];
      const matchesMap: Record<string, { today: number; ongoing: number }> = {};

      try {
        const response = await matchApi.getMatches({
          latitude: mapCenter.lat,
          longitude: mapCenter.lng,
          radius: DEFAULT_RADIUS,
          status: 'open',
        });

        response.data.forEach((match) => {
          const venueId = match.venue?.id;
          if (!venueId) return;
          if (!matchesMap[venueId]) {
            matchesMap[venueId] = { today: 0, ongoing: 0 };
          }

          const matchDate = match.match_date.split('T')[0];
          const isToday = matchDate === today;
          
          if (isToday) {
            matchesMap[venueId].today += 1;
            
            // Check if match is actually ongoing (happening now)
            // This requires checking if current time is between start_time and end_time
            // For now, we'll only count matches that are open and today as potentially ongoing
            // A more accurate implementation would check the actual time window
            if (match.status === 'open') {
              // Note: This is a simplified check. For true "ongoing" status,
              // you'd need to compare current time with match.start_time and match.end_time
              matchesMap[venueId].ongoing += 1;
            }
          }
        });

        setVenueMatches(matchesMap);
      } catch (error) {
        console.error('Error fetching matches:', error);
      }
    };

    fetchMatches();
  }, [nearbyVenues, mapCenter]);

  // Initialize Autocomplete when map is loaded
  useEffect(() => {
    if (isLoaded && searchInputRef.current && !autocompleteRef.current) {
      const autocomplete = new google.maps.places.Autocomplete(searchInputRef.current, {
        types: ['geocode'],
        componentRestrictions: { country: 'bd' },
      });

      const placeChangedListener = autocomplete.addListener('place_changed', () => {
        const place = autocomplete.getPlace();
        if (place.geometry?.location) {
          const newCenter = {
            lat: place.geometry.location.lat(),
            lng: place.geometry.location.lng(),
          };
          setMapCenter(newCenter);
          setMapZoom(15);
          fetchNearbyVenues(newCenter.lat, newCenter.lng, DEFAULT_RADIUS);
          setSearchQuery(place.formatted_address || '');
        }
      });

      autocompleteRef.current = autocomplete;

      // Cleanup: remove listener on unmount or when dependencies change
      return () => {
        if (placeChangedListener) {
          google.maps.event.removeListener(placeChangedListener);
        }
        autocompleteRef.current = null;
      };
    }
  }, [isLoaded, fetchNearbyVenues]);

  const mapRef = useRef<google.maps.Map | null>(null);
  const idleTimeoutRef = useRef<NodeJS.Timeout | null>(null);

  const handleMapLoad = useCallback((map: google.maps.Map) => {
    mapRef.current = map;
  }, []);

  // Debounced handler to prevent excessive API calls during zoom/pan operations
  const handleMapIdle = useCallback(() => {
    // Clear any pending timeout
    if (idleTimeoutRef.current) {
      clearTimeout(idleTimeoutRef.current);
    }

    // Debounce the API call - only fetch after map has been idle for 500ms
    idleTimeoutRef.current = setTimeout(() => {
      if (mapRef.current) {
        const center = mapRef.current.getCenter();
        if (center) {
          const newCenter = {
            lat: center.lat(),
            lng: center.lng(),
          };
          setMapCenter(newCenter);
          fetchNearbyVenues(newCenter.lat, newCenter.lng, DEFAULT_RADIUS);
        }
      }
    }, 500);
  }, [fetchNearbyVenues]);

  // Cleanup timeout on unmount
  useEffect(() => {
    return () => {
      if (idleTimeoutRef.current) {
        clearTimeout(idleTimeoutRef.current);
      }
    };
  }, []);

  const venuesWithMatches = useMemo(() => {
    return nearbyVenues.map((venue) => {
      const matches = venueMatches[venue.id] || { today: 0, ongoing: 0 };
      return {
        ...venue,
        matchesToday: matches.today,
        ongoingMatches: matches.ongoing,
      };
    });
  }, [nearbyVenues, venueMatches]);

  const markers = useMemo(() => {
    return venuesWithMatches.map((venue) => (
      <Marker
        key={venue.id}
        position={{
          lat: venue.latitude,
          lng: venue.longitude,
        }}
        onClick={() => setSelectedVenue(venue)}
        icon={{
          path: google.maps.SymbolPath.CIRCLE,
          scale: 8,
          fillColor: '#4CAF50',
          fillOpacity: 1,
          strokeColor: '#ffffff',
          strokeWeight: 2,
        }}
      />
    ));
  }, [venuesWithMatches]);

  if (!googleMapsApiKey || googleMapsApiKey.trim() === '') {
    return (
      <div className="w-full h-[600px] rounded-xl bg-gradient-to-br from-green-50 to-green-100 flex items-center justify-center">
        <div className="text-center p-4">
          <p className="text-muted-foreground mb-2">Google Maps API key is not configured</p>
          <p className="text-sm text-muted-foreground/70">
            Please add <code className="bg-muted px-2 py-1 rounded">VITE_GOOGLE_MAPS_API_KEY</code> to your <code className="bg-muted px-2 py-1 rounded">.env</code> file
          </p>
        </div>
      </div>
    );
  }

  if (!isLoaded) {
    return (
      <div className="w-full h-[600px] rounded-xl bg-gradient-to-br from-green-50 to-green-100 flex items-center justify-center">
        <p className="text-muted-foreground">Loading map...</p>
      </div>
    );
  }

  return (
    <div className="w-full">
      {/* Search and Action Bar */}
      <div className="flex flex-col sm:flex-row gap-3 mb-4">
        <div className="relative flex-1">
          <Search className="absolute left-3 top-1/2 -translate-y-1/2 text-muted-foreground w-4 h-4" />
          <Input
            ref={searchInputRef}
            type="text"
            placeholder="Search location, area..."
            value={searchQuery}
            onChange={(e) => setSearchQuery(e.target.value)}
            className="pl-10"
          />
        </div>
        <Button variant="outline" className="bg-green-50 border-green-200 text-green-700 hover:bg-green-100">
          <Filter className="w-4 h-4 mr-2" />
          Filter
        </Button>
        <Link to="/matches/create">
          <Button className="bg-primary text-white hover:bg-primary/90">
            <Plus className="w-4 h-4 mr-2" />
            Create Event
          </Button>
        </Link>
      </div>

      {/* Map Container */}
      <div className="relative">
        <GoogleMap
          mapContainerStyle={mapContainerStyle}
          center={mapCenter}
          zoom={mapZoom}
          options={mapOptions}
          onLoad={handleMapLoad}
          onDragEnd={handleMapIdle}
          onIdle={handleMapIdle}
        >
          {/* User Location Marker */}
          {latitude && longitude && (
            <Marker
              position={{ lat: latitude, lng: longitude }}
              icon={{
                path: google.maps.SymbolPath.CIRCLE,
                scale: 10,
                fillColor: '#2196F3',
                fillOpacity: 1,
                strokeColor: '#ffffff',
                strokeWeight: 3,
              }}
              title="You are here"
            />
          )}

          {/* Venue Markers */}
          {markers}

          {/* Info Window for Selected Venue */}
          {selectedVenue && (
            <InfoWindow
              position={{
                lat: selectedVenue.latitude,
                lng: selectedVenue.longitude,
              }}
              onCloseClick={() => setSelectedVenue(null)}
            >
              <div className="p-2 min-w-[200px]">
                <h3 className="font-semibold text-lg mb-1">{selectedVenue.name}</h3>
                {selectedVenue.matchesToday && selectedVenue.matchesToday > 0 ? (
                  <p className="text-sm text-muted-foreground">
                    {selectedVenue.matchesToday} {selectedVenue.matchesToday === 1 ? 'match' : 'matches'} today
                  </p>
                ) : selectedVenue.ongoingMatches && selectedVenue.ongoingMatches > 0 ? (
                  <p className="text-sm text-muted-foreground">
                    {selectedVenue.ongoingMatches} match ongoing
                  </p>
                ) : (
                  <p className="text-sm text-muted-foreground">No matches scheduled today</p>
                )}
              </div>
            </InfoWindow>
          )}
        </GoogleMap>

        {/* Venue Cards Overlay - positioned at top */}
        {venuesWithMatches.length > 0 && (
          <div className="absolute top-4 left-4 right-4 z-10 flex flex-wrap gap-2 pointer-events-none max-w-md">
            {venuesWithMatches.slice(0, 3).map((venue) => (
              <Card
                key={venue.id}
                className="pointer-events-auto bg-white p-3 shadow-lg hover:shadow-xl transition-all cursor-pointer border border-gray-100"
                onClick={() => {
                  setSelectedVenue(venue);
                  setMapCenter({ lat: venue.latitude, lng: venue.longitude });
                  setMapZoom(15);
                }}
              >
                <h4 className="font-semibold text-sm mb-1">{venue.name}</h4>
                {venue.matchesToday && venue.matchesToday > 0 ? (
                  <p className="text-xs text-muted-foreground">
                    {venue.matchesToday} {venue.matchesToday === 1 ? 'match' : 'matches'} today
                  </p>
                ) : venue.ongoingMatches && venue.ongoingMatches > 0 ? (
                  <p className="text-xs text-muted-foreground">{venue.ongoingMatches} match ongoing</p>
                ) : (
                  <p className="text-xs text-muted-foreground">No matches today</p>
                )}
              </Card>
            ))}
          </div>
        )}
      </div>
    </div>
  );
};


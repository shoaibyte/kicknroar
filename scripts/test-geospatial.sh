#!/bin/bash

# Geospatial Query Testing Script
# Requires PostgreSQL with PostGIS extension

echo "🗺️  Testing Geospatial Queries (PostGIS)"
echo "=========================================="
echo ""

# Check if DATABASE_URL is set
if [ -z "$DATABASE_URL" ]; then
    echo "❌ DATABASE_URL environment variable not set"
    echo "Set it to: postgres://user:pass@localhost:5432/dbname?sslmode=disable"
    exit 1
fi

echo "Testing PostGIS functions..."
echo ""

# Test find_nearby_venues function
echo "1. Testing find_nearby_venues() function"
psql "$DATABASE_URL" -c "
SELECT * FROM find_nearby_venues(
    23.8103,  -- Dhanmondi latitude
    90.4125,  -- Dhanmondi longitude
    5.0,      -- 5km radius
    NULL,     -- Any field type
    10        -- Limit 10
);
" || echo "❌ Function test failed"
echo ""

# Test find_nearby_matches function
echo "2. Testing find_nearby_matches() function"
psql "$DATABASE_URL" -c "
SELECT * FROM find_nearby_matches(
    23.8103,  -- Dhanmondi latitude
    90.4125,  -- Dhanmondi longitude
    5.0,      -- 5km radius
    CURRENT_DATE,  -- From today
    10        -- Limit 10
);
" || echo "❌ Function test failed"
echo ""

# Test PostGIS location trigger
echo "3. Testing PostGIS location auto-population"
psql "$DATABASE_URL" -c "
INSERT INTO venues (name, address, latitude, longitude, field_type, capacity)
VALUES ('Test Venue', 'Test Address', 23.8103, 90.4125, 'futsal', 10)
RETURNING id, name, ST_AsText(location) as location;
" || echo "❌ Trigger test failed"
echo ""

# Test spatial index
echo "4. Testing spatial index usage"
psql "$DATABASE_URL" -c "
EXPLAIN ANALYZE
SELECT id, name, 
    ST_Distance(
        location::geography,
        ST_SetSRID(ST_MakePoint(90.4125, 23.8103), 4326)::geography
    ) / 1000.0 as distance_km
FROM venues
WHERE ST_DWithin(
    location::geography,
    ST_SetSRID(ST_MakePoint(90.4125, 23.8103), 4326)::geography,
    5000
)
ORDER BY distance_km
LIMIT 10;
" || echo "❌ Spatial query test failed"
echo ""

echo "✅ Geospatial testing complete"
echo ""
echo "Note: These tests require:"
echo "  - PostgreSQL 15+ with PostGIS extension"
echo "  - Sample venue data in the database"
echo "  - Proper DATABASE_URL environment variable"


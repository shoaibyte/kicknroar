-- Enable PostGIS extension
CREATE EXTENSION IF NOT EXISTS postgis;

-- Verify PostGIS installation
SELECT PostGIS_version();

-- Create custom types (enums)
DO $$ BEGIN
CREATE TYPE skill_level_enum AS ENUM ('beginner', 'intermediate', 'advanced', 'professional');
EXCEPTION
    WHEN duplicate_object THEN null;
END $$;

-- Users table is automatically created by Ent
-- This migration file is for PostGIS setup and any manual customizations

-- Create indexes for better performance (if not already created by Ent)
-- These will be created by Ent, but documented here for reference

-- Additional PostGIS-specific configurations
-- Set up spatial reference system (WGS 84 - standard for GPS coordinates)
-- SRID 4326 is already default, but we document it here

-- Verify schema
DO $$
BEGIN
    ASSERT (SELECT EXISTS (
    SELECT FROM pg_extension WHERE extname = 'postgis'
)), 'PostGIS extension must be installed';
END $$;

-- PostGIS stored functions for nearby venue/match queries

-- Find nearby venues function
CREATE OR REPLACE FUNCTION find_nearby_venues(
    p_latitude DECIMAL,
    p_longitude DECIMAL,
    p_radius_km DECIMAL DEFAULT 5.0,
    p_field_type VARCHAR DEFAULT NULL,
    p_limit INTEGER DEFAULT 20
)
RETURNS TABLE (
    id UUID,
    name VARCHAR,
    distance_km DECIMAL,
    field_type VARCHAR,
    rating DECIMAL
) AS $$
BEGIN
    RETURN QUERY
    SELECT 
        v.id,
        v.name,
        ROUND(
            ST_Distance(
                v.location::geography,
                ST_SetSRID(ST_MakePoint(p_longitude, p_latitude), 4326)::geography
            ) / 1000.0,
            2
        ) as distance_km,
        v.field_type::VARCHAR,
        v.rating
    FROM venues v
    WHERE 
        v.is_active = true
        AND (p_field_type IS NULL OR v.field_type::VARCHAR = p_field_type)
        AND ST_DWithin(
            v.location::geography,
            ST_SetSRID(ST_MakePoint(p_longitude, p_latitude), 4326)::geography,
            p_radius_km * 1000
        )
    ORDER BY distance_km ASC
    LIMIT p_limit;
END;
$$ LANGUAGE plpgsql;

-- Find nearby matches function
CREATE OR REPLACE FUNCTION find_nearby_matches(
    p_latitude DECIMAL,
    p_longitude DECIMAL,
    p_radius_km DECIMAL DEFAULT 5.0,
    p_from_date DATE DEFAULT CURRENT_DATE,
    p_limit INTEGER DEFAULT 20
)
RETURNS TABLE (
    match_id UUID,
    match_title VARCHAR,
    venue_name VARCHAR,
    match_date DATE,
    start_time TIME,
    current_players INTEGER,
    max_players INTEGER,
    cost_per_player INTEGER,
    distance_km DECIMAL
) AS $$
BEGIN
    RETURN QUERY
    SELECT 
        m.id as match_id,
        m.title as match_title,
        v.name as venue_name,
        m.match_date,
        m.start_time,
        m.current_players,
        m.max_players,
        m.cost_per_player,
        ROUND(
            ST_Distance(
                v.location::geography,
                ST_SetSRID(ST_MakePoint(p_longitude, p_latitude), 4326)::geography
            ) / 1000.0,
            2
        ) as distance_km
    FROM matches m
    JOIN venues v ON m.venue_id = v.id
    WHERE 
        m.status = 'open'
        AND m.match_date >= p_from_date
        AND ST_DWithin(
            v.location::geography,
            ST_SetSRID(ST_MakePoint(p_longitude, p_latitude), 4326)::geography,
            p_radius_km * 1000
        )
    ORDER BY m.match_date ASC, m.start_time ASC, distance_km ASC
    LIMIT p_limit;
END;
$$ LANGUAGE plpgsql;

-- Trigger to auto-populate PostGIS location from lat/lng for venues
CREATE OR REPLACE FUNCTION update_venue_location()
RETURNS TRIGGER AS $$
BEGIN
    NEW.location = ST_SetSRID(
        ST_MakePoint(NEW.longitude, NEW.latitude), 
        4326
    );
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trigger_update_venue_location
    BEFORE INSERT OR UPDATE OF latitude, longitude ON venues
    FOR EACH ROW
    EXECUTE FUNCTION update_venue_location();

-- Trigger to increment player count when joining
CREATE OR REPLACE FUNCTION increment_match_players()
RETURNS TRIGGER AS $$
BEGIN
    UPDATE matches
    SET current_players = current_players + 1
    WHERE id = NEW.match_id;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trigger_increment_players
    AFTER INSERT ON match_participants
    FOR EACH ROW
    EXECUTE FUNCTION increment_match_players();

-- Trigger to decrement player count when leaving
CREATE OR REPLACE FUNCTION decrement_match_players()
RETURNS TRIGGER AS $$
BEGIN
    UPDATE matches
    SET current_players = current_players - 1
    WHERE id = OLD.match_id;
    RETURN OLD;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trigger_decrement_players
    AFTER DELETE ON match_participants
    FOR EACH ROW
    EXECUTE FUNCTION decrement_match_players();

-- Trigger to auto-update match status
CREATE OR REPLACE FUNCTION update_match_status()
RETURNS TRIGGER AS $$
BEGIN
    IF NEW.current_players >= NEW.max_players THEN
        NEW.status = 'full';
    ELSIF NEW.current_players < NEW.max_players AND NEW.status = 'full' THEN
        NEW.status = 'open';
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trigger_update_match_status
    BEFORE UPDATE OF current_players ON matches
    FOR EACH ROW
    EXECUTE FUNCTION update_match_status();

-- Success message
SELECT 'Initial schema migration completed successfully' AS status;
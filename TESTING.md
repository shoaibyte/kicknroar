# Testing Guide for Kick&Roar Backend

This document outlines the testing strategy and how to run tests for the backend API.

## Test Structure

Tests are organized by layer:
- `internal/util/*_test.go` - Unit tests for utilities
- `internal/pkg/*_test.go` - Unit tests for packages (JWT, validator, AWS)
- `internal/api/middleware/*_test.go` - Middleware tests (auth, CORS, rate limiting)
- `internal/api/handler/*_test.go` - Handler integration tests
- `internal/repository/*_test.go` - Repository tests
- `internal/service/*_test.go` - Service tests

## Running Tests

### Run All Tests
```bash
make test
# or
go test -v ./...
```

### Run Specific Test Package
```bash
go test -v ./internal/util/...
go test -v ./internal/api/middleware/...
go test -v ./internal/repository/...
```

### Run with Coverage
```bash
go test -v -cover ./...
go test -v -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

## Test Categories

### 1. Unit Tests

**Password Utilities** (`internal/util/password_test.go`)
- ✅ Password hashing with bcrypt
- ✅ Password verification
- ✅ Unique hash generation (salt)

**JWT Package** (`internal/pkg/jwt/jwt_test.go`)
- ✅ Access token generation
- ✅ Refresh token generation
- ✅ Token validation
- ✅ Invalid token handling
- ✅ Wrong secret handling

**Error Utilities** (`internal/util/errors_test.go`)
- ✅ Error code definitions
- ✅ Error constructors
- ✅ Error with details

### 2. Middleware Tests

**Authentication Middleware** (`internal/api/middleware/auth_test.go`)
- ✅ Valid token acceptance
- ✅ Missing Authorization header rejection
- ✅ Invalid token format rejection
- ✅ Expired token handling
- ✅ User ID extraction from context

**Rate Limiting** (`internal/api/middleware/rate_limit_test.go`)
- ✅ General rate limiting
- ✅ Auth-specific rate limiting
- ✅ Rate limiter cleanup
- ✅ Too many requests response

**CORS Middleware** (`internal/api/middleware/cors_test.go`)
- ✅ Allowed origin acceptance
- ✅ Invalid origin rejection
- ✅ CORS headers setting

### 3. Repository Tests

**Venue Repository** (`internal/repository/venue_repo_test.go`)
- ✅ Venue creation
- ✅ Find by ID
- ✅ List venues
- ✅ Update venue
- ✅ Not found handling

**Match Repository** (`internal/repository/match_repo_test.go`)
- ✅ Match creation
- ✅ Find by ID
- ✅ List matches with filters
- ✅ Status filtering

**Participant Repository** (`internal/repository/participant_repo_test.go`)
- ✅ Join match
- ✅ Leave match
- ✅ Duplicate join prevention
- ✅ Get participants

### 4. Handler Tests

**Auth Handler** (`internal/api/handler/auth_handler_test.go`)
- ✅ Signup with valid data
- ✅ Signup validation errors
- ✅ Duplicate email handling
- ✅ Login with valid credentials
- ✅ Login with invalid credentials
- ✅ Non-existent user handling

## Test Utilities

The `internal/testutil` package provides:
- `TestConfig()` - Test configuration
- `SetupTestDB()` - SQLite test database (fast)
- `SetupTestDBWithPostgres()` - PostgreSQL test database (for PostGIS)
- `TestJWTManager()` - JWT manager for testing
- `CreateTestUser()` - Helper to create test users
- `CreateTestVenue()` - Helper to create test venues
- `CreateTestMatch()` - Helper to create test matches

## Security Testing

### JWT Security
- ✅ Token generation with correct claims
- ✅ Token validation
- ✅ Expired token rejection
- ✅ Invalid signature rejection
- ✅ Wrong secret rejection

### Rate Limiting
- ✅ Auth endpoints: 5 requests/minute
- ✅ General endpoints: 100 requests/minute
- ✅ Upload endpoints: 10 requests/minute
- ✅ Rate limit cleanup after window

### CORS Security
- ✅ Only allowed origins accepted
- ✅ Credentials allowed
- ✅ Proper headers set

## Error Handling Tests

All handlers test:
- ✅ Validation errors (400)
- ✅ Authentication errors (401)
- ✅ Authorization errors (403)
- ✅ Not found errors (404)
- ✅ Conflict errors (409)
- ✅ Internal server errors (500)

## Geospatial Query Testing

For PostGIS queries, use `SetupTestDBWithPostgres()` with a real PostgreSQL database:

```go
func TestVenueRepository_FindNearby(t *testing.T) {
    // Requires PostgreSQL with PostGIS
    dsn := "postgres://user:pass@localhost:5432/test?sslmode=disable"
    client := testutil.SetupTestDBWithPostgres(t, dsn)
    defer testutil.CleanupTestDB(client)
    
    // Test nearby venue queries
    // ...
}
```

## Integration Testing

For full integration tests, you need:
1. Running PostgreSQL database with PostGIS
2. Set `DATABASE_URL` environment variable
3. Run tests with integration tag:

```bash
go test -v -tags=integration ./...
```

## Manual Testing Checklist

### Authentication Endpoints
- [ ] POST /api/v1/auth/signup - Create account
- [ ] POST /api/v1/auth/login - Login
- [ ] POST /api/v1/auth/refresh - Refresh token
- [ ] POST /api/v1/auth/logout - Logout

### User Endpoints
- [ ] GET /api/v1/users/me - Get profile
- [ ] PUT /api/v1/users/me - Update profile
- [ ] GET /api/v1/users/:id - Get user
- [ ] GET /api/v1/users/:id/stats - Get stats

### Venue Endpoints
- [ ] GET /api/v1/venues - List venues
- [ ] POST /api/v1/venues - Create venue
- [ ] GET /api/v1/venues/nearby - Find nearby (PostGIS)
- [ ] GET /api/v1/venues/:id - Get venue
- [ ] PUT /api/v1/venues/:id - Update venue

### Match Endpoints
- [ ] GET /api/v1/matches - List matches
- [ ] POST /api/v1/matches - Create match
- [ ] GET /api/v1/matches/:id - Get match
- [ ] PUT /api/v1/matches/:id - Update match
- [ ] DELETE /api/v1/matches/:id - Delete match
- [ ] POST /api/v1/matches/:id/join - Join match
- [ ] POST /api/v1/matches/:id/leave - Leave match
- [ ] GET /api/v1/matches/:id/participants - Get participants

### Upload Endpoints
- [ ] POST /api/v1/upload/avatar - Upload avatar
- [ ] POST /api/v1/venues/:id/upload - Upload venue image

## Test Coverage Goals

- Unit tests: >80% coverage
- Integration tests: Critical paths covered
- Security tests: All middleware tested
- Error handling: All error codes tested

## Continuous Integration

Tests should run on:
- Every commit
- Before merging PRs
- Before deployment

## Notes

- SQLite is used for fast unit tests (no PostGIS support)
- PostgreSQL required for geospatial query tests
- Use test fixtures for consistent test data
- Clean up test data after each test


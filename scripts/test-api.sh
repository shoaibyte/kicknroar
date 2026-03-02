#!/bin/bash

# API Testing Script for Kick&Roar Backend
# This script helps test all endpoints manually

BASE_URL="${API_URL:-http://localhost:8080}"
API_BASE="${BASE_URL}/api/v1"

# Colors for output
GREEN='\033[0;32m'
RED='\033[0;31m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

echo "🧪 Kick&Roar Backend API Testing"
echo "=================================="
echo "Base URL: ${BASE_URL}"
echo ""

# Test health check
echo -e "${YELLOW}Testing Health Check...${NC}"
response=$(curl -s -w "\n%{http_code}" "${BASE_URL}/api/v1/health")
http_code=$(echo "$response" | tail -n1)
body=$(echo "$response" | sed '$d')
if [ "$http_code" -eq 200 ]; then
    echo -e "${GREEN}✓ Health check passed${NC}"
else
    echo -e "${RED}✗ Health check failed (HTTP $http_code)${NC}"
fi
echo ""

# Test signup
echo -e "${YELLOW}Testing Signup...${NC}"
SIGNUP_RESPONSE=$(curl -s -X POST "${API_BASE}/auth/signup" \
    -H "Content-Type: application/json" \
    -d '{
        "email": "test@example.com",
        "password": "SecurePass123!",
        "full_name": "Test User",
        "phone": "+8801712345678"
    }')
echo "$SIGNUP_RESPONSE" | jq '.' 2>/dev/null || echo "$SIGNUP_RESPONSE"

# Extract token from signup response
TOKEN=$(echo "$SIGNUP_RESPONSE" | jq -r '.token' 2>/dev/null)
if [ "$TOKEN" != "null" ] && [ -n "$TOKEN" ]; then
    echo -e "${GREEN}✓ Signup successful, token obtained${NC}"
    export TOKEN
else
    echo -e "${RED}✗ Signup failed or token not found${NC}"
    echo "Trying login instead..."
    
    # Try login
    LOGIN_RESPONSE=$(curl -s -X POST "${API_BASE}/auth/login" \
        -H "Content-Type: application/json" \
        -d '{
            "email": "test@example.com",
            "password": "SecurePass123!"
        }')
    TOKEN=$(echo "$LOGIN_RESPONSE" | jq -r '.token' 2>/dev/null)
    if [ "$TOKEN" != "null" ] && [ -n "$TOKEN" ]; then
        echo -e "${GREEN}✓ Login successful${NC}"
        export TOKEN
    else
        echo -e "${RED}✗ Login also failed${NC}"
        exit 1
    fi
fi
echo ""

# Test protected endpoints
if [ -n "$TOKEN" ]; then
    echo -e "${YELLOW}Testing Protected Endpoints...${NC}"
    
    # Get current user
    echo "GET /users/me"
    curl -s -X GET "${API_BASE}/users/me" \
        -H "Authorization: Bearer ${TOKEN}" | jq '.' 2>/dev/null || echo "Failed"
    echo ""
    
    # List venues
    echo "GET /venues"
    curl -s -X GET "${API_BASE}/venues" \
        -H "Authorization: Bearer ${TOKEN}" | jq '.' 2>/dev/null || echo "Failed"
    echo ""
    
    # Find nearby venues
    echo "GET /venues/nearby?lat=23.8103&lng=90.4125&radius=5"
    curl -s -X GET "${API_BASE}/venues/nearby?lat=23.8103&lng=90.4125&radius=5" \
        -H "Authorization: Bearer ${TOKEN}" | jq '.' 2>/dev/null || echo "Failed"
    echo ""
    
    # List matches
    echo "GET /matches"
    curl -s -X GET "${API_BASE}/matches" \
        -H "Authorization: Bearer ${TOKEN}" | jq '.' 2>/dev/null || echo "Failed"
    echo ""
fi

# Test error handling
echo -e "${YELLOW}Testing Error Handling...${NC}"

# Invalid token
echo "Testing with invalid token..."
response=$(curl -s -w "\n%{http_code}" -X GET "${API_BASE}/users/me" \
    -H "Authorization: Bearer invalid-token")
http_code=$(echo "$response" | tail -n1)
if [ "$http_code" -eq 401 ]; then
    echo -e "${GREEN}✓ Invalid token correctly rejected (401)${NC}"
else
    echo -e "${RED}✗ Expected 401, got $http_code${NC}"
fi
echo ""

# Missing token
echo "Testing without token..."
response=$(curl -s -w "\n%{http_code}" -X GET "${API_BASE}/users/me")
http_code=$(echo "$response" | tail -n1)
if [ "$http_code" -eq 401 ]; then
    echo -e "${GREEN}✓ Missing token correctly rejected (401)${NC}"
else
    echo -e "${RED}✗ Expected 401, got $http_code${NC}"
fi
echo ""

# Invalid request data
echo "Testing invalid signup data..."
response=$(curl -s -w "\n%{http_code}" -X POST "${API_BASE}/auth/signup" \
    -H "Content-Type: application/json" \
    -d '{"email": "invalid-email"}')
http_code=$(echo "$response" | tail -n1)
if [ "$http_code" -eq 400 ]; then
    echo -e "${GREEN}✓ Invalid data correctly rejected (400)${NC}"
else
    echo -e "${RED}✗ Expected 400, got $http_code${NC}"
fi
echo ""

echo -e "${GREEN}✅ API Testing Complete${NC}"


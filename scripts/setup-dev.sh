#!/bin/bash

echo "🚀 Setting up Kick&Roar Backend Development Environment"
echo ""

# Check if .env exists
if [ ! -f .env ]; then
    echo "📝 Creating .env file from .env.example..."
    cp .env.example .env

    # Update DATABASE_URL for local development
    if [[ "$OSTYPE" == "darwin"* ]]; then
        # macOS
        sed -i '' 's|^# DATABASE_URL=postgres://kicknroar|DATABASE_URL=postgres://kicknroar|' .env
    else
        # Linux
        sed -i 's|^# DATABASE_URL=postgres://kicknroar|DATABASE_URL=postgres://kicknroar|' .env
    fi

    echo "✅ .env file created"
    echo "⚠️  Please update JWT_SECRET in .env file"
else
    echo "✅ .env file already exists"
fi

# Start Docker containers
echo ""
echo "🐳 Starting PostgreSQL with PostGIS..."
docker-compose up -d

# Wait for database to be ready
echo "⏳ Waiting for database to be ready..."
sleep 5

# Check database health
echo "🔍 Checking database connection..."
until docker exec kicknroar-db pg_isready -U kicknroar > /dev/null 2>&1; do
    echo "   Waiting for database..."
    sleep 2
done

echo "✅ Database is ready!"

# Install dependencies
echo ""
echo "📦 Installing Go dependencies..."
make deps

# Generate Ent code
echo ""
echo "🔧 Generating Ent code..."
make ent-gen

echo ""
echo "🎉 Setup complete!"
echo ""
echo "Next steps:"
echo "1. Update JWT_SECRET in .env file (generate a random string)"
echo "2. Run: make run"
echo "3. The server will create the database schema automatically"
echo ""
EOF

chmod +x scripts/setup-dev.sh

echo "✅ Setup script created!"
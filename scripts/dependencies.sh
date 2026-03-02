# Install core dependencies
go get github.com/labstack/echo/v4
go get github.com/labstack/echo/v4/middleware
go get entgo.io/ent/cmd/ent
go get github.com/lib/pq
go get github.com/golang-jwt/jwt/v5
go get golang.org/x/crypto/bcrypt
go get github.com/aws/aws-sdk-go-v2
go get github.com/aws/aws-sdk-go-v2/config
go get github.com/aws/aws-sdk-go-v2/service/s3
go get github.com/joho/godotenv
go get github.com/go-playground/validator/v10
go get github.com/google/uuid

# Tidy up dependencies
go mod tidy

echo "✅ Dependencies installed!"

# Display the project structure
echo "📁 Final Project Structure:"
tree -L 3 -I 'generated' --dirsfirst

echo ""
echo "🎉 Backend project structure is ready!"
echo ""
echo "Next steps:"
echo "1. Review the created files"
echo "2. Update .env with your credentials"
echo "3. Start implementing the Ent schemas"
echo "4. Generate Ent code: make ent-gen"
echo "5. Start building the handlers and services"
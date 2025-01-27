command_exists() {
    command -v "$1" >/dev/null 2>&1
}

if command_exists docker; then
    echo "✅ Docker is installed."
else
    echo "❌ Docker is not installed. Please install Docker and try again."
    exit 1
fi

if docker info >/dev/null 2>&1; then
    echo "✅ Docker is running."
else
    echo "❌ Docker is not running. Please start Docker and try again."
    exit 1
fi

if command_exists go; then
    echo "✅ Go is installed."
else
    echo "❌ Go is not installed. Please install Go and try again."
    exit 1
fi

echo "🚀 Starting Docker services..."
docker-compose up -d

# 5. Wait for the DB container to initialize
echo "⏳ Waiting for the database to be ready..."
# You can adjust the sleep duration or implement a more robust wait mechanism
sleep 10

# Optionally, you can implement a loop to wait until MySQL is ready
echo "🔍 Checking if MySQL is ready..."
while ! docker exec todo_db_mysql mysqladmin ping -h "localhost" --silent; do
    echo "Waiting for MySQL to be ready..."
    sleep 2
done
echo "✅ MySQL is ready."

echo "🔨 Building the Go project..."
go build -o todo-app main.go

if [ $? -ne 0 ]; then
    echo "❌ Failed to build the Go application."
    exit 1
fi
echo "✅ Go project built successfully."

echo "🚀 Starting the Todo App..."
./todo-app
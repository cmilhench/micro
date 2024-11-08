############################
# STEP 1 build
############################
FROM node:18-alpine AS builder
WORKDIR /app

# Download the dependencies
COPY package*.json ./
RUN npm ci --omit=dev

# Import the source files
COPY . .

############################
# STEP 2 build a small image
############################
FROM node:18-alpine AS runner

# Set working directory and user
WORKDIR /app

# Copy only the necessary files from builder
COPY --from=builder /app .

# Set environment variables
ENV NODE_ENV=production

# Start the service
CMD ["node", "index.js"]

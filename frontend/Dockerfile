FROM oven/bun:1.3.11-slim AS build

# Set the working directory
WORKDIR /app

# Copy the package.json and bun.lock to the working directory
COPY UnoFrontend/package.json /app/
COPY UnoFrontend/bun.lock /app/

# Install dependencies
RUN bun install

# Copy the entire project to the container
COPY UnoFrontend/ .

# Set environment variable explicitly
ARG API_BASE_URL
ENV API_BASE_URL=${API_BASE_URL}

# Build the Nuxt app for production
RUN bun run build

# Use a lightweight web server image for the final stage
FROM oven/bun:1.3.11-slim

# Set the working directory
WORKDIR /app

COPY --from=build /app/.output ./

# Expose the port that Nuxt will run on
EXPOSE 3000

ENV NODE_ENV=production

# Set the default command to run the Nuxt server
CMD ["bun", "run", "server/index.mjs"]
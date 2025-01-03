# Use the official Node.js image as the base image
FROM node:20-alpine AS build

# Set the working directory
WORKDIR /app

# Copy the package.json and package-lock.json to the working directory
COPY UnoFrontend/package.json /app/
COPY UnoFrontend/package-lock.json /app/

# Install dependencies
RUN npm install

# Copy the entire project to the container
COPY UnoFrontend/ .

# Build the Nuxt app for production
RUN npm run build

# Use a lightweight web server image for the final stage
FROM node:20-alpine

# Set the working directory
WORKDIR /app

# Copy the built files and dependencies from the build stage
COPY --from=build /app/.output /app/.output

# Expose the port that Nuxt will run on
EXPOSE 3000

# Set the default command to run the Nuxt server
CMD ["node", ".output/server/index.mjs"]
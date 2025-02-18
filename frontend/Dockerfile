# Use the official Node.js image
FROM node:lts-alpine3.21

# Set the working directory
WORKDIR /app

# Copy package.json and server.js
COPY package.json .
COPY server.js .

# Install dependencies
RUN npm install

# Expose the frontend port
EXPOSE 4000

# Run the frontend application
CMD ["node", "server.js"]
FROM node:20.9.0-alpine 

# Set the working directory in the container to /app
WORKDIR /app

# Copy package.json and pnpm-lock.yaml (if available) into the root directory
COPY package.json pnpm-lock.yaml* ./

# Install pnpm globally in the container
RUN npm install -g pnpm

# Install any needed packages specified in package.json
RUN pnpm install

# Bundle app source inside Docker image
COPY . .

# Make port 8082 available to the world outside this container
EXPOSE 8082

# Define environment variable for Vite
ENV PORT=8082

# Run the app when the container launches
# CMD ["pnpm", "run", "dev", "--host", "--port", "8082"]
CMD ["pnpm", "run", "dev"]

# Use a minimal base image
FROM debian:latest

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the Go executable to the container
COPY ./messagedistributor ./main

# Make the executable executable
RUN chmod +x ./main

# Expose port 8080 to the outside world
EXPOSE 8080

# Verify the copied file and its permissions
RUN ls -ltr ./main

# Command to run the executable
ENTRYPOINT ["/app/main"]

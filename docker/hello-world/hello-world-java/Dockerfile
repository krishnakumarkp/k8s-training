# Stage 1: Compile the Java program
FROM eclipse-temurin:21-jdk AS builder

WORKDIR /app

# Copy the Java source file to the container
COPY HelloWorld.java /app/

# Compile the Java program
RUN javac HelloWorld.java


# Stage 2: Create a minimal runtime image
FROM eclipse-temurin:21-jre-jammy

# Set the working directory to /app
WORKDIR /app

# Copy the compiled Java class from the builder stage to /app
COPY --from=builder /app/HelloWorld.class /app/

# Run the Java program
CMD ["java", "HelloWorld"]
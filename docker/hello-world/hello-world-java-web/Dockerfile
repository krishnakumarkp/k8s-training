# Stage 1: Build the application
FROM maven:3.8.7-openjdk-18 AS build

# Set the working directory in the container
WORKDIR /app

# Copy the Maven configuration files and source code
COPY pom.xml ./
COPY src ./src

# Run Maven build (clean and package the application)
RUN mvn clean package

# Stage 2: Run the application
FROM eclipse-temurin:21-jre-jammy

# Set the working directory in the container
WORKDIR /app

# Copy the built JAR file from the build stage
COPY --from=build /app/target/*.jar app.jar

# Expose the port the Spring Boot application listens on
EXPOSE 8080

# Command to run the application
ENTRYPOINT ["java", "-jar", "app.jar"]
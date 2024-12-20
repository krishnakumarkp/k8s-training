Write a Simple Java Program

```
public class HelloWorld {
    public static void main(String[] args) {
        System.out.println("Hello, World!");
    }
}

```

Compile the Java Program

```
javac HelloWorld.java

```

Create the Dockerfile

```
FROM eclipse-temurin:21-jre-jammy

# Copy the compiled Java class file into the container
COPY HelloWorld.class /app/

# Set the working directory to /app
WORKDIR /app

# Run the Java program
ENTRYPOINT ["java", "HelloWorld"]

```

Build the Docker Image

```
docker build -t java-hello-world .
```

Run the Docker Container

```
docker run --rm java-hello-world
```

I dont have java locally

Compile Locally in Another Image

```
docker run --rm -v "$(pwd)":/app -w /app eclipse-temurin:21-jdk javac HelloWorld.java

``

Multi-Stage Docker file

```
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

```

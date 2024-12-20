 Compile and Package the Application

```
mvn clean package

java -jar target/hello-world-web-0.0.1-SNAPSHOT.jar

```
http://localhost:8080/hello




Using docker

Maven Command to build
```
docker run --rm -v "$(pwd)":/app -w /app maven:3.8.7-openjdk-18 mvn clean install
```

Run

```
docker run --rm -v "$(pwd)":/app -w /app -p 8080:8080 eclipse-temurin:21-jdk java -jar target/hello-world-web-0.0.1-SNAPSHOT.jar

```

Using docker file

```
docker build -t spring-boot-helloworld .

```

```
docker run --rm -p 8080:8080 spring-boot-helloworld

```

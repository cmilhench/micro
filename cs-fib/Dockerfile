############################
# STEP 1 build executable binary
############################
FROM mcr.microsoft.com/dotnet/sdk:8.0-alpine AS build-env
WORKDIR /app

# Download the dependencies
COPY ./Fibonacci/*.csproj ./Fibonacci/
COPY ./FibonacciHttp/*.csproj ./FibonacciHttp/
RUN dotnet restore ./FibonacciHttp/FibonacciHttp.csproj

# Import the source files
COPY . ./

# Build the binary for the application
RUN dotnet publish ./FibonacciHttp/FibonacciHttp.csproj -c Release -o out


############################
# STEP 2 build a small image
############################
FROM mcr.microsoft.com/dotnet/aspnet:8.0-alpine

WORKDIR /app

COPY --from=build-env /app/out .

ENTRYPOINT ["dotnet", "FibonacciHttp.dll"]

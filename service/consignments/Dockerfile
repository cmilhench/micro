# Start from a small base
FROM alpine

# Our application requires no privileges
# so run it with a non-root user
ADD users /etc/passwd
USER nobody

# Our application runs on port 8080
# so allow hosts to bind to that port
EXPOSE 8080

# Add our application binary
ADD app /app

# Run our application!
ENTRYPOINT [ "/app" ]
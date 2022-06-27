## About

This repository is the home for a gateway http service using the
net/http package from the standard library which contains all functionalities
for the HTTP protocol.

Messages are sent to an amqp topic exchange with reply queue and correlationID
to achieve an asyncronus RPC micro-service gateway.

Dependent service listeners will answer request or the service will timeout
after 30 seconds.

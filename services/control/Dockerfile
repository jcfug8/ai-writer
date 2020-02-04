FROM alpine:latest

COPY control_binary /usr/bin/
COPY client /assets

CMD ["control_binary", "-persist=persist:50051", "-ai=ai:50051", "-a=0.0.0.0:80", "-assets=/assets"]
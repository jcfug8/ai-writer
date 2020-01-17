FROM alpine:latest

COPY services/control/control_binary /usr/bin/
COPY client /assets

CMD ["control_binary", "-persist=persist:50051", "-ai=ai:50051", "-a=0.0.0.0:80", "-assests=/assets"]
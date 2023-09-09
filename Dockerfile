FROM golang:alpine
RUN apk add \
          git\
          bash\
          curl
COPY install.sh install.sh
RUN ./install.sh
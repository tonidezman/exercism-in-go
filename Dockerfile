FROM golang:bullseye

COPY . .
RUN ./exercism configure --token=$EXERCISM_TOKEN
RUN ./exercism configure --workspace=/root/exercism

RUN echo "alias exercism='/workspace/exercism'" >> ~/.bashrc

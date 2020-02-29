FROM scratch

COPY ./build/terrain-linux-* /terrain
WORKDIR /module

CMD ["/terrain"]
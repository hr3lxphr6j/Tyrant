FROM golang as build_env
COPY . /tmp/tyrant/
RUN /tmp/tyrant/src/hack/build.sh /tmp/tyrant

FROM debian
COPY --from=build_env /tmp/tyrant/bin/Tyrant.linux /usr/bin/tyrant
VOLUME [ "/data" ]
ENTRYPOINT ["/usr/bin/tyrant"]
CMD ["-bind", ":6655", "-db", "/data/tyrant.db"]
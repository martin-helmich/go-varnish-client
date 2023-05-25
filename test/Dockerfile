ARG VARNISH_VERSION=6.0
FROM varnish:${VARNISH_VERSION}

# this is just for testing purposes
USER root
RUN echo "72be6aba-00c4-4908-a99f-0e4eb7cc86ca" > /etc/varnish/secret

ENTRYPOINT [ "/usr/sbin/varnishd"]

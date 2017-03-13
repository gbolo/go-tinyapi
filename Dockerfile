# Buidling A Tiny but Functional Docker Container
# For Educational Purposes only.
# Companion Blog Post: https://linuxctl.com/2017/03/building-tiny-secure-docker-containers/
FROM scratch
LABEL maintainer "George Bolo <gbolo@linuxctl.com>"

# Install Application
ENV TINYAPI_CFG_PATH /etc/tinyapi
ADD tinyapi /bin/tinyapi
ADD config.yml /etc/tinyapi/config.yml

# This application requires CA bundle for outbound HTTPS calls
ADD cacert-2017-01-18.pem /etc/pki/ca-trust/extracted/pem/tls-ca-bundle.pem

EXPOSE 8080
ENTRYPOINT ["/bin/tinyapi"]

FROM golang:1.23.5

WORKDIR /go/src
ENV PATH="/go/bin:${PATH}"

RUN go install github.com/golang/mock/mockgen@v1.5.0 && \
    go install github.com/spf13/cobra-cli@latest


RUN apt-get update && apt-get install sqlite3 -y

RUN usermod -u 1000 www-data
RUN ln -sf /go/bin/cobra-cli /go/bin/cobra
RUN mkdir -p /var/www/.cache
RUN chown -R www-data:www-data /go
RUN chown -R www-data:www-data /var/www/.cache
USER www-data

CMD ["tail", "-f", "/dev/null"]
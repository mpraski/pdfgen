FROM ubuntu

RUN apt-get -qq update && \
	apt-get -qqy install sudo

RUN mkdir /tmp/wkhtmltox
WORKDIR /tmp/wkhtmltox

ADD https://github.com/wkhtmltopdf/packaging/releases/download/0.12.6-1/wkhtmltox_0.12.6-1.focal_amd64.deb .
RUN sudo apt -qqyf install ./wkhtmltox_0.12.6-1.focal_amd64.deb

WORKDIR /
RUN rm -fr /tmp/wkhtmltox
COPY pdfgen /usr/local/bin/pdfgen

EXPOSE 8888
ENV PDFGEN_PORT=8888
ENV PDFGEN_ADDR=0.0.0.0

CMD ["pdfgen"]

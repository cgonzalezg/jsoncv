FROM python:alpine
WORKDIR /app
ADD . /app
CMD [ "python", "-m", "http.server", "5000"]

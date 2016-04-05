FROM python:alpine
WORDIR /app
ADD . /app
CMD [ "python", "-m", "http.server", "5000"]

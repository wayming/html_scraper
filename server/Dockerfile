# Python 服务端 Dockerfile
FROM python:3.9-slim

WORKDIR /app
COPY server/ /app

RUN pip install --no-cache-dir -r requirements.txt

CMD ["python", "server.py"]

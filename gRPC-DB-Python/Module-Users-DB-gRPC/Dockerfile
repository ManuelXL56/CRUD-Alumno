FROM python:3-alpine

WORKDIR /app

COPY ./ ./

RUN pip install --no-cache-dir -r requirements.txt

ENV PYTHONUNBUFFERED=1

CMD ["python", "main.py"]

EXPOSE 50052
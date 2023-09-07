FROM python:3.8@sha256:c3ac277830cfe6d4b092c9e58f6295ac79f6091e7aa602c505e7c887c6b6f513

WORKDIR /app

COPY requirements.txt ./
RUN pip install --no-cache-dir -r requirements.txt

COPY . .

CMD ["python", "manage.py", "runserver", "0.0.0.0:8000"]


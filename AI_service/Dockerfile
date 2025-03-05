FROM python:3.11
LABEL authors="29935"


WORKDIR /app
COPY requirements.txt .
RUN pip install -r requirements.txt -i https://mirrors.aliyun.com/pypi/simple/
COPY . .

CMD ["python", "app.py"]

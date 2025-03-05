FROM python:3.11
LABEL authors="29935"


WORKDIR /app
COPY requirements.txt .
RUN pip install -r requirements.txt -i https://pypi.mirrors.ustc.edu.cn/simple/
COPY . .

CMD ["python", "app.py"]

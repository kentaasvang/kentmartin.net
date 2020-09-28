FROM python:3.8-alpine
RUN mkdir /app
WORKDIR /app
ADD requirements.txt /app
ADD server.py /app
RUN pip3 install -r requirements.txt
CMD ["gunicorn", "-w 4", "-b", "0.0.0.0:8000", "server:create_app()"]
